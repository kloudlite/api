package domain

import (
	"fmt"
	nanoid "github.com/matoous/go-nanoid/v2"
	"kloudlite.io/apps/accounts/internal/entities"
	iamT "kloudlite.io/apps/iam/types"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/comms"
	"kloudlite.io/pkg/errors"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/repos"
)

func (d *domain) findInvitation(ctx AccountsContext, invitationId repos.ID) (*entities.Invitation, error) {
	inv, err := d.invitationRepo.FindOne(ctx, repos.Filter{
		"accountName": ctx.AccountName,
		"id":          invitationId,
	})
	if err != nil {
		return nil, err
	}

	if inv == nil {
		return nil, fmt.Errorf("no invitation found with id=%s", invitationId)
	}

	return inv, nil
}

func (d *domain) InviteUser(ctx AccountsContext, invitation entities.Invitation) (*entities.Invitation, error) {
	if err := d.checkAccountAccess(ctx, ctx.AccountName, ctx.UserId, iamT.InviteAccountMember); err != nil {
		return nil, err
	}

	_, err := d.findAccount(ctx, ctx.AccountName)
	if err != nil {
		return nil, err
	}

	invitation.InviteToken, err = nanoid.New(64)
	if err != nil {
		return nil, errors.NewEf(err, "failed to generate invite token")
	}

	inv, err := d.invitationRepo.Create(ctx, &invitation)
	if err != nil {
		return nil, errors.NewEf(err, "failed to create invitation")
	}

	if _, err := d.commsClient.SendAccountMemberInviteEmail(ctx, &comms.AccountMemberInviteEmailInput{
		AccountName:     ctx.AccountName,
		InvitationToken: inv.InviteToken,
		Email:           inv.UserEmail,
		Name:            inv.UserName,
	}); err != nil {
		return nil, err
	}

	return inv, nil
}

func (d *domain) ResendInviteEmail(ctx AccountsContext, invitationId repos.ID) (bool, error) {
	inv, err := d.findInvitation(ctx, invitationId)
	if err != nil {
		return false, err
	}

	action := iamT.InviteAccountMember
	if inv.UserRole == iamT.RoleAccountAdmin {
		action = iamT.InviteAccountAdmin
	}

	if err := d.checkAccountAccess(ctx, ctx.AccountName, ctx.UserId, action); err != nil {
		return false, err
	}

	return true, nil
}

func (d *domain) ListInvitations(ctx AccountsContext) ([]*entities.Invitation, error) {
	if err := d.checkAccountAccess(ctx, ctx.AccountName, ctx.UserId, iamT.ListAccountInvitations); err != nil {
		return nil, err
	}

	return d.invitationRepo.Find(ctx, repos.Query{Filter: repos.Filter{"accountName": ctx.AccountName}})
}

func (d *domain) DeleteInvitation(ctx AccountsContext, invitationId repos.ID) (bool, error) {
	if err := d.checkAccountAccess(ctx, ctx.AccountName, ctx.UserId, iamT.DeleteAccountInvitation); err != nil {
		return false, err
	}

	inv, err := d.findInvitation(ctx, invitationId)
	if err != nil {
		return false, err
	}

	if err := d.invitationRepo.DeleteById(ctx, inv.Id); err != nil {
		return false, err
	}
	return true, nil
}

func (d *domain) AcceptInvitation(ctx AccountsContext, invitationId repos.ID) (bool, error) {
	inv, err := d.findInvitation(ctx, invitationId)
	if err != nil {
		return false, err
	}

	inv.Accepted = fn.New(true)
	if _, err := d.invitationRepo.UpdateById(ctx, inv.Id, inv); err != nil {
		return false, err
	}

	return true, nil
}

func (d *domain) RejectInvitation(ctx AccountsContext, invitationId repos.ID) (bool, error) {
	inv, err := d.findInvitation(ctx, invitationId)
	if err != nil {
		return false, err
	}

	inv.Rejected = fn.New(true)
	if _, err := d.invitationRepo.UpdateById(ctx, inv.Id, inv); err != nil {
		return false, err
	}

	return true, nil
}
