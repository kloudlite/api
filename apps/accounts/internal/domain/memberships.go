package domain

import (
	"context"
	"kloudlite.io/apps/accounts/internal/entities"
	iamT "kloudlite.io/apps/iam/types"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam"
	"kloudlite.io/pkg/errors"
	"kloudlite.io/pkg/repos"
	"strings"
)

func (d *domain) addMembership(ctx context.Context, accountName string, userId repos.ID, resourceType iamT.ResourceType, role iamT.Role) error {
	if _, err := d.iamClient.AddMembership(ctx, &iam.AddMembershipIn{
		UserId:       string(userId),
		ResourceType: string(resourceType),
		ResourceRef:  iamT.NewResourceRef(accountName, iamT.ResourceAccount, accountName),
		Role:         string(role),
	}); err != nil {
		return err
	}

	return nil
}

func (d *domain) RemoveAccountMembership(ctx UserContext, accountName string, memberId repos.ID) (bool, error) {
	if err := d.checkAccountAccess(ctx, accountName, ctx.UserId, iamT.RemoveAccountMembership); err != nil {
		return false, err
	}

	account, err := d.findAccount(ctx, accountName)
	if err != nil {
		return false, err
	}

	if (account.IsActive != nil && !*account.IsActive) || account.IsMarkedForDeletion() {
		return false, errors.Newf("account %q is not active, or marked for deletion, aborting request", accountName)
	}

	out, err := d.iamClient.RemoveMembership(ctx, &iam.RemoveMembershipIn{
		UserId:      string(memberId),
		ResourceRef: iamT.NewResourceRef(accountName, iamT.ResourceAccount, accountName),
	})
	if err != nil {
		return false, err
	}

	return out.Result, nil
}

func (d *domain) UpdateAccountMembership(ctx UserContext, accountName string, memberId repos.ID, role iamT.Role) (bool, error) {
	if err := d.checkAccountAccess(ctx, accountName, ctx.UserId, iamT.RemoveAccountMembership); err != nil {
		return false, err
	}

	account, err := d.findAccount(ctx, accountName)
	if err != nil {
		return false, err
	}

	if (account.IsActive != nil && !*account.IsActive) || account.IsMarkedForDeletion() {
		return false, errors.Newf("account %q is not active, or marked for deletion, aborting request", accountName)
	}

	out, err := d.iamClient.RemoveMembership(ctx, &iam.RemoveMembershipIn{
		UserId:      string(memberId),
		ResourceRef: iamT.NewResourceRef(accountName, iamT.ResourceAccount, accountName),
	})
	if err != nil {
		return false, err
	}

	return out.Result, nil
}

func (d *domain) GetUserMemberships(ctx UserContext, accountName string, resourceRef string) ([]*entities.AccountMembership, error) {
	panic("not implemented yet")
}

func (d *domain) ListAccountMemberships(ctx UserContext) ([]*entities.AccountMembership, error) {
	out, err := d.iamClient.ListUserMemberships(ctx, &iam.UserMembershipsIn{
		UserId:       string(ctx.UserId),
		ResourceType: string(iamT.ResourceAccount),
	})
	if err != nil {
		return nil, err
	}

	memberships := make([]*entities.AccountMembership, len(out.RoleBindings))
	for i := range out.RoleBindings {
		memberships[i] = &entities.AccountMembership{
			AccountName: strings.Split(out.RoleBindings[i].ResourceRef, "/")[0],
			UserId:      repos.ID(out.RoleBindings[i].UserId),
			Role:        iamT.Role(out.RoleBindings[i].Role),
		}
	}

	return memberships, nil
}

func (d *domain) GetAccountMembership(ctx UserContext, accountName string) (*entities.AccountMembership, error) {
	membership, err := d.iamClient.GetMembership(
		ctx, &iam.GetMembershipIn{
			UserId:       string(ctx.UserId),
			ResourceType: string(iamT.ResourceAccount),
			ResourceRef:  iamT.NewResourceRef(accountName, iamT.ResourceAccount, accountName),
		},
	)
	if err != nil {
		return nil, err
	}
	return &entities.AccountMembership{
		AccountName: accountName,
		UserId:      repos.ID(membership.UserId),
		Role:        iamT.Role(membership.Role),
	}, nil
}
