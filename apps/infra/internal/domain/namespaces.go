package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func (d *domain) findNamespace(ctx InfraContext, clusterName string, namespaceName string) (*entities.Namespace, error) {
	ns, err := d.namespaceRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"clusterName":   clusterName,
		"metadata.name": namespaceName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	if ns == nil {
		return nil, errors.Newf("namespace with name %q not found", namespaceName)
	}
	return ns, nil
}

// GetNamespace implements Domain.
func (d *domain) GetNamespace(ctx InfraContext, clusterName string, namespace string) (*entities.Namespace, error) {
	//panic("unimplemented")
	ns, err := d.namespaceRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"clusterName":   clusterName,
		"metadata.name": namespace,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if ns == nil {
		return nil, errors.Newf("namespace with name %q not found", namespace)
	}
	return ns, nil
}

// ListNamespaces implements Domain.
func (d *domain) ListNamespaces(ctx InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Namespace], error) {
	//panic("unimplemented")
	filter := repos.Filter{
		"accountName": ctx.AccountName,
		"clusterName": clusterName,
	}
	return d.namespaceRepo.FindPaginated(ctx, d.namespaceRepo.MergeMatchFilters(filter, search), pagination)
}

// OnNamespaceDeleteMessage implements Domain.
func (d *domain) OnNamespaceDeleteMessage(ctx InfraContext, clusterName string, namespace entities.Namespace) error {
	//panic("unimplemented")
	if err := d.namespaceRepo.DeleteOne(ctx, repos.Filter{
		"metadata.name":      namespace.Name,
		"metadata.namespace": namespace.Namespace,
		"accountName":        ctx.AccountName,
		"clusterName":        clusterName,
	}); err != nil {
		return errors.NewE(err)
	}
	return nil
}

// OnNamespaceUpdateMessage implements Domain.
func (d *domain) OnNamespaceUpdateMessage(ctx InfraContext, clusterName string, namespace entities.Namespace, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	namespace.SyncStatus = t.SyncStatus{
		LastSyncedAt: opts.MessageTimestamp,
		State: func() t.SyncState {
			if status == types.ResourceStatusDeleting {
				return t.SyncStateDeletingAtAgent
			}
			return t.SyncStateUpdatedAtAgent
		}(),
	}

	_, err := d.namespaceRepo.Create(ctx, &namespace)
	if err != nil {
		return errors.NewE(err)
	}
	return nil
}
