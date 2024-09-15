package entities

import (
	"github.com/kloudlite/api/pkg/repos"
)

type ChargeStatus string

const (
	ChargeStatusPending ChargeStatus = "pending"
	ChargeStatusFailed  ChargeStatus = "failed"
	ChargeStatusSuccess ChargeStatus = "success"
)

type Charge struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	TeamId           string   `json:"teamId"`
	// WalletId         repos.ID `json:"walletId"`

	Amount   int    `json:"amount"`
	Currency string `json:"currency"`

	Description string `json:"description"`

	Status ChargeStatus `json:"status" graphql:"noinput"`
}

var ChargeIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
