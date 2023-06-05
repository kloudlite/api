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

func (d domain) GetEnv() *env.Env {
	return d.env
}

func (d domain) GetGRidFs() mongogridfs.GridFs {
	return d.gfs
}

func (d domain) GetTokenRepo() repos.DbRepo[*entities.Token] {
	return d.tokenRepo
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
