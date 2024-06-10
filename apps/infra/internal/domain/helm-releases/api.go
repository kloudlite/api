package helm_releases

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (d *Domain) findHelmRelease(ctx domainT.InfraContext, clusterName string, hrName string) (*entities.HelmRelease, error) {
	cluster, err := d.HelmReleaseRepo.FindOne(ctx, repos.Filter{
		fields.ClusterName:  clusterName,
		fields.AccountName:  ctx.AccountName,
		fields.MetadataName: hrName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cluster == nil {
		return nil, errors.Newf("helm release with name %q not found", hrName)
	}
	return cluster, nil
}

func (d *Domain) upsertHelmRelease(ctx domainT.InfraContext, clusterName string, hr *entities.HelmRelease) (*entities.HelmRelease, error) {
	cluster, err := d.HelmReleaseRepo.Upsert(ctx, repos.Filter{
		fields.ClusterName:  clusterName,
		fields.AccountName:  ctx.AccountName,
		fields.MetadataName: hr.Name,
	}, hr)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cluster == nil {
		return nil, errors.Newf("could not upsert helm release %s", hr.Name)
	}
	return cluster, nil
}

func (d *Domain) ListHelmReleases(ctx domainT.InfraContext, clusterName string, mf map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.HelmRelease], error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.ListHelmReleases); err != nil {
		return nil, errors.NewE(err)
	}

	f := repos.Filter{
		fields.ClusterName: clusterName,
		fields.AccountName: ctx.AccountName,
	}

	pr, err := d.HelmReleaseRepo.FindPaginated(ctx, d.HelmReleaseRepo.MergeMatchFilters(f, mf), pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return pr, nil
}

func (d *Domain) GetHelmRelease(ctx domainT.InfraContext, clusterName string, hrName string) (*entities.HelmRelease, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.GetHelmRelease); err != nil {
		return nil, errors.NewE(err)
	}

	c, err := d.GetHelmRelease(ctx, clusterName, hrName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return c, nil
}

func (d *Domain) applyHelmRelease(ctx domainT.InfraContext, hr *entities.HelmRelease) error {
	d.AddTrackingId(&hr.HelmChart, hr.Id)
	return d.ResDispatcher.ApplyToTargetCluster(ctx, ctx.AccountName, hr.ClusterName, &hr.HelmChart, hr.RecordVersion)
}

func (d *Domain) CreateHelmRelease(ctx domainT.InfraContext, clusterName string, hr entities.HelmRelease) (*entities.HelmRelease, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.CreateHelmRelease); err != nil {
		return nil, errors.NewE(err)
	}
	hr.EnsureGVK()
	if err := d.K8sClient.ValidateObject(ctx, &hr.HelmChart); err != nil {
		return nil, errors.NewE(err)
	}

	hr.IncrementRecordVersion()
	hr.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	hr.LastUpdatedBy = hr.CreatedBy

	existing, err := d.HelmReleaseRepo.FindOne(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: hr.Name,
		},
	)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if existing != nil {
		return nil, errors.Newf("helm release with name %q already exists", hr.Name)
	}

	hr.AccountName = ctx.AccountName
	hr.ClusterName = clusterName
	hr.SyncStatus = t.GenSyncStatus(t.SyncActionApply, hr.RecordVersion)

	nhr, err := d.HelmReleaseRepo.Create(ctx, &hr)
	if err != nil {
		return nil, errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, nhr.ClusterName, ports.ResourceTypeHelmRelease, nhr.Name, ports.PublishAdd)

	if err = d.ResDispatcher.ApplyToTargetCluster(ctx, ctx.AccountName, clusterName, &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: hr.Namespace,
		},
	}, hr.RecordVersion); err != nil {
		return nil, errors.NewE(err)
	}

	if err := d.applyHelmRelease(ctx, nhr); err != nil {
		return nil, errors.NewE(err)
	}

	return nhr, nil
}

func (d *Domain) UpdateHelmRelease(ctx domainT.InfraContext, clusterName string, hrIn entities.HelmRelease) (*entities.HelmRelease, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.UpdateHelmRelease); err != nil {
		return nil, errors.NewE(err)
	}

	hrIn.EnsureGVK()
	if err := d.K8sClient.ValidateObject(ctx, &hrIn); err != nil {
		return nil, errors.NewE(err)
	}

	patchForUpdate := common.PatchForUpdate(
		ctx,
		&hrIn,
		common.PatchOpts{
			XPatch: repos.Document{
				fc.HelmReleaseSpecChartVersion: hrIn.Spec.ChartVersion,
				fc.HelmReleaseSpecValues:       hrIn.Spec.Values,
			},
		})

	uphr, err := d.HelmReleaseRepo.Patch(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: hrIn.Name,
		},
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, uphr.ClusterName, ports.ResourceTypeHelmRelease, uphr.Name, ports.PublishUpdate)
	if err := d.applyHelmRelease(ctx, uphr); err != nil {
		return nil, errors.NewE(err)
	}
	return uphr, nil
}

func (d *Domain) DeleteHelmRelease(ctx domainT.InfraContext, clusterName string, name string) error {
	if err := d.CanPerformActionInAccount(ctx, iamT.DeleteHelmRelease); err != nil {
		return errors.NewE(err)
	}

	uphr, err := d.HelmReleaseRepo.Patch(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: name,
		},
		common.PatchForMarkDeletion(),
	)
	if err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, uphr.ClusterName, ports.ResourceTypeHelmRelease, uphr.Name, ports.PublishUpdate)

	return d.ResDispatcher.DeleteFromTargetCluster(ctx, ctx.AccountName, clusterName, &uphr.HelmChart)
}

func (d *Domain) OnHelmReleaseApplyError(ctx domainT.InfraContext, clusterName string, name string, errMsg string, opts domainT.UpdateAndDeleteOpts) error {
	uphr, err := d.HelmReleaseRepo.Patch(
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
	d.ResourceEventPublisher.PublishResourceEvent(ctx, uphr.ClusterName, ports.ResourceTypeHelmRelease, uphr.Name, ports.PublishUpdate)
	return errors.NewE(err)
}

func (d *Domain) OnHelmReleaseDeleteMessage(ctx domainT.InfraContext, clusterName string, hr entities.HelmRelease) error {
	err := d.HelmReleaseRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: hr.Name,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeHelmRelease, hr.Name, ports.PublishDelete)
	return err
}

func (d *Domain) OnHelmReleaseUpdateMessage(ctx domainT.InfraContext, clusterName string, hr entities.HelmRelease, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error {
	xhr, err := d.findHelmRelease(ctx, clusterName, hr.Name)
	if err != nil {
		return errors.NewE(err)
	}

	recordVersion, err := d.MatchRecordVersion(hr.Annotations, xhr.RecordVersion)
	if err != nil {
		return d.ResyncToTargetCluster(ctx, xhr.SyncStatus.Action, clusterName, xhr, xhr.RecordVersion)
	}

	uphr, err := d.HelmReleaseRepo.PatchById(
		ctx,
		xhr.Id,
		common.PatchForSyncFromAgent(&hr, recordVersion, status, common.PatchOpts{
			MessageTimestamp: opts.MessageTimestamp,
		}))
	if err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, uphr.ClusterName, ports.ResourceTypeHelmRelease, uphr.GetName(), ports.PublishUpdate)
	return nil
}
