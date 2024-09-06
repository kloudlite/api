package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"github.com/kloudlite/api/apps/finance/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/finance/internal/app/graph/model"
	"github.com/kloudlite/api/apps/finance/internal/domain"
	"github.com/kloudlite/api/apps/finance/internal/entities"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
)

// User is the resolver for the User field.
func (r *accountMembershipResolver) User(ctx context.Context, obj *entities.AccountMembership) (*model.User, error) {
	return &model.User{
		ID: obj.UserId,
	}, nil
}

// AccountsCreateAccount is the resolver for the accounts_createAccount field.
func (r *mutationResolver) AccountsCreateAccount(ctx context.Context, account entities.Account) (*entities.Account, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.domain.CreateAccount(uc, account)
}

// AccountsUpdateAccount is the resolver for the accounts_updateAccount field.
func (r *mutationResolver) AccountsUpdateAccount(ctx context.Context, account entities.Account) (*entities.Account, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.domain.UpdateAccount(uc, account)
}

// AccountsDeactivateAccount is the resolver for the accounts_deactivateAccount field.
func (r *mutationResolver) AccountsDeactivateAccount(ctx context.Context, accountName string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.domain.DeactivateAccount(uc, accountName)
}

// AccountsActivateAccount is the resolver for the accounts_activateAccount field.
func (r *mutationResolver) AccountsActivateAccount(ctx context.Context, accountName string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.domain.ActivateAccount(uc, accountName)
}

// AccountsDeleteAccount is the resolver for the accounts_deleteAccount field.
func (r *mutationResolver) AccountsDeleteAccount(ctx context.Context, accountName string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.domain.DeleteAccount(uc, accountName)
}

// AccountsInviteMember is the resolver for the accounts_inviteMember field.
func (r *mutationResolver) AccountsInviteMembers(ctx context.Context, accountName string, invitations []*entities.Invitation) ([]*entities.Invitation, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.domain.InviteMembers(uc, accountName, invitations)
}

// AccountsResendInviteMail is the resolver for the accounts_resendInviteMail field.
func (r *mutationResolver) AccountsResendInviteMail(ctx context.Context, accountName string, invitationID string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.domain.ResendInviteEmail(uc, accountName, repos.ID(invitationID))
}

// AccountsDeleteInvitation is the resolver for the accounts_deleteInvitation field.
func (r *mutationResolver) AccountsDeleteInvitation(ctx context.Context, accountName string, invitationID string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.domain.DeleteInvitation(uc, accountName, repos.ID(invitationID))
}

// AccountsAcceptInvitation is the resolver for the accounts_acceptInvitation field.
func (r *mutationResolver) AccountsAcceptInvitation(ctx context.Context, accountName string, inviteToken string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.domain.AcceptInvitation(uc, accountName, inviteToken)
}

// AccountsRejectInvitation is the resolver for the accounts_rejectInvitation field.
func (r *mutationResolver) AccountsRejectInvitation(ctx context.Context, accountName string, inviteToken string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	return r.domain.RejectInvitation(uc, accountName, inviteToken)
}

// AccountsRemoveAccountMembership is the resolver for the accounts_removeAccountMembership field.
func (r *mutationResolver) AccountsRemoveAccountMembership(ctx context.Context, accountName string, memberID repos.ID) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.domain.RemoveAccountMembership(uc, accountName, memberID)
}

// AccountsUpdateAccountMembership is the resolver for the accounts_updateAccountMembership field.
func (r *mutationResolver) AccountsUpdateAccountMembership(ctx context.Context, accountName string, memberID repos.ID, role iamT.Role) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.domain.UpdateAccountMembership(uc, accountName, memberID, iamT.Role(role))
}

// AccountsListAccounts is the resolver for the accounts_listAccounts field.
func (r *queryResolver) AccountsListAccounts(ctx context.Context) ([]*entities.Account, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	a, err := r.domain.ListAccounts(uc)
	if err != nil {
		return nil, errors.NewE(err)
	}
	if a == nil {
		return []*entities.Account{}, nil
	}
	return a, nil
}

