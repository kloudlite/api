package domain

import (
	"go.uber.org/fx"

	"kloudlite.io/apps/nodectrl/internal/domain/entities"
	"kloudlite.io/apps/nodectrl/internal/env"
	mongogridfs "kloudlite.io/pkg/mongo-gridfs"
	"kloudlite.io/pkg/repos"
)

type domain struct {
	env       *env.Env
	gfs       mongogridfs.GridFs
	tokenRepo repos.DbRepo[*entities.Token]
}

var Module = fx.Module("domain",
	fx.Provide(
		func(env *env.Env, gfs mongogridfs.GridFs, tokenRepo repos.DbRepo[*entities.Token]) Domain {
			return domain{
				env:       env,
				gfs:       gfs,
				tokenRepo: tokenRepo,
			}
		},
	),
	ProviderClientFx,
)
