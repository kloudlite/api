package entities

import (
	"github.com/kloudlite/api/pkg/repos"
)

type Wallet struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`

	TeamId    string `json:"teamId" graphql:"teamId"`
	Balance   int    `json:"balance" graphql:"balance"`
	Currency  string `json:"currency" graphql:"currency"`
	CreatedAt string `json:"createdAt" graphql:"createdAt"`
	UpdatedAt string `json:"updatedAt" graphql:"updatedAt"`
}

var WalletIndices = []repos.IndexField{
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
