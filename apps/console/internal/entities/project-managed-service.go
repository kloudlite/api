package entities

import (
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	corev1 "k8s.io/api/core/v1"
)

type ProjectManagedService struct {
	repos.BaseEntity             `json:",inline" graphql:"noinput"`
	crdsv1.ProjectManagedService `json:",inline"`

	AccountName string `json:"accountName" graphql:"noinput"`
	ProjectName string `json:"projectName" graphql:"noinput"`

	SyncedOutputSecretRef *corev1.Secret `json:"syncedOutputSecretRef" graphql:"ignore"`

	common.ResourceMetadata `json:",inline"`
	SyncStatus              t.SyncStatus `json:"syncStatus" graphql:"noinput"`
}

var ProjectManagedServiceIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "metadata.name", Value: repos.IndexAsc},
			{Key: "accountName", Value: repos.IndexAsc},
			{Key: "projectName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
