package domain

import (
	"github.com/kloudlite/api/apps/console/internal/entities"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
)

func (d *domain) addEnvMembership(ctx ConsoleContext, envName string, userId repos.ID, role iamT.Role) error {
	if err := d.checkEnvironmentAccess(ResourceContext{
		ConsoleContext:  ctx,
		EnvironmentName: envName,
	}, iamT.CreateEnvironmentMembership); err != nil {
		return errors.NewE(err)
	}

	if _, err := d.iamClient.AddMembership(ctx, &iam.AddMembershipIn{
		UserId:       string(userId),
		ResourceType: string(iamT.ResourceAccount),
		ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceEnvironment, envName),
		Role:         string(role),
	}); err != nil {
		return errors.NewE(err)
	}

	return nil
}

func (d *domain) RemoveEnvMembership(ctx ConsoleContext, envName string, memberId repos.ID) (bool, error) {
	if err := d.checkEnvironmentAccess(ResourceContext{
		ConsoleContext:  ctx,
		EnvironmentName: envName,
	}, iamT.RemoveEnvironmentMembership); err != nil {
		return false, errors.NewE(err)
	}

	out, err := d.iamClient.RemoveMembership(ctx, &iam.RemoveMembershipIn{
		UserId:      string(memberId),
		ResourceRef: iamT.NewResourceRef(ctx.AccountName, iamT.ResourceEnvironment, envName),
	})
	if err != nil {
		return false, errors.NewE(err)
	}

	return out.Result, nil
}

func (d *domain) UpdateEnvMembership(ctx ConsoleContext, envName string, memberId repos.ID, role iamT.Role) (bool, error) {
	if err := d.checkEnvironmentAccess(ResourceContext{
		ConsoleContext:  ctx,
		EnvironmentName: envName,
	}, iamT.UpdateEnvironmentMembership); err != nil {
		return false, errors.NewE(err)
	}

	out, err := d.iamClient.UpdateMembership(ctx, &iam.UpdateMembershipIn{
		UserId:       string(memberId),
		ResourceType: string(iamT.ResourceAccount),
		ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, envName),
		Role:         string(role),
	})

	if err != nil {
		return false, errors.NewE(err)
	}

	return out.Result, nil
}

func (d *domain) GetEnvMembership(ctx ConsoleContext, envName string) (*entities.EnvMembership, error) {
	membership, err := d.iamClient.GetMembership(
		ctx, &iam.GetMembershipIn{
			UserId:       string(ctx.UserId),
			ResourceType: string(iamT.ResourceAccount),
			ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, envName),
		},
	)

	if err != nil {
		return nil, errors.NewE(err)
	}
	return &entities.EnvMembership{
		EnvName: envName,
		UserId:  repos.ID(membership.UserId),
		Role:    iamT.Role(membership.Role),
	}, nil
}

// --- external
func (d *domain) ListMembershipsForEnv(ctx ConsoleContext, envName string) ([]*entities.EnvMembership, error) {
	if err := d.checkEnvironmentAccess(ResourceContext{
		ConsoleContext:  ctx,
		EnvironmentName: envName,
	}, iamT.ListEnvironments); err != nil {
		return nil, errors.NewE(err)
	}

	out, err := d.iamClient.ListMembershipsForResource(ctx, &iam.MembershipsForResourceIn{
		ResourceType: string(iamT.ResourceAccount),
		ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceEnvironment, envName),
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	memberships := make([]*entities.EnvMembership, len(out.RoleBindings))
	for i := range out.RoleBindings {
		memberships[i] = &entities.EnvMembership{
			EnvName: envName,
			UserId:  repos.ID(out.RoleBindings[i].UserId),
			Role:    iamT.Role(out.RoleBindings[i].Role),
		}
	}

	return memberships, nil
}

func (d *domain) AddCollaboratorToEnvironment(ctx ConsoleContext, envName string, userId repos.ID) (bool, error) {
	if err := d.checkEnvironmentAccess(ResourceContext{
		ConsoleContext:  ctx,
		EnvironmentName: envName,
	}, iamT.CreateEnvironmentMembership); err != nil {
		return false, errors.NewE(err)
	}

	out, err := d.iamClient.AddMembership(ctx, &iam.AddMembershipIn{
		UserId:       string(userId),
		ResourceType: string(iamT.ResourceAccount),
		ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceEnvironment, envName),
		Role:         string(iamT.RoleEnvironmentCollaborator),
	})

	if err != nil {
		return false, errors.NewE(err)
	}

	return out.Result, nil
}

func (d *domain) RemoveCollaboratorFromEnvironment(ctx ConsoleContext, envName string, userId repos.ID) (bool, error) {
	if err := d.checkEnvironmentAccess(ResourceContext{
		ConsoleContext:  ctx,
		EnvironmentName: envName,
	}, iamT.RemoveEnvironmentMembership); err != nil {
		return false, errors.NewE(err)
	}

	out, err := d.iamClient.RemoveMembership(ctx, &iam.RemoveMembershipIn{
		UserId:      string(userId),
		ResourceRef: iamT.NewResourceRef(ctx.AccountName, iamT.ResourceEnvironment, envName),
	})

	if err != nil {
		return false, errors.NewE(err)
	}

	return out.Result, nil
}
