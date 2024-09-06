package entities

import (
	"github.com/kloudlite/api/pkg/repos"
)

type Subscription struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`

	Seats  int    `json:"seats" graphql:"seats"`
	TeamId string `json:"teamId" graphql:"teamId"`
	Status string `json:"status" graphql:"status"`

	UpdatedAt string `json:"updatedAt" graphql:"updatedAt"`
	CreatedAt string `json:"createdAt" graphql:"createdAt"`
}

var SubscriptionIndices = []repos.IndexField{
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
			{Key: "targetNamespace", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
