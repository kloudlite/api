package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/kloudlite/api/apps/finance/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

// FinanceCreatePayment is the resolver for the finance_createPayment field.
func (r *mutationResolver) FinanceCreatePayment(ctx context.Context, payment entities.Payment) (*entities.Payment, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.domain.CreatePayment(uc, &payment)
}

// FinanceValidatePayment is the resolver for the finance_validatePayment field.
func (r *mutationResolver) FinanceValidatePayment(ctx context.Context, paymentID repos.ID) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, err
	}

	if err := r.domain.ValidatePayment(uc, paymentID); err != nil {
		return false, err
	}

	return true, nil
}

// FinanceGetWallet is the resolver for the finance_getWallet field.
func (r *queryResolver) FinanceGetWallet(ctx context.Context) (*entities.Wallet, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.domain.GetWallet(uc)
}

// FinanceListPayments is the resolver for the finance_listPayments field.
func (r *queryResolver) FinanceListPayments(ctx context.Context, walletID repos.ID) ([]*entities.Payment, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, err
	}

	p, err := r.domain.ListPayments(uc)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// FinanceListCharges is the resolver for the finance_listCharges field.
func (r *queryResolver) FinanceListCharges(ctx context.Context) ([]*entities.Charge, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.domain.ListCharges(uc)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
