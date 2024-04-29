package app

import (
	"context"

	"github.com/kloudlite/api/apps/comms/internal/domain"
	"github.com/kloudlite/api/pkg/logging"
)

func processNotification(ctx context.Context, d domain.Domain, consumer NotificationConsumer, logr logging.Logger) error {

	panic("not implemented")
}
