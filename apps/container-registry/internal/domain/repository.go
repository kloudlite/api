package domain

import (
	"fmt"
	"log"
	"net/url"
	"regexp"

	"github.com/kloudlite/container-registry-authorizer/admin"
	"kloudlite.io/apps/container-registry/internal/domain/entities"
	iamT "kloudlite.io/apps/iam/types"
	"kloudlite.io/common"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam"
	"kloudlite.io/pkg/docker"
	"kloudlite.io/pkg/repos"
)

// CreateRepository implements Domain.
func (d *Impl) CreateRepository(ctx RegistryContext, repoName string) (*entities.Repository, error) {

	pattern := `^[a-z0-9]+([._/-][a-z0-9]+)*$`

	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
	}

	if !re.MatchString(repoName) {
		return nil, fmt.Errorf("invalid repository name, must be lowercase alphanumeric with underscore")
	}

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateAccount),
	})

	if err != nil {
		return nil, err
	}

	if !co.Status {
		return nil, fmt.Errorf("unauthorized to create repository")
	}

	return d.repositoryRepo.Create(ctx, &entities.Repository{
		Name:        repoName,
		AccountName: ctx.AccountName,
		CreatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
	})
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

	res, err := d.digestRepo.Find(ctx, repos.Query{
		Filter: repos.Filter{
			"repository":  repoName,
			"accountName": ctx.AccountName,
		},
	})

	if err != nil {
		return err
	}

	if len(res) > 0 {
		return fmt.Errorf("repository %s is not empty, please delete all Digests first", repoName)
	}

	return d.repositoryRepo.DeleteOne(ctx, repos.Filter{"name": repoName, "accountName": ctx.AccountName})
}

func (d *Impl) DeleteRepositoryDigest(ctx RegistryContext, repoName string, digest string) error {

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
		return fmt.Errorf("unauthorized to delete repository digest")
	}

	e, err := d.digestRepo.FindOne(ctx, repos.Filter{
		"digest":      digest,
		"repository":  repoName,
		"accountName": ctx.AccountName,
	})

	if err != nil {
		return err
	}

	if e == nil {
		return fmt.Errorf("%s not found in repository %s", digest, repoName)
	}

	r_url, err := url.Parse(fmt.Sprintf("https://%s", d.envs.RegistryHost))
	if err != nil {
		return err
	}

	i, err := admin.GetExpirationTime(fmt.Sprintf("%d%s", 10, "s"))
	if err != nil {
		return err
	}

	token, err := admin.GenerateToken(KL_ADMIN, e.AccountName, string("read_write"), i, d.envs.RegistrySecretKey+e.AccountName)
	if err != nil {
		return err
	}

	r_url.User = url.UserPassword(KL_ADMIN, token)

	dockerCli := docker.NewDockerClient(r_url.String())
	if err := dockerCli.DeleteDigest(fmt.Sprintf("%s/%s", ctx.AccountName, repoName), e.Digest); err != nil {
		return err
	}

	e.Deleting = true

	_, err = d.digestRepo.UpdateById(ctx, e.Id, e)

	return err
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
	return d.repositoryRepo.FindPaginated(ctx, d.repositoryRepo.MergeMatchFilters(filter, search), pagination)
}

func (d *Impl) ListRepositoryDigests(ctx RegistryContext, repoName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Digest], error) {
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
		return nil, fmt.Errorf("unauthorized to list repository digests")
	}

	filter := repos.Filter{"accountName": ctx.AccountName, "repository": repoName}
	return d.digestRepo.FindPaginated(ctx, d.digestRepo.MergeMatchFilters(filter, search), pagination)
}