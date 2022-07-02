// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"kloudlite.io/pkg/repos"
)

type Account struct {
	ID           repos.ID             `json:"id"`
	Name         string               `json:"name"`
	Billing      *Billing             `json:"billing"`
	IsActive     bool                 `json:"isActive"`
	ContactEmail string               `json:"contactEmail"`
	ReadableID   repos.ID             `json:"readableId"`
	Memberships  []*AccountMembership `json:"memberships"`
	Created      string               `json:"created"`
}

func (Account) IsEntity() {}

type AccountMembership struct {
	User     *User    `json:"user"`
	Role     string   `json:"role"`
	Account  *Account `json:"account"`
	Accepted bool     `json:"accepted"`
}

type Billing struct {
	StripeCustomerID string                 `json:"stripeCustomerId"`
	CardholderName   string                 `json:"cardholderName"`
	Address          map[string]interface{} `json:"address"`
}

type BillingInput struct {
	StripeSetupIntentID string                 `json:"stripeSetupIntentId"`
	StripePaymentMethod string                 `json:"stripePaymentMethod"`
	CardholderName      string                 `json:"cardholderName"`
	Address             map[string]interface{} `json:"address"`
}

type ComputePlan struct {
	Name           string  `json:"name"`
	SharedPrice    float64 `json:"sharedPrice"`
	DedicatedPrice float64 `json:"dedicatedPrice"`
}

func (ComputePlan) IsEntity() {}

type User struct {
	ID                 repos.ID             `json:"id"`
	AccountMemberships []*AccountMembership `json:"accountMemberships"`
	AccountMembership  *AccountMembership   `json:"accountMembership"`
}

func (User) IsEntity() {}
