package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/kloudlite/api/pkg/errors"

	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/app/graph/model"
	"github.com/kloudlite/api/apps/console/internal/domain"
	"github.com/kloudlite/api/apps/console/internal/entities"
	fc "github.com/kloudlite/api/apps/console/internal/entities/field-constants"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	v11 "github.com/kloudlite/operator/apis/crds/v1"
	"github.com/kloudlite/operator/apis/wireguard/v1"
)

// Build is the resolver for the build field.
func (r *appResolver) Build(ctx context.Context, obj *entities.App) (*model.Build, error) {
	if obj.CIBuildId == nil {
		return nil, nil
	}
	return &model.Build{ID: *obj.CIBuildId}, nil
}

// CoreCreateEnvironment is the resolver for the core_createEnvironment field.
func (r *mutationResolver) CoreCreateEnvironment(ctx context.Context, env entities.Environment) (*entities.Environment, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateEnvironment(cc, env)
}

// CoreUpdateEnvironment is the resolver for the core_updateEnvironment field.
func (r *mutationResolver) CoreUpdateEnvironment(ctx context.Context, env entities.Environment) (*entities.Environment, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateEnvironment(cc, env)
}

// CoreDeleteEnvironment is the resolver for the core_deleteEnvironment field.
func (r *mutationResolver) CoreDeleteEnvironment(ctx context.Context, envName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteEnvironment(cc, envName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCloneEnvironment is the resolver for the core_cloneEnvironment field.
func (r *mutationResolver) CoreCloneEnvironment(ctx context.Context, clusterName string, sourceEnvName string, destinationEnvName string, displayName string, environmentRoutingMode v11.EnvironmentRoutingMode) (*entities.Environment, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.CloneEnvironment(cc, domain.CloneEnvironmentArgs{
		ClusterName:        clusterName,
		SourceEnvName:      sourceEnvName,
		DestinationEnvName: destinationEnvName,
		DisplayName:        displayName,
		EnvRoutingMode:     environmentRoutingMode,
	})
}

// CoreCreateImagePullSecret is the resolver for the core_createImagePullSecret field.
func (r *mutationResolver) CoreCreateImagePullSecret(ctx context.Context, pullSecret entities.ImagePullSecret) (*entities.ImagePullSecret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.CreateImagePullSecret(cc, pullSecret)
}

// CoreUpdateImagePullSecret is the resolver for the core_updateImagePullSecret field.
func (r *mutationResolver) CoreUpdateImagePullSecret(ctx context.Context, pullSecret entities.ImagePullSecret) (*entities.ImagePullSecret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.UpdateImagePullSecret(cc, pullSecret)
}

// CoreDeleteImagePullSecret is the resolver for the core_deleteImagePullSecret field.
func (r *mutationResolver) CoreDeleteImagePullSecret(ctx context.Context, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteImagePullSecret(cc, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateApp is the resolver for the core_createApp field.
func (r *mutationResolver) CoreCreateApp(ctx context.Context, envName string, app entities.App) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateApp(newResourceContext(cc, envName), app)
}

// CoreUpdateApp is the resolver for the core_updateApp field.
func (r *mutationResolver) CoreUpdateApp(ctx context.Context, envName string, app entities.App) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateApp(newResourceContext(cc, envName), app)
}

// CoreDeleteApp is the resolver for the core_deleteApp field.
func (r *mutationResolver) CoreDeleteApp(ctx context.Context, envName string, appName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteApp(newResourceContext(cc, envName), appName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreInterceptApp is the resolver for the core_interceptApp field.
func (r *mutationResolver) CoreInterceptApp(ctx context.Context, envName string, appname string, deviceName string, intercept bool, portMappings []*v11.AppInterceptPortMappings) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	pmappings := make([]v11.AppInterceptPortMappings, 0, len(portMappings))
	for i := range portMappings {
		if portMappings[i] != nil {
			pmappings = append(pmappings, *portMappings[i])
		}
	}

	return r.Domain.InterceptApp(newResourceContext(cc, envName), appname, deviceName, intercept, pmappings)
}

// CoreCreateExternalApp is the resolver for the core_createExternalApp field.
func (r *mutationResolver) CoreCreateExternalApp(ctx context.Context, envName string, externalApp entities.ExternalApp) (*entities.ExternalApp, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateExternalApp(newResourceContext(cc, envName), externalApp)
}

// CoreUpdateExternalApp is the resolver for the core_updateExternalApp field.
func (r *mutationResolver) CoreUpdateExternalApp(ctx context.Context, envName string, externalApp entities.ExternalApp) (*entities.ExternalApp, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateExternalApp(newResourceContext(cc, envName), externalApp)
}

// CoreDeleteExternalApp is the resolver for the core_deleteExternalApp field.
func (r *mutationResolver) CoreDeleteExternalApp(ctx context.Context, envName string, externalAppName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteExternalApp(newResourceContext(cc, envName), externalAppName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreInterceptExternalApp is the resolver for the core_interceptExternalApp field.
func (r *mutationResolver) CoreInterceptExternalApp(ctx context.Context, envName string, externalAppName string, deviceName string, intercept bool, portMappings []*v11.AppInterceptPortMappings) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	pmappings := make([]v11.AppInterceptPortMappings, 0, len(portMappings))
	for i := range portMappings {
		if portMappings[i] != nil {
			pmappings = append(pmappings, *portMappings[i])
		}
	}

	return r.Domain.InterceptExternalApp(newResourceContext(cc, envName), externalAppName, deviceName, intercept, pmappings)
}

// CoreCreateConfig is the resolver for the core_createConfig field.
func (r *mutationResolver) CoreCreateConfig(ctx context.Context, envName string, config entities.Config) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateConfig(newResourceContext(cc, envName), config)
}

// CoreUpdateConfig is the resolver for the core_updateConfig field.
func (r *mutationResolver) CoreUpdateConfig(ctx context.Context, envName string, config entities.Config) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateConfig(newResourceContext(cc, envName), config)
}

// CoreDeleteConfig is the resolver for the core_deleteConfig field.
func (r *mutationResolver) CoreDeleteConfig(ctx context.Context, envName string, configName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteConfig(newResourceContext(cc, envName), configName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateSecret is the resolver for the core_createSecret field.
func (r *mutationResolver) CoreCreateSecret(ctx context.Context, envName string, secret entities.Secret) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateSecret(newResourceContext(cc, envName), secret)
}

// CoreUpdateSecret is the resolver for the core_updateSecret field.
func (r *mutationResolver) CoreUpdateSecret(ctx context.Context, envName string, secret entities.Secret) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateSecret(newResourceContext(cc, envName), secret)
}

// CoreDeleteSecret is the resolver for the core_deleteSecret field.
func (r *mutationResolver) CoreDeleteSecret(ctx context.Context, envName string, secretName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteSecret(newResourceContext(cc, envName), secretName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateRouter is the resolver for the core_createRouter field.
func (r *mutationResolver) CoreCreateRouter(ctx context.Context, envName string, router entities.Router) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateRouter(newResourceContext(cc, envName), router)
}

// CoreUpdateRouter is the resolver for the core_updateRouter field.
func (r *mutationResolver) CoreUpdateRouter(ctx context.Context, envName string, router entities.Router) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateRouter(newResourceContext(cc, envName), router)
}

// CoreDeleteRouter is the resolver for the core_deleteRouter field.
func (r *mutationResolver) CoreDeleteRouter(ctx context.Context, envName string, routerName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteRouter(newResourceContext(cc, envName), routerName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateManagedResource is the resolver for the core_createManagedResource field.
func (r *mutationResolver) CoreCreateManagedResource(ctx context.Context, msvcName string, mres entities.ManagedResource) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateManagedResource(newMresContext(cc, &msvcName, nil), mres)
}

// CoreUpdateManagedResource is the resolver for the core_updateManagedResource field.
func (r *mutationResolver) CoreUpdateManagedResource(ctx context.Context, msvcName string, mres entities.ManagedResource) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateManagedResource(newMresContext(cc, &msvcName, nil), mres)
}

// CoreDeleteManagedResource is the resolver for the core_deleteManagedResource field.
func (r *mutationResolver) CoreDeleteManagedResource(ctx context.Context, msvcName string, mresName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteManagedResource(newMresContext(cc, &msvcName, nil), mresName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreImportManagedResource is the resolver for the core_importManagedResource field.
func (r *mutationResolver) CoreImportManagedResource(ctx context.Context, envName string, msvcName string, mresName string) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.ImportManagedResource(newMresContext(cc, &msvcName, &envName), mresName)
}

// CoreDeleteImportedManagedResource is the resolver for the core_deleteImportedManagedResource field.
func (r *mutationResolver) CoreDeleteImportedManagedResource(ctx context.Context, envName string, mresName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteImportedManagedResource(newResourceContext(cc, envName), mresName); err != nil {
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
func (r *mutationResolver) CoreUpdateVPNDeviceEnv(ctx context.Context, deviceName string, envName string) (bool, error) {
	if envName == "" {
		return false, fmt.Errorf("envName cannot be empty")
	}

	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ActivateVpnDeviceOnEnvironment(cc, deviceName, envName); err != nil {
		return false, errors.NewE(err)
	}

	return true, nil
}

// CoreUpdateVpnDeviceNs is the resolver for the core_updateVpnDeviceNs field.
func (r *mutationResolver) CoreUpdateVpnDeviceNs(ctx context.Context, deviceName string, ns string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ActivateVPNDeviceOnNamespace(cc, deviceName, ns); err != nil {
		return false, errors.NewE(err)
	}

	return true, nil
}

// CoreUpdateVpnClusterName is the resolver for the core_updateVpnClusterName field.
func (r *mutationResolver) CoreUpdateVpnClusterName(ctx context.Context, deviceName string, clusterName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ActivateVpnDeviceOnCluster(cc, deviceName, clusterName); err != nil {
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
func (r *queryResolver) CoreCheckNameAvailability(ctx context.Context, envName *string, resType entities.ResourceType, name string) (*domain.CheckNameAvailabilityOutput, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.CheckNameAvailability(ctx, cc.AccountName, envName, resType, name)
}

// CoreListEnvironments is the resolver for the core_listEnvironments field.
func (r *queryResolver) CoreListEnvironments(ctx context.Context, search *model.SearchEnvironments, pq *repos.CursorPagination) (*model.EnvironmentPaginatedRecords, error) {
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

	envs, err := r.Domain.ListEnvironments(cc, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.EnvironmentPaginatedRecords](envs)
}

// CoreGetEnvironment is the resolver for the core_getEnvironment field.
func (r *queryResolver) CoreGetEnvironment(ctx context.Context, name string) (*entities.Environment, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetEnvironment(cc, name)
}

// CoreResyncEnvironment is the resolver for the core_resyncEnvironment field.
func (r *queryResolver) CoreResyncEnvironment(ctx context.Context, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncEnvironment(cc, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListImagePullSecrets is the resolver for the infra_listImagePullSecrets field.
func (r *queryResolver) CoreListImagePullSecrets(ctx context.Context, search *model.SearchImagePullSecrets, pq *repos.CursorPagination) (*model.ImagePullSecretPaginatedRecords, error) {
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

	pullSecrets, err := r.Domain.ListImagePullSecrets(cc, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.ImagePullSecretPaginatedRecords](pullSecrets)
}

// InfraGetImagePullSecret is the resolver for the infra_getImagePullSecret field.
func (r *queryResolver) CoreGetImagePullSecret(ctx context.Context, name string) (*entities.ImagePullSecret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetImagePullSecret(cc, name)
}

// CoreResyncImagePullSecret is the resolver for the core_resyncImagePullSecret field.
func (r *queryResolver) CoreResyncImagePullSecret(ctx context.Context, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncImagePullSecret(cc, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListApps is the resolver for the core_listApps field.
func (r *queryResolver) CoreListApps(ctx context.Context, envName string, search *model.SearchApps, pq *repos.CursorPagination) (*model.AppPaginatedRecords, error) {
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

	pApps, err := r.Domain.ListApps(newResourceContext(cc, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.AppPaginatedRecords](pApps)
}

// CoreGetApp is the resolver for the core_getApp field.
func (r *queryResolver) CoreGetApp(ctx context.Context, envName string, name string) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetApp(newResourceContext(cc, envName), name)
}

// CoreResyncApp is the resolver for the core_resyncApp field.
func (r *queryResolver) CoreResyncApp(ctx context.Context, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncApp(newResourceContext(cc, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreRestartApp is the resolver for the core_restartApp field.
func (r *queryResolver) CoreRestartApp(ctx context.Context, envName string, appName string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.RestartApp(newResourceContext(cc, envName), appName); err != nil {
		return false, err
	}
	return true, nil
}

// CoreListExternalApps is the resolver for the core_listExternalApps field.
func (r *queryResolver) CoreListExternalApps(ctx context.Context, envName string, search *model.SearchExternalApps, pq *repos.CursorPagination) (*model.ExternalAppPaginatedRecords, error) {
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

	extApps, err := r.Domain.ListExternalApps(newResourceContext(cc, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.ExternalAppPaginatedRecords](extApps)
}

// CoreGetExternalApp is the resolver for the core_getExternalApp field.
func (r *queryResolver) CoreGetExternalApp(ctx context.Context, envName string, name string) (*entities.ExternalApp, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetExternalApp(newResourceContext(cc, envName), name)
}

// CoreResyncExternalApp is the resolver for the core_resyncExternalApp field.
func (r *queryResolver) CoreResyncExternalApp(ctx context.Context, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncExternalApp(newResourceContext(cc, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreGetConfigValues is the resolver for the core_getConfigValues field.
func (r *queryResolver) CoreGetConfigValues(ctx context.Context, envName string, queries []*domain.ConfigKeyRef) ([]*domain.ConfigKeyValueRef, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	m := make([]domain.ConfigKeyRef, len(queries))
	for i := range queries {
		m[i] = *queries[i]
	}

	return r.Domain.GetConfigEntries(newResourceContext(cc, envName), m)
}

// CoreListConfigs is the resolver for the core_listConfigs field.
func (r *queryResolver) CoreListConfigs(ctx context.Context, envName string, search *model.SearchConfigs, pq *repos.CursorPagination) (*model.ConfigPaginatedRecords, error) {
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

	pConfigs, err := r.Domain.ListConfigs(newResourceContext(cc, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.ConfigPaginatedRecords](pConfigs)
}

// CoreGetConfig is the resolver for the core_getConfig field.
func (r *queryResolver) CoreGetConfig(ctx context.Context, envName string, name string) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetConfig(newResourceContext(cc, envName), name)
}

// CoreResyncConfig is the resolver for the core_resyncConfig field.
func (r *queryResolver) CoreResyncConfig(ctx context.Context, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncConfig(newResourceContext(cc, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreGetSecretValues is the resolver for the core_getSecretValues field.
func (r *queryResolver) CoreGetSecretValues(ctx context.Context, envName string, queries []*domain.SecretKeyRef) ([]*domain.SecretKeyValueRef, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	m := make([]domain.SecretKeyRef, len(queries))
	for i := range queries {
		m[i] = *queries[i]
	}

	return r.Domain.GetSecretEntries(newResourceContext(cc, envName), m)
}

// CoreListSecrets is the resolver for the core_listSecrets field.
func (r *queryResolver) CoreListSecrets(ctx context.Context, envName string, search *model.SearchSecrets, pq *repos.CursorPagination) (*model.SecretPaginatedRecords, error) {
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

	pSecrets, err := r.Domain.ListSecrets(newResourceContext(cc, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.SecretPaginatedRecords](pSecrets)
}

// CoreGetSecret is the resolver for the core_getSecret field.
func (r *queryResolver) CoreGetSecret(ctx context.Context, envName string, name string) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetSecret(newResourceContext(cc, envName), name)
}

// CoreResyncSecret is the resolver for the core_resyncSecret field.
func (r *queryResolver) CoreResyncSecret(ctx context.Context, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncSecret(newResourceContext(cc, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListRouters is the resolver for the core_listRouters field.
func (r *queryResolver) CoreListRouters(ctx context.Context, envName string, search *model.SearchRouters, pq *repos.CursorPagination) (*model.RouterPaginatedRecords, error) {
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

	pRouters, err := r.Domain.ListRouters(newResourceContext(cc, envName), filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.RouterPaginatedRecords](pRouters)
}

// CoreGetRouter is the resolver for the core_getRouter field.
func (r *queryResolver) CoreGetRouter(ctx context.Context, envName string, name string) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetRouter(newResourceContext(cc, envName), name)
}

// CoreResyncRouter is the resolver for the core_resyncRouter field.
func (r *queryResolver) CoreResyncRouter(ctx context.Context, envName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncRouter(newResourceContext(cc, envName), name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreGetManagedResouceOutputKeys is the resolver for the core_getManagedResouceOutputKeys field.
func (r *queryResolver) CoreGetManagedResouceOutputKeys(ctx context.Context, msvcName *string, envName *string, name string) ([]string, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if msvcName == nil && envName == nil {
		return nil, errors.New("must specify either msvcName or envName")
	}

	return r.Domain.GetManagedResourceOutputKeys(newMresContext(cc, msvcName, envName), name)
}

// CoreGetManagedResouceOutputKeyValues is the resolver for the core_getManagedResouceOutputKeyValues field.
func (r *queryResolver) CoreGetManagedResouceOutputKeyValues(ctx context.Context, msvcName *string, envName *string, keyrefs []*domain.ManagedResourceKeyRef) ([]*domain.ManagedResourceKeyValueRef, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	m := make([]domain.ManagedResourceKeyRef, len(keyrefs))
	for i := range keyrefs {
		m[i] = *keyrefs[i]
	}

	if msvcName == nil && envName == nil {
		return nil, errors.New("must specify either msvcName or envName")
	}

	return r.Domain.GetManagedResourceOutputKVs(newMresContext(cc, msvcName, envName), m)
}

// CoreListManagedResources is the resolver for the core_listManagedResources field.
func (r *queryResolver) CoreListManagedResources(ctx context.Context, search *model.SearchManagedResources, pq *repos.CursorPagination) (*model.ManagedResourcePaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}

	if search == nil {
		return nil, errors.New("must specify search")
	}

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
			filter["managedServiceName"] = *search.ManagedServiceName
		}

		if search.EnvName != nil {
			filter["environmentName"] = *search.EnvName
		} else {
			filter[fc.ManagedResourceIsImported] = repos.MatchFilter{
				MatchType: repos.MatchTypeExact,
				Exact:     true,
			}
		}
	}

	if search.EnvName == nil && search.ManagedServiceName == nil {
		return nil, errors.New("either envName or managedServiceName must be specified")
	}

	pmsvcs, err := r.Domain.ListManagedResources(cc, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.ManagedResourcePaginatedRecords](pmsvcs)
}

// CoreGetManagedResource is the resolver for the core_getManagedResource field.
func (r *queryResolver) CoreGetManagedResource(ctx context.Context, msvcName *string, envName *string, name string) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetManagedResource(newMresContext(cc, msvcName, envName), name)
}

// CoreResyncManagedResource is the resolver for the core_resyncManagedResource field.
func (r *queryResolver) CoreResyncManagedResource(ctx context.Context, msvcName string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncManagedResource(cc, msvcName, name); err != nil {
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

	return fn.JsonConvertP[model.ConsoleVPNDevicePaginatedRecords](p)
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
