package app

import (
	"kloudlite.io/pkg/repos"
	"net/http"

	"go.uber.org/fx"
	"kloudlite.io/apps/finance/internal/app/graph"
	"kloudlite.io/apps/finance/internal/app/graph/generated"
	"kloudlite.io/apps/finance/internal/domain"
	"kloudlite.io/common"
	"kloudlite.io/pkg/cache"
	"kloudlite.io/pkg/config"
	httpServer "kloudlite.io/pkg/http-server"
)

type Env struct {
	CookieDomain string `env:"COOKIE_DOMAIN" required:"true"`
}

var Module = fx.Module(
	"application",
	fx.Provide(config.LoadEnv[Env]()),

	repos.NewFxMongoRepo[*domain.Account]("accounts", "acc", []string{}),
	fx.Invoke(func(
		server *http.ServeMux,
		d domain.Domain,
		env *Env,
		cacheClient cache.Client,
	) {
		schema := generated.NewExecutableSchema(
			generated.Config{Resolvers: graph.NewResolver(d)},
		)
		httpServer.SetupGQLServer(
			server,
			schema,
			cache.NewSessionRepo[*common.AuthSession](
				cacheClient,
				common.CookieName,
				env.CookieDomain,
				common.CacheSessionPrefix,
			),
		)
	}),
	domain.Module,
)
