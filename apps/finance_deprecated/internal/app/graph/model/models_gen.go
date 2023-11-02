// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"kloudlite.io/pkg/repos"
)

type Account struct {
	Name              string               `json:"name"`
	DisplayName       string               `json:"displayName"`
	Billing           *Billing             `json:"billing"`
	IsActive          bool                 `json:"isActive"`
	ContactEmail      string               `json:"contactEmail"`
	ReadableID        repos.ID             `json:"readableId"`
	Memberships       []*AccountMembership `json:"memberships"`
	Created           string               `json:"created"`
	OutstandingAmount float64              `json:"outstandingAmount"`
}

func (Account) IsEntity() {}

type AccountMembership struct {
	User     *User    `json:"user"`
	Role     string   `json:"role"`
	Account  *Account `json:"account"`
	Accepted bool     `json:"accepted"`
}

type Billing struct {
	CardholderName string                 `json:"cardholderName"`
	Address        map[string]interface{} `json:"address"`
}

type BillingInput struct {
	StripePaymentMethodID string                 `json:"stripePaymentMethodId"`
	CardholderName        string                 `json:"cardholderName"`
	Address               map[string]interface{} `json:"address"`
}

type User struct {
	ID                 repos.ID             `json:"id"`
	AccountMemberships []*AccountMembership `json:"accountMemberships"`
	AccountMembership  *AccountMembership   `json:"accountMembership"`
}

func (User) IsEntity() {}