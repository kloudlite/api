package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/generated"
	"github.com/kloudlite/api/common"
	fn "github.com/kloudlite/api/pkg/functions"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UserID is the resolver for the userId field.
func (r *github__com___kloudlite___api___common__CreatedOrUpdatedByResolver) UserID(ctx context.Context, obj *common.CreatedOrUpdatedBy) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("obj is nil")
	}

	return string(obj.UserId), nil
}

// Annotations is the resolver for the annotations field.
func (r *metadataResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	if obj == nil {
		return nil, fmt.Errorf("obj is nil")
	}
	var m map[string]interface{}
	if err := fn.JsonConversion(obj.Annotations, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return m, nil
}

// CreationTimestamp is the resolver for the creationTimestamp field.
func (r *metadataResolver) CreationTimestamp(ctx context.Context, obj *v1.ObjectMeta) (string, error) {
	if obj == nil {
		return "", errors.Newf("build-run/creation-time is nil")
	}
	return obj.CreationTimestamp.Format(time.RFC3339), nil
}

// DeletionTimestamp is the resolver for the deletionTimestamp field.
func (r *metadataResolver) DeletionTimestamp(ctx context.Context, obj *v1.ObjectMeta) (*string, error) {
	if obj == nil {
		return nil, errors.Newf("build-run/deletion-time is nil")
	}
	if obj.DeletionTimestamp == nil {
		return nil, nil
	}
	format := obj.DeletionTimestamp.Time.Format(time.RFC3339)
	return &format, nil
}

// Labels is the resolver for the labels field.
func (r *metadataResolver) Labels(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	if obj == nil {
		return nil, fmt.Errorf("obj is nil")
	}
	var m map[string]interface{}
	if err := fn.JsonConversion(obj.Labels, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return m, nil
}

// Annotations is the resolver for the annotations field.
func (r *metadataInResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}
	var m map[string]string
	if err := fn.JsonConversion(data, &m); err != nil {
		return errors.NewE(err)
	}
	obj.Annotations = m
	return nil
}

// Labels is the resolver for the labels field.
func (r *metadataInResolver) Labels(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	if obj == nil {
		return fmt.Errorf("obj is nil")
	}
	var m map[string]string
	if err := fn.JsonConversion(data, &m); err != nil {
		return errors.NewE(err)
	}
	obj.Labels = m
	return nil
}

// Github__com___kloudlite___api___common__CreatedOrUpdatedBy returns generated.Github__com___kloudlite___api___common__CreatedOrUpdatedByResolver implementation.
func (r *Resolver) Github__com___kloudlite___api___common__CreatedOrUpdatedBy() generated.Github__com___kloudlite___api___common__CreatedOrUpdatedByResolver {
	return &github__com___kloudlite___api___common__CreatedOrUpdatedByResolver{r}
}

// Metadata returns generated.MetadataResolver implementation.
func (r *Resolver) Metadata() generated.MetadataResolver { return &metadataResolver{r} }

// MetadataIn returns generated.MetadataInResolver implementation.
func (r *Resolver) MetadataIn() generated.MetadataInResolver { return &metadataInResolver{r} }

type github__com___kloudlite___api___common__CreatedOrUpdatedByResolver struct{ *Resolver }
type metadataResolver struct{ *Resolver }
type metadataInResolver struct{ *Resolver }
