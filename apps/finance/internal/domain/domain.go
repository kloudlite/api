package domain

import (
	"context"

	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/kloudlite/api/apps/finance/internal/env"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/accounts"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/auth"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/comms"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/repos"
	"go.uber.org/fx"
)

type PaymentService interface {
	CreatePayment(ctx UserContext, req *entities.Payment) (*entities.Payment, error)

	SyncPaymentStatus(ctx UserContext, paymentID repos.ID) (*entities.Payment, error)

	GetWallet(ctx UserContext) (*entities.Wallet, error)
	ListPayments(ctx UserContext) ([]*entities.Payment, error)
	ListCharges(ctx UserContext) ([]*entities.Charge, error)

	CreateCharge(ctx context.Context, req *entities.Charge) (*entities.Charge, error)
}

type Domain interface {
	PaymentService
}

type domain struct {
	authClient  auth.AuthClient
	iamClient   iam.IAMClient
	commsClient comms.CommsClient
	accountsCli accounts.AccountsClient

	paymentRepo      repos.DbRepo[*entities.Payment]
	invoiceRepo      repos.DbRepo[*entities.Invoice]
	walletRepo       repos.DbRepo[*entities.Wallet]
	chargeRepo       repos.DbRepo[*entities.Charge]
	subscriptionRepo repos.DbRepo[*entities.Subscription]

	Env    *env.Env
	logger logging.Logger
}

func NewDomain(
	iamCli iam.IAMClient,
	authClient auth.AuthClient,
	commsClient comms.CommsClient,
	accountsCli accounts.AccountsClient,

	paymentRepo repos.DbRepo[*entities.Payment],
	invoiceRepo repos.DbRepo[*entities.Invoice],
	walletRepo repos.DbRepo[*entities.Wallet],
	chargeRepo repos.DbRepo[*entities.Charge],
	subscriptionRepo repos.DbRepo[*entities.Subscription],

	ev *env.Env,

	logger logging.Logger,
) Domain {
	return &domain{
		authClient:  authClient,
		iamClient:   iamCli,
		commsClient: commsClient,
		accountsCli: accountsCli,

		paymentRepo:      paymentRepo,
		invoiceRepo:      invoiceRepo,
		walletRepo:       walletRepo,
		chargeRepo:       chargeRepo,
		subscriptionRepo: subscriptionRepo,

		Env: ev,

		logger: logger,
	}
}

var Module = fx.Module("domain", fx.Provide(NewDomain))
