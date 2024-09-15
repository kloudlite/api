package entities

import (
	"time"

	"github.com/kloudlite/api/pkg/repos"
)

type PaymentStatus string

const (
	PaymentStatusCreated       PaymentStatus = "created"
	PaymentStatusPartiallyPaid PaymentStatus = "partially_paid"
	PaymentStatusExpired       PaymentStatus = "expired"
	PaymentStatusCancelled     PaymentStatus = "cancelled"
	PaymentStatusPaid          PaymentStatus = "paid"
)

type PaymentLink struct {
	Id          string        `json:"id"`
	ReferenceId string        `json:"reference_id"`
	Status      PaymentStatus `json:"status"`
	ShortUrl    string        `json:"short_url"`
}

type Payment struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	CreatedAt        time.Time `json:"createdAt" graphql:"noinput"`
	UpdatedAt        time.Time `json:"updatedAt" graphql:"noinput"`

	TeamId string `json:"teamId" graphql:"noinput"`

	Amount   float64 `json:"amount"`
	Currency string  `json:"currency" graphql:"noinput"`

	Link *PaymentLink `json:"payment_link" graphql:"noinput"`
}

var PaymentIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
