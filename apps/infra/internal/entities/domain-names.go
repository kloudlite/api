package entities

import (
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/pkg/operator"
)

type DomainEntry struct {
	repos.BaseEntity        `json:",inline" graphql:"noinput"`
	common.ResourceMetadata `json:",inline"`

	DomainName string `json:"domainName"`

	AccountName string `json:"accountName" graphql:"noinput"`
	ClusterName string `json:"clusterName"`
}

func (d DomainEntry) GetDisplayName() string {
	return d.ResourceMetadata.DisplayName
}

func (d DomainEntry) GetStatus() operator.Status {
	return operator.Status{}
}

var DomainEntryIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
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
			{Key: "domainName", Value: repos.IndexAsc},
			{Key: "clusterName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
