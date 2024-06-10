package adapters

import (
	"fmt"

	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/nats"
)

type ResourceEventPublisherImpl struct {
	cli    *nats.Client
	logger logging.Logger
}

// PublishInfraEvent implements common.ResourceEventPublisher.
func (r *ResourceEventPublisherImpl) PublishInfraEvent(ctx types.InfraContext, resourceType ports.ResourceType, resName string, update ports.PublishMsg) {
	subject := fmt.Sprintf(
		"res-updates.account.%s.%s.%s",
		ctx.AccountName, resourceType, resName,
	)

	r.publish(subject, update)
}

// PublishResourceEvent implements common.ResourceEventPublisher.
func (r *ResourceEventPublisherImpl) PublishResourceEvent(ctx types.InfraContext, clusterName string, resourceType ports.ResourceType, resName string, update ports.PublishMsg) {
	subject := fmt.Sprintf(
		"res-updates.account.%s.cluster.%s.%s.%s",
		ctx.AccountName, clusterName, resourceType, resName,
	)

	r.publish(subject, update)
}

func (r *ResourceEventPublisherImpl) publish(subject string, msg ports.PublishMsg) {
	if err := r.cli.Conn.Publish(subject, []byte(msg)); err != nil {
		r.logger.Errorf(err, "failed to publish message to subject %q", subject)
	}
}

func NewResourceEventPublisher(cli *nats.Client, logger logging.Logger) ports.ResourceEventPublisher {
	return &ResourceEventPublisherImpl{
		cli,
		logger,
	}
}
