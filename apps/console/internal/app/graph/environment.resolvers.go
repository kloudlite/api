package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"

	"kloudlite.io/apps/console/internal/app/graph/generated"
	"kloudlite.io/apps/console/internal/app/graph/model"
	"kloudlite.io/apps/console/internal/domain/entities"
	fn "kloudlite.io/pkg/functions"
)

// Spec is the resolver for the spec field.
func (r *environmentResolver) Spec(ctx context.Context, obj *entities.Environment) (*model.EnvironmentSpec, error) {
	if obj == nil {
		return nil, nil
	}

	var m model.EnvironmentSpec
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// Spec is the resolver for the spec field.
func (r *environmentInResolver) Spec(ctx context.Context, obj *entities.Environment, data *model.EnvironmentSpecIn) error {
	if obj == nil {
		return nil
	}
	return fn.JsonConversion(data, &obj.Spec)
}

// Environment returns generated.EnvironmentResolver implementation.
func (r *Resolver) Environment() generated.EnvironmentResolver { return &environmentResolver{r} }

// EnvironmentIn returns generated.EnvironmentInResolver implementation.
func (r *Resolver) EnvironmentIn() generated.EnvironmentInResolver { return &environmentInResolver{r} }

type environmentResolver struct{ *Resolver }
type environmentInResolver struct{ *Resolver }
