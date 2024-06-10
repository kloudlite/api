package provider_secret

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	*common.Domain

	ProviderSecretRepo repos.DbRepo[*entities.CloudProviderSecret]

	ClusterDomain
}

type ClusterDomain struct {
	CountClustersWithProviderSecret func(ctx types.InfraContext, providerSecretName string) (int64, error)
}

type AWSAccessValidationOutput struct {
	Result          bool
	InstallationURL *string
}
