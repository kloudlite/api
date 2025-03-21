package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
)

func (d *domain) CreateWorkMachine(ctx InfraContext, workmachine entities.Workmachine) (*entities.Workmachine, error) {
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
		fields.DisplayName: workmachine.DisplayName,
		fields.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
	}

	upWorkmachine, err := d.workmachineRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName: ctx.AccountName,
		},
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return upWorkmachine, nil
}

func (d *domain) UpdateWorkmachineStatus(ctx InfraContext, status bool) (bool, error) {
	patchForUpdate := repos.Document{
		fc.WorkmachineMachineStatus: status,
		fields.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
	}

	_, err := d.workmachineRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName: ctx.AccountName,
		},
		patchForUpdate,
	)
	if err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}
