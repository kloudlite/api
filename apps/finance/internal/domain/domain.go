package domain

import (
	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/kloudlite/api/apps/finance/internal/env"
	"github.com/kloudlite/api/grpc-interfaces/container_registry"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/auth"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/comms"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/console"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/k8s"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/repos"
	"go.uber.org/fx"
)

type PaymentService interface {
}

type Domain interface {
	PaymentService
}

type domain struct {
	authClient              auth.AuthClient
	iamClient               iam.IAMClient
	consoleClient           console.ConsoleClient
	containerRegistryClient container_registry.ContainerRegistryClient
	commsClient             comms.CommsClient

	paymentRepo      repos.DbRepo[*entities.Payment]
	invoiceRepo      repos.DbRepo[*entities.Invoice]
	walletRepo       repos.DbRepo[*entities.Wallet]
	transactionRepo  repos.DbRepo[*entities.Transaction]
	subscriptionRepo repos.DbRepo[*entities.Subscription]

	// k8sClient k8s.Client

	Env *env.Env

	logger logging.Logger
}

func NewDomain(
	iamCli iam.IAMClient,
	consoleClient console.ConsoleClient,
	containerRegistryClient container_registry.ContainerRegistryClient,
	authClient auth.AuthClient,
	commsClient comms.CommsClient,

	k8sClient k8s.Client,

	paymentRepo repos.DbRepo[*entities.Payment],
	invoiceRepo repos.DbRepo[*entities.Invoice],
	walletRepo repos.DbRepo[*entities.Wallet],
	transactionRepo repos.DbRepo[*entities.Transaction],
	subscriptionRepo repos.DbRepo[*entities.Subscription],

	ev *env.Env,

	logger logging.Logger,
) Domain {
	return &domain{
		authClient:              authClient,
		iamClient:               iamCli,
		consoleClient:           consoleClient,
		commsClient:             commsClient,
		containerRegistryClient: containerRegistryClient,

		paymentRepo:      paymentRepo,
		invoiceRepo:      invoiceRepo,
		walletRepo:       walletRepo,
		transactionRepo:  transactionRepo,
		subscriptionRepo: subscriptionRepo,

		Env: ev,

		logger: logger,
	}
}

var Module = fx.Module("domain", fx.Provide(NewDomain))
