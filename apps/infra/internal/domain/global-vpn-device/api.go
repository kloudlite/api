package global_vpn_device

import (
	"fmt"
	"time"

	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/wgutils"

	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/iputils"

	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/repos"
	wgv1 "github.com/kloudlite/operator/apis/wireguard/v1"
)

func (d *Domain) claimNextFreeDeviceIP(ctx types.InfraContext, deviceName string, gvpnName string) (string, error) {
	for {
		freeIp, err := d.FreeDeviceIPRepo.FindOne(ctx, repos.Filter{
			fields.AccountName:           ctx.AccountName,
			fc.FreeDeviceIPGlobalVPNName: gvpnName,
		})
		if err != nil {
			return "", err
		}

		if freeIp == nil {
			gvpn, err := d.FindGlobalVPN(ctx, gvpnName)
			if err != nil {
				return "", err
			}

			ip, err := iputils.GetIPAddrInARange(gvpn.CIDR, gvpn.NumAllocatedDevices+1, gvpn.NumReservedIPsForNonClusterUse)
			if err != nil {
				return "", err
			}

			if _, err := d.FreeDeviceIPRepo.Create(ctx, &entities.FreeDeviceIP{
				AccountName:   ctx.AccountName,
				GlobalVPNName: gvpnName,
				IPAddr:        ip,
			}); err != nil {
				continue
			}

			if err := d.IncrementGlobalVPNAllocatedDevices(ctx, gvpn.Id, 1); err != nil {
				continue
			}

			// if _, err := d.gvpnRepo.PatchById(ctx, gvpn.Id, repos.Document{"$inc": map[string]any{fc.GlobalVPNNumAllocatedDevices: 1}}); err != nil {
			// 	continue
			// }

			continue
		}

		ipAddr := freeIp.IPAddr

		if _, err := d.ClaimDeviceIPRepo.Create(ctx, &entities.ClaimDeviceIP{
			AccountName:   ctx.AccountName,
			GlobalVPNName: gvpnName,
			IPAddr:        ipAddr,
			ClaimedBy:     deviceName,
		}); err != nil {
			d.Logger.Warnf("ip addr already claimed (err: %s), retrying again", err.Error())
			<-time.After(50 * time.Millisecond)
			continue
		}

		if err := d.FreeDeviceIPRepo.DeleteById(ctx, freeIp.Id); err != nil {
			return "", err
		}

		return ipAddr, nil
	}
}

func (d *Domain) UpdateGlobalVPNDevice(ctx types.InfraContext, device entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error) {
	panic("implement me")
}

func (d *Domain) deleteGlobalVPNDevice(ctx types.InfraContext, gvpn string, deviceName string) error {
	device, err := d.findGlobalVPNDevice(ctx, gvpn, deviceName)
	if err != nil {
		return err
	}

	if err := d.ClaimDeviceIPRepo.DeleteOne(ctx, repos.Filter{
		fc.AccountName:                  ctx.AccountName,
		fc.GlobalVPNDeviceGlobalVPNName: gvpn,
		fc.ClaimDeviceIPClaimedBy:       deviceName,
	}); err != nil {
		return err
	}

	if _, err := d.FreeDeviceIPRepo.Create(ctx, &entities.FreeDeviceIP{
		AccountName:   ctx.AccountName,
		GlobalVPNName: gvpn,
		IPAddr:        device.IPAddr,
	}); err != nil {
		return err
	}

	if err := d.GlobalVPNDeviceRepo.DeleteById(ctx, device.Id); err != nil {
		return err
	}

	// if err := d.reconGlobalVPNConnections(ctx, device.GlobalVPNName); err != nil {
	// 	return err
	// }

	if err := d.SyncGlobalVPNConnections(ctx, device.GlobalVPNName); err != nil {
		return err
	}

	return nil
}

func (d *Domain) DeleteGlobalVPNDevice(ctx types.InfraContext, gvpn string, deviceName string) error {
	return d.deleteGlobalVPNDevice(ctx, gvpn, deviceName)
}

func (d *Domain) ListGlobalVPNDevice(ctx types.InfraContext, gvpn string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.GlobalVPNDevice], error) {
	filter := d.GlobalVPNDeviceRepo.MergeMatchFilters(
		repos.Filter{
			fc.AccountName:                  ctx.AccountName,
			fc.GlobalVPNDeviceGlobalVPNName: gvpn,
		},
		map[string]repos.MatchFilter{
			fc.GlobalVPNDeviceCreationMethod: {
				MatchType:  repos.MatchTypeNotInArray,
				NotInArray: []any{types.GlobalVPNConnectionDeviceMethod, types.KloudliteGlobalVPNDeviceMethod},
			},
		},
		search,
	)
	return d.GlobalVPNDeviceRepo.FindPaginated(ctx, filter, pagination)
}

