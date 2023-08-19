package graph

import (
	"context"
	"fmt"

	"kloudlite.io/apps/accounts/internal/domain"
	"kloudlite.io/pkg/repos"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	domain domain.Domain
}

func NewResolver(domain domain.Domain) *Resolver {
	return &Resolver{
		domain: domain,
	}
}

func toUserContext(ctx context.Context) (domain.UserContext, error) {
	if userId, ok := ctx.Value("kloudlite-user-id").(string); ok {
		return domain.UserContext{Context: ctx, UserId: repos.ID(userId)}, nil
	}

	return domain.UserContext{}, fmt.Errorf("`kloudlite-user-id` not set in request context")
}
