package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"

	"github.com/kloudlite/api/pkg/errors"

	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/app/graph/model"
	"github.com/kloudlite/api/apps/console/internal/domain"
	"github.com/kloudlite/api/apps/console/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	v11 "github.com/kloudlite/operator/apis/crds/v1"
	"github.com/kloudlite/operator/apis/wireguard/v1"
)

// CoreCreateProject is the resolver for the core_createProject field.
func (r *mutationResolver) CoreCreateProject(ctx context.Context, project entities.Project) (*entities.Project, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateProject(cc, project)
}

// CoreUpdateProject is the resolver for the core_updateProject field.
func (r *mutationResolver) CoreUpdateProject(ctx context.Context, project entities.Project) (*entities.Project, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateProject(cc, project)
}

// CoreDeleteProject is the resolver for the core_deleteProject field.
func (r *mutationResolver) CoreDeleteProject(ctx context.Context, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteProject(cc, name); err != nil {
		return false, nil
	}
	return true, nil
}

// CoreCreateEnvironment is the resolver for the core_createEnvironment field.
func (r *mutationResolver) CoreCreateEnvironment(ctx context.Context, projectName string, env entities.Environment) (*entities.Environment, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateEnvironment(cc, projectName, env)
}

// CoreUpdateEnvironment is the resolver for the core_updateEnvironment field.
func (r *mutationResolver) CoreUpdateEnvironment(ctx context.Context, projectName string, env entities.Environment) (*entities.Environment, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateEnvironment(cc, projectName, env)
}

// CoreDeleteEnvironment is the resolver for the core_deleteEnvironment field.
func (r *mutationResolver) CoreDeleteEnvironment(ctx context.Context, projectName string, envName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteEnvironment(cc, projectName, envName); err != nil {
		return false, nil
	}
	return true, nil
}

// CoreCloneEnvironment is the resolver for the core_cloneEnvironment field.
func (r *mutationResolver) CoreCloneEnvironment(ctx context.Context, projectName string, sourceEnvName string, destinationEnvName string, displayName string, environmentRoutingMode v11.EnvironmentRoutingMode) (*entities.Environment, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CloneEnvironment(cc, projectName, sourceEnvName, destinationEnvName, displayName, environmentRoutingMode)
}

// CoreCreateImagePullSecret is the resolver for the core_createImagePullSecret field.
func (r *mutationResolver) CoreCreateImagePullSecret(ctx context.Context, projectName string, envName string, imagePullSecretIn entities.ImagePullSecret) (*entities.ImagePullSecret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.CreateImagePullSecret(newResourceContext(cc, projectName, envName), imagePullSecretIn)
}

// CoreDeleteImagePullSecret is the resolver for the core_deleteImagePullSecret field.
func (r *mutationResolver) CoreDeleteImagePullSecret(ctx context.Context, projectName string, envName string, secretName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteImagePullSecret(newResourceContext(cc, projectName, envName), secretName); err != nil {
		return false, nil
	}
	return true, nil
}

// CoreCreateApp is the resolver for the core_createApp field.
func (r *mutationResolver) CoreCreateApp(ctx context.Context, projectName string, envName string, app entities.App) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateApp(newResourceContext(cc, projectName, envName), app)
}

// CoreUpdateApp is the resolver for the core_updateApp field.
func (r *mutationResolver) CoreUpdateApp(ctx context.Context, projectName string, envName string, app entities.App) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateApp(newResourceContext(cc, projectName, envName), app)
}

// CoreDeleteApp is the resolver for the core_deleteApp field.
func (r *mutationResolver) CoreDeleteApp(ctx context.Context, projectName string, envName string, appName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteApp(newResourceContext(cc, projectName, envName), appName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreInterceptApp is the resolver for the core_interceptApp field.
func (r *mutationResolver) CoreInterceptApp(ctx context.Context, projectName string, envName string, appname string, deviceName string, intercept bool) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	return r.Domain.InterceptApp(newResourceContext(cc, projectName, envName), appname, deviceName, intercept)
}

// CoreCreateConfig is the resolver for the core_createConfig field.
func (r *mutationResolver) CoreCreateConfig(ctx context.Context, projectName string, envName string, config entities.Config) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateConfig(newResourceContext(cc, projectName, envName), config)
}

