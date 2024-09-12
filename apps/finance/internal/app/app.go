package app

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofiber/fiber/v2"
	"github.com/kloudlite/api/apps/finance/internal/app/graph"
	"github.com/kloudlite/api/apps/finance/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/finance/internal/domain"
	"github.com/kloudlite/api/apps/finance/internal/entities"
	"github.com/kloudlite/api/apps/finance/internal/env"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/constants"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/auth"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/comms"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/grpc"
	httpServer "github.com/kloudlite/api/pkg/http-server"
	"github.com/kloudlite/api/pkg/kv"
	"github.com/kloudlite/api/pkg/repos"
	"go.uber.org/fx"
)

type AuthCacheClient kv.Client
type AuthClient grpc.Client

type (
	CommsClient grpc.Client
	IAMClient   grpc.Client
)

var Module = fx.Module("app",
	repos.NewFxMongoRepo[*entities.Payment]("payments", "pmt", entities.PaymentIndices),
	repos.NewFxMongoRepo[*entities.Invoice]("invoices", "inv", entities.InvoiceIndices),
	repos.NewFxMongoRepo[*entities.Wallet]("wallets", "wlt", entities.InvoiceIndices),
	repos.NewFxMongoRepo[*entities.Charge]("charges", "chrg", entities.InvoiceIndices),
	repos.NewFxMongoRepo[*entities.Subscription]("subscriptions", "sbs", entities.InvoiceIndices),

	// fx.Provide(func(client AuthCacheClient) kv.Repo[*entities.Invitation] {
	// 	return kv.NewRepo[*entities.Invitation](client)
	// }),

	fx.Provide(func(conn IAMClient) iam.IAMClient {
		return iam.NewIAMClient(conn)
	}),

	fx.Provide(func(conn CommsClient) comms.CommsClient {
		return comms.NewCommsClient(conn)
	}),

	fx.Provide(func(conn AuthClient) auth.AuthClient {
		return auth.NewAuthClient(conn)
	}),

	domain.Module,

	fx.Invoke(
		func(server httpServer.Server, d domain.Domain, env *env.Env, sessionRepo kv.Repo[*common.AuthSession]) {
			gqlConfig := generated.Config{Resolvers: graph.NewResolver(d)}

			gqlConfig.Directives.IsLoggedInAndVerified = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
				sess := httpServer.GetSession[*common.AuthSession](ctx)
				if sess == nil {
					return nil, fiber.ErrUnauthorized
				}

				if !sess.UserVerified {
					return nil, fiber.ErrForbidden
				}

				return next(context.WithValue(ctx, "kloudlite-user-session", *sess))
			}

			gqlConfig.Directives.HasAccount = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
				sess := httpServer.GetSession[*common.AuthSession](ctx)
				if sess == nil {
					return nil, fiber.ErrUnauthorized
				}
				m := httpServer.GetHttpCookies(ctx)
				klAccount := m[env.AccountCookieName]
				if klAccount == "" {
					return nil, errors.Newf("no cookie named %q present in request", env.AccountCookieName)
				}

				nctx := context.WithValue(ctx, "user-session", sess)
				nctx = context.WithValue(nctx, "account-name", klAccount)
				return next(nctx)
			}

			schema := generated.NewExecutableSchema(gqlConfig)
			server.SetupGraphqlServer(schema,
				httpServer.NewSessionMiddleware(
					sessionRepo,
					constants.CookieName,
					env.CookieDomain,
					constants.CacheSessionPrefix,
				),
			)
		},
	),
)
