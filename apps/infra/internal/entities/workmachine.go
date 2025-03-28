package entities

import (
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/repos"
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
)

type Workmachine struct {
	repos.BaseEntity        `json:",inline" graphql:"noinput"`
	common.ResourceMetadata `json:",inline"`
	crdsv1.WorkMachine      `json:",inline"`

	AccountName string `json:"accountName" graphql:"noinput"`
	ClusterName string `json:"clusterName" graphql:"noinput"`
}

var WorkmachineIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: fields.Id, Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{
				Key:   fields.MetadataName,
				Value: repos.IndexAsc,
			},
			{
				Key:   fields.AccountName,
				Value: repos.IndexAsc,
			},
			{
				Key:   fields.ClusterName,
				Value: repos.IndexAsc,
			},
		},
		Unique: true,
	},
}
