package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
)

// CreatedAt is the resolver for the created_at field.
func (r *gitlabProjectResolver) CreatedAt(ctx context.Context, obj *entities.GitlabProject) (*string, error) {
	if obj == nil {
		return nil, fmt.Errorf("Object is nil")
	}
	if obj.CreatedAt == nil {
		return nil, nil
	}

	t := obj.CreatedAt.Format(time.RFC3339)
	return &t, nil
}

// LastActivityAt is the resolver for the last_activity_at field.
func (r *gitlabProjectResolver) LastActivityAt(ctx context.Context, obj *entities.GitlabProject) (*string, error) {
	if obj == nil {
		return nil, fmt.Errorf("Object is nil")
	}
	if obj.LastActivityAt == nil {
		return nil, nil
	}

	t := obj.LastActivityAt.Format(time.RFC3339)
	return &t, nil
}

// GitlabProject returns generated.GitlabProjectResolver implementation.
func (r *Resolver) GitlabProject() generated.GitlabProjectResolver { return &gitlabProjectResolver{r} }

type gitlabProjectResolver struct{ *Resolver }
