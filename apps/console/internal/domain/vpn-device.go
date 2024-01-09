package domain

import (
	"encoding/json"

	"github.com/kloudlite/api/apps/console/internal/entities"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/infra"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
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

	device.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	device.LastUpdatedBy = device.CreatedBy
	device.AccountName = ctx.AccountName

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

	clusterName, err := d.getClusterAttachedToProject(ctx, *device.ProjectName)
	if err != nil {
		return nil, errors.NewE(err)
	}
	if clusterName != nil {
		return nil, errors.NewE(errors.Newf("no cluster attached to project %s, so could not activate vpn device", *device.ProjectName))
	}

	deviceBytes, err := json.Marshal(nDevice.Device)
	if err != nil {
		return nil, errors.NewE(err)
	}

	resp, err := d.infraClient.UpsertVpnDevice(ctx, &infra.UpsertVpnDeviceIn{
		AccountName: ctx.AccountName,
		ClusterName: *clusterName,
		VpnDevice:   deviceBytes,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if err := json.Unmarshal(resp.VpnDevice, &nDevice.Device); err != nil {
		return nil, errors.NewE(err)
	}

	if err := json.Unmarshal(resp.WgConfig, &nDevice.WireguardConfig); err != nil {
		return nil, errors.NewE(err)
	}

	return nDevice, nil
}

func (d *domain) UpdateVPNDevice(ctx ConsoleContext, device entities.VPNDevice) (*entities.VPNDevice, error) {
	if err := d.canPerformActionInDevice(ctx, iamT.UpdateVPNDevice, device.Name); err != nil {
		return nil, errors.NewE(err)
	}

	currDevice, err := d.findVPNDevice(ctx, device.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	currDevice.LastUpdatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	currDevice.DisplayName = device.DisplayName
	currDevice.Spec.Ports = device.Spec.Ports

	nDevice, err := d.vpnDeviceRepo.UpdateById(ctx, device.Id, &device)
	if err != nil {
		return nil, errors.NewE(err)
	}
	d.resourceEventPublisher.PublishVpnDeviceEvent(nDevice, PublishUpdate)

	deviceBytes, err := json.Marshal(nDevice.Device)
	if err != nil {
		return nil, errors.NewE(err)
	}

	clusterName, err := d.getClusterAttachedToProject(ctx, *device.ProjectName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if clusterName == nil {
		return nil, errors.NewE(errors.Newf("no cluster attached to project %s, so could not activate vpn device", *device.ProjectName))
	}

	infraDevOut, err := d.infraClient.UpsertVpnDevice(ctx, &infra.UpsertVpnDeviceIn{
		AccountName: ctx.AccountName,
		ClusterName: *clusterName,
		VpnDevice:   deviceBytes,
	})

	if err := json.Unmarshal(infraDevOut.VpnDevice, &nDevice.Device); err != nil {
		return nil, errors.NewE(err)
	}

	if err := json.Unmarshal(infraDevOut.WgConfig, &nDevice.WireguardConfig); err != nil {
		return nil, errors.NewE(err)
	}

	return nDevice, nil
}

func (d *domain) DeleteVPNDevice(ctx ConsoleContext, name string) error {

	if err := d.canPerformActionInDevice(ctx, iamT.DeleteVPNDevice, name); err != nil {
		return errors.NewE(err)
	}

	device, err := d.findVPNDevice(ctx, name)
	if err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishVpnDeviceEvent(device, PublishDelete)

	clusterName, err := d.getClusterAttachedToProject(ctx, *device.ProjectName)
	if err != nil {
		return errors.NewE(err)
	}

	if clusterName != nil {
		_, err := d.infraClient.DeleteVpnDevice(ctx, &infra.DeleteVpnDeviceIn{
			AccountName: ctx.AccountName,
			Id:          string(device.Id),
		})
		if err != nil {
			return errors.NewE(err)
		}
	}

	if err := d.vpnDeviceRepo.DeleteById(ctx, device.Id); err != nil {
		return errors.NewE(err)
	}
	return nil
}
