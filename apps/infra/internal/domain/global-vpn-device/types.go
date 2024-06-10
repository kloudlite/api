package global_vpn_device

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
	wgv1 "github.com/kloudlite/operator/apis/wireguard/v1"
)

type Domain struct {
	*common.Domain

	GlobalVPNDeviceRepo repos.DbRepo[*entities.GlobalVPNDevice]
	FreeDeviceIPRepo    repos.DbRepo[*entities.FreeDeviceIP]
	ClaimDeviceIPRepo   repos.DbRepo[*entities.ClaimDeviceIP]

	GlobalVPNDomain
	GlobalVPNConnectionDomain
}

type GlobalVPNDomain struct {
	FindGlobalVPN func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error)

	IncrementGlobalVPNAllocatedDevices func(ctx types.InfraContext, value int) error
}

type GlobalVPNConnectionDomain struct {
	SyncGlobalVPNConnections      func(ctx types.InfraContext, gvpnName string) error
	ListGlobalVPNConnections      func(ctx types.InfraContext, gvpnName string) ([]*entities.GlobalVPNConnection, error)
	BuildGlobalVPNConnectionPeers func(vpns []*entities.GlobalVPNConnection) ([]wgv1.Peer, error)
}
