package byok_clusters

import (
	"encoding/base64"
	"fmt"

	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
)

var ErrClusterNotFound error = fmt.Errorf("cluster not found")

func (d *Domain) clusterAlreadyExists(ctx types.InfraContext, name string) (*bool, error) {
	exists, err := d.ClusterAlreadyExists(ctx, name)
	if err != nil {
		return nil, err
	}
	if exists {
		return fn.New(true), nil
	}

	existsBYOK, err := d.BYOKClusterRepo.FindOne(ctx, repos.Filter{
		fc.AccountName:  ctx.AccountName,
		fc.MetadataName: name,
	})
	if err != nil {
		return nil, err
	}

	if existsBYOK != nil {
		return fn.New(true), nil
	}

	return fn.New(false), nil
}

func (d *Domain) CreateBYOKCluster(ctx types.InfraContext, cluster entities.BYOKCluster) (*entities.BYOKCluster, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.CreateCluster); err != nil {
		return nil, errors.NewE(err)
	}

	exists, err := d.clusterAlreadyExists(ctx, cluster.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if exists != nil && *exists {
		return nil, errors.Newf("cluster/byok cluster with name (%s) already exists", cluster.Name)
	}

	accNs, err := d.GetAccNamespace(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	cluster.Namespace = accNs

	if cluster.GlobalVPN == "" {
		cluster.GlobalVPN = d.Env.DefaultGlobalVPN
	}

	ctoken, err := d.MsgOfficeSvc.GenerateClusterToken(ctx, ctx.AccountName, cluster.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	cluster.ClusterToken = ctoken

	cluster.MessageQueueTopicName = common.GetTenantClusterMessagingTopic(ctx.AccountName, cluster.Name)

	if _, err := d.EnsureGlobalVPN(ctx, cluster.GlobalVPN); err != nil {
		return nil, errors.NewE(err)
	}

	// clusterSvcCIDR, err := d.claimNextClusterSvcCIDR(ctx, cluster.Name, gvpn.Name)
	// if err != nil {
	// 	return nil, err
	// }

	gvpnConn, err := d.EnsureGlobalVPNConnection(ctx, cluster.Name, cluster.GlobalVPN, cluster.ClusterPublicEndpoint)
	if err != nil {
		return nil, errors.NewE(err)
	}

	cluster.ClusterSvcCIDR = gvpnConn.ClusterSvcCIDR

	cluster.AccountName = ctx.AccountName

	cluster.IncrementRecordVersion()
	cluster.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	cluster.LastUpdatedBy = cluster.CreatedBy

	cluster.SyncStatus = t.GenSyncStatus(t.SyncActionApply, 0)

	nCluster, err := d.BYOKClusterRepo.Create(ctx, &cluster)
	if err != nil {
		if d.BYOKClusterRepo.ErrAlreadyExists(err) {
			return nil, errors.NewEf(err, "cluster with name %q already exists in namespace %q", cluster.Name, cluster.Namespace)
		}
		return nil, errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishInfraEvent(ctx, ports.ResourceTypeCluster, nCluster.Name, ports.PublishAdd)

	return nCluster, nil
}

func (d *Domain) UpdateBYOKCluster(ctx types.InfraContext, clusterName string, displayName string) (*entities.BYOKCluster, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.UpdateCluster); err != nil {
		return nil, errors.NewE(err)
	}

	updated, err := d.BYOKClusterRepo.PatchOne(ctx, entities.UniqueBYOKClusterFilter(ctx.AccountName, clusterName), repos.Document{
		fc.DisplayName: displayName,
	})
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (d *Domain) ListBYOKCluster(ctx types.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.BYOKCluster], error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.ListClusters); err != nil {
		return nil, errors.NewE(err)
	}

	pRecords, err := d.BYOKClusterRepo.FindPaginated(ctx, d.BYOKClusterRepo.MergeMatchFilters(entities.ListBYOKClusterFilter(ctx.AccountName), search), pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return pRecords, nil
}

func (d *Domain) GetBYOKCluster(ctx types.InfraContext, name string) (*entities.BYOKCluster, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.GetCluster); err != nil {
		return nil, errors.NewE(err)
	}

	c, err := d.findBYOKCluster(ctx, name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return c, nil
}

func (d *Domain) GetBYOKClusterSetupInstructions(ctx types.InfraContext, name string) (*string, error) {
	cluster, err := d.findBYOKCluster(ctx, name)
	if err != nil {
		return nil, err
	}
	return fn.New(fmt.Sprintf(`helm upgrade --install kloudlite --namespace kloudlite --create-namespace kloudlite/kloudlite-agent --set accountName="%s" --set clusterName="%s" --set clusterToken="%s" --set messageOfficeGRPCAddr="%s" --set kloudliteRelease="%s" --set byok.enabled=true --set helmCharts.ingressNginx.enabled=true --set helmCharts.certManager.enabled=true`, ctx.AccountName, name, cluster.ClusterToken, d.Env.MessageOfficeExternalGrpcAddr, d.Env.KloudliteRelease)), nil
}

func (d *Domain) DeleteBYOKCluster(ctx types.InfraContext, name string) error {
	if err := d.CanPerformActionInAccount(ctx, iamT.DeleteCluster); err != nil {
		return errors.NewE(err)
	}

	cluster, err := d.findBYOKCluster(ctx, name)
	if err != nil {
		return errors.NewE(err)
	}

	if err := d.BYOKClusterRepo.DeleteOne(ctx, entities.UniqueBYOKClusterFilter(ctx.AccountName, name)); err != nil {
		return errors.NewE(err)
	}

	if cluster.GlobalVPN != "" {
		if err := d.DeleteGlobalVPNConnection(ctx, cluster.Name, cluster.GlobalVPN); err != nil {
			return errors.NewE(err)
		}
		// if err := d.claimClusterSvcCIDRRepo.DeleteOne(ctx, repos.Filter{
		// 	fc.ClaimClusterSvcCIDRClaimedByCluster: cluster.Name,
		// 	fc.AccountName:                         ctx.AccountName,
		// 	fc.ClaimClusterSvcCIDRGlobalVPNName:    cluster.GlobalVPN,
		// }); err != nil {
		// 	return errors.NewE(err)
		// }
		//
		// if _, err := d.freeClusterSvcCIDRRepo.Create(ctx, &entities.FreeClusterSvcCIDR{
		// 	AccountName:    ctx.AccountName,
		// 	GlobalVPNName:  cluster.GlobalVPN,
		// 	ClusterSvcCIDR: cluster.ClusterSvcCIDR,
		// }); err != nil {
		// 	return errors.NewE(err)
		// }
	}

	return nil
}

func (d *Domain) findBYOKCluster(ctx types.InfraContext, clusterName string) (*entities.BYOKCluster, error) {
	cluster, err := d.BYOKClusterRepo.FindOne(ctx, entities.UniqueBYOKClusterFilter(ctx.AccountName, clusterName))
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cluster == nil {
		return nil, ErrClusterNotFound
	}
	return cluster, nil
}

func (d *Domain) UpsertBYOKClusterKubeconfig(ctx types.InfraContext, clusterName string, kubeconfig []byte) error {
	byokCluster, err := d.findBYOKCluster(ctx, clusterName)
	if err != nil {
		return err
	}

	if _, err := d.BYOKClusterRepo.PatchById(ctx, byokCluster.Id, repos.Document{
		fc.BYOKClusterKubeconfig: t.EncodedString{
			Value:    base64.StdEncoding.EncodeToString(kubeconfig),
			Encoding: "base64",
		},
	}); err != nil {
		return err
	}

	return nil
}

func (d *Domain) isBYOKCluster(ctx types.InfraContext, name string) bool {
	cluster, err := d.findBYOKCluster(ctx, name)
	if err != nil {
		return false
	}

	return cluster != nil
}
