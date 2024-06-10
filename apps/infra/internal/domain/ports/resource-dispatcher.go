package ports

import (
	"context"

	client "sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceDispatcher interface {
	ApplyToTargetCluster(ctx context.Context, accountName, clusterName string, obj client.Object, recordVersion int) error
	DeleteFromTargetCluster(ctx context.Context, accountName, clusterName string, obj client.Object) error
}
