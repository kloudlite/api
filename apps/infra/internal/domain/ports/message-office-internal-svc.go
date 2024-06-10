package ports

import (
	"context"
)

type MessageOfficeInternalSvc interface {
	GenerateClusterToken(ctx context.Context, accountName string, clusterName string) (string, error)
	GetClusterToken(ctx context.Context, accountName string, clusterName string) (string, error)
}
