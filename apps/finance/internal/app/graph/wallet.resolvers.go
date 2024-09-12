package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/kloudlite/api/apps/finance/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/finance/internal/entities"
)

// CreationTime is the resolver for the creationTime field.
func (r *walletResolver) CreationTime(ctx context.Context, obj *entities.Wallet) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("not found: Wallet")
	}
	return obj.CreationTime.String(), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *walletResolver) UpdateTime(ctx context.Context, obj *entities.Wallet) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("not found: Wallet")
	}
	return obj.UpdateTime.String(), nil
}

// Wallet returns generated.WalletResolver implementation.
func (r *Resolver) Wallet() generated.WalletResolver { return &walletResolver{r} }

type walletResolver struct{ *Resolver }
