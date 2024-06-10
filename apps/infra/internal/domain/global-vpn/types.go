package global_vpn

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	*common.Domain
	GlobalVPNRepo repos.DbRepo[*entities.GlobalVPN]

	GlobalVPNDevice
	ClusterDomain
}

type GlobalVPNDevice struct {
	CreateGlobalVPNDevice func(types.InfraContext, entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error)
}

type ClusterDomain struct {
	CountClustersInGlobalVPN func(ctx types.InfraContext, gvpnName string) (int64, error)
}
