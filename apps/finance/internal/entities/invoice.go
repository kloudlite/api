package entities

import (
	"github.com/kloudlite/api/pkg/repos"
)

type Invoice struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`

	TeamId    string `json:"teamId" graphql:"teamId"`
	Amount    int    `json:"amount" graphql:"amount"`
	Currency  string `json:"currency" graphql:"currency"`
	DueDate   string `json:"dueDate" graphql:"dueDate"`
	Status    string `json:"status" graphql:"status"`
	CreatedAt string `json:"createdAt" graphql:"createdAt"`
	UpdatedAt string `json:"updatedAt" graphql:"updatedAt"`
}

var InvoiceIndices = []repos.IndexField{
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
