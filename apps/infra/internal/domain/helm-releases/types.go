package helm_releases

import (
	common "github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	*common.Domain

	HelmReleaseRepo repos.DbRepo[*entities.HelmRelease]
}
