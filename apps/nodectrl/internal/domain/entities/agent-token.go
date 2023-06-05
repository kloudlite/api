package entities

import (
	"kloudlite.io/pkg/repos"
)

type Token struct {
	repos.BaseEntity `json:",inline"`

	JoinToken   string `json:"join%oken"`
	EndpointUrl string `json:"endpointUrl" yaml:"endPointUrl"`
	KubeConfig  string `json:"kubeConfig" yaml:"kubeConfig"`

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
