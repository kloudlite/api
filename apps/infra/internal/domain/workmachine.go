package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
)

func (d *domain) findWorkmachine(ctx InfraContext, name string) (*entities.Workmachine, error) {
	wm, err := d.workmachineRepo.FindOne(ctx, repos.Filter{
		fc.AccountName:  ctx.AccountName,
		fc.MetadataName: name,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	if wm == nil {
		return nil, errors.Newf("no workmachine for account=%q found", ctx.AccountName)
	}
	return wm, nil
}

func (d *domain) UpsertWorkMachine(ctx InfraContext, workmachine entities.Workmachine) (*entities.Workmachine, error) {
	workmachine.AccountName = ctx.AccountName
	workmachine.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	workmachine.LastUpdatedBy = workmachine.CreatedBy

	wm, err := d.workmachineRepo.Create(ctx, &workmachine)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return wm, nil
}

func (d *domain) UpdateWorkMachine(ctx InfraContext, workmachine entities.Workmachine) (*entities.Workmachine, error) {
	patchForUpdate := repos.Document{
		fc.DisplayName: workmachine.DisplayName,
		fc.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
	}

	upWorkmachine, err := d.workmachineRepo.Patch(
		ctx,
		repos.Filter{
			fc.AccountName:  ctx.AccountName,
			fc.MetadataName: workmachine.Name,
		},
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return upWorkmachine, nil
}

func (d *domain) UpdateWorkmachineStatus(ctx InfraContext, status bool, name string) (bool, error) {
	patchForUpdate := repos.Document{
		fc.WorkmachineSpecState: status,
		fc.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
	}

	_, err := d.workmachineRepo.Patch(
		ctx,
		repos.Filter{
			fc.AccountName:  ctx.AccountName,
			fc.MetadataName: name,
		},
		patchForUpdate,
	)
	if err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

func (d *domain) GetWorkmachine(ctx InfraContext, name string) (*entities.Workmachine, error) {
	return d.findWorkmachine(ctx, name)
}
