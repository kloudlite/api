package common

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/apps/infra/internal/env"
	"github.com/kloudlite/api/pkg/k8s"
	"github.com/kloudlite/api/pkg/logging"
)

type Domain struct {
	Logger logging.Logger
	Env    *env.Env

	K8sClient k8s.Client

	IAMSvc                 ports.IAMSvc
	AccountsSvc            ports.AccountsSvc
	ResDispatcher          ports.ResourceDispatcher
	ResourceEventPublisher ports.ResourceEventPublisher
}
