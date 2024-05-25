package volume_attachment

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/common"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type Domain struct {
	*common.Domain
	VolumeAttachmentRepo repos.DbRepo[*entities.VolumeAttachment]
}
