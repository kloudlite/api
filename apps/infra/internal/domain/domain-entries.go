package domain

import (
	"context"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
)

func (d *domain) ListDomainEntries(ctx InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.DomainEntry], error) {
	if err := d.canPerformActionInAccount(ctx, iamT.ListDomainEntries); err != nil {
		return nil, errors.NewE(err)
	}

	filters := map[string]any{
		fields.AccountName: ctx.AccountName,
	}
	return d.domainEntryRepo.FindPaginated(ctx, d.domainEntryRepo.MergeMatchFilters(filters, search), pagination)
}

func (d *domain) GetDomainEntry(ctx InfraContext, domainName string) (*entities.DomainEntry, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.GetDomainEntry); err != nil {
		return nil, errors.NewE(err)
	}
	return d.findDomainEntry(ctx, ctx.AccountName, domainName)
}

func (d *domain) CreateDomainEntry(ctx InfraContext, de entities.DomainEntry) (*entities.DomainEntry, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.CreateDomainEntry); err != nil {
		return nil, errors.NewE(err)
	}
	de.AccountName = ctx.AccountName
	de.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	de.LastUpdatedBy = de.CreatedBy

	nde, err := d.domainEntryRepo.Create(ctx, &de)
	if err != nil {
		return nil, errors.NewE(err)
	}
	d.resourceEventPublisher.PublishResourceEvent(ctx, nde.ClusterName, ResourceTypeDomainEntries, nde.DomainName, PublishAdd)

	return nde, nil
}

func (d *domain) UpdateDomainEntry(ctx InfraContext, de entities.DomainEntry) (*entities.DomainEntry, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.UpdateDomainEntry); err != nil {
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

	newDe, err := d.domainEntryRepo.UpdateById(ctx, existing.Id, existing)
	if err != nil {
		return nil, errors.NewE(err)
	}
	d.resourceEventPublisher.PublishResourceEvent(ctx, newDe.ClusterName, ResourceTypeDomainEntries, newDe.DomainName, PublishUpdate)
	return newDe, nil
}

func (d *domain) DeleteDomainEntry(ctx InfraContext, domainName string) error {
	if err := d.canPerformActionInAccount(ctx, iamT.DeleteDomainEntry); err != nil {
		return errors.NewE(err)
	}

	existing, err := d.findDomainEntry(ctx, ctx.AccountName, domainName)
	if err != nil {
		return errors.NewE(err)
	}

	err = d.domainEntryRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: domainName,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishResourceEvent(ctx, existing.ClusterName, ResourceTypeDomainEntries, domainName, PublishDelete)
	return nil
}

func (d *domain) findDomainEntry(ctx context.Context, accountName string, domainName string) (*entities.DomainEntry, error) {
	filters := repos.Filter{
		fields.AccountName:       accountName,
		fc.DomainEntryDomainName: domainName,
	}
	one, err := d.domainEntryRepo.FindOne(ctx, filters)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if one == nil {
		return nil, errors.Newf("domainName entry not found")
	}

	return one, nil
}
