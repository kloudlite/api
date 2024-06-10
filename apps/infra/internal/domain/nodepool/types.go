package nodepool

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
	corev1 "k8s.io/api/core/v1"
)

type Domain struct {
	*common.Domain

	NodepoolRepo repos.DbRepo[*entities.NodePool]

	ClusterDomain
	ProviderSecretDomain
}

type ClusterDomain struct {
	FindCluster func(ctx types.InfraContext, name string) (*entities.Cluster, error)
}

type ProviderSecretDomain struct {
	FindProviderSecret           func(ctx types.InfraContext, name string) (*entities.CloudProviderSecret, error)
	GetProviderSecretAsK8sSecret func(ctx types.InfraContext, name string) (*corev1.Secret, error)
}
