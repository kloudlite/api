package adapters

import (
	"context"

	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
)

type iamSvc struct {
	IAMClient iam.IAMClient
}

// Can implements ports.IAMSvc.
func (i *iamSvc) Can(ctx context.Context, in *iam.CanIn) (*iam.CanOut, error) {
	return i.IAMClient.Can(ctx, in)
}

func NewIAMSvc(iamClient iam.IAMClient) ports.IAMSvc {
	return &iamSvc{IAMClient: iamClient}
}
