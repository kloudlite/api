package domain

import (
	"context"

	fWebsocket "github.com/gofiber/websocket/v2"
	"kloudlite.io/apps/console/internal/domain/entities"
	"kloudlite.io/apps/console/internal/domain/entities/localenv"
	op_crds "kloudlite.io/apps/console/internal/domain/op-crds"
	"kloudlite.io/pkg/cache"
	"kloudlite.io/pkg/repos"
)

type Domain interface {
	GetEdgeRegions(ctx context.Context, providerId repos.ID) ([]*entities.EdgeRegion, error)
	CreateEdgeRegion(ctx context.Context, providerId repos.ID, region *entities.EdgeRegion) error
	DeleteEdgeRegion(ctx context.Context, edgeId repos.ID) error
	UpdateEdgeRegion(ctx context.Context, edgeId repos.ID, region *EdgeRegionUpdate) error

	GetCloudProviders(ctx context.Context, accountId repos.ID) ([]*entities.CloudProvider, error)
	CreateCloudProvider(ctx context.Context, accountId *repos.ID, provider *entities.CloudProvider) error
	DeleteCloudProvider(ctx context.Context, providerId repos.ID) error
	UpdateCloudProvider(ctx context.Context, providerId repos.ID, provider *CloudProviderUpdate) error

	GetComputePlan(ctx context.Context, name string) (*entities.ComputePlan, error)
	GetComputePlans(ctx context.Context) ([]entities.ComputePlan, error)
	GetStoragePlans(ctx context.Context) ([]entities.StoragePlan, error)

	GetDevice(ctx context.Context, id repos.ID) (*entities.Device, error)
	GetDeviceConfig(ctx context.Context, id repos.ID) (map[string]any, error)
	AddDevice(
		ctx context.Context,
		deviceName string, accountId repos.ID, userId repos.ID) (dev *entities.Device, e error)
	UpdateDevice(ctx context.Context, deviceId repos.ID, deviceName *string, region *string, ports []entities.Port) (done bool, e error)
	RemoveDevice(ctx context.Context, deviceId repos.ID) error
	ListAccountDevices(ctx context.Context, accountId repos.ID) ([]*entities.Device, error)
	ListUserDevices(ctx context.Context, userId repos.ID) ([]*entities.Device, error)

	CreateProject(ctx context.Context, ownerId repos.ID, accountId repos.ID, projectName string, displayName string, logo *string, regionId *repos.ID, description *string) (*entities.Project, error)
	GetAccountProjects(ctx context.Context, id repos.ID) ([]*entities.Project, error)
	GetProjectWithID(ctx context.Context, projectId repos.ID) (*entities.Project, error)
	OnUpdateProject(ctx context.Context, response *op_crds.StatusUpdate) error
	OnDeleteProject(ctx context.Context, response *op_crds.StatusUpdate) error

	CreateConfig(ctx context.Context, id repos.ID, configName string, desc *string, configData []*entities.Entry) (*entities.Config, error)
	UpdateConfig(ctx context.Context, configId repos.ID, desc *string, configData []*entities.Entry) (bool, error)

	GetConfigs(ctx context.Context, projectId repos.ID) ([]*entities.Config, error)
	GetConfig(ctx context.Context, configId repos.ID) (*entities.Config, error)
	DeleteConfig(ctx context.Context, configId repos.ID) (bool, error)
	//OnUpdateConfig(ctx context.Context, configId repos.ID) error

	CreateSecret(ctx context.Context, projectId repos.ID, secretName string, desc *string, secretData []*entities.Entry) (*entities.Secret, error)
	UpdateSecret(ctx context.Context, secretId repos.ID, desc *string, secretData []*entities.Entry) (bool, error)

	DeleteSecret(ctx context.Context, secretId repos.ID) (bool, error)
	GetSecrets(ctx context.Context, projectId repos.ID) ([]*entities.Secret, error)
	GetSecret(ctx context.Context, secretId repos.ID) (*entities.Secret, error)
	//OnUpdateSecret(ctx context.Context, secretId repos.ID) error

	GetRouters(ctx context.Context, projectID repos.ID) ([]*entities.Router, error)
	GetRouter(ctx context.Context, routerID repos.ID) (*entities.Router, error)
	DeleteRouter(ctx context.Context, routerID repos.ID) (bool, error)
	CreateRouter(ctx context.Context, id repos.ID, routerName string, domains []string, routes []*entities.Route) (*entities.Router, error)
	UpdateRouter(ctx context.Context, id repos.ID, domains []string, entries []*entities.Route) (bool, error)
	OnUpdateRouter(ctx context.Context, r *op_crds.StatusUpdate) error
	OnDeleteRouter(ctx context.Context, r *op_crds.StatusUpdate) error

	GetManagedSvc(ctx context.Context, managedSvcID repos.ID) (*entities.ManagedService, error)
	GetManagedSvcOutput(ctx context.Context, managedSvcID repos.ID) (map[string]any, error)
	GetManagedSvcs(ctx context.Context, projectID repos.ID) ([]*entities.ManagedService, error)
	InstallManagedSvc(ctx context.Context, projectID repos.ID, templateID repos.ID, name string, values map[string]interface{}) (*entities.ManagedService, error)
	UpdateManagedSvc(ctx context.Context, managedServiceId repos.ID, values map[string]interface{}) (bool, error)
	UnInstallManagedSvc(ctx context.Context, managedServiceId repos.ID) (bool, error)
	OnUpdateManagedSvc(ctx context.Context, r *op_crds.StatusUpdate) error
	OnDeleteManagedService(todo context.Context, o *op_crds.StatusUpdate) error

	GetManagedRes(ctx context.Context, managedResID repos.ID) (*entities.ManagedResource, error)
	GetManagedResOutput(ctx context.Context, managedResID repos.ID) (map[string]any, error)
	GetManagedResources(ctx context.Context, projectID repos.ID) ([]*entities.ManagedResource, error)
	GetManagedResourcesOfService(ctx context.Context, installationId repos.ID) ([]*entities.ManagedResource, error)
	OnDeleteManagedResource(todo context.Context, o *op_crds.StatusUpdate) error

	InstallManagedRes(
		ctx context.Context,
		installationId repos.ID,
		name string,
		resourceType string,
		values map[string]string,
	) (*entities.ManagedResource, error)
	UpdateManagedRes(ctx context.Context, managedResID repos.ID, values map[string]string) (bool, error)
	UnInstallManagedRes(ctx context.Context, managedResID repos.ID) (bool, error)
	OnUpdateManagedRes(ctx context.Context, r *op_crds.StatusUpdate) error

	GetApps(ctx context.Context, projectId repos.ID) ([]*entities.App, error)
	GetInterceptedApps(ctx context.Context, deviceId repos.ID) ([]*entities.App, error)
	FreezeApp(ctx context.Context, appId repos.ID) error
	UnFreezeApp(ctx context.Context, appId repos.ID) error
	RestartApp(ctx context.Context, appId repos.ID) error
	GetApp(ctx context.Context, projectID repos.ID) (*entities.App, error)
	DeleteApp(ctx context.Context, appID repos.ID) (bool, error)
	OnUpdateApp(ctx context.Context, r *op_crds.StatusUpdate) error
	OnDeleteApp(ctx context.Context, r *op_crds.StatusUpdate) error
	GetManagedServiceTemplates(ctx context.Context) ([]*entities.ManagedServiceCategory, error)
	InstallApp(
		ctx context.Context,
		projectId repos.ID,
		app entities.App,
	) (*entities.App, error)
	UpdateApp(
		ctx context.Context,
		appId repos.ID,
		app entities.App,
	) (*entities.App, error)

	GetProjectMemberships(ctx context.Context, projectID repos.ID) ([]*entities.ProjectMembership, error)
	InviteProjectMember(ctx context.Context, projectID repos.ID, email string, role string) (bool, error)
	RemoveProjectMember(ctx context.Context, projectId repos.ID, userId repos.ID) error

	SetupAccount(ctx context.Context, accountId repos.ID) (bool, error)

	DeviceByNameExists(ctx context.Context, accountId repos.ID, name string) (bool, error)
	DeleteProject(ctx context.Context, id repos.ID) (bool, error)
	GetDockerCredentials(ctx context.Context, id repos.ID) (username string, password string, err error)
	GenerateEnv(ctx context.Context, klfile localenv.KLFile) (map[string]string, map[string]string, error)
	InterceptApp(ctx context.Context, appId repos.ID, deviceId repos.ID) error
	CloseIntercept(ctx context.Context, appId repos.ID) error

	GetSocketCtx(
		conn *fWebsocket.Conn,
		cacheClient AuthCacheClient,
		cookieName,
		cookieDomain string,
		sessionKeyPrefix string,
	) context.Context
}

type AuthCacheClient cache.Client

type InfraActionMessage interface {
	entities.SetupClusterAccountAction | entities.SetupClusterAction | entities.DeleteClusterAction | entities.UpdateClusterAction | entities.AddPeerAction | entities.DeletePeerAction
}

type InfraMessenger interface {
	SendAction(action any) error
}

type WorkloadMessenger interface {
	SendAction(action string, resId string, res any) error
}

func SendAction[T InfraActionMessage](i InfraMessenger, action T) error {
	return i.SendAction(action)
}
