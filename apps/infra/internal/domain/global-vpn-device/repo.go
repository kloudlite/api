package global_vpn_device

import (
	"fmt"
	"time"

	"github.com/kloudlite/api/apps/infra/internal/domain/templates"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/iputils"
	"github.com/kloudlite/api/pkg/k8s"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/api/pkg/wgutils"
	wgv1 "github.com/kloudlite/operator/apis/wireguard/v1"
)

type ClaimNextFreeDeviceIPArgs struct {
	Logger logging.Logger

	GlobalVPNName     string
	DeviceName        string
	FreeDeviceIPRepo  repos.DbRepo[*entities.FreeDeviceIP]
	ClaimDeviceIPRepo repos.DbRepo[*entities.ClaimDeviceIP]

	FindGlobalVPN func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error)

	IncrementGlobalVPNAllocatedDevices func(ctx types.InfraContext, gvpnId repos.ID, incrementBy int) error
}

func ClaimNextFreeDeviceIP(ctx types.InfraContext, args ClaimNextFreeDeviceIPArgs) (string, error) {
	for {
		freeIp, err := args.FreeDeviceIPRepo.FindOne(ctx, repos.Filter{
			fc.AccountName:               ctx.AccountName,
			fc.FreeDeviceIPGlobalVPNName: args.GlobalVPNName,
		})
		if err != nil {
			return "", err
		}

		if freeIp == nil {
			gvpn, err := args.FindGlobalVPN(ctx, args.GlobalVPNName)
			if err != nil {
				return "", err
			}

			ip, err := iputils.GetIPAddrInARange(gvpn.CIDR, gvpn.NumAllocatedDevices+1, gvpn.NumReservedIPsForNonClusterUse)
			if err != nil {
				return "", err
			}

			if _, err := args.FreeDeviceIPRepo.Create(ctx, &entities.FreeDeviceIP{
				AccountName:   ctx.AccountName,
				GlobalVPNName: args.GlobalVPNName,
				IPAddr:        ip,
			}); err != nil {
				continue
			}

			if err := args.IncrementGlobalVPNAllocatedDevices(ctx, gvpn.Id, 1); err != nil {
				continue
			}

			// if _, err := d.gvpnRepo.PatchById(ctx, gvpn.Id, repos.Document{"$inc": map[string]any{fc.GlobalVPNNumAllocatedDevices: 1}}); err != nil {
			// 	continue
			// }

			continue
		}

		ipAddr := freeIp.IPAddr

		if _, err := args.ClaimDeviceIPRepo.Create(ctx, &entities.ClaimDeviceIP{
			AccountName:   ctx.AccountName,
			GlobalVPNName: args.GlobalVPNName,
			IPAddr:        ipAddr,
			ClaimedBy:     args.DeviceName,
		}); err != nil {
			args.Logger.Warnf("ip addr already claimed (err: %s), retrying again", err.Error())
			<-time.After(50 * time.Millisecond)
			continue
		}

		if err := args.FreeDeviceIPRepo.DeleteById(ctx, freeIp.Id); err != nil {
			return "", err
		}

		return ipAddr, nil
	}
}

type CreateGlobalVPNDeviceArgs struct {
	Logger logging.Logger
	Device entities.GlobalVPNDevice

	FreeDeviceIPRepo  repos.DbRepo[*entities.FreeDeviceIP]
	ClaimDeviceIPRepo repos.DbRepo[*entities.ClaimDeviceIP]

	GlobalVPNDeviceRepo repos.DbRepo[*entities.GlobalVPNDevice]

	GetGlobalVPN                       func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error)
	IncrementGlobalVPNAllocatedDevices func(ctx types.InfraContext, gvpnId repos.ID, incrementBy int) error

	SyncGlobalVPNConnections func(ctx types.InfraContext, gvpnName string) error
}

