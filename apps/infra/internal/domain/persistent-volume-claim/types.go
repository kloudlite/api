package persistent_volume_claim

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	*common.Domain
	PersistentVolumeClaimRepo repos.DbRepo[*entities.PersistentVolumeClaim]
}
