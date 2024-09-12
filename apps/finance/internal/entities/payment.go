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

type PaymentLink struct {
	Id          string `json:"id"`
	ReferenceId string `json:"reference_id"`
	Status      string `json:"status"`
	ShortUrl    string `json:"short_url"`
}

type Payment struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	CreatedAt        time.Time `json:"createdAt" graphql:"noinput"`
	UpdatedAt        time.Time `json:"updatedAt" graphql:"noinput"`

	TeamId string `json:"teamId" graphql:"noinput"`

	Amount   int    `json:"amount"`
	Currency string `json:"currency" graphql:"noinput"`

	// Status PaymentStatus `json:"status" graphql:"noinput"`

	Link *PaymentLink `json:"payment_link"`
}

var PaymentIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
