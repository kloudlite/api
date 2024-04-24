package domain

import (
	"context"
	fc "github.com/kloudlite/api/apps/iot-console/internal/entities/field-constants"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
)

type ResourceType string

const (
	ResourceTypeIOTProject         ResourceType = "iot_project"
	ResourceTypeIOTDeviceBlueprint ResourceType = "iot_device_blueprint"
	ResourceTypeIOTDeployment      ResourceType = "iot_deployment"
	ResourceTypeIOTDevice          ResourceType = "iot_device"
)

type CheckNameAvailabilityOutput struct {
	Result         bool     `json:"result"`
	SuggestedNames []string `json:"suggestedNames,omitempty"`
}

func checkResourceName[T repos.Entity](ctx context.Context, filters repos.Filter, repo repos.DbRepo[T]) (*CheckNameAvailabilityOutput, error) {
	res, err := repo.FindOne(ctx, filters)
	if err != nil {
		return &CheckNameAvailabilityOutput{Result: false}, errors.NewE(err)
	}

	if fn.IsNil(res) {
		return &CheckNameAvailabilityOutput{Result: true}, nil
	}

	return &CheckNameAvailabilityOutput{
		Result:         false,
		SuggestedNames: fn.GenValidK8sResourceNames(filters["name"].(string), 3),
	}, nil
}

func checkAppResourceName[T repos.Entity](ctx context.Context, filters repos.Filter, repo repos.DbRepo[T]) (*CheckNameAvailabilityOutput, error) {
	res, err := repo.FindOne(ctx, filters)
	if err != nil {
		return &CheckNameAvailabilityOutput{Result: false}, errors.NewE(err)
	}

	if fn.IsNil(res) {
		return &CheckNameAvailabilityOutput{Result: true}, nil
	}

	return &CheckNameAvailabilityOutput{
		Result:         false,
		SuggestedNames: fn.GenValidK8sResourceNames(filters[fields.MetadataName].(string), 3),
	}, nil
}

func (d *domain) CheckNameAvailability(ctx IotResourceContext, deviceBlueprintName *string, deploymentName *string, resourceType ResourceType, name string) (*CheckNameAvailabilityOutput, error) {

	switch resourceType {
	case ResourceTypeIOTProject:
		{
			return checkResourceName(ctx, repos.Filter{
				fields.AccountName: ctx.AccountName,
				fc.IOTProjectName:  name,
			}, d.iotProjectRepo)
		}
	case ResourceTypeIOTDeviceBlueprint:
		{
			return checkResourceName(ctx, repos.Filter{
				fields.AccountName:        ctx.AccountName,
				fields.ProjectName:        ctx.ProjectName,
				fc.IOTDeviceBlueprintName: name,
			}, d.iotDeviceBlueprintRepo)
		}
	case ResourceTypeIOTDeployment:
		{
			return checkResourceName(ctx, repos.Filter{
				fields.AccountName:   ctx.AccountName,
				fields.ProjectName:   ctx.ProjectName,
				fc.IOTDeploymentName: name,
			}, d.iotDeploymentRepo)
		}
	case ResourceTypeIOTDevice:
		{
			return checkResourceName(ctx, repos.Filter{
				fields.AccountName:         ctx.AccountName,
				fields.ProjectName:         ctx.ProjectName,
				fc.IOTDeviceDeploymentName: deploymentName,
				fc.IOTDeviceName:           name,
			}, d.iotDeviceRepo)
		}
	default:
		{
			return &CheckNameAvailabilityOutput{Result: false}, errors.Newf("unknown resource type provided: %q", resourceType)
		}
	}

}
