package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"

	"kloudlite.io/apps/accounts/internal/app/graph/generated"
	"kloudlite.io/apps/accounts/internal/entities"
)

// Role is the resolver for the role field.
func (r *accountMembershipResolver) Role(ctx context.Context, obj *entities.AccountMembership) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("membership is nil")
	}
	return string(obj.Role), nil
}

// UserID is the resolver for the userId field.
func (r *accountMembershipResolver) UserID(ctx context.Context, obj *entities.AccountMembership) (string, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// AccountMembership returns generated.AccountMembershipResolver implementation.
func (r *Resolver) AccountMembership() generated.AccountMembershipResolver {
	return &accountMembershipResolver{r}
}

type accountMembershipResolver struct{ *Resolver }
