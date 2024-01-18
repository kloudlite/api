package domain

import (
	"time"

	"github.com/kloudlite/api/apps/console/internal/entities"
	fc "github.com/kloudlite/api/apps/console/internal/entities/field-constants"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	wgv1 "github.com/kloudlite/operator/apis/wireguard/v1"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func (d *domain) findVPNDevice(ctx ConsoleContext, name string) (*entities.ConsoleVPNDevice, error) {
	device, err := d.vpnDeviceRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.MetadataName: name,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if device == nil {
		return nil, errors.Newf("no vpn device with name=%q found", name)
	}

	return device, nil
}

func (d *domain) getClusterFromDevice(ctx ConsoleContext, device *entities.ConsoleVPNDevice) (string, error) {
	if device == nil {
		return "", errors.Newf("device is nil")
	}

	if device.ProjectName == nil {
		return "", errors.NewE(errors.Newf("project name is nil"))
	}

	cluster, err := d.getClusterAttachedToProject(ctx, *device.ProjectName)
	if err != nil {
		return "", errors.NewE(err)
	}
	if cluster == nil {
		return "", errors.NewE(errors.Newf("no cluster attached to project %s", *device.ProjectName))
	}
	return *cluster, nil
}

func (d *domain) updateVpnOnCluster(ctx ConsoleContext, ndev, xdev *entities.ConsoleVPNDevice) error {

	ndev.Namespace = d.envVars.DeviceNamespace
	ndev.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &ndev.Device); err != nil {
		return errors.NewE(err)
	}

	if ndev.ProjectName != nil && ndev.EnvironmentName != nil {
		if err := d.applyK8sResource(ctx, *ndev.ProjectName, &ndev.Device, ndev.RecordVersion); err != nil {
			return errors.NewE(err)
		}
	}

	if (xdev.ProjectName != nil) && (*xdev.ProjectName != *ndev.ProjectName) {
		xdev.Spec.Disabled = true
		if err := d.applyK8sResource(ctx, *xdev.ProjectName, &xdev.Device, xdev.RecordVersion); err != nil {
			return errors.NewE(err)
		}
	}

	return nil
}

func (d *domain) ListVPNDevices(ctx ConsoleContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.ConsoleVPNDevice], error) {
	if err := d.canPerformActionInAccount(ctx, iamT.ListVPNDevices); err != nil {
		return nil, errors.NewE(err)
	}

	filter := repos.Filter{"accountName": ctx.AccountName}
	return d.vpnDeviceRepo.FindPaginated(ctx, d.vpnDeviceRepo.MergeMatchFilters(filter, search), pagination)
}

func (d *domain) ListVPNDevicesForUser(ctx ConsoleContext) ([]*entities.ConsoleVPNDevice, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.ListVPNDevices); err != nil {
		return nil, errors.NewE(err)
	}

	return d.vpnDeviceRepo.Find(ctx, repos.Query{
		Filter: repos.Filter{
			"createdBy.userId": ctx.UserId,
		},
	})
}

func (d *domain) GetVPNDevice(ctx ConsoleContext, name string) (*entities.ConsoleVPNDevice, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.GetVPNDevice); err != nil {
		return nil, errors.NewE(err)
	}

	return d.findVPNDevice(ctx, name)
}

