package global_vpn_connection

import (
	"fmt"

	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"

	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	wgv1 "github.com/kloudlite/operator/apis/wireguard/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListGlobalVPNConnectionsArgs struct {
	GlobalVPNName           string
	GlobalVPNConnectionRepo repos.DbRepo[*entities.GlobalVPNConnection]
}

func listGlobalVPNConnections(ctx types.InfraContext, args ListGlobalVPNConnectionsArgs) ([]*entities.GlobalVPNConnection, error) {
	return args.GlobalVPNConnectionRepo.Find(ctx, repos.Query{
		Filter: repos.Filter{
			fc.AccountName:  ctx.AccountName,
			fc.MetadataName: args.GlobalVPNName,
		},
	})
}

func ListGlobalVPNConnections(ctx types.InfraContext, args ListGlobalVPNConnectionsArgs) ([]*entities.GlobalVPNConnection, error) {
	return listGlobalVPNConnections(ctx, args)
}

func GenerateGlobalVPNConnectionPeers(ctx types.InfraContext, conns []*entities.GlobalVPNConnection) ([]wgv1.Peer, error) {
	peers := make([]wgv1.Peer, 0, len(conns))
	for _, c := range conns {
		if c.ParsedWgParams != nil {
			if c.ParsedWgParams.WgPublicKey == "" {
				continue
			}

			if c.ParsedWgParams.NodePort == nil {
				ctx.Logger.Infof("nodeport not available yet, for gvpn %s", c.Name)
				continue
			}

			peers = append(peers, wgv1.Peer{
				ClusterName: c.ClusterName,
				IP:          c.ParsedWgParams.IP,
				PublicKey:   c.ParsedWgParams.WgPublicKey,
				Endpoint:    fmt.Sprintf("%s:%s", c.ClusterPublicEndpoint, *c.ParsedWgParams.NodePort),
				AllowedIPs:  []string{c.ClusterSvcCIDR},
			})
		}
	}

	return peers, nil
}

type ApplyGlobalVPNConnectionArgs struct {
	ResDispatcher       ports.ResourceDispatcher
	GlobalVPNConnection *entities.GlobalVPNConnection
}

func applyGlobalVPNConnection(ctx types.InfraContext, args ApplyGlobalVPNConnectionArgs) error {
	gvpn := args.GlobalVPNConnection
	if err := args.ResDispatcher.ApplyToTargetCluster(ctx, ctx.AccountName, gvpn.ClusterName, &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      gvpn.Spec.WgRef.Name,
			Namespace: gvpn.Spec.WgRef.Namespace,
		},
		StringData: map[string]string{
			"ip": gvpn.DeviceRef.IPAddr,
		},
	}, 0); err != nil {
		return err
	}
	return args.ResDispatcher.ApplyToTargetCluster(ctx, gvpn.AccountName, gvpn.ClusterName, &gvpn.GlobalVPN, gvpn.RecordVersion)
}

type SyncGlobalVPNConnectionsArgs struct {
	GlobalVPNName           string
	GlobalVPNConnectionRepo repos.DbRepo[*entities.GlobalVPNConnection]

	GeneratePeersFromGlobalVPNDevices func(ctx types.InfraContext, vpnName string) (publicPeers []wgv1.Peer, privatePeers []wgv1.Peer, err error)

	ResDispatcher ports.ResourceDispatcher
}

func SyncGlobalVPNConnections(ctx types.InfraContext, args SyncGlobalVPNConnectionsArgs) error {
	gvpnConns, err := listGlobalVPNConnections(ctx, ListGlobalVPNConnectionsArgs{
		GlobalVPNName:           args.GlobalVPNName,
		GlobalVPNConnectionRepo: args.GlobalVPNConnectionRepo,
	})
	if err != nil {
		return errors.NewE(err)
	}

	peers, err := GenerateGlobalVPNConnectionPeers(ctx, gvpnConns)
	if err != nil {
		return err
	}

	publicPeers, privatePeers, err := args.GeneratePeersFromGlobalVPNDevices(ctx, args.GlobalVPNName)
	if err != nil {
		return err
	}

	peers = append(peers, publicPeers...)
	peers = append(peers, privatePeers...)

	for _, xcc := range gvpnConns {
		if fmt.Sprintf("%#v", xcc.Spec.Peers) == fmt.Sprintf("%#v", peers) {
			continue
		}

		xcc.Spec.Peers = peers
		unp, err := args.GlobalVPNConnectionRepo.Patch(
			ctx,
			repos.Filter{
				fc.AccountName:  ctx.AccountName,
				fc.ClusterName:  xcc.ClusterName,
				fc.MetadataName: xcc.Name,
			},
			common.PatchForUpdate(ctx, xcc, common.PatchOpts{
				XPatch: map[string]any{
					fc.GlobalVPNConnectionSpecPeers: peers,
				},
			}),
		)
		if err != nil {
			return errors.NewE(err)
		}

		if err := applyGlobalVPNConnection(ctx, ApplyGlobalVPNConnectionArgs{
			ResDispatcher:       args.ResDispatcher,
			GlobalVPNConnection: unp,
		}); err != nil {
			return errors.NewE(err)
		}
	}

	return nil
}

type GetGlobalVPNConnectionArgs struct{
  GlobalVPNConnectionName string
  GlobalVPNConnectionRepo repos.DbRepo[*entities.GlobalVPNConnection]
}

func GetGlobalVPNConnection(ctx types.InfraContext, args GetGlobalVPNConnectionArgs) (*entities.GlobalVPNConnection, error) {
  args.GlobalVPNConnectionRepo.FindOne(ctx, repos.Filter)
}
