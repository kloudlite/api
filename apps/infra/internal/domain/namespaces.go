package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

// GetNamespace implements Domain.
func (*domain) GetNamespace(ctx InfraContext, clusterName string, pvcName string) (*entities.PersistentVolumeClaim, error) {
	panic("unimplemented")
}

// ListNamespaces implements Domain.
func (*domain) ListNamespaces(ctx InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.PersistentVolumeClaim], error) {
	panic("unimplemented")
}

// OnNamespaceDeleteMessage implements Domain.
func (*domain) OnNamespaceDeleteMessage(ctx InfraContext, clusterName string, pvc entities.PersistentVolumeClaim) error {
	panic("unimplemented")
}

// OnNamespaceUpdateMessage implements Domain.
func (*domain) OnNamespaceUpdateMessage(ctx InfraContext, clusterName string, pvc entities.PersistentVolumeClaim, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	panic("unimplemented")
}
