package app

import (
	"encoding/json"
	"fmt"

	"github.com/kloudlite/api/apps/iot-console/internal/domain"
	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	t "github.com/kloudlite/api/apps/tenant-agent/types"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/messaging"
	msgTypes "github.com/kloudlite/api/pkg/messaging/types"
	"github.com/kloudlite/operator/pkg/constants"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SendTargetClusterMessagesProducer messaging.Producer

type resourceDispatcherImpl struct {
	producer messaging.Producer
}

func NewResourceDispatcher(producer SendTargetClusterMessagesProducer) domain.ResourceDispatcher {
	return &resourceDispatcherImpl{
		producer,
	}
}

func (a *resourceDispatcherImpl) ApplyToTargetDevice(ctx domain.IotConsoleContext, deviceName string, obj client.Object, recordVersion int) error {
	clusterName := entities.GetClusterName(deviceName)

	ann := obj.GetAnnotations()
	if ann == nil {
		ann = make(map[string]string, 1)
	}
	ann[constants.RecordVersionKey] = fmt.Sprintf("%d", recordVersion)
	obj.SetAnnotations(ann)

	m, err := fn.K8sObjToMap(obj)
	if err != nil {
		return errors.NewE(err)
	}

	b, err := json.Marshal(t.AgentMessage{
		AccountName: ctx.AccountName,
		ClusterName: clusterName,
		Action:      t.ActionApply,
		Object:      m,
	})
	if err != nil {
		return errors.NewE(err)
	}

	err = a.producer.Produce(ctx, msgTypes.ProduceMsg{
		Subject: common.GetTenantClusterMessagingTopic(ctx.AccountName, clusterName),
		Payload: b,
	})

	return errors.NewE(err)
}

func (d *resourceDispatcherImpl) DeleteFromTargetDevice(ctx domain.IotConsoleContext, deviceName string, obj client.Object) error {
	clusterName := entities.GetClusterName(deviceName)
	m, err := fn.K8sObjToMap(obj)
	if err != nil {
		return errors.NewE(err)
	}

	b, err := json.Marshal(t.AgentMessage{
		AccountName: ctx.AccountName,
		ClusterName: clusterName,
		Action:      t.ActionDelete,
		Object:      m,
	})
	if err != nil {
		return errors.NewE(err)
	}

	err = d.producer.Produce(ctx, msgTypes.ProduceMsg{
		Subject: common.GetTenantClusterMessagingTopic(ctx.AccountName, clusterName),
		Payload: b,
	})

	return errors.NewE(err)
}
