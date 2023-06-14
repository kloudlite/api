package entities

import (
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

type MRes struct {
	repos.BaseEntity       `json:",inline"`
	crdsv1.ManagedResource `json:",inline" json-schema:"uri=k8s://managedresources.crds.kloudlite.io"`
	AccountName            string       `json:"accountName"`
	ClusterName            string       `json:"clusterName"`
	SyncStatus             t.SyncStatus `json:"syncStatus" graphql:"common=true"`
}

var MresIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "metadata.name", Value: repos.IndexAsc},
			{Key: "metadata.namespace", Value: repos.IndexAsc},
			{Key: "clusterName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "accountName", Value: repos.IndexAsc},
		},
	},
}
