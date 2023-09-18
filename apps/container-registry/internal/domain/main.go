package domain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kloudlite/container-registry-authorizer/admin"
	"go.uber.org/fx"
	"kloudlite.io/apps/container-registry/internal/domain/entities"
	"kloudlite.io/apps/container-registry/internal/env"
	"kloudlite.io/pkg/docker"
	"kloudlite.io/pkg/repos"
)

type Impl struct {
	repositoryRepo repos.DbRepo[*entities.Repository]
	credentialRepo repos.DbRepo[*entities.Credential]
	dockerCli      docker.DockerCli
}

func getExpirationTime(expiration string) (time.Time, error) {
	now := time.Now()

	// Split the string into the numeric and duration type parts
	length := len(expiration)
	if length < 2 {
		return now, fmt.Errorf("invalid expiration format")
	}

	durationValStr := expiration[:length-1]
	durationVal, err := strconv.Atoi(durationValStr)
	if err != nil {
		return now, fmt.Errorf("invalid duration value: %v", err)
	}

	durationType := expiration[length-1]

	switch durationType {
	case 'h':
		return now.Add(time.Duration(durationVal) * time.Hour), nil
	case 'd':
		return now.AddDate(0, 0, durationVal), nil
	case 'w':
		return now.AddDate(0, 0, durationVal*7), nil
	case 'm':
		return now.AddDate(0, durationVal, 0), nil
	case 'y':
		return now.AddDate(durationVal, 0, 0), nil
	default:
		return now, fmt.Errorf("invalid duration type: %v", durationType)
	}
}

// CreateCredential implements Domain.
func (d *Impl) CreateCredential(ctx RegistryContext, credName string, username string, access entities.RepoAccess, expiration string) error {

	i, err := getExpirationTime(expiration)
	if err != nil {
		return err
	}

	_, err = d.credentialRepo.Create(ctx, &entities.Credential{
		Name:        credName,
		Token:       admin.GenerateToken(username, ctx.accountName, string(access), i),
		Access:      access,
		AccountName: ctx.accountName,
	})
	if err != nil {
		return err
	}
	return nil
}

// ListCredentials implements Domain.
func (d *Impl) ListCredentials(ctx RegistryContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Credential], error) {

	filter := repos.Filter{"accountName": ctx.accountName}
	return d.credentialRepo.FindPaginated(ctx, d.credentialRepo.MergeMatchFilters(filter, search), pagination)
}

// DeleteCredential implements Domain.
func (d *Impl) DeleteCredential(ctx RegistryContext, credName string) error {
	return d.credentialRepo.DeleteMany(ctx, repos.Filter{"name": credName})
}

// CreateRepository implements Domain.
func (d *Impl) CreateRepository(ctx RegistryContext, repoName string) error {
	_, err := d.repositoryRepo.Create(ctx, &entities.Repository{
		Name:        repoName,
		AccountName: ctx.accountName,
	})
	return err
}

// DeleteRepository implements Domain.
func (d *Impl) DeleteRepository(ctx RegistryContext, repoName string) error {
	return d.repositoryRepo.DeleteMany(ctx, repos.Filter{"name": repoName})
}

// DeleteRepositoryTag implements Domain.
func (d *Impl) DeleteRepositoryTag(ctx RegistryContext, repoName string, tag Tag) error {
	return d.dockerCli.DeleteRepositoryTag(repoName, string(tag))
}

// ListRepositories implements Domain.
func (d *Impl) ListRepositories(ctx RegistryContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Repository], error) {

	filter := repos.Filter{"accountName": ctx.accountName}
	return d.repositoryRepo.FindPaginated(ctx, d.credentialRepo.MergeMatchFilters(filter, search), pagination)
}

// ListRepositoryTags implements Domain.
func (d *Impl) ListRepositoryTags(ctx RegistryContext, repoName string, limit *int, after *string) ([]Tag, error) {
	s, err := d.dockerCli.ListRepositoryTags(repoName, limit, after)
	if err != nil {
		return nil, err
	}

	res := make([]Tag, len(s))
	for i, v := range s {
		res[i] = Tag(v)
	}

	return res, nil
}

var Module = fx.Module(
	"domain",
	fx.Provide(
		func(e *env.Env,
			repositoryRepo repos.DbRepo[*entities.Repository],
			credentialRepo repos.DbRepo[*entities.Credential],
		) (Domain, error) {
			return &Impl{
				repositoryRepo: repositoryRepo,
				credentialRepo: credentialRepo,
				dockerCli:      docker.NewDockerCli(e.RegistryUrl),
			}, nil
		}),
)
