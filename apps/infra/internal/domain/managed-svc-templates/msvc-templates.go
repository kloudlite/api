package managed_svc_templates

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
)

func (d *Domain) ListManagedSvcTemplates() ([]*entities.MsvcTemplate, error) {
	return d.ManagedSvcTemplates, nil
}

func (d *Domain) GetManagedSvcTemplate(category string, name string) (*entities.MsvcTemplateEntry, error) {
	return d.ManagedSvcTemplatesMap[category][name], nil
}
