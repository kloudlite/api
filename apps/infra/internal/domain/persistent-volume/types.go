package persistent_volume

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	*common.Domain
	PersistentVolumeRepo repos.DbRepo[*entities.PersistentVolume]
}
