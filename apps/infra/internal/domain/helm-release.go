package domain

import (
	"time"

	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
)

func (d *domain) findHelmRelease(ctx InfraContext, clusterName string, hrName string) (*entities.HelmRelease, error) {
	cluster, err := d.helmReleaseRepo.FindOne(ctx, repos.Filter{
		"clusterName":   clusterName,
		"accountName":   ctx.AccountName,
		"metadata.name": hrName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cluster == nil {
		return nil, errors.Newf("cluster with name %q not found", clusterName)
	}
	return cluster, nil
}

func (d *domain) ListHelmReleases(ctx InfraContext, clusterName string, mf map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.HelmRelease], error) {

	if err := d.canPerformActionInAccount(ctx, iamT.ListHelmReleases); err != nil {
		return nil, errors.NewE(err)
	}

	f := repos.Filter{
		"clusterName": clusterName,
		"accountName": ctx.AccountName,
	}

	pr, err := d.helmReleaseRepo.FindPaginated(ctx, d.helmReleaseRepo.MergeMatchFilters(f, mf), pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return pr, nil

}

func (d *domain) GetHelmRelease(ctx InfraContext, clusterName string, hrName string) (*entities.HelmRelease, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.GetHelmRelease); err != nil {
		return nil, errors.NewE(err)
	}

	c, err := d.GetHelmRelease(ctx, clusterName, hrName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return c, nil
}

func (d *domain) CreateHelmRelease(ctx InfraContext, clusterName string, hr entities.HelmRelease) (*entities.HelmRelease, error) {

	if err := d.canPerformActionInAccount(ctx, iamT.CreateHelmRelease); err != nil {
		return nil, errors.NewE(err)
	}

	hr.IncrementRecordVersion()

	// ctx.AccountName

	hr.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	hr.LastUpdatedBy = hr.CreatedBy

	existing, err := d.helmReleaseRepo.FindOne(ctx, repos.Filter{
		"clusterName":   clusterName,
		"accountName":   ctx.AccountName,
		"metadata.name": hr.Name,
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	if existing != nil {
		return nil, errors.Newf("helm release with name %q already exists", hr.Name)
	}

	hr.AccountName = ctx.AccountName
	hr.ClusterName = clusterName

	if err := d.resDispatcher.ApplyToTargetCluster(ctx, clusterName, &hr.HelmChart, 1); err != nil {
		return nil, errors.NewE(err)
	}

	if cms, err := d.helmReleaseRepo.Create(ctx, &hr); err != nil {
		return nil, errors.NewE(err)
	} else {
		d.resourceEventPublisher.PublishHelmReleaseEvent(&hr, PublishAdd)
		return cms, nil
	}
}

func (d *domain) UpdateHelmRelease(ctx InfraContext, clusterName string, hr entities.HelmRelease) (*entities.HelmRelease, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.UpdateHelmRelease); err != nil {
		return nil, errors.NewE(err)
	}

	hr.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &hr); err != nil {
		return nil, errors.NewE(err)
	}

	cms, err := d.findHelmRelease(ctx, clusterName, hr.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cms.IsMarkedForDeletion() {
		return nil, errors.Newf("helm release with name %q is marked for deletion", hr.Name)
	}

	cms.IncrementRecordVersion()
	cms.LastUpdatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	cms.Labels = hr.Labels
	cms.Annotations = hr.Annotations

	cms.SyncStatus = t.GenSyncStatus(t.SyncActionApply, cms.RecordVersion)

	unp, err := d.helmReleaseRepo.UpdateById(ctx, cms.Id, cms)
	if err != nil {
		return nil, errors.NewE(err)
	}

	d.resourceEventPublisher.PublishHelmReleaseEvent(unp, PublishUpdate)

	if err := d.resDispatcher.ApplyToTargetCluster(ctx, clusterName, &unp.HelmChart, unp.RecordVersion); err != nil {
		return nil, errors.NewE(err)
	}

	return unp, nil
}

func (d *domain) DeleteHelmRelease(ctx InfraContext, clusterName string, name string) error {

	if err := d.canPerformActionInAccount(ctx, iamT.DeleteHelmRelease); err != nil {
		return errors.NewE(err)
	}

	svc, err := d.findHelmRelease(ctx, clusterName, name)
	if err != nil {
		return errors.NewE(err)
	}

	if svc.IsMarkedForDeletion() {
		return errors.Newf("helm release with name %q is marked for deletion", name)
	}

	svc.MarkedForDeletion = fn.New(true)
	svc.SyncStatus = t.GetSyncStatusForDeletion(svc.Generation)
	upC, err := d.helmReleaseRepo.UpdateById(ctx, svc.Id, svc)
	if err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishHelmReleaseEvent(upC, PublishUpdate)

	return d.resDispatcher.DeleteFromTargetCluster(ctx, clusterName, &upC.HelmChart)

}

func (d *domain) OnHelmReleaseApplyError(ctx InfraContext, clusterName string, name string, errMsg string) error {

	svc, err := d.findHelmRelease(ctx, clusterName, name)
	if err != nil {
		return errors.NewE(err)
	}

	svc.SyncStatus.State = t.SyncStateErroredAtAgent
	svc.SyncStatus.LastSyncedAt = time.Now()
	svc.SyncStatus.Error = &errMsg

	_, err = d.helmReleaseRepo.UpdateById(ctx, svc.Id, svc)
	d.resourceEventPublisher.PublishHelmReleaseEvent(svc, PublishUpdate)
	return errors.NewE(err)

}
func (d *domain) OnHelmReleaseDeleteMessage(ctx InfraContext, clusterName string, hr entities.HelmRelease) error {
	svc, _ := d.findHelmRelease(ctx, clusterName, hr.Name)
	if svc == nil {
		// does not exist, (maybe already deleted)
		return nil
	}

	if err := d.matchRecordVersion(hr.Annotations, svc.RecordVersion); err != nil {
		return d.resyncToTargetCluster(ctx, svc.SyncStatus.Action, clusterName, svc, svc.RecordVersion)
	}

	err := d.clusterManagedServiceRepo.DeleteById(ctx, svc.Id)
	d.resourceEventPublisher.PublishHelmReleaseEvent(svc, PublishDelete)
	return err

}
func (d *domain) OnHelmReleaseUpdateMessage(ctx InfraContext, clusterName string, hr entities.HelmRelease) error {
	svc, err := d.findHelmRelease(ctx, clusterName, hr.Name)
	if err != nil {
		hr.CreatedBy = common.CreatedOrUpdatedBy{
			UserId:    repos.ID(common.CreatedOnTenantClusterUserId),
			UserName:  common.CreatedOnTenantClusterUserName,
			UserEmail: common.CreatedOnTenantClusterUserEmail,
		}
		hr.LastUpdatedBy = hr.CreatedBy
		_, err := d.helmReleaseRepo.Create(ctx, &hr)
		if err != nil {
			return errors.NewE(err)
		}
	}

	if err := d.matchRecordVersion(hr.Annotations, svc.RecordVersion); err != nil {
		return d.resyncToTargetCluster(ctx, svc.SyncStatus.Action, clusterName, svc, svc.RecordVersion)
	}

	svc.Status = hr.Status

	svc.SyncStatus.State = t.SyncStateReceivedUpdateFromAgent
	svc.SyncStatus.LastSyncedAt = time.Now()
	svc.SyncStatus.Error = nil
	svc.SyncStatus.RecordVersion = svc.RecordVersion

	if _, err := d.helmReleaseRepo.UpdateById(ctx, svc.Id, svc); err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishHelmReleaseEvent(svc, PublishUpdate)
	return nil

}
