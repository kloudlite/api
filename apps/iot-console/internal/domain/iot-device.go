package domain

import (
	"context"
	"fmt"
	"sync"

	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	fc "github.com/kloudlite/api/apps/iot-console/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	message_office_internal "github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/message-office-internal"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/seancfoley/ipaddress-go/ipaddr"
)

func getRemoteDeviceIp(deviceOffcet int64, IpBase string) ([]byte, error) {
	deviceRange := ipaddr.NewIPAddressString(fmt.Sprintf("%s/16", IpBase))

	if address, addressError := deviceRange.ToAddress(); addressError == nil {
		increment := address.Increment(deviceOffcet)
		return []byte(ipaddr.NewIPAddressString(increment.GetNetIP().String()).String()), nil
	} else {
		return nil, addressError
	}
}

type lck struct {
	mu    *sync.Mutex
	count int
}
type mlock map[string]*lck

var (
	createmu = make(mlock)
)

func (c mlock) lock(key string) func() {
	if c[key] == nil {
		c[key] = &lck{
			mu:    &sync.Mutex{},
			count: 0,
		}
	}

	c[key].count++

	c[key].mu.Lock()
	return func() {
		c[key].count--
		c[key].mu.Unlock()

		if c[key].count == 0 {
			delete(c, key)
		}
	}
}

func getCidrRanges(index int) (*string, error) {
	if index < 2 || index > 255 {
		return nil, fmt.Errorf("ip range can only be between 2 and 255")
	}

	return functions.New(fmt.Sprintf("10.%d.0.0/16", index)), nil
}

// TODO: IAM Checks needs to be implemented

func (d *domain) generateClusterMoToken(ctx IotResourceContext, dev entities.IOTDevice) (*string, error) {
	gcto, err := d.messageOfficeInternalClient.GenerateClusterToken(ctx, &message_office_internal.GenerateClusterTokenIn{
		AccountName: ctx.AccountName,
		ClusterName: dev.GetClusterName(),
	})
	if err != nil {
		return nil, err
	}

	return &gcto.ClusterToken, nil
}

func (d *domain) findDevices(ctx IotResourceContext, deploymentName string) ([]*entities.IOTDevice, error) {
	filter := ctx.IOTConsoleDBFilters()
	filter.Add(fc.IOTDeviceDeploymentName, deploymentName)
	devs, err := d.iotDeviceRepo.Find(ctx, repos.Query{
		Filter: filter,
		Sort:   map[string]interface{}{fc.IOTDeviceIndex: 1},
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	return devs, nil
}

func (d *domain) findDevice(ctx IotResourceContext, name string, deploymentName string) (*entities.IOTDevice, error) {
	filter := ctx.IOTConsoleDBFilters()
	filter.Add(fc.IOTDeviceDeploymentName, deploymentName)
	filter.Add(fc.IOTDeviceName, name)
	dev, err := d.iotDeviceRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.NewE(err)
	}
	if dev == nil {
		return nil, errors.Newf("no device with name=%q found", name)
	}
	return dev, nil
}

func (d *domain) ListDevices(ctx IotResourceContext, deploymentName string, search map[string]repos.MatchFilter, pq repos.CursorPagination) (*repos.PaginatedRecord[*entities.IOTDevice], error) {
	filter := ctx.IOTConsoleDBFilters()
	filter.Add(fc.IOTDeviceDeploymentName, deploymentName)
	return d.iotDeviceRepo.FindPaginated(ctx, d.iotDeviceRepo.MergeMatchFilters(filter, search), pq)
}

func (d *domain) GetDevice(ctx IotResourceContext, name string, deploymentName string) (*entities.IOTDevice, error) {
	return d.findDevice(ctx, name, deploymentName)
}

func (d *domain) GetPublicKeyDevice(ctx context.Context, publicKey string) (*entities.DeviceWithServices, error) {
	filter := repos.Filter{
		fc.IOTDevicePublicKey: publicKey,
	}
	dev, err := d.iotDeviceRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.NewE(err)
	}
	if dev == nil {
		return nil, errors.Newf("no device with publickey=%q found", publicKey)
	}

	dep, err := d.iotDeploymentRepo.FindOne(ctx, repos.Filter{
		fc.IOTDeploymentName: dev.DeploymentName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	deviceWithServices := &entities.DeviceWithServices{
		IOTDevice:      dev,
		ExposedDomains: dep.ExposedDomains,
		ExposedIps:     dep.ExposedIps,
	}

	return deviceWithServices, nil
}

func (d *domain) CreateDevice(ctx IotResourceContext, deploymentName string, device entities.IOTDevice) (*entities.IOTDevice, error) {

	// as devices will be dependent on other resources, we need to lock the resource
	unlock := createmu.lock(deploymentName)
	defer unlock()

	devs, err := d.findDevices(ctx, deploymentName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	index := 2

	for _, dev := range devs {
		if dev.Index != index {
			break
		}
		index++
	}

	device.Index = index

	svcCidr, err := getCidrRanges(index)
	if err != nil {
		return nil, errors.NewE(err)
	}

	device.ServiceCIDR = *svcCidr

	ip, err := getRemoteDeviceIp(int64(index), "10.0.0.0")
	if err != nil {
		return nil, errors.NewE(err)
	}

	device.IP = string(ip)

	// TODO: validate for name and public key

	// TODO: check access of account and if have access then only create device

	device.ProjectName = ctx.ProjectName
	device.AccountName = ctx.AccountName
	device.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	device.LastUpdatedBy = device.CreatedBy
	device.DeploymentName = deploymentName

	// TODO: validate device name
	// TODO: Generate Device SVC CIDR
	// TODO: Generate Device Ip for cluster-group

	device.PodCIDR = "10.1.0.0/16"

	ctoken, err := d.generateClusterMoToken(ctx, device)
	if err != nil {
		return nil, errors.NewE(err)
	}

	device.ClusterToken = *ctoken

	nDevice, err := d.iotDeviceRepo.Create(ctx, &device)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return nDevice, nil
}

func (d *domain) UpdateDevice(ctx IotResourceContext, deploymentName string, device entities.IOTDevice) (*entities.IOTDevice, error) {
	patchForUpdate := repos.Document{
		fields.DisplayName: device.DisplayName,
		fields.LastUpdatedBy: common.CreatedOrUpdatedBy{
			UserId:    ctx.GetUserId(),
			UserName:  ctx.GetUserName(),
			UserEmail: ctx.GetUserEmail(),
		},
	}

	patchFilter := ctx.IOTConsoleDBFilters()
	patchFilter.Add(fc.IOTDeviceDeploymentName, deploymentName)
	patchFilter.Add(fc.IOTDeviceName, device.Name)

	upDev, err := d.iotDeviceRepo.Patch(
		ctx,
		patchFilter,
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return upDev, nil
}

func (d *domain) DeleteDevice(ctx IotResourceContext, deploymentName string, name string) error {
	filter := ctx.IOTConsoleDBFilters()
	filter.Add(fc.IOTDeviceDeploymentName, deploymentName)
	filter.Add(fc.IOTDeviceName, name)
	err := d.iotDeviceRepo.DeleteOne(
		ctx,
		filter,
	)
	if err != nil {
		return errors.NewE(err)
	}
	return nil
}
