package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/accounts/internal/app/graph/generated"
	"kloudlite.io/apps/accounts/internal/entities"
	fn "kloudlite.io/pkg/functions"
)

// CreationTime is the resolver for the creationTime field.
func (r *accountResolver) CreationTime(ctx context.Context, obj *entities.Account) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("account is nil")
	}
	return obj.BaseEntity.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *accountResolver) ID(ctx context.Context, obj *entities.Account) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return string(obj.Id), nil
}

// Spec is the resolver for the spec field.
func (r *accountResolver) Spec(ctx context.Context, obj *entities.Account) (map[string]interface{}, error) {
	if obj == nil {
		return nil, fmt.Errorf("resource is nil")
	}
	m := map[string]any{}
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *accountResolver) UpdateTime(ctx context.Context, obj *entities.Account) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *accountInResolver) Metadata(ctx context.Context, obj *entities.Account, data *v1.ObjectMeta) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}
	return fn.JsonConversion(data, &obj.ObjectMeta)
}

// Spec is the resolver for the spec field.
func (r *accountInResolver) Spec(ctx context.Context, obj *entities.Account, data map[string]interface{}) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}
	return fn.JsonConversion(data, &obj.Spec)
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// AccountIn returns generated.AccountInResolver implementation.
func (r *Resolver) AccountIn() generated.AccountInResolver { return &accountInResolver{r} }

type accountResolver struct{ *Resolver }
type accountInResolver struct{ *Resolver }

