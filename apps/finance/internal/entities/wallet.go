package entities

import (
	"github.com/kloudlite/api/pkg/repos"
	"time"
)

type Wallet struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`

	TeamId    string    `json:"teamId"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var WalletIndices = []repos.IndexField{
	{
		Field:  []repos.IndexKey{{Key: "id", Value: repos.IndexAsc}},
		Unique: true,
	},
	{
		Field:  []repos.IndexKey{{Key: "teamId", Value: repos.IndexAsc}},
		Unique: true,
	},
}
