package domain

import (
	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	fc "github.com/kloudlite/api/apps/iot-console/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func newIOTResourceContext(ctx IotConsoleContext, projectName string) IotResourceContext {
	return IotResourceContext{
		IotConsoleContext: ctx,
		ProjectName:       projectName,
	}
}

func (d *domain) findDeviceBlueprint(ctx IotResourceContext, name string) (*entities.IOTDeviceBlueprint, error) {
	filter := ctx.IOTConsoleDBFilters()
	filter.Add(fc.IOTDeviceBlueprintName, name)
	devBlueprint, err := d.iotDeviceBlueprintRepo.FindOne(ctx, ctx.IOTConsoleDBFilters().Add("name", name))
	if err != nil {
		return nil, errors.NewE(err)
	}
	if devBlueprint == nil {
		return nil, errors.Newf("no device Blueprint with name=%q found", name)
	}
	return devBlueprint, nil
}

func (d *domain) ListDeviceBlueprints(ctx IotResourceContext, search map[string]repos.MatchFilter, pq repos.CursorPagination) (*repos.PaginatedRecord[*entities.IOTDeviceBlueprint], error) {
	filter := ctx.IOTConsoleDBFilters()
	return d.iotDeviceBlueprintRepo.FindPaginated(ctx, d.iotDeviceBlueprintRepo.MergeMatchFilters(filter, search), pq)
}

func (d *domain) GetDeviceBlueprint(ctx IotResourceContext, name string) (*entities.IOTDeviceBlueprint, error) {
	return d.findDeviceBlueprint(ctx, name)
}

func (d *domain) CreateDeviceBlueprint(ctx IotResourceContext, deviceBlueprint entities.IOTDeviceBlueprint) (*entities.IOTDeviceBlueprint, error) {
	deviceBlueprint.ProjectName = ctx.ProjectName
	deviceBlueprint.AccountName = ctx.AccountName
	deviceBlueprint.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	deviceBlueprint.LastUpdatedBy = deviceBlueprint.CreatedBy

	if deviceBlueprint.BluePrintType == "singleton_blueprint" {
		deviceBlueprint.BluePrintType = entities.SingletonBlueprint
	} else {
		deviceBlueprint.BluePrintType = entities.GroupBlueprint
	}

	nDeviceBlueprint, err := d.iotDeviceBlueprintRepo.Create(ctx, &deviceBlueprint)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return nDeviceBlueprint, nil
}

func (d *domain) UpdateDeviceBlueprint(ctx IotResourceContext, deviceBlueprint entities.IOTDeviceBlueprint) (*entities.IOTDeviceBlueprint, error) {
	patchForUpdate := repos.Document{
		fields.DisplayName: deviceBlueprint.DisplayName,
		fields.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.GetUserId(),
			UserName:  ctx.GetUserName(),
			UserEmail: ctx.GetUserEmail(),
		},
	}

	patchFilter := ctx.IOTConsoleDBFilters().Add(fc.IOTDeviceBlueprintName, deviceBlueprint.Name)

	upDevBlueprint, err := d.iotDeviceBlueprintRepo.Patch(
		ctx,
		patchFilter,
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return upDevBlueprint, nil
}

func (d *domain) DeleteDeviceBlueprint(ctx IotResourceContext, name string) error {
	err := d.iotDeviceBlueprintRepo.DeleteOne(
		ctx,
		ctx.IOTConsoleDBFilters().Add(fc.IOTDeviceBlueprintName, name),
	)
	if err != nil {
		return errors.NewE(err)
	}
	return nil
}

func (d *domain) onDeviceUpdate(ctx IotConsoleContext, clusterName string, bp entities.IOTDeviceBlueprint, status types.ResourceStatus, opts UpdateAndDeleteOpts, newBp *entities.IOTDeviceBlueprint) error {
	devName, err := entities.ExtractDeviceName(clusterName)
	if err != nil {
		return errors.NewE(err)
	}

	xDev, err := d.findDevice(newIOTResourceContext(ctx, bp.ProjectName), *devName, bp.DeploymentName)
	if err != nil {
		return errors.NewE(err)
	}

	if xDev == nil {
		return errors.Newf("no apps found")
	}
	recordVersion, err := d.MatchRecordVersion(bp.Annotations, xDev.RecordVersion)
	if err != nil {
		return errors.NewE(err)
	}

	xDev.CurrentBlueprint = newBp

	uapp, err := d.iotDeviceRepo.PatchById(
		ctx,
		xDev.Id,
		common.PatchForSyncFromAgent(xDev, recordVersion, status, common.PatchOpts{
			MessageTimestamp: opts.MessageTimestamp,
		}))

	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeIOTDevice, uapp.GetName(), PublishUpdate)
	return errors.NewE(err)

}

func (d *domain) OnBlueprintUpdateMessage(ctx IotConsoleContext, clusterName string, bp entities.IOTDeviceBlueprint, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	return d.onDeviceUpdate(ctx, clusterName, bp, status, opts, &bp)
}

func (d *domain) OnBlueprintDeleteMessage(ctx IotConsoleContext, clusterName string, bp entities.IOTDeviceBlueprint) error {
	return d.onDeviceUpdate(ctx, clusterName, bp, types.ResourceStatus(types.ResourceStatusDeleted.String()), UpdateAndDeleteOpts{}, nil)
}

func (d *domain) OnBlueprintApplyError(ctx IotConsoleContext, clusterName string, bpName string, errMsg string, blueprint entities.IOTDeviceBlueprint, opts UpdateAndDeleteOpts) error {
	uapp, err := d.iotDeviceRepo.Patch(
		ctx,
		newIOTResourceContext(ctx, blueprint.ProjectName).IOTConsoleDBFilters().Add(fields.MetadataName, bpName),
		common.PatchForErrorFromAgent(
			errMsg,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
			},
		),
	)
	if err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishResourceEvent(ctx, clusterName, ResourceTypeIOTDevice, uapp.GetName(), PublishUpdate)
	return errors.NewE(err)
}
