package adapters

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
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
	producer SendTargetClusterMessagesProducer
}

// ApplyToTargetCluster implements common.ResourceDispatcher.
func (rd *resourceDispatcherImpl) ApplyToTargetCluster(ctx context.Context, accountName string, clusterName string, obj client.Object, recordVersion int) error {
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
		AccountName: accountName,
		ClusterName: clusterName,
		Action:      t.ActionApply,
		Object:      m,
	})
	if err != nil {
		return errors.NewE(err)
	}

	err = rd.producer.Produce(ctx, msgTypes.ProduceMsg{
		Subject: common.GetTenantClusterMessagingTopic(accountName, clusterName),
		Payload: b,
	})

	return errors.NewE(err)
}

// DeleteFromTargetCluster implements common.ResourceDispatcher.
func (rd *resourceDispatcherImpl) DeleteFromTargetCluster(ctx context.Context, accountName string, clusterName string, obj client.Object) error {
	m, err := fn.K8sObjToMap(obj)
	if err != nil {
		return errors.NewE(err)
	}

	b, err := json.Marshal(t.AgentMessage{
		AccountName: accountName,
		ClusterName: clusterName,
		Action:      t.ActionDelete,
		Object:      m,
	})
	if err != nil {
		return errors.NewE(err)
	}

	err = rd.producer.Produce(ctx, msgTypes.ProduceMsg{
		Subject: common.GetTenantClusterMessagingTopic(accountName, clusterName),
		Payload: b,
	})

	return errors.NewE(err)
}

func NewResourceDispatcher(producer SendTargetClusterMessagesProducer) ports.ResourceDispatcher {
	return &resourceDispatcherImpl{
		producer,
	}
}