// CoreUpdateConfig is the resolver for the core_updateConfig field.
func (r *mutationResolver) CoreUpdateConfig(ctx context.Context, projectName string, envName string, config entities.Config) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateConfig(newResourceContext(cc, projectName, envName), config)
}

// CoreDeleteConfig is the resolver for the core_deleteConfig field.
func (r *mutationResolver) CoreDeleteConfig(ctx context.Context, projectName string, envName string, configName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteConfig(newResourceContext(cc, projectName, envName), configName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateSecret is the resolver for the core_createSecret field.
func (r *mutationResolver) CoreCreateSecret(ctx context.Context, projectName string, envName string, secret entities.Secret) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateSecret(newResourceContext(cc, projectName, envName), secret)
}

// CoreUpdateSecret is the resolver for the core_updateSecret field.
func (r *mutationResolver) CoreUpdateSecret(ctx context.Context, projectName string, envName string, secret entities.Secret) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateSecret(newResourceContext(cc, projectName, envName), secret)
}

// CoreDeleteSecret is the resolver for the core_deleteSecret field.
func (r *mutationResolver) CoreDeleteSecret(ctx context.Context, projectName string, envName string, secretName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteSecret(newResourceContext(cc, projectName, envName), secretName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateRouter is the resolver for the core_createRouter field.
func (r *mutationResolver) CoreCreateRouter(ctx context.Context, projectName string, envName string, router entities.Router) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateRouter(newResourceContext(cc, projectName, envName), router)
}

// CoreUpdateRouter is the resolver for the core_updateRouter field.
func (r *mutationResolver) CoreUpdateRouter(ctx context.Context, projectName string, envName string, router entities.Router) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateRouter(newResourceContext(cc, projectName, envName), router)
}

// CoreDeleteRouter is the resolver for the core_deleteRouter field.
func (r *mutationResolver) CoreDeleteRouter(ctx context.Context, projectName string, envName string, routerName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteRouter(newResourceContext(cc, projectName, envName), routerName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateManagedResource is the resolver for the core_createManagedResource field.
func (r *mutationResolver) CoreCreateManagedResource(ctx context.Context, projectName string, envName string, mres entities.ManagedResource) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateManagedResource(newResourceContext(cc, projectName, envName), mres)
}

// CoreUpdateManagedResource is the resolver for the core_updateManagedResource field.
func (r *mutationResolver) CoreUpdateManagedResource(ctx context.Context, projectName string, envName string, mres entities.ManagedResource) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateManagedResource(newResourceContext(cc, projectName, envName), mres)
}

