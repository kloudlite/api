package persistent_volume

import (
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

// GetPV implements Domain.
func (d *Domain) GetPV(ctx domainT.InfraContext, clusterName string, pvName string) (*entities.PersistentVolume, error) {
	pv, err := d.PersistentVolumeRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: pvName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if pv == nil {
		return nil, errors.Newf("persistent volume with name %q not found", pvName)
	}
	return pv, nil
}

// ListPVs implements Domain.
func (d *Domain) ListPVs(ctx domainT.InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.PersistentVolume], error) {
	filter := repos.Filter{
		fields.AccountName: ctx.AccountName,
		fields.ClusterName: clusterName,
	}
	return d.PersistentVolumeRepo.FindPaginated(ctx, d.PersistentVolumeRepo.MergeMatchFilters(filter, search), pagination)
}

func (d *Domain) DeletePV(ctx domainT.InfraContext, clusterName string, pvName string) error {
	// FIXME: (IAM role binding for DeletePV)
	if err := d.CanPerformActionInAccount(ctx, iamT.DeleteNodepool); err != nil {
		return errors.NewE(err)
	}

	upv, err := d.PersistentVolumeRepo.Patch(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: pvName,
		},
		common.PatchForMarkDeletion(),
	)
	if err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeNodePool, upv.Name, ports.PublishUpdate)
	return d.ResDispatcher.DeleteFromTargetCluster(ctx, ctx.AccountName, clusterName, &upv.PersistentVolume)
}

// OnPVDeleteMessage implements Domain.
func (d *Domain) OnPVDeleteMessage(ctx domainT.InfraContext, clusterName string, pv entities.PersistentVolume) error {
	if err := d.PersistentVolumeRepo.DeleteOne(ctx, repos.Filter{
		fields.MetadataName: pv.Name,
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
	}); err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypePV, pv.Name, ports.PublishDelete)
	return nil
}

// OnPVUpdateMessage implements Domain.
func (d *Domain) OnPVUpdateMessage(ctx domainT.InfraContext, clusterName string, pv entities.PersistentVolume, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error {
	pv.SyncStatus.LastSyncedAt = opts.MessageTimestamp
	pv.SyncStatus.State = func() t.SyncState {
		if status == types.ResourceStatusDeleting {
			return t.SyncStateDeletingAtAgent
		}
		return t.SyncStateUpdatedAtAgent
	}()
	pv.AccountName = ctx.AccountName
	pv.ClusterName = clusterName
	upsert, err := d.PersistentVolumeRepo.Upsert(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: pv.Name,
	}, &pv)
	if err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypePV, upsert.Name, ports.PublishUpdate)
	return nil
}
