package app

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/kloudlite/api/pkg/errors"

	"github.com/kloudlite/api/apps/iot-console/internal/domain"
	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/kloudlite/api/pkg/messaging"
	msgTypes "github.com/kloudlite/api/pkg/messaging/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ReceiveResourceUpdatesConsumer messaging.Consumer

func gvk(obj client.Object) string {
	val := obj.GetObjectKind().GroupVersionKind().String()
	return val
}

var (
	// helmreleaseGVK = fn.GVK("crds.kloudlite.io/v1", "HelmChart")
	// pvcGVK              = fn.GVK("v1", "PersistentVolumeClaim")
	// pvGVK               = fn.GVK("v1", "PersistentVolume")
	// volumeAttachmentGVK = fn.GVK("storage.k8s.io/v1", "VolumeAttachment")
	// namespaceGVK = fn.GVK("v1", "Namespace")
	blueprintGVK = fn.GVK("crds.kloudlite.io/v1", "Blueprint")
)

func processResourceUpdates(consumer ReceiveResourceUpdatesConsumer, d domain.Domain, logger logging.Logger) {
	readMsg := func(msg *msgTypes.ConsumeMsg) error {
		logger.Debugf("processing msg timestamp %s", msg.Timestamp.Format(time.RFC3339))

		var su types.ResourceUpdate
		if err := json.Unmarshal(msg.Payload, &su); err != nil {
			logger.Errorf(err, "parsing into status update")
			return nil
		}

		if su.Object == nil {
			logger.Infof("message does not contain 'object', so won't be able to find a resource uniquely, thus ignoring ...")
			return nil
		}

		if len(strings.TrimSpace(su.AccountName)) == 0 {
			logger.Infof("message does not contain 'accountName', so won't be able to find a resource uniquely, thus ignoring ...")
			return nil
		}

		dctx := domain.IotConsoleContext{Context: context.TODO(), UserId: "sys-user-process-iot-console-updates", AccountName: su.AccountName}

		obj := unstructured.Unstructured{Object: su.Object}
		gvkStr := obj.GetObjectKind().GroupVersionKind().String()

		resStatus, err := func() (types.ResourceStatus, error) {
			v, ok := su.Object[types.ResourceStatusKey]
			if !ok {
				return "", errors.NewE(fmt.Errorf("field %s not found in object", types.ResourceStatusKey))
			}
			s, ok := v.(string)
			if !ok {
				return "", errors.NewE(fmt.Errorf("field value %v is not a string", v))
			}

			return types.ResourceStatus(s), nil
		}()
		if err != nil {
			return err
		}

		deviceName, err := entities.ExtractDeviceName(su.ClusterName)
		if err != nil {
			return errors.NewE(err)
		}

		mLogger := logger.WithKV(
			"gvk", obj.GetObjectKind().GroupVersionKind(),
			"NN", fmt.Sprintf("%s/%s", obj.GetNamespace(), obj.GetName()),
			"resource-status", resStatus,
			"accountName/clusterName", fmt.Sprintf("%s/%s", su.AccountName, su.ClusterName),
			"deviceName", deviceName,
		)

		mLogger.Infof("received message")
		defer func() {
			mLogger.Infof("processed message")
		}()

		switch gvkStr {
		case blueprintGVK.String():
			{
				var bp entities.IOTDeviceBlueprint
				if err := fn.JsonConversion(su.Object, &bp); err != nil {
					return errors.NewE(err)
				}

				if resStatus == types.ResourceStatusDeleted {
					return d.OnBlueprintDeleteMessage(dctx, su.ClusterName, bp)
				}

				return d.OnBlueprintUpdateMessage(dctx, su.ClusterName, bp, resStatus, domain.UpdateAndDeleteOpts{MessageTimestamp: msg.Timestamp})
			}
		default:
			{
				mLogger.Infof("iot-console status updates consumer does not acknowledge the gvk %s", gvk(&obj))
				return nil
			}
		}
	}

	if err := consumer.Consume(readMsg, msgTypes.ConsumeOpts{
		OnError: func(err error) error {
			logger.Errorf(err, "error while consuming message")
			return nil
		},
	}); err != nil {
		logger.Errorf(err, "error while consuming messages")
	}
}