func (d *domain) CreateVPNDevice(ctx ConsoleContext, device entities.ConsoleVPNDevice) (*entities.ConsoleVPNDevice, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.CreateVPNDevice); err != nil {
		return nil, errors.NewE(err)
	}

	device.Namespace = d.envVars.DeviceNamespace

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
	device.LinkedClusters = []string{}

	device.SyncStatus = t.GenSyncStatus(t.SyncActionApply, device.RecordVersion)

	if device.ProjectName != nil && device.EnvironmentName != nil {
		s, err := d.envTargetNamespace(ctx, *device.ProjectName, *device.EnvironmentName)
		if err != nil {
			return nil, errors.NewE(err)
		}

		device.Spec.ActiveNamespace = &s

		clusterName, err := d.getClusterFromDevice(ctx, &device)
		if err != nil {
			return nil, errors.NewE(err)
		}

		device.LinkedClusters = append(device.LinkedClusters, clusterName)
	}

	if _, err := d.iamClient.AddMembership(ctx, &iam.AddMembershipIn{
		UserId:       string(ctx.UserId),
		ResourceType: string(iamT.ResourceConsoleVPNDevice),
		ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceConsoleVPNDevice, device.Name),
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

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeVPNDevice, nDevice.Name, PublishAdd)

	if device.ProjectName != nil && device.EnvironmentName != nil {
		d.applyK8sResource(ctx, *nDevice.ProjectName, &nDevice.Device, nDevice.RecordVersion)
	}

	return nDevice, nil
}

func (d *domain) UpdateVPNDevice(ctx ConsoleContext, device entities.ConsoleVPNDevice) (*entities.ConsoleVPNDevice, error) {
	if err := d.canPerformActionInDevice(ctx, iamT.UpdateVPNDevice, device.Name); err != nil {
		return nil, errors.NewE(err)
	}

	device.Namespace = d.envVars.DeviceNamespace
	xdevice, err := d.findVPNDevice(ctx, device.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if device.ProjectName != nil && device.EnvironmentName != nil {
		s, err := d.envTargetNamespace(ctx, *device.ProjectName, *device.EnvironmentName)
		if err != nil {
			return nil, errors.NewE(err)
		}

		device.Spec.ActiveNamespace = &s

		if xdevice.ProjectName != nil && *xdevice.ProjectName != *device.ProjectName {
			clusterName, err := d.getClusterFromDevice(ctx, &device)
			if err != nil {
				return nil, errors.NewE(err)
			}

			// is cluster already linked?
			linked := false
			for _, v := range xdevice.LinkedClusters {
				if v == clusterName {
					linked = true
					break
				}
			}

			if !linked {
				device.LinkedClusters = append(xdevice.LinkedClusters, clusterName)
			}
		}
	}

<<<<<<< HEAD
	patchForUpdate := common.PatchForUpdate(
		ctx,
		&device,
		common.PatchOpts{
			XPatch: repos.Document{
				fc.ConsoleVPNDeviceSpec: device.Spec,
				fields.ProjectName:      device.ProjectName,
				fields.EnvironmentName:  device.EnvironmentName,
			},
		})

	upDevice, err := d.vpnDeviceRepo.Patch(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.MetadataName: device.Name,
	}, patchForUpdate)
=======
	patch := repos.Document{
		"displayName":     device.DisplayName,
		"spec":            device.Spec,
		"projectName":     device.ProjectName,
		"environmentName": device.EnvironmentName,
		"lastUpdatedBy": common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
		"linkedClusters": device.LinkedClusters,
	}
>>>>>>> origin/release-1.0.5

	if err != nil {
		return nil, errors.NewE(err)
	}

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeVPNDevice, device.Name, PublishUpdate)

	if err := d.updateVpnOnCluster(ctx, upDevice, xdevice); err != nil {
		return nil, errors.NewE(err)
	}

	return upDevice, nil
}

func (d *domain) DeleteVPNDevice(ctx ConsoleContext, name string) error {
	if err := d.canPerformActionInDevice(ctx, iamT.DeleteVPNDevice, name); err != nil {
		return errors.NewE(err)
	}

	upDevice, err := d.vpnDeviceRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: name,
		},
		common.PatchForMarkDeletion(),
	)

	if err != nil {
		return errors.NewE(err)
	}