func CreateGlobalVPNDevice(ctx types.InfraContext, args CreateGlobalVPNDeviceArgs) (*entities.GlobalVPNDevice, error) {
	args.Device.AccountName = ctx.AccountName
	args.Device.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	args.Device.LastUpdatedBy = args.Device.CreatedBy

	privateKey, publicKey, err := wgutils.GenerateKeyPair()
	if err != nil {
		return nil, err
	}

	args.Device.PrivateKey = privateKey
	args.Device.PublicKey = publicKey

	ip, err := ClaimNextFreeDeviceIP(ctx, ClaimNextFreeDeviceIPArgs{
		Logger:            args.Logger,
		GlobalVPNName:     args.Device.GlobalVPNName,
		DeviceName:        args.Device.Name,
		FreeDeviceIPRepo:  args.FreeDeviceIPRepo,
		ClaimDeviceIPRepo: args.ClaimDeviceIPRepo,
		FindGlobalVPN: func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
			return args.GetGlobalVPN(ctx, gvpnName)
		},
		IncrementGlobalVPNAllocatedDevices: func(ctx types.InfraContext, gvpnId repos.ID, incrementBy int) error {
			return args.IncrementGlobalVPNAllocatedDevices(ctx, gvpnId, incrementBy)
		},
	})
	if err != nil {
		return nil, err
	}

	args.Device.IPAddr = ip

	gv, err := args.GlobalVPNDeviceRepo.Create(ctx, &args.Device)
	if err != nil {
		return nil, err
	}

	if err := args.SyncGlobalVPNConnections(ctx, args.Device.GlobalVPNName); err != nil {
		return nil, err
	}

	return gv, nil
}

type GeneratePeersFromGlobalVPNDevicesArgs struct {
	GlobalVPNName       string
	GlobalVPNDeviceRepo repos.DbRepo[*entities.GlobalVPNDevice]
}

