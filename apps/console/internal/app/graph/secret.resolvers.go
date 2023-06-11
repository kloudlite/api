package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/console/internal/app/graph/generated"
	"kloudlite.io/apps/console/internal/domain/entities"
	fn "kloudlite.io/pkg/functions"
)

// Data is the resolver for the data field.
func (r *secretResolver) Data(ctx context.Context, obj *entities.Secret) (map[string]interface{}, error) {
	var m map[string]any
	if err := fn.JsonConversion(obj.Data, &m); err != nil {
		return m, err
	}
	return m, nil
}

// StringData is the resolver for the stringData field.
func (r *secretResolver) StringData(ctx context.Context, obj *entities.Secret) (map[string]interface{}, error) {
	var m map[string]any
	if err := fn.JsonConversion(obj.StringData, &m); err != nil {
		return m, err
	}
	return m, nil
}

// Type is the resolver for the type field.
func (r *secretResolver) Type(ctx context.Context, obj *entities.Secret) (*string, error) {
	return (*string)(&obj.Type), nil
}

// Data is the resolver for the data field.
func (r *secretInResolver) Data(ctx context.Context, obj *entities.Secret, data map[string]interface{}) error {
	if data != nil {
		return fn.JsonConversion(obj, data)
	}
	return fmt.Errorf("data is nil")
}

// Metadata is the resolver for the metadata field.
func (r *secretInResolver) Metadata(ctx context.Context, obj *entities.Secret, data *v1.ObjectMeta) error {
	if data != nil {
		obj.ObjectMeta = *data
		return nil
	}
	return fmt.Errorf("data is nil")
}

// StringData is the resolver for the stringData field.
func (r *secretInResolver) StringData(ctx context.Context, obj *entities.Secret, data map[string]interface{}) error {
	return fn.JsonConversion(data, &obj.StringData)
}

// Type is the resolver for the type field.
func (r *secretInResolver) Type(ctx context.Context, obj *entities.Secret, data *string) error {
	if data != nil {
		obj.Type = corev1.SecretType(*data)
		return nil
	}
	return fmt.Errorf("secret type is nil")
}

// Secret returns generated.SecretResolver implementation.
func (r *Resolver) Secret() generated.SecretResolver { return &secretResolver{r} }

// SecretIn returns generated.SecretInResolver implementation.
func (r *Resolver) SecretIn() generated.SecretInResolver { return &secretInResolver{r} }

type secretResolver struct{ *Resolver }
type secretInResolver struct{ *Resolver }
