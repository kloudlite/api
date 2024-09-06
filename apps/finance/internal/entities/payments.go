package entities

import (
	"github.com/kloudlite/api/pkg/repos"
)

type Payment struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`

	TeamId        string `json:"teamId"`
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
	Method        string `json:"method"`
	TransactionId string `json:"transactionId"`
	CreatedAt     string `json:"createdAt"`
}

var PaymentIndices = []repos.IndexField{
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
