package domain

import (
	"kloudlite.io/apps/console/internal/domain/entities"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

func (d *domain) ListImagePullSecrets(ctx ConsoleContext, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.ImagePullSecret], error) {
	if err := d.canReadSecretsFromAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return nil, err
	}

	return d.ipsRepo.FindPaginated(ctx, repos.Filter{"accountName": ctx.AccountName}, pagination)
}

func (d *domain) GetImagePullSecret(ctx ConsoleContext, name string) (*entities.ImagePullSecret, error) {
	if err := d.canReadSecretsFromAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return nil, err
	}

	return d.ipsRepo.FindOne(ctx, repos.Filter{"accountName": ctx.AccountName, "name": name})
}

func (d *domain) CreateImagePullSecret(ctx ConsoleContext, secret entities.ImagePullSecret) (*entities.ImagePullSecret, error) {
	if err := d.canMutateSecretsInAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return nil, err
	}
	return d.ipsRepo.Create(ctx, &secret)
}

func (d *domain) DeleteImagePullSecret(ctx ConsoleContext, name string) error {
	if err := d.canMutateSecretsInAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return err
	}
	return d.ipsRepo.DeleteOne(ctx, repos.Filter{"accountName": ctx.AccountName, "name": name})
}
