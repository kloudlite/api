package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"time"

	v11 "github.com/kloudlite/operator/apis/crds/v1"
	"github.com/kloudlite/operator/pkg/operator"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/console/internal/app/graph/generated"
	"kloudlite.io/apps/console/internal/app/graph/model"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/types"
)

func (r *metadataResolver) Labels(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	m := make(map[string]any, len(obj.Labels))
	if obj.Labels == nil {
		return nil, nil
	}
	if err := fn.JsonConversion(obj.Labels, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (r *metadataResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	m := make(map[string]any, len(obj.Labels))
	if obj.Labels == nil {
		return nil, nil
	}
	if err := fn.JsonConversion(obj.Labels, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (r *metadataResolver) CreationTimestamp(ctx context.Context, obj *v1.ObjectMeta) (string, error) {
	return obj.GetCreationTimestamp().Format(time.RFC3339), nil
}

func (r *metadataResolver) DeletionTimestamp(ctx context.Context, obj *v1.ObjectMeta) (*string, error) {
	d := obj.GetDeletionTimestamp()
	if d == nil {
		return nil, nil
	}
	return fn.New(d.Format(time.RFC3339)), nil
}

func (r *overridesResolver) Patches(ctx context.Context, obj *v11.JsonPatch) ([]*model.Patch, error) {
	m := make([]*model.Patch, len(obj.Patches))
	if err := fn.JsonConversion(obj.Patches, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (r *statusResolver) Checks(ctx context.Context, obj *operator.Status) (map[string]interface{}, error) {
	m := make(map[string]any, len(obj.Checks))
	if err := fn.JsonConversion(obj.Checks, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (r *statusResolver) DisplayVars(ctx context.Context, obj *operator.Status) (map[string]interface{}, error) {
	var m map[string]any
	b, err := obj.DisplayVars.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (r *syncStatusResolver) SyncScheduledAt(ctx context.Context, obj *types.SyncStatus) (string, error) {
	return obj.SyncScheduledAt.Format(time.RFC3339), nil
}

func (r *syncStatusResolver) LastSyncedAt(ctx context.Context, obj *types.SyncStatus) (*string, error) {
	return fn.New(obj.LastSyncedAt.Format(time.RFC3339)), nil
}

func (r *metadataInResolver) Labels(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	if obj.Labels == nil {
		obj.Labels = make(map[string]string, len(data))
	}
	return fn.JsonConversion(data, &obj.Labels)
}

func (r *metadataInResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	if obj.Labels == nil {
		obj.Labels = make(map[string]string, len(data))
	}
	return fn.JsonConversion(data, &obj.Labels)
}

func (r *overridesInResolver) Patches(ctx context.Context, obj *v11.JsonPatch, data []*model.PatchIn) error {
	return fn.JsonConversion(data, &obj.Patches)
}

// Metadata returns generated.MetadataResolver implementation.
func (r *Resolver) Metadata() generated.MetadataResolver { return &metadataResolver{r} }

// Overrides returns generated.OverridesResolver implementation.
func (r *Resolver) Overrides() generated.OverridesResolver { return &overridesResolver{r} }

// Status returns generated.StatusResolver implementation.
func (r *Resolver) Status() generated.StatusResolver { return &statusResolver{r} }

// SyncStatus returns generated.SyncStatusResolver implementation.
func (r *Resolver) SyncStatus() generated.SyncStatusResolver { return &syncStatusResolver{r} }

// MetadataIn returns generated.MetadataInResolver implementation.
func (r *Resolver) MetadataIn() generated.MetadataInResolver { return &metadataInResolver{r} }

// OverridesIn returns generated.OverridesInResolver implementation.
func (r *Resolver) OverridesIn() generated.OverridesInResolver { return &overridesInResolver{r} }

type metadataResolver struct{ *Resolver }
type overridesResolver struct{ *Resolver }
type statusResolver struct{ *Resolver }
type syncStatusResolver struct{ *Resolver }
type metadataInResolver struct{ *Resolver }
type overridesInResolver struct{ *Resolver }