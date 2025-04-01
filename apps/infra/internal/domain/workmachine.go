package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func (d *domain) applyWorkmachine(ctx InfraContext, wm *entities.Workmachine) error {
	addTrackingId(&wm.WorkMachine, wm.Id)
	return d.resDispatcher.ApplyToTargetCluster(ctx, wm.DispatchAddr, &wm.WorkMachine, wm.RecordVersion)
}

func (d *domain) findWorkmachine(ctx InfraContext, clusterName string, name string) (*entities.Workmachine, error) {
	wm, err := d.workmachineRepo.FindOne(ctx, repos.Filter{
		fc.AccountName:  ctx.AccountName,
		fc.MetadataName: name,
		fc.ClusterName:  clusterName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	// if wm == nil {
	// 	return nil, errors.Newf("no workmachine for account=%q found", ctx.AccountName)
	// }
	return wm, nil
}

func (d *domain) CreateWorkMachine(ctx InfraContext, clusterName string, workmachine entities.Workmachine) (*entities.Workmachine, error) {
	workmachine.AccountName = ctx.AccountName
	workmachine.ClusterName = clusterName

	workmachine.DispatchAddr = &entities.DispatchAddr{
		AccountName: ctx.AccountName,
		ClusterName: clusterName}

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

	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeWorkmachine, wm.Name, PublishAdd)

	if err := d.applyWorkmachine(ctx, wm); err != nil {
		return nil, errors.NewE(err)
	}

	return wm, nil
}

func (d *domain) UpdateWorkMachine(ctx InfraContext, clusterName string, workmachine entities.Workmachine) (*entities.Workmachine, error) {
	patchForUpdate := repos.Document{
		fc.DisplayName:     workmachine.DisplayName,
		fc.WorkmachineSpec: workmachine.Spec,
		fc.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
	}

	upWorkmachine, err := d.workmachineRepo.Patch(
		ctx,
		repos.Filter{
			fc.AccountName:     ctx.AccountName,
			fc.MetadataName:    workmachine.Name,
			fields.ClusterName: clusterName,
		},
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}

	d.resourceEventPublisher.PublishResourceEvent(ctx, workmachine.ClusterName, ResourceTypeWorkmachine, upWorkmachine.Name, PublishUpdate)

	if err := d.applyWorkmachine(ctx, upWorkmachine); err != nil {
		return nil, errors.NewE(err)
	}

	return upWorkmachine, nil
}

func (d *domain) UpdateWorkmachineStatus(ctx InfraContext, clusterName string, status bool, name string) (bool, error) {
	machineStatus := "OFF"
	if status {
		machineStatus = "ON"
	}

	patchForUpdate := repos.Document{
		fc.WorkmachineSpecState: machineStatus,
		// fc.WorkmachineMachineStatus: status,
		fc.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
	}

	upWorkmachine, err := d.workmachineRepo.Patch(
		ctx,
		repos.Filter{
			fc.AccountName:     ctx.AccountName,
			fc.MetadataName:    name,
			fields.ClusterName: clusterName,
		},
		patchForUpdate,
	)
	if err != nil {
		return false, errors.NewE(err)
	}

	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeWorkmachine, upWorkmachine.Name, PublishUpdate)

	if err := d.applyWorkmachine(ctx, upWorkmachine); err != nil {
		return false, errors.NewE(err)
	}

	return true, nil
}

func (d *domain) GetWorkmachine(ctx InfraContext, clusterName string, name string) (*entities.Workmachine, error) {
	return d.findWorkmachine(ctx, clusterName, name)
}

func (d *domain) OnWorkmachineDeleteMessage(ctx InfraContext, clusterName string, workmachine entities.Workmachine) error {
	err := d.workmachineRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.ClusterName:  clusterName,
			fields.MetadataName: workmachine.Name,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeWorkmachine, workmachine.Name, PublishDelete)
	return nil
}

func (d *domain) OnWorkmachineUpdateMessage(ctx InfraContext, clusterName string, workmachine entities.Workmachine, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	wm, err := d.findWorkmachine(ctx, clusterName, workmachine.Name)
	if err != nil {
		return errors.NewE(err)
	}

	if wm == nil {
		workmachine.AccountName = ctx.AccountName
		workmachine.ClusterName = clusterName

		workmachine.CreatedBy = common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		}

		workmachine.LastUpdatedBy = workmachine.CreatedBy

		wm, err = d.workmachineRepo.Create(ctx, &workmachine)
		if err != nil {
			return errors.NewE(err)
		}
	}

	upWm, err := d.workmachineRepo.PatchById(
		ctx,
		wm.Id,
		common.PatchForSyncFromAgent(
			&workmachine,
			workmachine.RecordVersion,
			status,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
			}))
	if err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeWorkmachine, upWm.Name, PublishUpdate)
	return nil
}
