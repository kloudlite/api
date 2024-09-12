package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/api/apps/finance/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/finance/internal/app/graph/model"
	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

// CreatedAt is the resolver for the createdAt field.
func (r *chargeResolver) CreatedAt(ctx context.Context, obj *entities.Charge) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("object is nil")
	}

	return obj.CreatedAt.Format(time.RFC3339), nil
}

// CreationTime is the resolver for the creationTime field.
func (r *chargeResolver) CreationTime(ctx context.Context, obj *entities.Charge) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("object is nil")
	}

	return obj.CreationTime.Format(time.RFC3339), nil
}

// Status is the resolver for the status field.
func (r *chargeResolver) Status(ctx context.Context, obj *entities.Charge) (model.GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus, error) {
	if obj == nil {
		return model.GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus(""), fmt.Errorf("object is nil")
	}

	return model.GithubComKloudliteAPIAppsFinanceInternalEntitiesChargeStatus(obj.Status), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *chargeResolver) UpdatedAt(ctx context.Context, obj *entities.Charge) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("object is nil")
	}

	return obj.UpdatedAt.Format(time.RFC3339), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *chargeResolver) UpdateTime(ctx context.Context, obj *entities.Charge) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("object is nil")
	}

	return obj.UpdateTime.Format(time.RFC3339), nil
}

// WalletID is the resolver for the walletId field.
func (r *chargeResolver) WalletID(ctx context.Context, obj *entities.Charge) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("object is nil")
	}

	return string(obj.WalletId), nil
}

// WalletID is the resolver for the walletId field.
func (r *chargeInResolver) WalletID(ctx context.Context, obj *entities.Charge, data string) error {
	if obj == nil {
		return fmt.Errorf("object is nil")
	}

	obj.WalletId = repos.ID(data)
	return nil
}

// Charge returns generated.ChargeResolver implementation.
func (r *Resolver) Charge() generated.ChargeResolver { return &chargeResolver{r} }

// ChargeIn returns generated.ChargeInResolver implementation.
func (r *Resolver) ChargeIn() generated.ChargeInResolver { return &chargeInResolver{r} }

type chargeResolver struct{ *Resolver }
type chargeInResolver struct{ *Resolver }
