package cluster_managed_service

import (
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func (d *Domain) ListClusterManagedServices(ctx domainT.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.ClusterManagedService], error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.ListClusterManagedServices); err != nil {
		return nil, errors.NewE(err)
	}

	f := repos.Filter{
		fields.AccountName: ctx.AccountName,
	}

	pr, err := d.ClusterManagedServiceRepo.FindPaginated(ctx, d.ClusterManagedServiceRepo.MergeMatchFilters(f, search), pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return pr, nil
}

func (d *Domain) FindClusterManagedService(ctx domainT.InfraContext, name string) (*entities.ClusterManagedService, error) {
	cmsvc, err := d.ClusterManagedServiceRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.MetadataName: name,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cmsvc == nil {
		return nil, errors.Newf("cmsvc with name %q not found", name)
	}
	return cmsvc, nil
}

func (d *Domain) GetClusterManagedService(ctx domainT.InfraContext, serviceName string) (*entities.ClusterManagedService, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.GetClusterManagedService); err != nil {
		return nil, errors.NewE(err)
	}

	c, err := d.FindClusterManagedService(ctx, serviceName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return c, nil
}

func (d *Domain) applyClusterManagedService(ctx domainT.InfraContext, cmsvc *entities.ClusterManagedService) error {
	d.AddTrackingId(&cmsvc.ClusterManagedService, cmsvc.Id)
	return d.ResDispatcher.ApplyToTargetCluster(ctx, ctx.AccountName, cmsvc.ClusterName, &cmsvc.ClusterManagedService, cmsvc.RecordVersion)
}

func (d *Domain) CreateClusterManagedService(ctx domainT.InfraContext, cmsvc entities.ClusterManagedService) (*entities.ClusterManagedService, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.CreateClusterManagedService); err != nil {
		return nil, errors.NewE(err)
	}

	cmsvc.IncrementRecordVersion()

	cmsvc.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	cmsvc.LastUpdatedBy = cmsvc.CreatedBy

	existing, err := d.ClusterManagedServiceRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.MetadataName: cmsvc.Name,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if existing != nil {
		return nil, errors.Newf("cluster managed service with name %q already exists", cmsvc.ClusterName)
	}

	cmsvc.AccountName = ctx.AccountName
	cmsvc.SyncStatus = t.GenSyncStatus(t.SyncActionApply, cmsvc.RecordVersion)

	// cmsvc.Spec.SharedSecret = fn.New(fn.CleanerNanoid(40))

	cmsvc.EnsureGVK()

	if err := d.K8sClient.ValidateObject(ctx, &cmsvc.ClusterManagedService); err != nil {
		return nil, errors.NewE(err)
	}

	ncms, err := d.ClusterManagedServiceRepo.Create(ctx, &cmsvc)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if err := d.applyClusterManagedService(ctx, &cmsvc); err != nil {
		return nil, errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, cmsvc.ClusterName, ports.ResourceTypeClusterManagedService, ncms.Name, ports.PublishAdd)

	return ncms, nil
}

func (d *Domain) UpdateClusterManagedService(ctx domainT.InfraContext, cmsvc entities.ClusterManagedService) (*entities.ClusterManagedService, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.UpdateClusterManagedService); err != nil {
		return nil, errors.NewE(err)
	}

	cmsvc.EnsureGVK()
	if err := d.K8sClient.ValidateObject(ctx, &cmsvc.ClusterManagedService); err != nil {
		return nil, errors.NewE(err)
	}

	patchForUpdate := common.PatchForUpdate(
		ctx,
		&cmsvc,
		common.PatchOpts{
			XPatch: repos.Document{
				fc.ClusterManagedServiceSpecMsvcSpec: cmsvc.Spec.MSVCSpec,
			},
		})

	ucmsvc, err := d.ClusterManagedServiceRepo.Patch(ctx, repos.Filter{fields.AccountName: ctx.AccountName, fields.MetadataName: cmsvc.Name}, patchForUpdate)
	if err != nil {
		return nil, errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, ucmsvc.ClusterName, ports.ResourceTypeClusterManagedService, ucmsvc.Name, ports.PublishUpdate)

	if err := d.applyClusterManagedService(ctx, ucmsvc); err != nil {
		return nil, errors.NewE(err)
	}

	return ucmsvc, nil
}

func (d *Domain) DeleteClusterManagedService(ctx domainT.InfraContext, name string) error {
	if err := d.CanPerformActionInAccount(ctx, iamT.DeleteClusterManagedService); err != nil {
		return errors.NewE(err)
	}

	ucmsvc, err := d.ClusterManagedServiceRepo.Patch(ctx, repos.Filter{fields.AccountName: ctx.AccountName, fields.MetadataName: name}, common.PatchForMarkDeletion())
	if err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, ucmsvc.ClusterName, ports.ResourceTypeClusterManagedService, ucmsvc.Name, ports.PublishUpdate)

	return d.ResDispatcher.DeleteFromTargetCluster(ctx, ctx.AccountName, ucmsvc.ClusterName, &ucmsvc.ClusterManagedService)
}

func (d *Domain) OnClusterManagedServiceApplyError(ctx domainT.InfraContext, clusterName, name, errMsg string, opts domainT.UpdateAndDeleteOpts) error {
	ucmsvc, err := d.ClusterManagedServiceRepo.Patch(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
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

	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeClusterManagedService, ucmsvc.Name, ports.PublishDelete)
	return errors.NewE(err)
}

func (d *Domain) OnClusterManagedServiceDeleteMessage(ctx domainT.InfraContext, clusterName string, service entities.ClusterManagedService) error {
	err := d.ClusterManagedServiceRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: service.Name,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeClusterManagedService, service.Name, ports.PublishDelete)
	return err
}

func (d *Domain) OnClusterManagedServiceUpdateMessage(ctx domainT.InfraContext, clusterName string, service entities.ClusterManagedService, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error {
	xService, err := d.FindClusterManagedService(ctx, service.Name)
	if err != nil {
		return errors.NewE(err)
	}

	if xService == nil {
		return errors.Newf("no cluster manage service found")
	}

	if _, err := d.MatchRecordVersion(service.Annotations, xService.RecordVersion); err != nil {
		return d.ResyncToTargetCluster(ctx, xService.SyncStatus.Action, clusterName, xService, xService.RecordVersion)
	}

	recordVersion, err := d.MatchRecordVersion(service.Annotations, xService.RecordVersion)
	if err != nil {
		return errors.NewE(err)
	}

	patch := repos.Document{
		fc.ClusterManagedServiceSpecTargetNamespace: service.Spec.TargetNamespace,
	}

	ucmsvc, err := d.ClusterManagedServiceRepo.PatchById(ctx, xService.Id, common.PatchForSyncFromAgent(&service, recordVersion, status, common.PatchOpts{
		MessageTimestamp: opts.MessageTimestamp,
		XPatch:           patch,
	}))
	if err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeClusterManagedService, ucmsvc.GetName(), ports.PublishUpdate)
	return nil
}
