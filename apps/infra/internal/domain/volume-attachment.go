package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

// GetVolumeAttachment implements Domain.
func (d *domain) GetVolumeAttachment(ctx InfraContext, clusterName string, volAttachmentName string) (*entities.VolumeAttachment, error) {
	//panic("unimplemented")
	volatt, err := d.volumeAttachmentRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"clusterName":   clusterName,
		"metadata.name": volAttachmentName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if volatt == nil {
		return nil, errors.Newf("persistent volume claim with name %q not found", volAttachmentName)
	}
	return volatt, nil
}

// ListVolumeAttachments implements Domain.
func (d *domain) ListVolumeAttachments(ctx InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.VolumeAttachment], error) {
	//panic("unimplemented")
	filter := repos.Filter{
		"accountName": ctx.AccountName,
		"clusterName": clusterName,
	}
	return d.volumeAttachmentRepo.FindPaginated(ctx, d.nodePoolRepo.MergeMatchFilters(filter, search), pagination)
}

// OnVolumeAttachmentDeleteMessage implements Domain.
func (d *domain) OnVolumeAttachmentDeleteMessage(ctx InfraContext, clusterName string, volumeAttachment entities.VolumeAttachment) error {
	//panic("unimplemented")
	if err := d.pvcRepo.DeleteOne(ctx, repos.Filter{
		"metadata.name":      volumeAttachment.Name,
		"metadata.namespace": volumeAttachment.Namespace,
		"accountName":        ctx.AccountName,
		"clusterName":        clusterName,
	}); err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishVolumeAttachmentEvent(&volumeAttachment, PublishDelete)
	//d.resourceEventPublisher.PublishPvcResEvent(&pvc, PublishDelete)
	return nil
}

// OnVolumeAttachmentUpdateMessage implements Domain.
func (d *domain) OnVolumeAttachmentUpdateMessage(ctx InfraContext, clusterName string, volumeAttachment entities.VolumeAttachment, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	//panic("unimplemented")
	volumeAttachment.SyncStatus = t.SyncStatus{
		LastSyncedAt: opts.MessageTimestamp,
		State: func() t.SyncState {
			if status == types.ResourceStatusDeleting {
				return t.SyncStateDeletingAtAgent
			}
			return t.SyncStateUpdatedAtAgent
		}(),
	}

	_, err := d.volumeAttachmentRepo.Create(ctx, &volumeAttachment)
	if err != nil {
		return errors.NewE(err)
	}
	return nil
}
