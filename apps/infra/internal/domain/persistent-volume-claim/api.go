package persistent_volume_claim

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func (d *Domain) ListPVCs(ctx domainT.InfraContext, clusterName string, matchFilters map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.PersistentVolumeClaim], error) {
	filter := repos.Filter{
		fields.AccountName: ctx.AccountName,
		fields.ClusterName: clusterName,
	}
	return d.PersistentVolumeClaimRepo.FindPaginated(ctx, d.PersistentVolumeClaimRepo.MergeMatchFilters(filter, matchFilters), pagination)
}

func (d *Domain) GetPVC(ctx domainT.InfraContext, clusterName string, buildRunName string) (*entities.PersistentVolumeClaim, error) {
	pvc, err := d.PersistentVolumeClaimRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: buildRunName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if pvc == nil {
		return nil, errors.Newf("persistent volume claim with name %q not found", buildRunName)
	}
	return pvc, nil
}

func (d *Domain) OnPVCUpdateMessage(ctx domainT.InfraContext, clusterName string, pvc entities.PersistentVolumeClaim, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error {
	xpvc, err := d.PersistentVolumeClaimRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: pvc.Name,
	})
	if err != nil {
		return err
	}

	if xpvc == nil {
		pvc.AccountName = ctx.AccountName
		pvc.ClusterName = clusterName

		pvc.CreatedBy = common.CreatedOrUpdatedBy{
			UserId:    repos.ID(common.CreatedByResourceSyncUserId),
			UserName:  common.CreatedByResourceSyncUsername,
			UserEmail: common.CreatedByResourceSyncUserEmail,
		}
		pvc.LastUpdatedBy = pvc.CreatedBy
		xpvc, err = d.PersistentVolumeClaimRepo.Create(ctx, &pvc)
		if err != nil {
			return errors.NewE(err)
		}
	}

	upvc, err := d.PersistentVolumeClaimRepo.PatchById(
		ctx,
		xpvc.Id,
		common.PatchForSyncFromAgent(&pvc, pvc.RecordVersion, status, common.PatchOpts{
			MessageTimestamp: opts.MessageTimestamp,
		}))
	if err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypePVC, upvc.Name, ports.PublishUpdate)
	return nil
}

func (d *Domain) OnPVCDeleteMessage(ctx domainT.InfraContext, clusterName string, pvc entities.PersistentVolumeClaim) error {
	if err := d.PersistentVolumeClaimRepo.DeleteOne(ctx, repos.Filter{
		fields.MetadataName:      pvc.Name,
		fields.MetadataNamespace: pvc.Namespace,
		fields.AccountName:       ctx.AccountName,
		fields.ClusterName:       clusterName,
	}); err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypePVC, pvc.Name, ports.PublishDelete)
	return nil
}
