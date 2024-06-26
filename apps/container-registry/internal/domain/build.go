package domain

import (
	"context"
	"slices"

	"github.com/google/uuid"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
	fc "github.com/kloudlite/api/apps/container-registry/internal/domain/entities/field-constants"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	dbv1 "github.com/kloudlite/operator/apis/distribution/v1"
)

func (d *Impl) AddBuild(ctx RegistryContext, build entities.Build) (*entities.Build, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.CreateBuildIntegration),
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to add build")
	}

	if err := validateBuild(build); err != nil {
		return nil, errors.NewE(err)
	}

	var webhookId *int

	if build.Source.Provider == "gitlab" {
		webhookId, err = d.GitlabAddWebhook(ctx, ctx.UserId, d.gitlab.GetRepoId(build.Source.Repository))
		if err != nil {
			return nil, errors.NewE(err)
		}
	}

	return d.buildRepo.Create(ctx, &entities.Build{
		Spec: func() dbv1.BuildRunSpec {
			build.Spec.AccountName = ctx.AccountName
			return build.Spec
		}(),
		Name:             build.Name,
		BuildClusterName: build.BuildClusterName,
		CreatedBy:        common.CreatedOrUpdatedBy{UserId: ctx.UserId, UserName: ctx.UserName, UserEmail: ctx.UserEmail},
		LastUpdatedBy:    common.CreatedOrUpdatedBy{},
		Source:           entities.GitSource{Repository: build.Source.Repository, Branch: build.Source.Branch, Provider: build.Source.Provider, WebhookId: webhookId},
		CredUser:         common.CreatedOrUpdatedBy{UserId: ctx.UserId, UserName: ctx.UserName, UserEmail: ctx.UserEmail},
		ErrorMessages:    map[string]string{},
		Status:           entities.BuildStatusIdle,
	})
}

func (d *Impl) UpdateBuild(ctx RegistryContext, id repos.ID, build entities.Build) (*entities.Build, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateBuildIntegration),
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to update build")
	}

	if err := validateBuild(build); err != nil {
		return nil, errors.NewE(err)
	}

	if build.Spec.AccountName == "" {
		build.Spec.AccountName = ctx.AccountName
	}

	patchDoc := repos.Document{
		fc.BuildName:             build.Name,
		fc.BuildBuildClusterName: build.BuildClusterName,
		fields.LastUpdatedBy:     common.CreatedOrUpdatedBy{UserId: ctx.UserId, UserName: ctx.UserName, UserEmail: ctx.UserEmail},
		fc.BuildSource:           build.Source,
		fc.BuildSpec:             build.Spec,
	}

	return d.buildRepo.Patch(ctx, repos.Filter{fields.Id: id}, patchDoc)
}

func (d *Impl) UpdateBuildInternal(ctx context.Context, build *entities.Build) (*entities.Build, error) {
	return d.buildRepo.UpdateById(ctx, build.Id, build)
}

func (d *Impl) ListBuildsByGit(ctx context.Context, repoUrl, branch, provider string) ([]*entities.Build, error) {
	filter := repos.Filter{
		fc.BuildSourceRepository: repoUrl,
		fc.BuildSourceBranch:     branch,
		fc.BuildSourceProvider:   provider,
	}

	b, err := d.buildRepo.Find(ctx, repos.Query{
		Filter: filter,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	return b, nil
}

func (d *Impl) ListBuilds(ctx RegistryContext, repoName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Build], error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.ListBuildIntegrations),
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to list builds")
	}

	filter := repos.Filter{
		fc.BuildSpecAccountName:      ctx.AccountName,
		fc.BuildSpecRegistryRepoName: repoName,
	}

	return d.buildRepo.FindPaginated(ctx, d.buildRepo.MergeMatchFilters(filter, search), pagination)
}

func (d *Impl) GetBuild(ctx RegistryContext, buildId repos.ID) (*entities.Build, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.GetBuildIntegration),
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to get build")
	}

	b, err := d.buildRepo.FindOne(ctx, repos.Filter{
		"spec.accountName": ctx.AccountName,
		"id":               buildId,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if b == nil {
		return nil, errors.Newf("build (id=%s) not found", buildId)
	}

	return b, nil
}

func (d *Impl) DeleteBuild(ctx RegistryContext, buildId repos.ID) error {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.DeleteBuildIntegration),
	})
	if err != nil {
		return errors.NewE(err)
	}

	if !co.Status {
		return errors.Newf("unauthorized to delete build")
	}

	b, err := d.buildRepo.FindById(ctx, buildId)
	if err != nil {
		return errors.NewE(err)
	}

	if err = d.buildRepo.DeleteOne(ctx, repos.Filter{
		fc.BuildSpecAccountName: ctx.AccountName,
		fields.Id:               buildId,
	}); err != nil {
		return errors.NewE(err)
	}

	if b.Source.Provider == "gitlab" {
		at, err := d.getAccessTokenByUserId(ctx, "gitlab", ctx.UserId)
		if err != nil {
			return nil
		}

		if err := d.gitlab.DeleteWebhook(ctx, at, string(b.Source.Repository), entities.GitlabWebhookId(*b.Source.WebhookId)); err != nil {
			d.logger.Errorf(err, "error while deleting webhook")
		}
	}

	return nil
}

func (d *Impl) TriggerBuild(ctx RegistryContext, buildId repos.ID) error {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.CreateBuildRun),
	})
	if err != nil {
		return errors.NewE(err)
	}

	if !co.Status {
		return errors.Newf("unauthorized to trigger build")
	}

	b, err := d.buildRepo.FindById(ctx, buildId)
	if err != nil {
		return errors.NewE(err)
	}
	if b == nil {
		return errors.Newf("build not found")
	}

	var pullToken string
	var commitHash string

	if !slices.Contains([]string{"github", "gitlab"}, string(b.Source.Provider)) {
		return errors.Newf("provider %s not supported", b.Source.Provider)
	}

	// at, err := d.getAccessTokenByUserId(ctx, string(b.Source.Provider), ctx.UserId)
	at, err := d.getAccessTokenByUserId(ctx, string(b.Source.Provider), b.CreatedBy.UserId)
	if err != nil {
		return errors.NewE(err)
	}

	switch b.Source.Provider {
	case "gitlab":
		pullToken, err = d.GitlabPullToken(ctx, ctx.UserId)
		if err != nil {
			return errors.NewE(err)
		}

		commitHash, err = d.gitlab.GetLatestCommit(ctx, at, b.Source.Repository, b.Source.Branch)
		if err != nil {
			return errors.NewE(err)
		}

	case "github":

		pullToken, err = d.GithubInstallationToken(ctx, b.Source.Repository)
		if err != nil {
			return errors.NewE(err)
		}

		commitHash, err = d.github.GetLatestCommit(ctx, at, b.Source.Repository, b.Source.Branch)
		if err != nil {
			return errors.NewE(err)
		}
	default:
		return errors.Newf("provider %s not supported", b.Source.Provider)
	}

	b.Name = string(buildId)

	uid := uuid.NewString()

	if err := d.CreateBuildRun(ctx, b, &GitWebhookPayload{
		GitProvider: string(b.Source.Provider),
		RepoUrl:     b.Source.Repository,
		GitBranch:   b.Source.Branch,
		CommitHash:  commitHash,
	}, pullToken, uid); err != nil {
		return errors.NewE(err)
	}

	return nil
}
