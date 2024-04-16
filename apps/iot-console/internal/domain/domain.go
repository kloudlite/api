package domain

import (
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	"github.com/kloudlite/api/apps/iot-console/internal/env"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	message_office_internal "github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/message-office-internal"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/k8s"
	"github.com/kloudlite/api/pkg/kv"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/repos"
	"go.uber.org/fx"
)

type domain struct {
	k8sClient k8s.Client
	logger    logging.Logger

	messageOfficeInternalClient message_office_internal.MessageOfficeInternalClient
	iamClient                   iam.IAMClient

	iotProjectRepo         repos.DbRepo[*entities.IOTProject]
	iotDeploymentRepo      repos.DbRepo[*entities.IOTDeployment]
	iotDeviceRepo          repos.DbRepo[*entities.IOTDevice]
	iotDeviceBlueprintRepo repos.DbRepo[*entities.IOTDeviceBlueprint]
	iotAppRepo             repos.DbRepo[*entities.IOTApp]

	envVars *env.Env
}

func (d *domain) canPerformActionInAccount(ctx IotConsoleContext, action iamT.Action) error {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(action),
	})
	if err != nil {
		return errors.NewE(err)
	}
	if !co.Status {
		return errors.Newf("unauthorized to perform action %q in account %q", action, ctx.AccountName)
	}
	return nil
}

type IOTConsoleCacheStore kv.BinaryDataRepo

var Module = fx.Module("domain",
	fx.Provide(func(
		k8sClient k8s.Client,
		logger logging.Logger,

		messageOfficeInternalClient message_office_internal.MessageOfficeInternalClient,
		iamClient iam.IAMClient,

		iotProjectRepo repos.DbRepo[*entities.IOTProject],
		iotDeploymentRepo repos.DbRepo[*entities.IOTDeployment],
		iotDeviceRepo repos.DbRepo[*entities.IOTDevice],
		iotDeviceBlueprintRepo repos.DbRepo[*entities.IOTDeviceBlueprint],
		iotAppRepo repos.DbRepo[*entities.IOTApp],

		ev *env.Env,
	) Domain {
		return &domain{
			k8sClient:                   k8sClient,
			logger:                      logger,
			iotProjectRepo:              iotProjectRepo,
			iotDeploymentRepo:           iotDeploymentRepo,
			iotDeviceRepo:               iotDeviceRepo,
			iotDeviceBlueprintRepo:      iotDeviceBlueprintRepo,
			iotAppRepo:                  iotAppRepo,
			envVars:                     ev,
			messageOfficeInternalClient: messageOfficeInternalClient,
			iamClient:                   iamClient,
		}
	}),
)
