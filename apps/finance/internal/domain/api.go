package domain

import (
	"context"

	iamT "kloudlite.io/apps/iam/types"
	"kloudlite.io/pkg/repos"
)

type FinanceContext struct {
	context.Context
	UserId repos.ID
}

type Domain interface {
	// CRUD
	CreateAccount(ctx FinanceContext, name string) (*Account, error)
	GetAccount(ctx FinanceContext, name string) (*Account, error)
	UpdateAccount(ctx FinanceContext, name string, email *string) (*Account, error)
	DeleteAccount(ctx FinanceContext, name string) (bool, error)

	DeactivateAccount(ctx FinanceContext, name string) (bool, error)
	ActivateAccount(ctx FinanceContext, name string) (bool, error)

	// Memberships
	AddAccountMember(ctx FinanceContext, accountName string, email string, role iamT.Role) (bool, error)

	RemoveAccountMember(ctx FinanceContext, accountName string, userId repos.ID) (bool, error)

	UpdateAccountMember(ctx FinanceContext, accountName string, userId repos.ID, role string) (bool, error)

	GetUserMemberships(ctx FinanceContext, resourceRef string) ([]*Membership, error)
	GetAccountMemberships(ctx FinanceContext, userId repos.ID) ([]*Membership, error)
	GetAccountMembership(ctx FinanceContext, userId repos.ID, accountName string) (*Membership, error)
}
