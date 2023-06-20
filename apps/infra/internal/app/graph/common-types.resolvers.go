package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	fn "kloudlite.io/pkg/functions"
	"time"

	"github.com/kloudlite/operator/pkg/operator"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/infra/internal/app/graph/generated"
	"kloudlite.io/apps/infra/internal/app/graph/model"
	"kloudlite.io/pkg/types"
)

// Checks is the resolver for the checks field.
func (r *github_com__kloudlite__operator__pkg__operator_StatusResolver) Checks(ctx context.Context, obj *operator.Status) (map[string]interface{}, error) {
	var m map[string]any
	if err := fn.JsonConversion(obj.Checks, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// LastReconcileTime is the resolver for the lastReconcileTime field.
func (r *github_com__kloudlite__operator__pkg__operator_StatusResolver) LastReconcileTime(ctx context.Context, obj *operator.Status) (*string, error) {
	if obj == nil {
		return nil, fmt.Errorf("syncStatus is nil")
	}
	if obj.LastReconcileTime == nil {
		return nil, nil
	}
	return fn.New(obj.LastReconcileTime.Format(time.RFC3339)), nil
}

// Message is the resolver for the message field.
func (r *github_com__kloudlite__operator__pkg__operator_StatusResolver) Message(ctx context.Context, obj *operator.Status) (*model.GithubComKloudliteOperatorPkgRawJSONRawJSON, error) {
	if obj == nil {
		return nil, fmt.Errorf("syncStatus is nil")
	}
	if obj.Message == nil {
		return nil, nil
	}
	return &model.GithubComKloudliteOperatorPkgRawJSONRawJSON{
		RawMessage: obj.Message.RawMessage,
	}, nil
}

// Resources is the resolver for the resources field.
func (r *github_com__kloudlite__operator__pkg__operator_StatusResolver) Resources(ctx context.Context, obj *operator.Status) ([]*model.GithubComKloudliteOperatorPkgOperatorResourceRef, error) {
	if obj == nil {
		return nil, fmt.Errorf("syncStatus is nil")
	}
	m := make([]*model.GithubComKloudliteOperatorPkgOperatorResourceRef, len(obj.Resources))
	if err := fn.JsonConversion(obj.Resources, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// Action is the resolver for the action field.
func (r *kloudlite_io__pkg__types_SyncStatusResolver) Action(ctx context.Context, obj *types.SyncStatus) (model.KloudliteIoPkgTypesSyncStatusAction, error) {
	if obj == nil {
		return "", fmt.Errorf("syncStatus is nil")
	}
	return model.KloudliteIoPkgTypesSyncStatusAction(obj.Action), nil
}

// LastSyncedAt is the resolver for the lastSyncedAt field.
func (r *kloudlite_io__pkg__types_SyncStatusResolver) LastSyncedAt(ctx context.Context, obj *types.SyncStatus) (*string, error) {
	if obj == nil {
		return nil, fmt.Errorf("syncStatus is nil")
	}
	return fn.New(obj.LastSyncedAt.Format(time.RFC3339)), nil
}

// State is the resolver for the state field.
func (r *kloudlite_io__pkg__types_SyncStatusResolver) State(ctx context.Context, obj *types.SyncStatus) (*model.KloudliteIoPkgTypesSyncStatusState, error) {
	if obj == nil {
		return nil, fmt.Errorf("syncStatus is nil")
	}
	return fn.New(model.KloudliteIoPkgTypesSyncStatusState(obj.State)), nil
}

// SyncScheduledAt is the resolver for the syncScheduledAt field.
func (r *kloudlite_io__pkg__types_SyncStatusResolver) SyncScheduledAt(ctx context.Context, obj *types.SyncStatus) (*string, error) {
	if obj == nil {
		return nil, fmt.Errorf("syncStatus is nil")
	}
	return fn.New(obj.SyncScheduledAt.Format(time.RFC3339)), nil
}

// Annotations is the resolver for the annotations field.
func (r *metadataResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	var m map[string]any
	if err := fn.JsonConversion(obj.Annotations, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// Labels is the resolver for the labels field.
func (r *metadataResolver) Labels(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	var m map[string]any
	if err := fn.JsonConversion(obj.Labels, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// Annotations is the resolver for the annotations field.
func (r *metadataInResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	var m map[string]string
	if err := fn.JsonConversion(data, &m); err != nil {
		return err
	}
	obj.SetAnnotations(m)
	return nil
}

// Labels is the resolver for the labels field.
func (r *metadataInResolver) Labels(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	var m map[string]string
	if err := fn.JsonConversion(data, &m); err != nil {
		return err
	}
	obj.SetLabels(m)
	return nil
}

// Github_com__kloudlite__operator__pkg__operator_Status returns generated.Github_com__kloudlite__operator__pkg__operator_StatusResolver implementation.
func (r *Resolver) Github_com__kloudlite__operator__pkg__operator_Status() generated.Github_com__kloudlite__operator__pkg__operator_StatusResolver {
	return &github_com__kloudlite__operator__pkg__operator_StatusResolver{r}
}

// Kloudlite_io__pkg__types_SyncStatus returns generated.Kloudlite_io__pkg__types_SyncStatusResolver implementation.
func (r *Resolver) Kloudlite_io__pkg__types_SyncStatus() generated.Kloudlite_io__pkg__types_SyncStatusResolver {
	return &kloudlite_io__pkg__types_SyncStatusResolver{r}
}

// Metadata returns generated.MetadataResolver implementation.
func (r *Resolver) Metadata() generated.MetadataResolver { return &metadataResolver{r} }

// MetadataIn returns generated.MetadataInResolver implementation.
func (r *Resolver) MetadataIn() generated.MetadataInResolver { return &metadataInResolver{r} }

type github_com__kloudlite__operator__pkg__operator_StatusResolver struct{ *Resolver }
type kloudlite_io__pkg__types_SyncStatusResolver struct{ *Resolver }
type metadataResolver struct{ *Resolver }
type metadataInResolver struct{ *Resolver }
