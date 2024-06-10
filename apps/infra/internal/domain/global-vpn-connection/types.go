package global_vpn_connection

import (
	"time"

	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"

	wgv1 "github.com/kloudlite/operator/apis/wireguard/v1"
)

type Domain struct {
	*common.Domain

	GlobalVPNClusterConnectionRepo repos.DbRepo[*entities.GlobalVPNConnection]
	FreeClusterSvcCIDRRepo         repos.DbRepo[*entities.FreeClusterSvcCIDR]
	ClaimClusterSvcCIDRRepo        repos.DbRepo[*entities.ClaimClusterSvcCIDR]

	GlobalVPNDomain
	GlobalVPNDeviceDomain
	BYOKDomain
}

type GlobalVPNDomain struct {
	FindGlobalVPN func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error)
}

type GlobalVPNDeviceDomain struct {
	CreateGlobalVPNDevice func(ctx types.InfraContext, device entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error)
	DeleteGlobalVPNDevice func(ctx types.InfraContext, gvpnName string, deviceName string) error

	BuildPeersFromGlobalVPNDevices func(ctx types.InfraContext, gvpnName string) (publicPeers []wgv1.Peer, privatePeer []wgv1.Peer, err error)

  SyncKloudliteDeviceOnPlatform func(ctx types.InfraContext, gvpnName string) error
}

type BYOKDomain struct {
	IsBYOKCluster        func(ctx types.InfraContext, name string) bool
	MarkBYOKClusterReady func(ctx types.InfraContext, name string, timestamp time.Time) error
}
