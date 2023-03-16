package entities

import (
	cmgrV1 "github.com/kloudlite/cluster-operator/apis/cmgr/v1"
	"kloudlite.io/pkg/repos"
)

type Cluster struct {
	repos.BaseEntity `bson:",inline" json:",inline"`
	cmgrV1.Cluster   `json:",inline"`
	AccountName      string `json:"accountName"`
}

var ClusterIndices = []repos.IndexField{
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
}
