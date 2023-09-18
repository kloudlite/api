package domain

import (
	"context"

	"kloudlite.io/apps/container-registry/internal/domain/entities"
	"kloudlite.io/pkg/repos"
)

type Tag string

func NewRegistryContext(parent context.Context, userId repos.ID, accountName string) RegistryContext {
	return RegistryContext{
		Context:     parent,
		userId:      userId,
		accountName: accountName,
	}
}

type Domain interface {
	// registry
	ListRepositories(ctx RegistryContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Repository], error)
	CreateRepository(ctx RegistryContext, repoName string) error
	DeleteRepository(ctx RegistryContext, repoName string) error

	// tags
	ListRepositoryTags(ctx RegistryContext, repoName string, limit *int, after *string) ([]Tag, error)
	DeleteRepositoryTag(ctx RegistryContext, repoName string, tag Tag) error

	// credential
	ListCredentials(ctx RegistryContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Credential], error)
	CreateCredential(ctx RegistryContext, credName string, username string, access entities.RepoAccess, expiration string) error
	DeleteCredential(ctx RegistryContext, credName string) error
}
