package ingress_hosts

import (
	"context"

	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
	networkingv1 "k8s.io/api/networking/v1"
)

func (d *Domain) ListDomainEntries(ctx domainT.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.DomainEntry], error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.ListDomainEntries); err != nil {
		return nil, errors.NewE(err)
	}

	filters := map[string]any{
		fields.AccountName: ctx.AccountName,
	}
	return d.DomainEntryRepo.FindPaginated(ctx, d.DomainEntryRepo.MergeMatchFilters(filters, search), pagination)
}

func (d *Domain) GetDomainEntry(ctx domainT.InfraContext, DomainName string) (*entities.DomainEntry, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.GetDomainEntry); err != nil {
		return nil, errors.NewE(err)
	}
	return d.findDomainEntry(ctx, ctx.AccountName, DomainName)
}

func (d *Domain) CreateDomainEntry(ctx domainT.InfraContext, de entities.DomainEntry) (*entities.DomainEntry, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.CreateDomainEntry); err != nil {
		return nil, errors.NewE(err)
	}
	de.AccountName = ctx.AccountName
	de.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	de.LastUpdatedBy = de.CreatedBy

	nde, err := d.DomainEntryRepo.Create(ctx, &de)
	if err != nil {
		return nil, errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, nde.ClusterName, ports.ResourceTypeDomainEntries, nde.DomainName, ports.PublishAdd)

	return nde, nil
}

func (d *Domain) UpdateDomainEntry(ctx domainT.InfraContext, de entities.DomainEntry) (*entities.DomainEntry, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.UpdateDomainEntry); err != nil {
		return nil, errors.NewE(err)
	}

	existing, err := d.findDomainEntry(ctx, ctx.AccountName, de.DomainName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	existing.DisplayName = de.DisplayName
	existing.LastUpdatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	newDe, err := d.DomainEntryRepo.UpdateById(ctx, existing.Id, existing)
	if err != nil {
		return nil, errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, newDe.ClusterName, ports.ResourceTypeDomainEntries, newDe.DomainName, ports.PublishUpdate)
	return newDe, nil
}

func (d *Domain) DeleteDomainEntry(ctx domainT.InfraContext, DomainName string) error {
	if err := d.CanPerformActionInAccount(ctx, iamT.DeleteDomainEntry); err != nil {
		return errors.NewE(err)
	}

	existing, err := d.findDomainEntry(ctx, ctx.AccountName, DomainName)
	if err != nil {
		return errors.NewE(err)
	}

	err = d.DomainEntryRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: DomainName,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, existing.ClusterName, ports.ResourceTypeDomainEntries, DomainName, ports.PublishDelete)
	return nil
}

func (d *Domain) findDomainEntry(ctx context.Context, accountName string, DomainName string) (*entities.DomainEntry, error) {
	filters := repos.Filter{
		fields.AccountName:       accountName,
		fc.DomainEntryDomainName: DomainName,
	}
	one, err := d.DomainEntryRepo.FindOne(ctx, filters)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if one == nil {
		return nil, errors.Newf("DomainName entry not found")
	}

	return one, nil
}

func (d *Domain) OnIngressUpdateMessage(ctx domainT.InfraContext, clusterName string, ingress networkingv1.Ingress, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error {
	for i := range ingress.Spec.Rules {
		DomainName := ingress.Spec.Rules[i].Host
		de, err := d.DomainEntryRepo.Upsert(ctx, repos.Filter{
			fields.AccountName:       ctx.AccountName,
			fields.ClusterName:       clusterName,
			fc.DomainEntryDomainName: DomainName,
		}, &entities.DomainEntry{
			ResourceMetadata: common.ResourceMetadata{
				DisplayName:   DomainName,
				CreatedBy:     common.CreatedOrUpdatedByResourceSync,
				LastUpdatedBy: common.CreatedOrUpdatedByResourceSync,
			},
			DomainName:  DomainName,
			AccountName: ctx.AccountName,
			ClusterName: clusterName,
		})
		if err != nil {
			return err
		}

		d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeDomainEntries, de.DomainName, ports.PublishUpdate)
	}

	return nil
}

func (d *Domain) OnIngressDeleteMessage(ctx domainT.InfraContext, clusterName string, ingress networkingv1.Ingress) error {
	DomainNames := make([]any, 0, len(ingress.Spec.Rules))
	for i := range ingress.Spec.Rules {
		DomainNames = append(DomainNames, ingress.Spec.Rules[i].Host)
	}

	filter := repos.Filter{
		fields.AccountName: ctx.AccountName,
		fields.ClusterName: clusterName,
	}

	filters := d.DomainEntryRepo.MergeMatchFilters(filter, map[string]repos.MatchFilter{
		fc.DomainEntryDomainName: {
			MatchType: repos.MatchTypeArray,
			Array:     DomainNames,
		},
	})

	err := d.DomainEntryRepo.DeleteMany(ctx, filters)
	if err != nil {
		return err
	}

	for i := range DomainNames {
		d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeDomainEntries, DomainNames[i].(string), ports.PublishDelete)
	}
	return nil
}
