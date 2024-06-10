package ports

import (
	"context"

	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
)

type IAMSvc interface {
	Can(ctx context.Context, in *iam.CanIn) (*iam.CanOut, error)
}
