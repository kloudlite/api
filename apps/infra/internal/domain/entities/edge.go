package entities

import (
	infraV1 "github.com/kloudlite/cluster-operator/apis/infra/v1"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

type Edge struct {
	repos.BaseEntity `bson:",inline"`
	infraV1.Edge     `json:",inline" bson:",inline"`
	AccountName      string     `json:"accountName"`
	ClusterName      string     `json:"clusterName"`
	SyncStatus       t.SyncStatus `json:"syncStatus"`
}

var EdgeIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "metadata.name", Value: repos.IndexAsc},
		},
		Unique: true,
	},

	{
		Field: []repos.IndexKey{
			{Key: "accountName", Value: repos.IndexAsc},
		},
	},
	{
		Field: []repos.IndexKey{
			{Key: "clusterName", Value: repos.IndexAsc},
		},
	},
}