// AccountsGetAccount is the resolver for the accounts_getAccount field.
func (r *queryResolver) AccountsGetAccount(ctx context.Context, accountName string) (*entities.Account, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.domain.GetAccount(uc, accountName)
}

// AccountsResyncAccount is the resolver for the accounts_resyncAccount field.
func (r *queryResolver) AccountsResyncAccount(ctx context.Context, accountName string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.domain.ResyncAccount(uc, accountName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// AccountsListInvitations is the resolver for the accounts_listInvitations field.
func (r *queryResolver) AccountsListInvitations(ctx context.Context, accountName string) ([]*entities.Invitation, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	invs, err := r.domain.ListInvitations(uc, accountName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if invs == nil {
		return []*entities.Invitation{}, nil
	}
	return invs, nil
}

// AccountsGetInvitation is the resolver for the accounts_getInvitation field.
func (r *queryResolver) AccountsGetInvitation(ctx context.Context, accountName string, invitationID string) (*entities.Invitation, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.domain.GetInvitation(uc, accountName, repos.ID(invitationID))
}

// AccountsListInvitationsForUser is the resolver for the accounts_listInvitationsForUser field.
func (r *queryResolver) AccountsListInvitationsForUser(ctx context.Context, onlyPending bool) ([]*entities.Invitation, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	invitations, err := r.domain.ListInvitationsForUser(uc, onlyPending)
	if err != nil {
		return nil, errors.NewE(err)
	}
	if invitations == nil {
		return make([]*entities.Invitation, 0), nil
	}

	return invitations, nil
}

// AccountsCheckNameAvailability is the resolver for the accounts_checkNameAvailability field.
func (r *queryResolver) AccountsCheckNameAvailability(ctx context.Context, name string) (*domain.CheckNameAvailabilityOutput, error) {
	return r.domain.CheckNameAvailability(ctx, name)
}

// AccountsListMembershipsForUser is the resolver for the accounts_listMembershipsForUser field.
func (r *queryResolver) AccountsListMembershipsForUser(ctx context.Context) ([]*entities.AccountMembership, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.domain.ListMembershipsForUser(uc)
}

// AccountsListMembershipsForAccount is the resolver for the accounts_listMembershipsForAccount field.
func (r *queryResolver) AccountsListMembershipsForAccount(ctx context.Context, accountName string, role *iamT.Role) ([]*entities.AccountMembership, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.domain.ListMembershipsForAccount(uc, accountName, nil)
}

// AccountsGetAccountMembership is the resolver for the accounts_getAccountMembership field.
func (r *queryResolver) AccountsGetAccountMembership(ctx context.Context, accountName string) (*entities.AccountMembership, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.domain.GetAccountMembership(uc, accountName)
}

// AccountsEnsureKloudliteRegistryPullSecrets is the resolver for the accounts_ensureKloudliteRegistryPullSecrets field.
func (r *queryResolver) AccountsEnsureKloudliteRegistryPullSecrets(ctx context.Context, accountName string) (bool, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.domain.EnsureKloudliteRegistryCredentials(uc, accountName); err != nil {
		return false, err
	}
	return true, nil
}

// AccountsAvailableKloudliteRegions is the resolver for the accounts_availableKloudliteRegions field.
func (r *queryResolver) AccountsAvailableKloudliteRegions(ctx context.Context) ([]*domain.AvailableKloudliteRegion, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.domain.AvailableKloudliteRegions(uc)
}

// Accounts is the resolver for the accounts field.
func (r *userResolver) Accounts(ctx context.Context, obj *model.User) ([]*entities.AccountMembership, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.domain.ListMembershipsForUser(uc)
}

// AccountInvitations is the resolver for the accountInvitations field.
func (r *userResolver) AccountInvitations(ctx context.Context, obj *model.User, onlyPending bool) ([]*entities.Invitation, error) {
	uc, err := toUserContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.domain.ListInvitationsForUser(uc, onlyPending)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
