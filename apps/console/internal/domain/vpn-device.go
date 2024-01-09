package domain

import (
	"github.com/kloudlite/api/apps/console/internal/entities"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func (d *domain) findVPNDevice(ctx ConsoleContext, name string) (*entities.VPNDevice, error) {
	device, err := d.vpnDeviceRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"metadata.name": name,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if device == nil {
		return nil, errors.Newf("no vpn device with name=%q found", name)
	}

	return device, nil
}

func (d *domain) ListVPNDevices(ctx ConsoleContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.VPNDevice], error) {
	if err := d.canPerformActionInAccount(ctx, iamT.CreateVPNDevice); err != nil {
		return nil, errors.NewE(err)
	}

	filter := repos.Filter{"accountName": ctx.AccountName}
	return d.vpnDeviceRepo.FindPaginated(ctx, d.vpnDeviceRepo.MergeMatchFilters(filter, search), pagination)
}

func (d *domain) GetVPNDevice(ctx ConsoleContext, name string) (*entities.VPNDevice, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.GetVPNDevice); err != nil {
		return nil, errors.NewE(err)
	}
	return d.findVPNDevice(ctx, name)
}

func (d *domain) CreateVPNDevice(ctx ConsoleContext, device entities.VPNDevice) (*entities.VPNDevice, error) {

	if err := d.canPerformActionInAccount(ctx, iamT.CreateVPNDevice); err != nil {
		return nil, errors.NewE(err)
	}

	device.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &device.Device); err != nil {
		return nil, errors.NewE(err)
	}

	device.IncrementRecordVersion()
	device.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	device.LastUpdatedBy = device.CreatedBy

	device.AccountName = ctx.AccountName
	device.SyncStatus = t.GenSyncStatus(t.SyncActionApply, device.RecordVersion)

	if _, err := d.iamClient.AddMembership(ctx, &iam.AddMembershipIn{
		UserId:       string(ctx.UserId),
		ResourceType: string(iamT.ResourceVPNDevice),
		ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceVPNDevice, device.Name),
		Role:         string(iamT.RoleResourceOwner),
	}); err != nil {
		return nil, errors.NewE(err)
	}

	nDevice, err := d.vpnDeviceRepo.Create(ctx, &device)
	if err != nil {
		if d.vpnDeviceRepo.ErrAlreadyExists(err) {
			// TODO: better insights into error, when it is being caused by duplicated indexes
			return nil, errors.NewE(err)
		}
		return nil, errors.NewE(err)
	}

	d.resourceEventPublisher.PublishVpnDeviceEvent(&device, PublishAdd)

	// TODO: push device to infra using grpc
	// considring cluster of project is same for each environment of that cluster
	// if device.ProjectName != nil {
	// 	device.Spec.Disabled = false
	// 	if err := d.applyK8sResource(ctx, *device.ProjectName, &device.Device, device.RecordVersion); err != nil {
	// 		return nil, errors.NewE(err)
	// 	}
	// }
	//
	return nDevice, nil
}

func (d *domain) UpdateVPNDevice(ctx ConsoleContext, device entities.VPNDevice) (*entities.VPNDevice, error) {
	if err := d.canPerformActionInDevice(ctx, iamT.UpdateVPNDevice, device.Name); err != nil {
		return nil, errors.NewE(err)
	}

	device.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &device.Device); err != nil {
		return nil, errors.NewE(err)
	}

	currDevice, err := d.findVPNDevice(ctx, device.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	currDevice.IncrementRecordVersion()
	currDevice.LastUpdatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	currDevice.DisplayName = device.DisplayName

	currDevice.Labels = device.Labels
	currDevice.Annotations = device.Annotations

	currDevice.Spec.Ports = device.Spec.Ports

	currDevice.SyncStatus = t.GenSyncStatus(t.SyncActionApply, currDevice.RecordVersion)

	nDevice, err := d.vpnDeviceRepo.UpdateById(ctx, device.Id, &device)
	if err != nil {
		return nil, errors.NewE(err)
	}
	d.resourceEventPublisher.PublishVpnDeviceEvent(nDevice, PublishUpdate)

	// TODO: push device to infra using grpc
	// if device.ProjectName != nil && *device.ProjectName != "" {
	// 	if err := d.applyK8sResource(ctx, *device.ProjectName, &device.Device, device.RecordVersion); err != nil {
	// 		return nil, errors.NewE(err)
	// 	}
	//
	// 	if currDevice.ProjectName != nil && *currDevice.ProjectName != *device.ProjectName {
	//
	// 		currDevice.Spec.Disabled = true
	//
	// 		if err := d.applyK8sResource(ctx, *currDevice.ProjectName, &currDevice.Device, currDevice.RecordVersion); err != nil {
	// 			return nil, errors.NewE(err)
	// 		}
	// 	}
	//
	// }

	return nDevice, nil
}

func (d *domain) DeleteVPNDevice(ctx ConsoleContext, name string) error {
	panic("not implemented")
}

func (d *domain) OnVPNDeviceApplyError(ctx ConsoleContext, name, errMsg string, opts UpdateAndDeleteOpts) error {
	panic("not implemented")
}

func (d *domain) OnVPNDeviceDeleteMessage(ctx ConsoleContext, device entities.VPNDevice) error {

	panic("not implemented")
}

func (d *domain) OnVPNDeviceUpdateMessage(ctx ConsoleContext, device entities.VPNDevice, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	panic("not implemented")
}

func (d *domain) ResyncVPNDevice(ctx ConsoleContext, name string) error {
	panic("not implemented")
}
