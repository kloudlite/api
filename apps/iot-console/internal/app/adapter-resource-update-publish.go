package app

import (
	"fmt"

	"github.com/kloudlite/api/apps/iot-console/internal/domain"
	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/nats"
)

type ResourceEventPublisherImpl struct {
	cli    *nats.Client
	logger logging.Logger
}

func (r *ResourceEventPublisherImpl) PublishInfraEvent(ctx domain.IotConsoleContext, resourceType domain.ResourceType, resName string, update domain.PublishMsg) {
	subject := fmt.Sprintf(
		"res-updates.account.%s.%s.%s",
		ctx.AccountName, resourceType, resName,
	)

	r.publish(subject, update)
}

func (r *ResourceEventPublisherImpl) PublishResourceEvent(ctx domain.IotConsoleContext, deviceName string, resourceType domain.ResourceType, resName string, update domain.PublishMsg) {
	clusterName := entities.GetClusterName(deviceName)

	subject := fmt.Sprintf(
		"res-updates.account.%s.cluster.%s.%s.%s",
		ctx.AccountName, clusterName, resourceType, resName,
	)

	r.publish(subject, update)
}

func (r *ResourceEventPublisherImpl) publish(subject string, msg domain.PublishMsg) {
	if err := r.cli.Conn.Publish(subject, []byte(msg)); err != nil {
		r.logger.Errorf(err, "failed to publish message to subject %q", subject)
	}
}

func NewResourceEventPublisher(cli *nats.Client, logger logging.Logger) domain.ResourceEventPublisher {
	return &ResourceEventPublisherImpl{
		cli,
		logger,
	}
}
