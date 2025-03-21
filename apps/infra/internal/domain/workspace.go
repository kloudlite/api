package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
)

func (d *domain) findWorkspace(ctx InfraContext, name string) (*entities.Workspace, error) {
	ws, err := d.workspaceRepo.FindOne(ctx, repos.Filter{
		fields.AccountName: ctx.AccountName,
		fc.WorkspaceName:   name,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	if ws == nil {
		return nil, errors.Newf("no workspace with name=%q found", name)
	}
	return ws, nil
}

func (d *domain) CreateWorkspace(ctx InfraContext, workspace entities.Workspace) (*entities.Workspace, error) {
	workspace.AccountName = ctx.AccountName
	workspace.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	workspace.LastUpdatedBy = workspace.CreatedBy

	ws, err := d.workspaceRepo.Create(ctx, &workspace)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return ws, nil
}

func (d *domain) UpdateWorkspace(ctx InfraContext, workspace entities.Workspace) (*entities.Workspace, error) {
	patchForUpdate := repos.Document{
		fields.DisplayName: workspace.Name,
		fields.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
	}

	upWorkspace, err := d.workspaceRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName: ctx.AccountName,
			fc.WorkspaceName:   workspace.Name,
		},
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return upWorkspace, nil
}

func (d *domain) DeleteWorkspace(ctx InfraContext, name string) error {
	err := d.workspaceRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.AccountName: ctx.AccountName,
			fc.WorkspaceName:   name,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}
	return nil
}

func (d *domain) GetWorkspace(ctx InfraContext, name string) (*entities.Workspace, error) {
	return d.findWorkspace(ctx, name)
}

func (d *domain) ListWorkspaces(ctx InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Workspace], error) {
	filter := repos.Filter{
		fields.AccountName: ctx.AccountName,
	}
	return d.workspaceRepo.FindPaginated(ctx, d.workspaceRepo.MergeMatchFilters(filter, search), pagination)
}
