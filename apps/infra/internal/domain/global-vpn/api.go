package global_vpn

import (
	"math"

	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (d *Domain) CreateGlobalVPN(ctx types.InfraContext, gvpn entities.GlobalVPN) (*entities.GlobalVPN, error) {
	return d.createGlobalVPN(ctx, gvpn)
}

const (
	kloudliteGlobalVPNDevice = "kloudlite-global-vpn-device"
)

func (d *Domain) createGlobalVPN(ctx types.InfraContext, gvpn entities.GlobalVPN) (*entities.GlobalVPN, error) {
	if gvpn.CIDR == "" {
		gvpn.CIDR = d.Env.BaseCIDR
	}

	if gvpn.AllocatableCIDRSuffix == 0 {
		gvpn.AllocatableCIDRSuffix = d.Env.AllocatableCIDRSuffix
	}

	if gvpn.NumReservedIPsForNonClusterUse == 0 {
		gvpn.NumReservedIPsForNonClusterUse = d.Env.ClustersOffset * int(math.Pow(2, float64(32-gvpn.AllocatableCIDRSuffix)))
	}

	if gvpn.WgInterface == "" {
		gvpn.WgInterface = "kl0"
	}

	gv, err := d.GlobalVPNRepo.Create(ctx, &gvpn)
	if err != nil {
		return nil, err
	}

	device, err := d.CreateGlobalVPNDevice(ctx, entities.GlobalVPNDevice{
		ObjectMeta: metav1.ObjectMeta{
			Name: kloudliteGlobalVPNDevice,
		},
		ResourceMetadata: common.ResourceMetadata{
			DisplayName:   kloudliteGlobalVPNDevice,
			CreatedBy:     common.CreatedOrUpdatedByKloudlite,
			LastUpdatedBy: common.CreatedOrUpdatedByKloudlite,
		},
		AccountName:    ctx.AccountName,
		GlobalVPNName:  gv.Name,
		PublicEndpoint: nil,
		CreationMethod: kloudliteGlobalVPNDevice,
	})
	if err != nil {
		return nil, err
	}

	return d.GlobalVPNRepo.PatchById(ctx, gv.Id, repos.Document{fc.GlobalVPNKloudliteDeviceName: device.Name, fc.GlobalVPNKloudliteDeviceIpAddr: device.IPAddr})
}

func (d *Domain) EnsureGlobalVPN(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
	gvpn, err := d.GlobalVPNRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.MetadataName: gvpnName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if gvpn != nil {
		return gvpn, nil
	}

	return d.createGlobalVPN(ctx, entities.GlobalVPN{
		ObjectMeta: metav1.ObjectMeta{
			Name: gvpnName,
		},
		ResourceMetadata: common.ResourceMetadata{
			DisplayName:   gvpnName,
			CreatedBy:     common.CreatedOrUpdatedByKloudlite,
			LastUpdatedBy: common.CreatedOrUpdatedByKloudlite,
		},
		AccountName: ctx.AccountName,
	})
}

func (d *Domain) UpdateGlobalVPN(ctx types.InfraContext, cgIn entities.GlobalVPN) (*entities.GlobalVPN, error) {
	return nil, errors.New("not implemented")
}

func (d *Domain) DeleteGlobalVPN(ctx types.InfraContext, name string) error {
	if err := d.CanPerformActionInAccount(ctx, iamT.DeleteCluster); err != nil {
		return errors.NewE(err)
	}

	cCount, err := d.CountClustersInGlobalVPN(ctx, name)
	if err != nil {
		return errors.NewE(err)
	}
	if cCount != 0 {
		return errors.Newf("delete clusters first, aborting cluster group deletion")
	}

	ucg, err := d.GlobalVPNRepo.Patch(
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

	d.ResourceEventPublisher.PublishInfraEvent(ctx, ports.ResourceTypeClusterGroup, ucg.Name, ports.PublishUpdate)
	return nil
}

func (d *Domain) ListGlobalVPN(ctx types.InfraContext, mf map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.GlobalVPN], error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.ListClusters); err != nil {
		return nil, errors.NewE(err)
	}

	f := repos.Filter{
		fields.AccountName: ctx.AccountName,
	}

	pr, err := d.GlobalVPNRepo.FindPaginated(ctx, d.GlobalVPNRepo.MergeMatchFilters(f, mf), pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return pr, nil
}

func (d *Domain) GetGlobalVPN(ctx types.InfraContext, name string) (*entities.GlobalVPN, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.GetCluster); err != nil {
		return nil, errors.NewE(err)
	}

	c, err := d.FindGlobalVPN(ctx, name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return c, nil
}

func (d *Domain) FindGlobalVPN(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
	cg, err := d.GlobalVPNRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.MetadataName: gvpnName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cg == nil {
		return nil, errors.Newf("GlobalVPN with name=%s not found", gvpnName)
	}
	return cg, nil
}