<<<<<<< HEAD
	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeVPNDevice, name, PublishUpdate)

	if err := d.deleteK8sResource(ctx, *upDevice.ProjectName, &upDevice.Device); err != nil {
		return errors.NewE(err)
=======
	if device.IsMarkedForDeletion() {
		return errors.Newf("vpnDevice %q is already marked for deletion", name)
	}

	if _, err := d.vpnDeviceRepo.PatchById(ctx, device.Id, repos.Document{
		"markedForDeletion": fn.New(true),
		"lastUpdatedBy": common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
		"syncStatus.lastSyncedAt": time.Now(),
		"syncStatus.action":       t.SyncActionDelete,
		"syncStatus.state":        t.SyncStateInQueue,
	}); err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishVpnDeviceEvent(device, PublishDelete)

	for _, v := range device.LinkedClusters {
		if err := d.deleteK8sResourceOfCluster(ctx, v, &device.Device); err != nil {
			return errors.NewE(err)
		}
>>>>>>> origin/release-1.0.5
	}

	return nil
}

func (d *domain) UpdateVpnDevicePorts(ctx ConsoleContext, devName string, ports []*wgv1.Port) error {

	if err := d.canPerformActionInDevice(ctx, iamT.UpdateVPNDevice, devName); err != nil {
		return errors.NewE(err)
	}

	xdevice, err := d.findVPNDevice(ctx, devName)
	if err != nil {
		return errors.NewE(err)
	}

	var prt []wgv1.Port
	for _, p := range ports {
		if p != nil {
			prt = append(prt, *p)
		}
	}

	nDevice, err := d.vpnDeviceRepo.PatchById(
		ctx,
		xdevice.Id,
		repos.Document{
			fields.LastUpdatedBy: common.CreatedOrUpdatedBy{
				UserId:    ctx.UserId,
				UserName:  ctx.UserName,
				UserEmail: ctx.UserEmail,
			},
			fc.ConsoleVPNDeviceSpecPorts: prt,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeVPNDevice, nDevice.Name, PublishUpdate)

	if err := d.updateVpnOnCluster(ctx, nDevice, xdevice); err != nil {
		return errors.NewE(err)
	}

	return nil
}

func (d *domain) UpdateVpnDeviceEnvironment(ctx ConsoleContext, devName string, projectName string, envName string) error {
	xdevice, err := d.findVPNDevice(ctx, devName)
	if err != nil {
		return errors.NewE(err)
	}

	envNamesapce, err := d.envTargetNamespace(ctx, projectName, envName)
	if err != nil {
		return errors.NewE(err)
	}

<<<<<<< HEAD
	ndevice, err := d.vpnDeviceRepo.PatchById(
		ctx,
		xdevice.Id,
		repos.Document{
			fields.ProjectName:                     projectName,
			fields.EnvironmentName:                 envName,
			fc.ConsoleVPNDeviceSpecActiveNamespace: envNamesapce,
			fields.LastUpdatedBy: common.CreatedOrUpdatedBy{
				UserId:    ctx.UserId,
				UserName:  ctx.UserName,
				UserEmail: ctx.UserEmail,
			},
		},
	)
=======
	linkedClusters := xdevice.LinkedClusters

	if xdevice.ProjectName != nil && *xdevice.ProjectName != projectName {
		clusterName, err := d.getClusterAttachedToProject(ctx, projectName)
		if err != nil {
			return errors.NewE(err)
		}

		// is cluster already linked?
		linked := false
		for _, v := range xdevice.LinkedClusters {
			if v == *clusterName {
				linked = true
				break
			}
		}

		if !linked {
			linkedClusters = append(xdevice.LinkedClusters, *clusterName)
		}
	}

	patch := repos.Document{
		"projectName":          projectName,
		"environmentName":      envName,
		"spec.activeNamespace": envNamesapce,
		"lastUpdatedBy": common.CreatedOrUpdatedBy{
			UserId:    ctx.UserId,
			UserName:  ctx.UserName,
			UserEmail: ctx.UserEmail,
		},
		"linkedClusters": linkedClusters,
	}
>>>>>>> origin/release-1.0.5

	if err != nil {
		return errors.NewE(err)
	}
<<<<<<< HEAD
	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeVPNDevice, ndevice.Name, PublishUpdate)
=======

	d.resourceEventPublisher.PublishVpnDeviceEvent(ndevice, PublishUpdate)
>>>>>>> origin/release-1.0.5

	if err := d.updateVpnOnCluster(ctx, ndevice, xdevice); err != nil {
		return errors.NewE(err)
	}

	return nil
}

