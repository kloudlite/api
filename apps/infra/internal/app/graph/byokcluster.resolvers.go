package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"time"

	"github.com/kloudlite/api/pkg/errors"

	"github.com/kloudlite/api/apps/infra/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/infra/internal/app/graph/model"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *bYOKClusterResolver) CreationTime(ctx context.Context, obj *entities.BYOKCluster) (string, error) {
	if obj == nil {
		return "", errors.Newf("cluster obj is nil")
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *bYOKClusterResolver) ID(ctx context.Context, obj *entities.BYOKCluster) (repos.ID, error) {
	if obj == nil {
		return "", errors.Newf("cluster obj is nil")
	}
	return obj.Id, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *bYOKClusterResolver) UpdateTime(ctx context.Context, obj *entities.BYOKCluster) (string, error) {
	if obj == nil {
		return "", errors.Newf("cluster is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Visibility is the resolver for the visibility field.
func (r *bYOKClusterResolver) Visibility(ctx context.Context, obj *entities.BYOKCluster) (*model.GithubComKloudliteAPIAppsInfraInternalEntitiesClusterVisbility, error) {
	if obj == nil {
		return nil, errors.Newf("cluster is nil")
	}
	return fn.JsonConvertP[model.GithubComKloudliteAPIAppsInfraInternalEntitiesClusterVisbility](obj.Visibility)
}

// Metadata is the resolver for the metadata field.
func (r *bYOKClusterInResolver) Metadata(ctx context.Context, obj *entities.BYOKCluster, data *v1.ObjectMeta) error {
	if obj == nil {
		return errors.Newf("cluster is nil")
	}
	return fn.JsonConversion(data, &obj.ObjectMeta)
}

// Visibility is the resolver for the visibility field.
func (r *bYOKClusterInResolver) Visibility(ctx context.Context, obj *entities.BYOKCluster, data *model.GithubComKloudliteAPIAppsInfraInternalEntitiesClusterVisbilityIn) error {
	if obj == nil {
		return errors.Newf("cluster is nil")
	}
	return fn.JsonConversion(data, &obj.Visibility)
}

// BYOKCluster returns generated.BYOKClusterResolver implementation.
func (r *Resolver) BYOKCluster() generated.BYOKClusterResolver { return &bYOKClusterResolver{r} }

// BYOKClusterIn returns generated.BYOKClusterInResolver implementation.
func (r *Resolver) BYOKClusterIn() generated.BYOKClusterInResolver { return &bYOKClusterInResolver{r} }

type (
	bYOKClusterResolver   struct{ *Resolver }
	bYOKClusterInResolver struct{ *Resolver }
)
