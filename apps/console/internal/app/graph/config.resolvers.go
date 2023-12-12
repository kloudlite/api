package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *configResolver) CreationTime(ctx context.Context, obj *entities.Config) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return obj.BaseEntity.CreationTime.Format(time.RFC3339), nil
}

// Data is the resolver for the data field.
func (r *configResolver) Data(ctx context.Context, obj *entities.Config) (map[string]interface{}, error) {
	m := make(map[string]any, len(obj.Data))
	if err := fn.JsonConversion(obj.Data, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// ID is the resolver for the id field.
func (r *configResolver) ID(ctx context.Context, obj *entities.Config) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return string(obj.Id), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *configResolver) UpdateTime(ctx context.Context, obj *entities.Config) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return obj.BaseEntity.UpdateTime.Format(time.RFC3339), nil
}

// Data is the resolver for the data field.
func (r *configInResolver) Data(ctx context.Context, obj *entities.Config, data map[string]interface{}) error {
	return fn.JsonConversion(data, &obj.Data)
}

// Metadata is the resolver for the metadata field.
func (r *configInResolver) Metadata(ctx context.Context, obj *entities.Config, data *v1.ObjectMeta) error {
	obj.ObjectMeta = *data
	return nil
}

// Config returns generated.ConfigResolver implementation.
func (r *Resolver) Config() generated.ConfigResolver { return &configResolver{r} }

// ConfigIn returns generated.ConfigInResolver implementation.
func (r *Resolver) ConfigIn() generated.ConfigInResolver { return &configInResolver{r} }

type configResolver struct{ *Resolver }
type configInResolver struct{ *Resolver }
