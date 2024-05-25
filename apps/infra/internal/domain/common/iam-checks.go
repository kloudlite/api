package common

import (
	"fmt"

	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
)

func (d *Domain) CanPerformActionInAccount(ctx domainT.InfraContext, action iamT.Action) error {
	co, err := d.IAMSvc.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(action),
	})
	if err != nil {
		return types.ErrGRPCCall{Err: err}
	}
	if !co.Status {
		return types.ErrIAMUnauthorized{
			UserId:   string(ctx.UserId),
			Resource: fmt.Sprintf("account: %s", ctx.AccountName),
			Action:   string(action),
		}
	}
	return nil
}

func (d *Domain) CanPerformActionInDevice(ctx domainT.InfraContext, action iamT.Action, devName string) error {
	co, err := d.IAMSvc.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceInfraVPNDevice, devName),
		},
		Action: string(action),
	})
	if err != nil {
		return types.ErrGRPCCall{Err: err}
	}
	if !co.Status {
		return types.ErrIAMUnauthorized{
			UserId:   string(ctx.UserId),
			Resource: fmt.Sprintf("device: %s", devName),
			Action:   string(action),
		}
	}
	return nil
}
