package clusters

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	*common.Domain

	ClusterRepo  repos.DbRepo[*entities.Cluster]
	MsgOfficeSvc ports.MessageOfficeInternalSvc

	GlobalVPNDomain
	GlobalVPNConnectionDomain

	GlobalVPNDeviceDomain

	BYOKClusterDomain
	ProviderSecretDomain

	NodepoolDomain
	PVDomain

	HelmReleaseDomain
}

type GlobalVPNDomain struct {
	EnsureGlobalVPN func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error)
	FindGlobalVPN   func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error)
}

type GlobalVPNConnectionDomain struct {
	EnsureGlobalVPNConnection func(ctx types.InfraContext, clusterName string, groupName string, clusterPublicEndpoint string) (*entities.GlobalVPNConnection, error)
	FindGlobalVPNConnection   func(ctx types.InfraContext, gvpnName string, clusterName string) (*entities.GlobalVPNConnection, error)
	DeleteGlobalVPNConnection func(ctx types.InfraContext, gvpnName string, clusterName string) error
}

type BYOKClusterDomain struct {
	FindBYOKCluster func(ctx types.InfraContext, clusterName string) (*entities.BYOKCluster, error)
}

type ProviderSecretDomain struct {
	FindProviderSecret func(ctx types.InfraContext, clusterName string) (*entities.CloudProviderSecret, error)
}

type NodepoolDomain struct {
	CountNodepools func(ctx types.InfraContext, clusterName string) (int, error)
}

type PVDomain struct {
	CountPVs func(ctx types.InfraContext, clusterName string) (int, error)
}

type GlobalVPNDeviceDomain struct {
	GetGlobalVPNDeviceWgConfig func(ctx types.InfraContext, gvpnName string, deviceName string) (string, error)
}

type HelmReleaseDomain struct {
	UpsertKloudliteHelmRelease func(ctx types.InfraContext, clusterName string, hr *entities.HelmRelease) (*entities.HelmRelease, error)
}
