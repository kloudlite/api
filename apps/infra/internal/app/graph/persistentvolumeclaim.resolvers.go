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
)

// CreationTime is the resolver for the creationTime field.
func (r *persistentVolumeClaimResolver) CreationTime(ctx context.Context, obj *entities.PersistentVolumeClaim) (string, error) {
	if obj == nil {
		return "", errors.Newf("persistent-volume-claim/creation-time is nil")
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *persistentVolumeClaimResolver) ID(ctx context.Context, obj *entities.PersistentVolumeClaim) (string, error) {
	if obj == nil {
		return "", errors.Newf("persitent volume claim is nil")
	}
	return string(obj.Id), nil
}

// Spec is the resolver for the spec field.
func (r *persistentVolumeClaimResolver) Spec(ctx context.Context, obj *entities.PersistentVolumeClaim) (*model.K8sIoAPICoreV1PersistentVolumeClaimSpec, error) {
	var m model.K8sIoAPICoreV1PersistentVolumeClaimSpec
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// Status is the resolver for the status field.
func (r *persistentVolumeClaimResolver) Status(ctx context.Context, obj *entities.PersistentVolumeClaim) (*model.K8sIoAPICoreV1PersistentVolumeClaimStatus, error) {
	var m model.K8sIoAPICoreV1PersistentVolumeClaimStatus
	if err := fn.JsonConversion(obj.Status, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *persistentVolumeClaimResolver) UpdateTime(ctx context.Context, obj *entities.PersistentVolumeClaim) (string, error) {
	if obj == nil || obj.UpdateTime.IsZero() {
		return "", errors.Newf("persistent-volume-claim/update-time is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// PersistentVolumeClaim returns generated.PersistentVolumeClaimResolver implementation.
func (r *Resolver) PersistentVolumeClaim() generated.PersistentVolumeClaimResolver {
	return &persistentVolumeClaimResolver{r}
}

type persistentVolumeClaimResolver struct{ *Resolver }
