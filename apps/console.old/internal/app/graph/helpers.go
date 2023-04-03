package graph

import (
	"kloudlite.io/apps/console.old/internal/app/graph/model"
	"kloudlite.io/apps/console.old/internal/domain/entities"
	"kloudlite.io/pkg/repos"
)

func projectModelFromEntity(projectEntity *entities.Project) *model.Project {
	return &model.Project{
		ID:          projectEntity.Id,
		Name:        projectEntity.Name,
		DisplayName: projectEntity.DisplayName,
		ReadableID:  projectEntity.ReadableId,
		Logo:        projectEntity.Logo,
		Description: projectEntity.Description,
		RegionID: func() repos.ID {
			if projectEntity.RegionId != nil {
				return *projectEntity.RegionId
			}
			return ""
		}(),
		Account: &model.Account{
			ID: projectEntity.AccountId,
		},
		Status: string(projectEntity.Status),
	}
}

func configModelFromEntity(configEntity *entities.Config) *model.Config {
	entries := make([]*model.CSEntry, 0)
	for _, e := range configEntity.Data {
		entries = append(entries, &model.CSEntry{
			Key:   e.Key,
			Value: e.Value,
		})
	}
	return &model.Config{
		ID:      configEntity.Id,
		Name:    configEntity.Name,
		Project: &model.Project{ID: configEntity.ProjectId},
		Entries: entries,
		Status:  string(configEntity.Status),
	}
}

func routerModelFromEntity(routerEntity *entities.Router) *model.Router {
	entries := make([]*model.Route, 0)
	for _, e := range routerEntity.Routes {
		entries = append(entries, &model.Route{
			Path:    e.Path,
			AppName: e.AppName,
			Port: func() *int {
				i := int(e.Port)
				return &i
			}(),
		})
	}
	d := routerEntity.Domains
	if d == nil {
		d = []string{}
	}
	return &model.Router{
		ID:      routerEntity.Id,
		Name:    routerEntity.Name,
		Project: &model.Project{ID: routerEntity.ProjectId},
		Domains: d,
		Routes:  entries,
		Status:  string(routerEntity.Status),
	}
}

func secretModelFromEntity(secretEntity *entities.Secret) *model.Secret {
	entries := make([]*model.CSEntry, 0)
	for _, e := range secretEntity.Data {
		entries = append(entries, &model.CSEntry{
			Key:   e.Key,
			Value: e.Value,
		})
	}
	return &model.Secret{
		ID:      secretEntity.Id,
		Name:    secretEntity.Name,
		Project: &model.Project{ID: secretEntity.ProjectId},
		Entries: entries,
		Status:  string(secretEntity.Status),
	}
}

func managedSvcModelFromEntity(svcEntity *entities.ManagedService) *model.ManagedSvc {
	return &model.ManagedSvc{
		CreatedAt: svcEntity.CreationTime.String(),
		UpdatedAt: func() *string {
			s := svcEntity.UpdateTime.String()
			return &s
		}(),
		ID:      svcEntity.Id,
		Name:    svcEntity.Name,
		Project: &model.Project{ID: svcEntity.ProjectId},
		Source:  string(svcEntity.ServiceType),
		Values:  svcEntity.Values,
		Status:  string(svcEntity.Status),
	}
}

func managedResourceModelFromEntity(resEntity *entities.ManagedResource) *model.ManagedRes {
	kvs := make(map[string]any, 0)
	for k, v := range resEntity.Values {
		kvs[k] = v
	}
	return &model.ManagedRes{
		ID:           resEntity.Id,
		Name:         resEntity.Name,
		ResourceType: string(resEntity.ResourceType),
		Installation: &model.ManagedSvc{
			ID: resEntity.ServiceId,
		},
		Values:    kvs,
		Status:    string(resEntity.Status),
		CreatedAt: resEntity.CreationTime.String(),
		UpdatedAt: func() *string {
			s := resEntity.UpdateTime.String()
			return &s
		}(),
	}
}