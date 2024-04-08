package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func (d *domain) findClusterConn(ctx InfraContext, clusterName string, connName string) (*entities.ClusterConnection, error) {
	cc, err := d.clusterConnRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: connName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	if cc == nil {
		return nil, errors.Newf("cluster connection with name %q not found", clusterName)
	}
	return cc, nil
}

func (d *domain) OnClusterConnDeleteMessage(ctx InfraContext, clusterName string, clusterConn entities.ClusterConnection) error {
	err := d.clusterConnRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.ClusterName:  clusterName,
			fields.MetadataName: clusterConn.Name,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeClusterConnection, clusterConn.Name, PublishDelete)
	return err
}

func (d *domain) OnClusterConnUpdateMessage(ctx InfraContext, clusterName string, clusterConn entities.ClusterConnection, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	xnp, err := d.findClusterConn(ctx, clusterName, clusterConn.Name)
	if err != nil {
		return errors.NewE(err)
	}

	if xnp == nil {
		return errors.Newf("no cluster connection found")
	}

	if _, err := d.matchRecordVersion(clusterConn.Annotations, xnp.RecordVersion); err != nil {
		return d.resyncToTargetCluster(ctx, xnp.SyncStatus.Action, clusterName, &xnp.ClusterConnection, xnp.RecordVersion)
	}

	recordVersion, err := d.matchRecordVersion(clusterConn.Annotations, xnp.RecordVersion)
	if err != nil {
		return errors.NewE(err)
	}

	unp, err := d.clusterConnRepo.PatchById(
		ctx,
		xnp.Id,
		common.PatchForSyncFromAgent(&clusterConn,
			recordVersion, status,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
			}))
	if err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeClusterConnection, unp.GetName(), PublishUpdate)
	return nil
}
func (d *domain) OnClusterConnApplyError(ctx InfraContext, clusterName string, name string, errMsg string, opts UpdateAndDeleteOpts) error {
	unp, err := d.clusterConnRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.ClusterName:  clusterName,
			fields.MetadataName: name,
		},
		common.PatchForErrorFromAgent(
			errMsg,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
			},
		),
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeClusterConnection, unp.Name, PublishUpdate)
	return errors.NewE(err)
}