// CoreDeleteManagedResource is the resolver for the core_deleteManagedResource field.
func (r *mutationResolver) CoreDeleteManagedResource(ctx context.Context, projectName string, envName string, mresName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteManagedResource(newResourceContext(cc, projectName, envName), mresName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateProjectManagedService is the resolver for the core_createProjectManagedService field.
func (r *mutationResolver) CoreCreateProjectManagedService(ctx context.Context, projectName string, pmsvc entities.ProjectManagedService) (*entities.ProjectManagedService, error) {
	ictx, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateProjectManagedService(ictx, projectName, pmsvc)
}

// CoreUpdateProjectManagedService is the resolver for the core_updateProjectManagedService field.
func (r *mutationResolver) CoreUpdateProjectManagedService(ctx context.Context, projectName string, pmsvc entities.ProjectManagedService) (*entities.ProjectManagedService, error) {
	ictx, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateProjectManagedService(ictx, projectName, pmsvc)
}

// CoreDeleteProjectManagedService is the resolver for the core_deleteProjectManagedService field.
func (r *mutationResolver) CoreDeleteProjectManagedService(ctx context.Context, projectName string, pmsvcName string) (bool, error) {
	ictx, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	err = r.Domain.DeleteProjectManagedService(ictx, projectName, pmsvcName)
	if err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateVPNDevice is the resolver for the core_createVPNDevice field.
func (r *mutationResolver) CoreCreateVPNDevice(ctx context.Context, vpnDevice entities.ConsoleVPNDevice) (*entities.ConsoleVPNDevice, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.CreateVPNDevice(cc, vpnDevice)
}

// CoreUpdateVPNDevice is the resolver for the core_updateVPNDevice field.
func (r *mutationResolver) CoreUpdateVPNDevice(ctx context.Context, vpnDevice entities.ConsoleVPNDevice) (*entities.ConsoleVPNDevice, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.UpdateVPNDevice(cc, vpnDevice)
}

// CoreUpdateVPNDevicePorts is the resolver for the core_updateVPNDevicePorts field.
func (r *mutationResolver) CoreUpdateVPNDevicePorts(ctx context.Context, deviceName string, ports []*v1.Port) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.UpdateVpnDevicePorts(cc, deviceName, ports); err != nil {
		return false, errors.NewE(err)
	}

	return true, nil
}

// CoreUpdateVPNDeviceEnv is the resolver for the core_updateVPNDeviceEnv field.
func (r *mutationResolver) CoreUpdateVPNDeviceEnv(ctx context.Context, deviceName string, projectName string, envName string) (bool, error) {
	if projectName == "" && envName == "" {
		return false, fmt.Errorf("projectName and envName cannot be empty")
	}

	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.UpdateVpnDeviceEnvironment(cc, deviceName, projectName, envName); err != nil {
		return false, errors.NewE(err)
	}

	return true, nil
}

// CoreDeleteVPNDevice is the resolver for the core_deleteVPNDevice field.
func (r *mutationResolver) CoreDeleteVPNDevice(ctx context.Context, deviceName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.DeleteVPNDevice(cc, deviceName); err != nil {
		return false, errors.NewE(err)
	}

	return true, nil
}

// CoreCheckNameAvailability is the resolver for the core_checkNameAvailability field.
func (r *queryResolver) CoreCheckNameAvailability(ctx context.Context, projectName string, envName *string, resType entities.ResourceType, name string) (*domain.CheckNameAvailabilityOutput, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.CheckNameAvailability(ctx, cc.AccountName, projectName, envName, resType, name)
}

// CoreListProjects is the resolver for the core_listProjects field.
func (r *queryResolver) CoreListProjects(ctx context.Context, search *model.SearchProjects, pq *repos.CursorPagination) (*model.ProjectPaginatedRecords, error) {
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	cc, err := toConsoleContext(ctx)
	if err != nil {
		// if cc.UserId == "" || cc.AccountName == "" {
		// }
		return nil, errors.NewE(err)
	}

	p, err := r.Domain.ListProjects(ctx, cc.UserId, cc.AccountName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.ProjectPaginatedRecords](p)
}

// CoreGetProject is the resolver for the core_getProject field.
func (r *queryResolver) CoreGetProject(ctx context.Context, name string) (*entities.Project, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetProject(cc, name)
}

// CoreResyncProject is the resolver for the core_resyncProject field.
func (r *queryResolver) CoreResyncProject(ctx context.Context, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncProject(cc, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListEnvironments is the resolver for the core_listEnvironments field.
func (r *queryResolver) CoreListEnvironments(ctx context.Context, projectName string, search *model.SearchEnvironments, pq *repos.CursorPagination) (*model.EnvironmentPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.ProjectName != nil {
			filter["spec.projectName"] = *search.ProjectName
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	envs, err := r.Domain.ListEnvironments(cc, projectName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.EnvironmentPaginatedRecords](envs)
}

// CoreGetEnvironment is the resolver for the core_getEnvironment field.
func (r *queryResolver) CoreGetEnvironment(ctx context.Context, projectName string, name string) (*entities.Environment, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetEnvironment(cc, projectName, name)
}

// CoreResyncEnvironment is the resolver for the core_resyncEnvironment field.
func (r *queryResolver) CoreResyncEnvironment(ctx context.Context, projectName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncEnvironment(cc, projectName, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListImagePullSecrets is the resolver for the infra_listImagePullSecrets field.
func (r *queryResolver) CoreListImagePullSecrets(ctx context.Context, projectName string, envName string, search *model.SearchImagePullSecrets, pq *repos.CursorPagination) (*model.ImagePullSecretPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	pullSecrets, err := r.Domain.ListImagePullSecrets(newResourceContext(cc, projectName, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.ImagePullSecretPaginatedRecords](pullSecrets)
}

// InfraGetImagePullSecret is the resolver for the infra_getImagePullSecret field.
func (r *queryResolver) CoreGetImagePullSecret(ctx context.Context, projectName string, envName string, name string) (*entities.ImagePullSecret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetImagePullSecret(newResourceContext(cc, projectName, envName), name)
}

// CoreResyncImagePullSecret is the resolver for the core_resyncImagePullSecret field.
func (r *queryResolver) CoreResyncImagePullSecret(ctx context.Context, projectName string, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncImagePullSecret(newResourceContext(cc, projectName, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListApps is the resolver for the core_listApps field.
func (r *queryResolver) CoreListApps(ctx context.Context, projectName string, envName string, search *model.SearchApps, pq *repos.CursorPagination) (*model.AppPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	pApps, err := r.Domain.ListApps(newResourceContext(cc, projectName, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.AppPaginatedRecords](pApps)
}

// CoreGetApp is the resolver for the core_getApp field.
func (r *queryResolver) CoreGetApp(ctx context.Context, projectName string, envName string, name string) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetApp(newResourceContext(cc, projectName, envName), name)
}

// CoreResyncApp is the resolver for the core_resyncApp field.
func (r *queryResolver) CoreResyncApp(ctx context.Context, projectName string, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncApp(newResourceContext(cc, projectName, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreGetConfigValues is the resolver for the core_getConfigValues field.
func (r *queryResolver) CoreGetConfigValues(ctx context.Context, projectName string, envName string, queries []*domain.ConfigKeyRef) ([]*domain.ConfigKeyValueRef, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	m := make([]domain.ConfigKeyRef, len(queries))
	for i := range queries {
		m[i] = *queries[i]
	}

	return r.Domain.GetConfigEntries(newResourceContext(cc, projectName, envName), m)
}

// CoreListConfigs is the resolver for the core_listConfigs field.
func (r *queryResolver) CoreListConfigs(ctx context.Context, projectName string, envName string, search *model.SearchConfigs, pq *repos.CursorPagination) (*model.ConfigPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	pConfigs, err := r.Domain.ListConfigs(newResourceContext(cc, projectName, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.ConfigPaginatedRecords](pConfigs)
}

// CoreGetConfig is the resolver for the core_getConfig field.
func (r *queryResolver) CoreGetConfig(ctx context.Context, projectName string, envName string, name string) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetConfig(newResourceContext(cc, projectName, envName), name)
}

// CoreResyncConfig is the resolver for the core_resyncConfig field.
func (r *queryResolver) CoreResyncConfig(ctx context.Context, projectName string, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncConfig(newResourceContext(cc, projectName, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreGetSecretValues is the resolver for the core_getSecretValues field.
func (r *queryResolver) CoreGetSecretValues(ctx context.Context, projectName string, envName string, queries []*domain.SecretKeyRef) ([]*domain.SecretKeyValueRef, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	m := make([]domain.SecretKeyRef, len(queries))
	for i := range queries {
		m[i] = *queries[i]
	}

	return r.Domain.GetSecretEntries(newResourceContext(cc, projectName, envName), m)
}

// CoreListSecrets is the resolver for the core_listSecrets field.
func (r *queryResolver) CoreListSecrets(ctx context.Context, projectName string, envName string, search *model.SearchSecrets, pq *repos.CursorPagination) (*model.SecretPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	pSecrets, err := r.Domain.ListSecrets(newResourceContext(cc, projectName, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.SecretPaginatedRecords](pSecrets)
}

// CoreGetSecret is the resolver for the core_getSecret field.
func (r *queryResolver) CoreGetSecret(ctx context.Context, projectName string, envName string, name string) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetSecret(newResourceContext(cc, projectName, envName), name)
}

// CoreResyncSecret is the resolver for the core_resyncSecret field.
func (r *queryResolver) CoreResyncSecret(ctx context.Context, projectName string, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncSecret(newResourceContext(cc, projectName, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListRouters is the resolver for the core_listRouters field.
func (r *queryResolver) CoreListRouters(ctx context.Context, projectName string, envName string, search *model.SearchRouters, pq *repos.CursorPagination) (*model.RouterPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	pRouters, err := r.Domain.ListRouters(newResourceContext(cc, projectName, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.RouterPaginatedRecords](pRouters)
}

// CoreGetRouter is the resolver for the core_getRouter field.
func (r *queryResolver) CoreGetRouter(ctx context.Context, projectName string, envName string, name string) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetRouter(newResourceContext(cc, projectName, envName), name)
}

// CoreResyncRouter is the resolver for the core_resyncRouter field.
func (r *queryResolver) CoreResyncRouter(ctx context.Context, projectName string, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncRouter(newResourceContext(cc, projectName, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreGetManagedResouceOutputKeys is the resolver for the core_getManagedResouceOutputKeys field.
func (r *queryResolver) CoreGetManagedResouceOutputKeys(ctx context.Context, projectName string, envName string, name string) ([]string, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetManagedResourceOutputKeys(newResourceContext(cc, projectName, envName), name)
}

// CoreGetManagedResouceOutputKeyValues is the resolver for the core_getManagedResouceOutputKeyValues field.
func (r *queryResolver) CoreGetManagedResouceOutputKeyValues(ctx context.Context, projectName string, envName string, keyrefs []*domain.ManagedResourceKeyRef) ([]*domain.ManagedResourceKeyValueRef, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	m := make([]domain.ManagedResourceKeyRef, len(keyrefs))
	for i := range keyrefs {
		m[i] = *keyrefs[i]
	}

	return r.Domain.GetManagedResourceOutputKVs(newResourceContext(cc, projectName, envName), m)
}

// CoreListManagedResources is the resolver for the core_listManagedResources field.
func (r *queryResolver) CoreListManagedResources(ctx context.Context, projectName string, envName string, search *model.SearchManagedResources, pq *repos.CursorPagination) (*model.ManagedResourcePaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}

		if search.ManagedServiceName != nil {
			filter["spec.msvcRef.name"] = *search.ManagedServiceName
		}
	}

	pmsvcs, err := r.Domain.ListManagedResources(newResourceContext(cc, projectName, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.ManagedResourcePaginatedRecords](pmsvcs)
}

// CoreGetManagedResource is the resolver for the core_getManagedResource field.
func (r *queryResolver) CoreGetManagedResource(ctx context.Context, projectName string, envName string, name string) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetManagedResource(newResourceContext(cc, projectName, envName), name)
}

// CoreResyncManagedResource is the resolver for the core_resyncManagedResource field.
func (r *queryResolver) CoreResyncManagedResource(ctx context.Context, projectName string, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncManagedResource(newResourceContext(cc, projectName, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListProjectManagedServices is the resolver for the core_listProjectManagedServices field.
func (r *queryResolver) CoreListProjectManagedServices(ctx context.Context, projectName string, search *model.SearchProjectManagedService, pq *repos.CursorPagination) (*model.ProjectManagedServicePaginatedRecords, error) {
	ictx, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if pq == nil {
		pq = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}

		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pmsvcs, err := r.Domain.ListProjectManagedServices(ictx, projectName, filter, *pq)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.ProjectManagedServicePaginatedRecords](pmsvcs)
}

// CoreGetProjectManagedService is the resolver for the core_getProjectManagedService field.
func (r *queryResolver) CoreGetProjectManagedService(ctx context.Context, projectName string, name string) (*entities.ProjectManagedService, error) {
	ictx, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetProjectManagedService(ictx, projectName, name)
}

// CoreResyncProjectManagedService is the resolver for the core_resyncProjectManagedService field.
func (r *queryResolver) CoreResyncProjectManagedService(ctx context.Context, projectName string, name string) (bool, error) {
	ictx, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncProjectManagedService(ictx, projectName, name); err != nil {
		return false, errors.NewE(err)
	}

	return true, nil
}

// CoreListVPNDevices is the resolver for the core_listVPNDevices field.
func (r *queryResolver) CoreListVPNDevices(ctx context.Context, search *model.CoreSearchVPNDevices, pq *repos.CursorPagination) (*model.ConsoleVPNDevicePaginatedRecords, error) {
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	p, err := r.Domain.ListVPNDevices(cc, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvert[model.ConsoleVPNDevicePaginatedRecords](p)
}

// CoreListVPNDevicesForUser is the resolver for the core_listVPNDevicesForUser field.
func (r *queryResolver) CoreListVPNDevicesForUser(ctx context.Context) ([]*entities.ConsoleVPNDevice, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.ListVPNDevicesForUser(cc)
}

// CoreGetVPNDevice is the resolver for the core_getVPNDevice field.
func (r *queryResolver) CoreGetVPNDevice(ctx context.Context, name string) (*entities.ConsoleVPNDevice, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetVPNDevice(cc, name)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
