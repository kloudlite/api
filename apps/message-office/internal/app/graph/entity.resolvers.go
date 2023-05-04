package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"

	"kloudlite.io/apps/message-office/internal/app/graph/generated"
	"kloudlite.io/apps/message-office/internal/app/graph/model"
)

// FindBYOCClusterByMetadataNameAndSpecAccountName is the resolver for the findBYOCClusterByMetadataNameAndSpecAccountName field.
func (r *entityResolver) FindBYOCClusterByMetadataNameAndSpecAccountName(ctx context.Context, metadataName string, specAccountName string) (*model.BYOCCluster, error) {
	return &model.BYOCCluster{
		Metadata: &model.Metadata{
			Name: metadataName,
		},
		Spec: &model.BYOCClusterSpec{
			AccountName: specAccountName,
		},
		ClusterToken: "",
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }