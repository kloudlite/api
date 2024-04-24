package app

import (
	"context"
	"encoding/json"

	"github.com/kloudlite/api/apps/iot-console/internal/domain"
	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	t "github.com/kloudlite/api/apps/tenant-agent/types"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/messaging"
	"github.com/kloudlite/api/pkg/messaging/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ErrorOnApplyConsumer messaging.Consumer

func processErrorOnApply(consumer ErrorOnApplyConsumer, logger logging.Logger, d domain.Domain) {
	counter := 0
	processMsg := func(msg *types.ConsumeMsg) error {
		counter += 1

		var errObj t.AgentErrMessage
		if err := json.Unmarshal(msg.Payload, &errObj); err != nil {
			return errors.NewE(err)
		}

		obj := unstructured.Unstructured{Object: errObj.Object}

		mLogger := logger.WithKV(
			"gvk", obj.GroupVersionKind(),
			"accountName", errObj.AccountName,
			"clusterName", errObj.ClusterName,
		)

		mLogger.Infof("[%d] received message", counter)
		defer func() {
			mLogger.Infof("[%d] processed message", counter)
		}()

		dctx := domain.IotConsoleContext{Context: context.TODO(), UserId: "sys-user:error-on-apply-worker", UserEmail: "", UserName: "", AccountName: errObj.AccountName}

		opts := domain.UpdateAndDeleteOpts{MessageTimestamp: msg.Timestamp}

		gvkstr := obj.GroupVersionKind().String()
		switch gvkstr {
		case blueprintGVK.String():
			{
				cc, err := fn.JsonConvert[entities.IOTDeviceBlueprint](obj.Object)
				if err != nil {
					return err
				}

				if errObj.Action == t.ActionApply {
					return d.OnBlueprintApplyError(dctx, errObj.ClusterName, obj.GetName(), errObj.Error, cc, opts)
				}
				return d.OnBlueprintDeleteMessage(dctx, errObj.ClusterName, cc)
			}
		// case helmreleaseGVK.String():
		// 	{
		// 		helmRelease, err := fn.JsonConvert[entities.HelmRelease](obj.Object)
		// 		if err != nil {
		// 			return err
		// 		}
		//
		// 		if errObj.Action == t.ActionApply {
		// 			return d.OnHelmReleaseApplyError(dctx, errObj.ClusterName, obj.GetName(), errObj.Error, opts)
		// 		}
		// 		return d.OnHelmReleaseDeleteMessage(dctx, errObj.ClusterName, helmRelease)
		// 	}
		default:
			{
				return errors.Newf("infra error-on-apply reader does not acknowledge resource with gvk (%s)", gvkstr)
			}
		}
	}

	if err := consumer.Consume(processMsg, types.ConsumeOpts{
		OnError: func(err error) error {
			return nil
		},
	}); err != nil {
		logger.Errorf(err, "when setting up error-on-apply consumer")
	}
}