func (d *Domain) CreateGlobalVPNDevice(ctx types.InfraContext, gvpnDevice entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error) {
	return d.createGlobalVPNDevice(ctx, gvpnDevice)
}

func (d *Domain) createGlobalVPNDevice(ctx types.InfraContext, gvpnDevice entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error) {
	gvpnDevice.AccountName = ctx.AccountName
	gvpnDevice.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	gvpnDevice.LastUpdatedBy = gvpnDevice.CreatedBy

	privateKey, publicKey, err := wgutils.GenerateKeyPair()
	if err != nil {
		return nil, err
	}

	gvpnDevice.PrivateKey = privateKey
	gvpnDevice.PublicKey = publicKey

	ip, err := d.claimNextFreeDeviceIP(ctx, gvpnDevice.Name, gvpnDevice.GlobalVPNName)
	if err != nil {
		return nil, err
	}

	gvpnDevice.IPAddr = ip

	gv, err := d.GlobalVPNDeviceRepo.Create(ctx, &gvpnDevice)
	if err != nil {
		return nil, err
	}

	if err := d.SyncGlobalVPNConnections(ctx, gvpnDevice.GlobalVPNName); err != nil {
		return nil, err
	}

	return gv, nil
}

func (d *Domain) BuildPeersFromGlobalVPNDevices(ctx types.InfraContext, gvpn string) (publicPeers []wgv1.Peer, privatePeers []wgv1.Peer, err error) {
	devices, err := d.GlobalVPNDeviceRepo.Find(ctx, repos.Query{
		Filter: map[string]any{
			fc.AccountName:                  ctx.AccountName,
			fc.GlobalVPNDeviceGlobalVPNName: gvpn,
			fc.GlobalVPNDeviceCreationMethod: map[string]any{
				"$ne": types.GlobalVPNConnectionDeviceMethod,
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	publicPeers = make([]wgv1.Peer, 0, 10) // 10 is just a random low number
	privatePeers = make([]wgv1.Peer, 0, len(devices))
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

func (d *Domain) GetGlobalVPNDevice(ctx types.InfraContext, gvpn string, gvpnDevice string) (*entities.GlobalVPNDevice, error) {
	if gvpn == "" || gvpnDevice == "" {
		return nil, errors.New("invalid global vpn or device")
	}

	return d.findGlobalVPNDevice(ctx, gvpn, gvpnDevice)
}

func (d *Domain) GetGlobalVPNDeviceWgConfig(ctx types.InfraContext, gvpn string, gvpnDevice string) (string, error) {
	return d.getGlobalVPNDeviceWgConfig(ctx, gvpn, gvpnDevice, nil)
}

func (d *Domain) getGlobalVPNDeviceWgConfig(ctx types.InfraContext, gvpn string, gvpnDevice string, postUp []string) (string, error) {
	device, err := d.findGlobalVPNDevice(ctx, gvpn, gvpnDevice)
	if err != nil {
		return "", err
	}

	gvpnConns, err := d.ListGlobalVPNConnections(ctx, gvpn)
	if err != nil {
		return "", err
	}

	gvpnConnPeers, err := d.BuildGlobalVPNConnectionPeers(ctx, gvpnConns)
	if err != nil {
		return "", err
	}

	pubPeers, privPeers, err := d.BuildPeersFromGlobalVPNDevices(ctx, gvpn)
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
		return "", errors.Newf("no DNS server found for global VPN device %s", gvpn)
	}

	config, err := wgutils.GenerateWireguardConfig(wgutils.WgConfigParams{
		IPAddr:     device.IPAddr,
		PrivateKey: device.PrivateKey,
		DNS:        dnsServer,
		PostUp:     postUp,
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

func (d *Domain) findGlobalVPNDevice(ctx types.InfraContext, gvpn string, gvpnDevice string) (*entities.GlobalVPNDevice, error) {
	device, err := d.GlobalVPNDeviceRepo.FindOne(ctx, repos.Filter{
		fc.AccountName:                  ctx.AccountName,
		fc.GlobalVPNDeviceGlobalVPNName: gvpn,
		fc.MetadataName:                 gvpnDevice,
	})
	if err != nil {
		return nil, err
	}

	if device == nil {
		return nil, errors.Newf("no global vpn device (name=%s) found", gvpnDevice)
	}
	return device, nil
}
