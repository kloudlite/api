package app

import (
	"context"
	"encoding/json"

	t "github.com/kloudlite/api/apps/tenant-agent/types"
	"github.com/kloudlite/api/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/kloudlite/api/apps/console/internal/domain"
	"github.com/kloudlite/api/apps/console/internal/entities"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/messaging"
	msgTypes "github.com/kloudlite/api/pkg/messaging/types"
)

type ErrorOnApplyConsumer messaging.Consumer

func ProcessErrorOnApply(consumer ErrorOnApplyConsumer, d domain.Domain, logger logging.Logger) {
	counter := 0

	getResourceContext := func(ctx domain.ConsoleContext, obj unstructured.Unstructured) (domain.ResourceContext, error) {
		mapping, err := d.GetResourceMapping(ctx, entities.ResourceTypeApp, obj.GetNamespace(), obj.GetName())
		if err != nil {
			return domain.ResourceContext{}, err
		}

		return newResourceContext(ctx, mapping.ProjectName, mapping.EnvironmentName), nil
	}

	msgReader := func(msg *msgTypes.ConsumeMsg) error {
		counter += 1
		logger.Debugf("received message [%d]", counter)
		var errMsg t.AgentErrMessage
		if err := json.Unmarshal(msg.Payload, &errMsg); err != nil {
			return errors.NewE(err)
		}

		obj := unstructured.Unstructured{Object: errMsg.Object}

		mLogger := logger.WithKV(
			"gvk", obj.GroupVersionKind(),
			"accountName", errMsg.AccountName,
			"clusterName", errMsg.ClusterName,
		)

		mLogger.Infof("received message")
		defer func() {
			mLogger.Infof("processed message")
		}()

		kind := obj.GroupVersionKind().Kind
		dctx := domain.NewConsoleContext(context.TODO(), "sys-user:apply-on-error-worker", errMsg.AccountName)

		opts := domain.UpdateAndDeleteOpts{MessageTimestamp: msg.Timestamp}

		switch kind {
		case "Project":
			{
				return d.OnProjectApplyError(dctx, errMsg.Error, obj.GetName(), opts)
			}
		case "Environment":
			{
				return d.OnEnvironmentApplyError(dctx, errMsg.Error, obj.GetNamespace(), obj.GetName(), opts)
			}
		case "App":
			{
				rctx, err := getResourceContext(dctx, obj)
				if err != nil {
					return errors.NewE(err)
				}

				return d.OnAppApplyError(rctx, errMsg.Error, obj.GetName(), opts)
			}
		case "Config":
			{
				rctx, err := getResourceContext(dctx, obj)
				if err != nil {
					return errors.NewE(err)
				}

				return d.OnConfigApplyError(rctx, errMsg.Error, obj.GetName(), opts)
			}
		case "Secret":
			{
				rctx, err := getResourceContext(dctx, obj)
				if err != nil {
					return errors.NewE(err)
				}

				return d.OnSecretApplyError(rctx, errMsg.Error, obj.GetName(), opts)
			}
		case "Router":
			{
				rctx, err := getResourceContext(dctx, obj)
				if err != nil {
					return errors.NewE(err)
				}

				return d.OnRouterApplyError(rctx, errMsg.Error, obj.GetName(), opts)
			}
		case "ManagedResource":
			{
				rctx, err := getResourceContext(dctx, obj)
				if err != nil {
					return errors.NewE(err)
				}

				return d.OnManagedResourceApplyError(rctx, errMsg.Error, obj.GetName(), opts)
			}
		default:
			{
				return errors.Newf("console apply error reader does not acknowledge resource with kind (%s)", kind)
			}
		}
	}

	if err := consumer.Consume(msgReader, msgTypes.ConsumeOpts{
		OnError: func(err error) error {
			logger.Errorf(err, "received while reading messages, ignoring it")
			return nil
		},
	}); err != nil {
		logger.Errorf(err, "error while consuming messages")
	}
}
