package app

import (
	"context"
	"encoding/json"

	"go.uber.org/fx"
	"kloudlite.io/apps/console.old/internal/domain"
	"kloudlite.io/pkg/redpanda"
)

type workloadMessengerImpl struct {
	producer redpanda.Producer
}

func (i *workloadMessengerImpl) SendAction(action string, kafkaTopic string, resId string, res any) error {
	marshal, err := json.Marshal(
		map[string]any{
			"action":  action,
			"payload": res,
		},
	)
	if err != nil {
		return err
	}
	if _, err := i.producer.Produce(context.TODO(), kafkaTopic, resId, marshal); err != nil {
		return err
	}
	return nil
}

func fxWorkloadMessenger() fx.Option {
	return fx.Provide(
		func(p redpanda.Producer) domain.WorkloadMessenger {
			return &workloadMessengerImpl{
				producer: p,
			}
		},
	)
}