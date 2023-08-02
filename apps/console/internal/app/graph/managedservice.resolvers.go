package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/console/internal/app/graph/generated"
	"kloudlite.io/apps/console/internal/app/graph/model"
	"kloudlite.io/apps/console/internal/domain/entities"
	fn "kloudlite.io/pkg/functions"
)

// CreationTime is the resolver for the creationTime field.
func (r *managedServiceResolver) CreationTime(ctx context.Context, obj *entities.ManagedService) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return obj.BaseEntity.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *managedServiceResolver) ID(ctx context.Context, obj *entities.ManagedService) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return string(obj.Id), nil
}

// Spec is the resolver for the spec field.
func (r *managedServiceResolver) Spec(ctx context.Context, obj *entities.ManagedService) (*model.GithubComKloudliteOperatorApisCrdsV1ManagedServiceSpec, error) {
	m := &model.GithubComKloudliteOperatorApisCrdsV1ManagedServiceSpec{}
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *managedServiceResolver) UpdateTime(ctx context.Context, obj *entities.ManagedService) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return obj.BaseEntity.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *managedServiceInResolver) Metadata(ctx context.Context, obj *entities.ManagedService, data *v1.ObjectMeta) error {
	obj.ObjectMeta = *data
	return nil
}

// Spec is the resolver for the spec field.
func (r *managedServiceInResolver) Spec(ctx context.Context, obj *entities.ManagedService, data *model.GithubComKloudliteOperatorApisCrdsV1ManagedServiceSpecIn) error {
	if obj == nil {
		return fmt.Errorf("resource is nil")
	}
	return fn.JsonConversion(data, obj.Spec)
}

// ManagedService returns generated.ManagedServiceResolver implementation.
func (r *Resolver) ManagedService() generated.ManagedServiceResolver {
	return &managedServiceResolver{r}
}

// ManagedServiceIn returns generated.ManagedServiceInResolver implementation.
func (r *Resolver) ManagedServiceIn() generated.ManagedServiceInResolver {
	return &managedServiceInResolver{r}
}

type managedServiceResolver struct{ *Resolver }
type managedServiceInResolver struct{ *Resolver }