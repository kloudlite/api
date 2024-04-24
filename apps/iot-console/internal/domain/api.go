package domain

import (
	"context"
	"time"

	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

type IotConsoleContext struct {
	context.Context
	AccountName string

	UserId    repos.ID
	UserEmail string
	UserName  string
}

type IotResourceContext struct {
	IotConsoleContext
	ProjectName string
}

type UpdateAndDeleteOpts struct {
	MessageTimestamp time.Time
}

func (r IotResourceContext) IOTConsoleDBFilters() repos.Filter {
	return repos.Filter{
		fields.AccountName: r.AccountName,
		fields.ProjectName: r.ProjectName,
	}
}

func (i IotConsoleContext) GetUserId() repos.ID { return i.UserId }

func (i IotConsoleContext) GetUserEmail() string {
	return i.UserEmail
}

func (i IotConsoleContext) GetUserName() string {
	return i.UserName
}
func (i IotConsoleContext) GetAccountName() string { return i.AccountName }

const (
	PublishAdd    PublishMsg = "added"
	PublishDelete PublishMsg = "deleted"
	PublishUpdate PublishMsg = "updated"
)

type Domain interface {
	CheckNameAvailability(ctx IotResourceContext, deviceBlueprintName *string, deploymentName *string, resourceType ResourceType, name string) (*CheckNameAvailabilityOutput, error)
	ListProjects(ctx IotConsoleContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.IOTProject], error)
	GetProject(ctx IotConsoleContext, name string) (*entities.IOTProject, error)

	CreateProject(ctx IotConsoleContext, project entities.IOTProject) (*entities.IOTProject, error)
	UpdateProject(ctx IotConsoleContext, project entities.IOTProject) (*entities.IOTProject, error)
	DeleteProject(ctx IotConsoleContext, name string) error

	ListDeployments(ctx IotResourceContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.IOTDeployment], error)
	GetDeployment(ctx IotResourceContext, name string) (*entities.IOTDeployment, error)

	CreateDeployment(ctx IotResourceContext, deployment entities.IOTDeployment) (*entities.IOTDeployment, error)
	UpdateDeployment(ctx IotResourceContext, deployment entities.IOTDeployment) (*entities.IOTDeployment, error)
	DeleteDeployment(ctx IotResourceContext, name string) error

	ListDevices(ctx IotResourceContext, deploymentName string, search map[string]repos.MatchFilter, pq repos.CursorPagination) (*repos.PaginatedRecord[*entities.IOTDevice], error)
	GetDevice(ctx IotResourceContext, name string, deploymentName string) (*entities.IOTDevice, error)

	GetPublicKeyDevice(ctx context.Context, publicKey string) (*entities.DeviceWithServices, error)

	CreateDevice(ctx IotResourceContext, deploymentName string, device entities.IOTDevice) (*entities.IOTDevice, error)
	UpdateDevice(ctx IotResourceContext, deploymentName string, device entities.IOTDevice) (*entities.IOTDevice, error)
	DeleteDevice(ctx IotResourceContext, deploymentName string, name string) error

	ListDeviceBlueprints(ctx IotResourceContext, search map[string]repos.MatchFilter, pq repos.CursorPagination) (*repos.PaginatedRecord[*entities.IOTDeviceBlueprint], error)
	GetDeviceBlueprint(ctx IotResourceContext, name string) (*entities.IOTDeviceBlueprint, error)

	CreateDeviceBlueprint(ctx IotResourceContext, deviceBlueprint entities.IOTDeviceBlueprint) (*entities.IOTDeviceBlueprint, error)
	UpdateDeviceBlueprint(ctx IotResourceContext, deviceBlueprint entities.IOTDeviceBlueprint) (*entities.IOTDeviceBlueprint, error)
	DeleteDeviceBlueprint(ctx IotResourceContext, name string) error

	OnBlueprintApplyError(ctx IotConsoleContext, clusterName string, name string, errMsg string, blueprint entities.IOTDeviceBlueprint, opts UpdateAndDeleteOpts) error
	OnBlueprintDeleteMessage(ctx IotConsoleContext, clusterName string, blueprint entities.IOTDeviceBlueprint) error
	OnBlueprintUpdateMessage(ctx IotConsoleContext, clusterName string, blueprint entities.IOTDeviceBlueprint, status types.ResourceStatus, opts UpdateAndDeleteOpts) error
}
