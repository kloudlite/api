package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	return wm, nil
}

func (d *domain) UpsertWorkMachine(ctx InfraContext, clusterName string, workmachineName string, sshPublicKeys []string, machineType string, running bool) (*entities.Workmachine, error) {

	wm, err := d.findWorkmachine(
		ctx,
		clusterName,
		workmachineName,
	)

	if err != nil {
		return nil, errors.NewE(err)
	}

	if wm == nil {
		wm = &entities.Workmachine{
			AccountName:   ctx.AccountName,
			ClusterName:   clusterName,
			SshPublicKeys: sshPublicKeys,
			MachineType:   machineType,
			Running:       running,
			WorkMachine: crdsv1.WorkMachine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      workmachineName,
					Namespace: clusterName,
				},
			},
		}
	} else {
		wm.SshPublicKeys = sshPublicKeys
		wm.MachineType = machineType
		wm.Running = running

	}

	upWorkmachine, err := d.workmachineRepo.Upsert(
		ctx,
		repos.Filter{
			fc.AccountName:     ctx.AccountName,
			fc.MetadataName:    workmachineName,
			fields.ClusterName: clusterName,
		},
		wm,
	)

	if err != nil {
		return nil, errors.NewE(err)
	}

	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeWorkmachine, wm.Name, PublishAdd)

	if err := d.applyWorkmachine(ctx, wm); err != nil {
		return nil, errors.NewE(err)
	}

	return upWorkmachine, nil

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
