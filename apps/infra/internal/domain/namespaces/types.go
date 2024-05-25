package namespace

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	NamespaceRepo repos.DbRepo[*entities.Namespace]
}
