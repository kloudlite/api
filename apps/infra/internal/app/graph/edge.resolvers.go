package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"kloudlite.io/apps/infra/internal/app/graph/generated"
	"kloudlite.io/apps/infra/internal/app/graph/model"
	"kloudlite.io/apps/infra/internal/domain/entities"
	fn "kloudlite.io/pkg/functions"
)

func (r *edgeResolver) Spec(ctx context.Context, obj *entities.Edge) (*model.EdgeSpec, error) {
	var m model.EdgeSpec
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *edgeInResolver) Spec(ctx context.Context, obj *entities.Edge, data *model.EdgeSpecIn) error {
	if obj == nil {
		return nil
	}
	return fn.JsonConversion(data, &obj.Spec)
}

// Edge returns generated.EdgeResolver implementation.
func (r *Resolver) Edge() generated.EdgeResolver { return &edgeResolver{r} }

// EdgeIn returns generated.EdgeInResolver implementation.
func (r *Resolver) EdgeIn() generated.EdgeInResolver { return &edgeInResolver{r} }

type edgeResolver struct{ *Resolver }
type edgeInResolver struct{ *Resolver }
