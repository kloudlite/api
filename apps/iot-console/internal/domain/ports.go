package domain

import (
	"context"

	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/accounts"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AccountsSvc interface {
	GetAccount(ctx context.Context, userId string, accountName string) (*accounts.GetAccountOut, error)
}

type PublishMsg string

type ResourceDispatcher interface {
	ApplyToTargetDevice(ctx IotConsoleContext, clusterName string, obj client.Object, recordVersion int) error
	DeleteFromTargetDevice(ctx IotConsoleContext, clusterName string, obj client.Object) error
}

type ResourceEventPublisher interface {
	PublishInfraEvent(ctx IotConsoleContext, resourceType ResourceType, resName string, update PublishMsg)
	PublishResourceEvent(ctx IotConsoleContext, clusterName string, resourceType ResourceType, resName string, update PublishMsg)
}
