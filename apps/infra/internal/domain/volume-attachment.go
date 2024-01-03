package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

// GetVolumeAttachment implements Domain.
func (*domain) GetVolumeAttachment(ctx InfraContext, clusterName string, pvcName string) (*entities.PersistentVolumeClaim, error) {
	panic("unimplemented")
}

// ListVolumeAttachments implements Domain.
func (*domain) ListVolumeAttachments(ctx InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.PersistentVolumeClaim], error) {
	panic("unimplemented")
}

// OnVolumeAttachmentDeleteMessage implements Domain.
func (*domain) OnVolumeAttachmentDeleteMessage(ctx InfraContext, clusterName string, pvc entities.PersistentVolumeClaim) error {
	panic("unimplemented")
}

// OnVolumeAttachmentUpdateMessage implements Domain.
func (*domain) OnVolumeAttachmentUpdateMessage(ctx InfraContext, clusterName string, pvc entities.PersistentVolumeClaim, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	panic("unimplemented")
}
