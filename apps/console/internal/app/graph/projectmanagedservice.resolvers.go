package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/app/graph/model"
	"github.com/kloudlite/api/apps/console/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *projectManagedServiceResolver) CreationTime(ctx context.Context, obj *entities.ProjectManagedService) (string, error) {
	if obj == nil {
		return "", errNilProjectManagedService
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *projectManagedServiceResolver) ID(ctx context.Context, obj *entities.ProjectManagedService) (string, error) {
	if obj == nil {
		return "", errNilProjectManagedService
	}
	return string(obj.Id), nil
}

// Spec is the resolver for the spec field.
func (r *projectManagedServiceResolver) Spec(ctx context.Context, obj *entities.ProjectManagedService) (*model.GithubComKloudliteOperatorApisCrdsV1ProjectManagedServiceSpec, error) {
	if obj == nil {
		return nil, errNilProjectManagedService
	}

	m := &model.GithubComKloudliteOperatorApisCrdsV1ProjectManagedServiceSpec{}
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return m, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *projectManagedServiceResolver) UpdateTime(ctx context.Context, obj *entities.ProjectManagedService) (string, error) {
	if obj == nil {
		return "", errNilProjectManagedService
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *projectManagedServiceInResolver) Metadata(ctx context.Context, obj *entities.ProjectManagedService, data *v1.ObjectMeta) error {
	if obj == nil {
		return errNilProjectManagedService
	}

	if data != nil {
		obj.ObjectMeta = *data
	}

	return nil
}

// Spec is the resolver for the spec field.
func (r *projectManagedServiceInResolver) Spec(ctx context.Context, obj *entities.ProjectManagedService, data *model.GithubComKloudliteOperatorApisCrdsV1ProjectManagedServiceSpecIn) error {
	if obj == nil {
		return errNilProjectManagedService
	}
	return fn.JsonConversion(data, &obj.Spec)
}

// ProjectManagedService returns generated.ProjectManagedServiceResolver implementation.
func (r *Resolver) ProjectManagedService() generated.ProjectManagedServiceResolver {
	return &projectManagedServiceResolver{r}
}

// ProjectManagedServiceIn returns generated.ProjectManagedServiceInResolver implementation.
func (r *Resolver) ProjectManagedServiceIn() generated.ProjectManagedServiceInResolver {
	return &projectManagedServiceInResolver{r}
}

type projectManagedServiceResolver struct{ *Resolver }
type projectManagedServiceInResolver struct{ *Resolver }