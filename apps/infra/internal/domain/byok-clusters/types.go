package byok_clusters

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	*common.Domain

	BYOKClusterRepo repos.DbRepo[*entities.BYOKCluster]

	MsgOfficeSvc ports.MessageOfficeInternalSvc

	GlobalVPNDomain
	GlobalVPNConnectionDomain

	ClusterDomain
}

type GlobalVPNDomain struct {
	EnsureGlobalVPN func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error)
	FindGlobalVPN   func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error)
}

type ClusterDomain struct {
	ClusterAlreadyExists func(ctx types.InfraContext, name string) (bool, error)
}

type GlobalVPNConnectionDomain struct {
	EnsureGlobalVPNConnection func(ctx types.InfraContext, clusterName string, groupName string, clusterPublicEndpoint string) (*entities.GlobalVPNConnection, error)
	FindGlobalVPNConnection   func(ctx types.InfraContext, gvpnName string, clusterName string) (*entities.GlobalVPNConnection, error)
	DeleteGlobalVPNConnection func(ctx types.InfraContext, gvpnName string, clusterName string) error
}
