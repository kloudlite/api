package app

import (
	"encoding/json"
	"fmt"
	"github.com/kloudlite/api/apps/infra/internal/domain"
	"github.com/kloudlite/api/common"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/messaging"
	msgTypes "github.com/kloudlite/api/pkg/messaging/types"
	t "github.com/kloudlite/operator/agent/types"
	"github.com/kloudlite/operator/pkg/constants"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SendTargetClusterMessagesProducer messaging.Producer

type resourceDispatcherImpl struct {
	producer messaging.Producer
}

func NewResourceDispatcher(producer SendTargetClusterMessagesProducer) domain.ResourceDispatcher{
	return &resourceDispatcherImpl{
		producer,
	}
}

func (a *resourceDispatcherImpl) ApplyToTargetCluster(ctx domain.InfraContext, clusterName string, obj client.Object, recordVersion int) error {
	ann := obj.GetAnnotations()
	if ann == nil {
		ann = make(map[string]string, 1)
	}
	ann[constants.RecordVersionKey] = fmt.Sprintf("%d", recordVersion)
	obj.SetAnnotations(ann)

	m, err := fn.K8sObjToMap(obj)
	if err != nil {
		return err
	}

	b, err := json.Marshal(t.AgentMessage{
		AccountName: ctx.AccountName,
		ClusterName: clusterName,
		Action:      t.ActionApply,
		Object:      m,
	})
	if err != nil {
		return err
	}

	err = a.producer.Produce(ctx, msgTypes.ProduceMsg{
		Subject: common.GetTenantClusterMessagingTopic(ctx.AccountName, clusterName),
		Payload: b,
	})

	return err
}


func (d *resourceDispatcherImpl) DeleteFromTargetCluster(ctx domain.InfraContext, clusterName string, obj client.Object) error {
	m, err := fn.K8sObjToMap(obj)
	if err != nil {
		return err
	}

	b, err := json.Marshal(t.AgentMessage{
		AccountName: ctx.AccountName,
		ClusterName: clusterName,
		Action:      t.ActionDelete,
		Object:      m,
	})
	if err != nil {
		return err
	}

	err = d.producer.Produce(ctx, msgTypes.ProduceMsg{
		Subject: common.GetTenantClusterMessagingTopic(ctx.AccountName, clusterName),
		Payload: b,
	})

	return err
}