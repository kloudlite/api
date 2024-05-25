package managed_svc_templates

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
)

type Domain struct {
	ManagedSvcTemplates    []*entities.MsvcTemplate
	ManagedSvcTemplatesMap map[string]map[string]*entities.MsvcTemplateEntry
}
