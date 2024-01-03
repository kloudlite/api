package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

// GetPV implements Domain.
func (*domain) GetPV(ctx InfraContext, clusterName string, pvcName string) (*entities.PersistentVolumeClaim, error) {
	panic("unimplemented")
}

// ListPVs implements Domain.
func (*domain) ListPVs(ctx InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.PersistentVolumeClaim], error) {
	panic("unimplemented")
}

// OnPVDeleteMessage implements Domain.
func (*domain) OnPVDeleteMessage(ctx InfraContext, clusterName string, pvc entities.PersistentVolumeClaim) error {
	panic("unimplemented")
}

// OnPVUpdateMessage implements Domain.
func (*domain) OnPVUpdateMessage(ctx InfraContext, clusterName string, pvc entities.PersistentVolumeClaim, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	panic("unimplemented")
}
