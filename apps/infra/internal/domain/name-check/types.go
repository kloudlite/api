package name_check

import "github.com/kloudlite/api/apps/infra/internal/domain/types"

type Domain struct {
	ClusterDomain
	GlobalVPNDeviceDomain
	NodepoolDomain
	HelmReleaseDomain
	ClusterManagedServiceDomain
	ProviderSecretDomain
}

type ClusterDomain struct {
	IsClusterNameAvailable func(ctx types.InfraContext, name string) (bool, error)
}

type GlobalVPNDeviceDomain struct {
	IsGlobalVPNDeviceNameAvailable func(ctx types.InfraContext, name string) (bool, error)
}

type NodepoolDomain struct {
	IsNodepoolNameAvailable func(ctx types.InfraContext, clusterName string, name string) (bool, error)
}

type HelmReleaseDomain struct {
	IsHelmReleaseNameAvailable func(ctx types.InfraContext, clusterName string, name string) (bool, error)
}

type ClusterManagedServiceDomain struct {
	IsClusterManagedSvcNameAvailable func(ctx types.InfraContext, clusterName string, name string) (bool, error)
}

type ProviderSecretDomain struct {
	IsProviderSecretNameAvailable func(ctx types.InfraContext, name string) (bool, error)
}
