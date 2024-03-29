package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

// GetNamespace implements Domain.
func (d *domain) GetNamespace(ctx InfraContext, clusterName string, namespace string) (*entities.Namespace, error) {
	ns, err := d.namespaceRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: namespace,
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
	filter := repos.Filter{
		fields.AccountName: ctx.AccountName,
		fields.ClusterName: clusterName,
	}
	return d.namespaceRepo.FindPaginated(ctx, d.namespaceRepo.MergeMatchFilters(filter, search), pagination)
}

// OnNamespaceDeleteMessage implements Domain.
func (d *domain) OnNamespaceDeleteMessage(ctx InfraContext, clusterName string, namespace entities.Namespace) error {
	if err := d.namespaceRepo.DeleteOne(ctx, repos.Filter{
		fields.MetadataName:      namespace.Name,
		fields.MetadataNamespace: namespace.Namespace,
		fields.AccountName:       ctx.AccountName,
		fields.ClusterName:       clusterName,
	}); err != nil {
		return errors.NewE(err)
	}
	return nil
}

// OnNamespaceUpdateMessage implements Domain.
func (d *domain) OnNamespaceUpdateMessage(ctx InfraContext, clusterName string, namespace entities.Namespace, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	ns, err := d.namespaceRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: namespace.Name,
	})
	if err != nil {
		return err
	}

	if ns == nil {
		namespace.AccountName = ctx.AccountName
		namespace.ClusterName = clusterName

		namespace.CreatedBy = common.CreatedOrUpdatedBy{
			UserId:    repos.ID(common.CreatedByResourceSyncUserId),
			UserName:  common.CreatedByResourceSyncUsername,
			UserEmail: common.CreatedByResourceSyncUserEmail,
		}
		namespace.LastUpdatedBy = namespace.CreatedBy

		ns, err = d.namespaceRepo.Create(ctx, &namespace)
		if err != nil {
			return errors.NewE(err)
		}
	}

	_, err = d.namespaceRepo.PatchById(
		ctx,
		ns.Id,
		common.PatchForSyncFromAgent(&namespace, ns.RecordVersion, status, common.PatchOpts{
			MessageTimestamp: opts.MessageTimestamp,
		}))
	if err != nil {
		return errors.NewE(err)
	}
	return nil
}
