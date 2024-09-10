package entities

import (
	"github.com/kloudlite/api/pkg/repos"
)

type SubscriptionStatus string

const (
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusPending   SubscriptionStatus = "pending"
	SubscriptionStatusSuspended SubscriptionStatus = "suspended"
)

type Subscription struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`

	Seats  int                `json:"seats"`
	TeamId string             `json:"teamId"`
	Status SubscriptionStatus `json:"status"`

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