func GeneratePeersFromGlobalVPNDevices(ctx types.InfraContext, args GeneratePeersFromGlobalVPNDevicesArgs) ([]wgv1.Peer, []wgv1.Peer, error) {
	devices, err := args.GlobalVPNDeviceRepo.Find(ctx, repos.Query{
		Filter: map[string]any{
			fc.AccountName:                  ctx.AccountName,
			fc.GlobalVPNDeviceGlobalVPNName: args.GlobalVPNName,
			fc.GlobalVPNDeviceCreationMethod: map[string]any{
				"$ne": types.GlobalVPNConnectionDeviceMethod,
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	publicPeers := make([]wgv1.Peer, 0, 10) // 10 is just a random low number
	privatePeers := make([]wgv1.Peer, 0, len(devices))
	for i := range devices {
		if devices[i].PublicEndpoint != nil {
			publicPeers = append(publicPeers, wgv1.Peer{
				PublicKey:  devices[i].PublicKey,
				Endpoint:   *devices[i].PublicEndpoint,
				IP:         devices[i].IPAddr,
				DeviceName: devices[i].Name,
				AllowedIPs: []string{fmt.Sprintf("%s/32", devices[i].IPAddr)},
			})
			continue
		}

		privatePeers = append(privatePeers, wgv1.Peer{
			PublicKey:  devices[i].PublicKey,
			IP:         devices[i].IPAddr,
			DeviceName: devices[i].Name,
			AllowedIPs: []string{fmt.Sprintf("%s/32", devices[i].IPAddr)},
		})
	}

	return publicPeers, privatePeers, nil
}

type SyncKloudliteGlobalVPNDeviceArgs struct {
	GlobalVPNName       string
	GlobalVPNDeviceRepo repos.DbRepo[*entities.GlobalVPNDevice]

	KubeReverseProxyImage string

	GetAccountNamespace func(ctx types.InfraContext) (string, error)
	FindGlobalVPN       func(ctx types.InfraContext, name string) (*entities.GlobalVPN, error)

	K8sClient k8s.Client
}

func SyncKloudliteGlobalVPNDeviceOnPlatform(ctx types.InfraContext, args SyncKloudliteGlobalVPNDeviceArgs) error {
	b, err := templates.Read(templates.GlobalVPNKloudliteDeviceTemplate)
	if err != nil {
		return errors.NewE(err)
	}
	accNs, err := args.GetAccountNamespace(ctx)
	if err != nil {
		return errors.NewE(err)
	}

	gv, err := args.FindGlobalVPN(ctx, args.GlobalVPNName)
	if err != nil {
		return err
	}

	if gv.KloudliteDevice.Name == "" {
		return nil
	}

	// 2. Grab wireguard config from that device
	// wgConfig, err := d.getGlobalVPNDeviceWgConfig(ctx, gv.Name, gv.KloudliteDevice.Name, nil)
	wgConfig, err := GetGlobalVPNDeviceWgConfig(ctx, GetGlobalVPNDeviceWgConfigArgs{
		GlobalVPNDeviceRepo: nil,
		GlobalVPN:           gv.Name,
		GlobalVPNDevice:     gv.KloudliteDevice.Name,
	})
	if err != nil {
		return err
	}

	deploymentBytes, err := templates.ParseBytes(b, templates.GVPNKloudliteDeviceTemplateVars{
		Name:                  fmt.Sprintf("kloudlite-device-proxy-%s", gv.Name),
		Namespace:             accNs,
		WgConfig:              wgConfig,
		KubeReverseProxyImage: args.KubeReverseProxyImage,
	})
	if err != nil {
		return err
	}

	if err := args.K8sClient.ApplyYAML(ctx, deploymentBytes); err != nil {
		return errors.NewE(err)
	}

	return nil
}

type findGlobalVPNDeviceArgs struct {
	GlobalVPNDeviceRepo repos.DbRepo[*entities.GlobalVPNDevice]

	GlobalVPN       string
	GlobalVPNDevice string
}

func findGlobalVPNDevice(ctx types.InfraContext, args findGlobalVPNDeviceArgs) (*entities.GlobalVPNDevice, error) {
	return args.GlobalVPNDeviceRepo.FindOne(ctx, entities.UniqueGlobalVPNDevice(ctx.AccountName, args.GlobalVPN, args.GlobalVPNDevice))
}

type GetGlobalVPNDeviceWgConfigArgs struct {
	GlobalVPNDeviceRepo repos.DbRepo[*entities.GlobalVPNDevice]

	GlobalVPN       string
	GlobalVPNDevice string

	ListGlobalVPNConnections    func(ctx types.InfraContext, gvpn string) ([]*entities.GlobalVPNConnection, error)
	GenGlobalVPNConnectionPeers func(ctx types.InfraContext, conns []*entities.GlobalVPNConnection) ([]*wgv1.Peer, error)
	GenGlobalVPNDevicePeers     func(ctx types.InfraContext, gvpn string) (public []*wgv1.Peer, private []*wgv1.Peer, err error)
}

func GetGlobalVPNDeviceWgConfig(ctx types.InfraContext, args GetGlobalVPNDeviceWgConfigArgs) (string, error) {
	device, err := findGlobalVPNDevice(ctx, findGlobalVPNDeviceArgs{
		GlobalVPNDeviceRepo: args.GlobalVPNDeviceRepo,
		GlobalVPN:           args.GlobalVPN,
		GlobalVPNDevice:     args.GlobalVPNDevice,
	})
	if err != nil {
		return "", err
	}

	gvpnConns, err := args.ListGlobalVPNConnections(ctx, args.GlobalVPN)
	if err != nil {
		return "", err
	}

	gvpnConnPeers, err := args.GenGlobalVPNConnectionPeers(ctx, gvpnConns)
	if err != nil {
		return "", err
	}

	pubPeers, privPeers, err := args.GenGlobalVPNDevicePeers(ctx, args.GlobalVPN)
	if err != nil {
		return "", err
	}

	pubPeers = append(gvpnConnPeers, pubPeers...)

	publicPeers := make([]wgutils.PublicPeer, 0, len(pubPeers))
	for _, peer := range pubPeers {
		publicPeers = append(publicPeers, wgutils.PublicPeer{
			PublicKey:  peer.PublicKey,
			AllowedIPs: peer.AllowedIPs,
			Endpoint:   peer.Endpoint,
			IPAddr:     peer.IP,
		})
	}

	privatePeers := make([]wgutils.PrivatePeer, 0, len(privPeers))
	for _, peer := range privatePeers {
		privatePeers = append(privatePeers, wgutils.PrivatePeer{
			PublicKey:  peer.PublicKey,
			AllowedIPs: peer.AllowedIPs,
		})
	}

	dnsServer := ""
	for i := range gvpnConns {
		if gvpnConns[i].ParsedWgParams != nil && gvpnConns[i].ParsedWgParams.DNSServer != nil {
			dnsServer = *gvpnConns[i].ParsedWgParams.DNSServer
		}
	}

	if dnsServer == "" {
		return "", errors.Newf("no DNS server found for global VPN %s", args.GlobalVPN)
	}

	config, err := wgutils.GenerateWireguardConfig(wgutils.WgConfigParams{
		IPAddr:     device.IPAddr,
		PrivateKey: device.PrivateKey,
		DNS:        dnsServer,
		PostUp:     nil,
		// PostUp: []string{
		// 	"sudo iptables -A INPUT -i wg0 -j DROP",
		// },
		PublicPeers:  publicPeers,
		PrivatePeers: privatePeers,
	})
	if err != nil {
		return "", err
	}

	return config, nil
}
