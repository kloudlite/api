package entities

import (
	"time"

	"github.com/kloudlite/api/pkg/repos"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusSuccess PaymentStatus = "success"
)

type Payment struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	CreatedAt        time.Time `json:"createdAt" graphql:"noinput"`
	UpdatedAt        time.Time `json:"updatedAt" graphql:"noinput"`

	TeamId   string   `json:"teamId"`
	WalletId repos.ID `json:"walletId"`

	Amount   int    `json:"amount"`
	Currency string `json:"currency"`

	Status PaymentStatus `json:"status" graphql:"noinput"`
}

var PaymentIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
