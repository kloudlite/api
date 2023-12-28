package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/infra/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/infra/internal/app/graph/model"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *clusterManagedServiceResolver) CreationTime(ctx context.Context, obj *entities.ClusterManagedService) (string, error) {
	if obj == nil {
		return "", errors.Newf("clusterManagedService obj is nil")
	}

	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *clusterManagedServiceResolver) ID(ctx context.Context, obj *entities.ClusterManagedService) (string, error) {
	if obj == nil {
		return "", errors.Newf("clusterManagedService obj is nil")
	}

	return string(obj.Id), nil
}

// Spec is the resolver for the spec field.
func (r *clusterManagedServiceResolver) Spec(ctx context.Context, obj *entities.ClusterManagedService) (*model.GithubComKloudliteOperatorApisCrdsV1ClusterManagedServiceSpec, error) {
	if obj == nil {
		return nil, errors.Newf("clusterManagedService is nil")
	}

	var spec model.GithubComKloudliteOperatorApisCrdsV1ClusterManagedServiceSpec

	if err := fn.JsonConversion(&obj.Spec, &spec); err != nil {
		return nil, errors.NewE(err)
	}

	return &spec, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *clusterManagedServiceResolver) UpdateTime(ctx context.Context, obj *entities.ClusterManagedService) (string, error) {
	if obj == nil {
		return "", errors.Newf("clusterManagedService is nil")
	}

	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *clusterManagedServiceInResolver) Metadata(ctx context.Context, obj *entities.ClusterManagedService, data *v1.ObjectMeta) error {
	if obj == nil {
		return errors.Newf("clusterManagedService is nil")
	}

	return fn.JsonConversion(data, &obj.ObjectMeta)
}

// Spec is the resolver for the spec field.
func (r *clusterManagedServiceInResolver) Spec(ctx context.Context, obj *entities.ClusterManagedService, data *model.GithubComKloudliteOperatorApisCrdsV1ClusterManagedServiceSpecIn) error {
	if obj == nil {
		return errors.Newf("clusterManagedService is nil")
	}

	return fn.JsonConversion(data, &obj.Spec)
}

// ClusterManagedService returns generated.ClusterManagedServiceResolver implementation.
func (r *Resolver) ClusterManagedService() generated.ClusterManagedServiceResolver {
	return &clusterManagedServiceResolver{r}
}

// ClusterManagedServiceIn returns generated.ClusterManagedServiceInResolver implementation.
func (r *Resolver) ClusterManagedServiceIn() generated.ClusterManagedServiceInResolver {
	return &clusterManagedServiceInResolver{r}
}

type clusterManagedServiceResolver struct{ *Resolver }
type clusterManagedServiceInResolver struct{ *Resolver }