func (d *domain) OnVPNDeviceUpdateMessage(ctx ConsoleContext, device entities.ConsoleVPNDevice, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	xdevice, err := d.findVPNDevice(ctx, device.Name)

	if err != nil {
		return errors.NewE(err)
	}

	if err := d.MatchRecordVersion(device.Annotations, xdevice.RecordVersion); err != nil {
		if xdevice.ProjectName != nil {
			return d.resyncK8sResource(ctx, *xdevice.ProjectName, xdevice.SyncStatus.Action, &xdevice.Device, xdevice.RecordVersion)
		}
	}

	upDevice, err := d.vpnDeviceRepo.PatchById(
		ctx,
		xdevice.Id,
		common.PatchForSyncFromAgent(
			&device,
			status,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
				XPatch: repos.Document{
					fc.ConsoleVPNDeviceWireguardConfig: device.WireguardConfig,
				},
			}))

	if err != nil {
		return errors.NewE(err)
	}

<<<<<<< HEAD
	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeVPNDevice, upDevice.Name, PublishUpdate)

	return nil
}

func (d *domain) OnVPNDeviceDeleteMessage(ctx ConsoleContext, device entities.ConsoleVPNDevice) error {
	err := d.vpnDeviceRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: device.Name,
		},
	)

	if err != nil {
=======
	// ensure that the device is not linked to any cluster, if it is, unlink it
	{
		clusterName, err := d.getClusterFromDevice(ctx, &device)
		if err != nil {
			return errors.NewE(err)
		}

		xdevice.LinkedClusters = func() []string {
			var clusters []string
			for _, c := range xdevice.LinkedClusters {
				if c != clusterName {
					clusters = append(clusters, c)
				}
			}
			return clusters
		}()

		if len(xdevice.LinkedClusters) >= 0 {
			_, err := d.vpnDeviceRepo.PatchById(ctx, xdevice.Id, repos.Document{
				"linkedClusters": xdevice.LinkedClusters,
			})

			if err != nil {
				return errors.NewE(err)
			}

			d.resourceEventPublisher.PublishVpnDeviceEvent(xdevice, PublishUpdate)
			return nil
		}
	}

	if _, err := d.iamClient.RemoveMembership(ctx, &iam.RemoveMembershipIn{
		UserId:      string(ctx.UserId),
		ResourceRef: iamT.NewResourceRef(ctx.AccountName, iamT.ResourceConsoleVPNDevice, device.Name),
	}); err != nil {
		return errors.NewE(err)
	}

	if err = d.vpnDeviceRepo.DeleteById(ctx, xdevice.Id); err != nil {
>>>>>>> origin/release-1.0.5
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeVPNDevice, device.Name, PublishDelete)

	return nil
}

func (d *domain) OnVPNDeviceApplyError(ctx ConsoleContext, errMsg string, name string, opts UpdateAndDeleteOpts) error {
	device, err := d.findVPNDevice(ctx, name)
	if err != nil {
		return errors.NewE(err)
	}

	patch := repos.Document{
		"syncStatus.state":        t.SyncStateErroredAtAgent,
		"syncStatus.lastSyncedAt": opts.MessageTimestamp,
		"syncStatus.error":        errMsg,
	}

	udevice, err := d.vpnDeviceRepo.PatchById(ctx, device.Id, patch)
	if err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishVpnDeviceEvent(udevice, PublishUpdate)

	return errors.NewE(err)
}
