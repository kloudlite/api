package domain

import (
	"kloudlite.io/apps/nodectrl/internal/domain/entities"
	"kloudlite.io/apps/nodectrl/internal/env"
	mongogridfs "kloudlite.io/pkg/mongo-gridfs"
	"kloudlite.io/pkg/repos"
)

type Domain interface {
	GetEnv() *env.Env

	GetGRidFs() mongogridfs.GridFs
	GetTokenRepo() repos.DbRepo[*entities.Token]
}
