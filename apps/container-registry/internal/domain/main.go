package domain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kloudlite/container-registry-authorizer/admin"
	"go.uber.org/fx"
	"kloudlite.io/apps/container-registry/internal/domain/entities"
	"kloudlite.io/apps/container-registry/internal/env"
	iamT "kloudlite.io/apps/iam/types"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam"
	"kloudlite.io/pkg/docker"
	"kloudlite.io/pkg/repos"
)

type Impl struct {
	repositoryRepo repos.DbRepo[*entities.Repository]
	credentialRepo repos.DbRepo[*entities.Credential]
	dockerCli      docker.DockerCli
	iamClient      iam.IAMClient
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
func (d *Impl) CreateCredential(ctx RegistryContext, credential entities.Credential) error {

	i, err := getExpirationTime(fmt.Sprintf("%d%s", credential.Expiration.Value, credential.Expiration.Unit))
	if err != nil {
		return err
	}

	_, err = d.credentialRepo.Create(ctx, &entities.Credential{
		Name:        credential.Name,
		Token:       admin.GenerateToken(credential.UserName, ctx.AccountName, string(credential.Access), i),
		Access:      credential.Access,
		AccountName: ctx.AccountName,
	})
	if err != nil {
		return err
	}
	return nil
}

// ListCredentials implements Domain.
func (d *Impl) ListCredentials(ctx RegistryContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Credential], error) {

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.GetAccount),
	})

	if err != nil {
		return nil, err
	}

	if !co.Status {
		return nil, fmt.Errorf("unauthorized to get credentials")
	}

	filter := repos.Filter{"accountName": ctx.AccountName}
	return d.credentialRepo.FindPaginated(ctx, d.credentialRepo.MergeMatchFilters(filter, search), pagination)
}

// DeleteCredential implements Domain.
func (d *Impl) DeleteCredential(ctx RegistryContext, credName string) error {

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateAccount),
	})

	if err != nil {
		return err
	}

	if !co.Status {
		return fmt.Errorf("unauthorized to delete credentials")
	}

	return d.credentialRepo.DeleteMany(ctx, repos.Filter{"name": credName})
}

// CreateRepository implements Domain.
func (d *Impl) CreateRepository(ctx RegistryContext, repoName string) error {

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateAccount),
	})

	if err != nil {
		return err
	}

	if !co.Status {
		return fmt.Errorf("unauthorized to create repository")
	}

	_, err = d.repositoryRepo.Create(ctx, &entities.Repository{
		Name:        repoName,
		AccountName: ctx.AccountName,
	})
	return err
}

// DeleteRepository implements Domain.
func (d *Impl) DeleteRepository(ctx RegistryContext, repoName string) error {

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateAccount),
	})

	if err != nil {
		return err
	}

	if !co.Status {
		return fmt.Errorf("unauthorized to delete repository")
	}

	if _, err = d.repositoryRepo.FindOne(ctx, repos.Filter{
		"name":        repoName,
		"accountName": ctx.AccountName,
	}); err != nil {
		return err
	}

	if err := d.dockerCli.DeleteRepository(repoName); err != nil {
		return err
	}

	return d.repositoryRepo.DeleteOne(ctx, repos.Filter{"name": repoName, "accountName": ctx.AccountName})
}

// DeleteRepositoryTag implements Domain.
func (d *Impl) DeleteRepositoryTag(ctx RegistryContext, repoName string, tag Tag) error {

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateAccount),
	})

	if err != nil {
		return err
	}

	if !co.Status {
		return fmt.Errorf("unauthorized to delete repository tag")
	}

	return d.dockerCli.DeleteRepositoryTag(repoName, string(tag))
}

// ListRepositories implements Domain.
func (d *Impl) ListRepositories(ctx RegistryContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Repository], error) {

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.GetAccount),
	})

	if err != nil {
		return nil, err
	}

	if !co.Status {
		return nil, fmt.Errorf("unauthorized to list repositories")
	}

	filter := repos.Filter{"accountName": ctx.AccountName}
	return d.repositoryRepo.FindPaginated(ctx, d.credentialRepo.MergeMatchFilters(filter, search), pagination)
}

// ListRepositoryTags implements Domain.
func (d *Impl) ListRepositoryTags(ctx RegistryContext, repoName string, limit *int, after *string) ([]Tag, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.GetAccount),
	})

	if err != nil {
		return nil, err
	}

	if !co.Status {
		return nil, fmt.Errorf("unauthorized to list repository tags")
	}

	if repoName == "" {
		return nil, fmt.Errorf("invalid repository name")
	}

	if _, err = d.repositoryRepo.FindOne(ctx, repos.Filter{
		"name":        repoName,
		"accountName": ctx.AccountName,
	}); err != nil {
		return nil, err
	}

	// repoName is of the form <account-name>/<repo-name>
	s, err := d.dockerCli.ListRepositoryTags(fmt.Sprintf("%s/%s", ctx.AccountName, repoName), limit, after)
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
			iamClient iam.IAMClient,
		) (Domain, error) {
			return &Impl{
				repositoryRepo: repositoryRepo,
				credentialRepo: credentialRepo,
				dockerCli:      docker.NewDockerCli(e.RegistryUrl),
				iamClient:      iamClient,
			}, nil
		}),
)
