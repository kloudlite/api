package entities

import (
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"

	"kloudlite.io/pkg/repos"
)

type Token struct {
	repos.BaseEntity `json:",inline"`
	crdsv1.Secret    `json:",inline"`

	Token       string `json:"token"`
	NodeId      string `json:"nodeId"`
	AccountName string `json:"accountName"`
	ClusterName string `json:"clusterName"`
}

var TokenIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "nodeId", Value: repos.IndexAsc},
			{Key: "accountName", Value: repos.IndexAsc},
			{Key: "clusterName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
