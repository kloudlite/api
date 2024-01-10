package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"encoding/base64"
	"github.com/kloudlite/api/apps/infra/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/infra/internal/app/graph/model"
	"github.com/kloudlite/api/apps/infra/internal/domain"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/apis/wireguard/v1"
)

// AdminKubeconfig is the resolver for the adminKubeconfig field.
func (r *clusterResolver) AdminKubeconfig(ctx context.Context, obj *entities.Cluster) (*model.EncodedValue, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	s, err := r.Domain.GetClusterAdminKubeconfig(ictx, obj.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if s == nil {
		return nil, errors.Newf("kubeconfig could not be found")
	}

	return &model.EncodedValue{
		Value:    base64.StdEncoding.EncodeToString([]byte(*s)),
		Encoding: "base64",
	}, nil
}

// InfraCreateCluster is the resolver for the infra_createCluster field.
func (r *mutationResolver) InfraCreateCluster(ctx context.Context, cluster entities.Cluster) (*entities.Cluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateCluster(ictx, cluster)
}

// InfraUpdateCluster is the resolver for the infra_updateCluster field.
func (r *mutationResolver) InfraUpdateCluster(ctx context.Context, cluster entities.Cluster) (*entities.Cluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.UpdateCluster(ictx, cluster)
}

// InfraDeleteCluster is the resolver for the infra_deleteCluster field.
func (r *mutationResolver) InfraDeleteCluster(ctx context.Context, name string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteCluster(ictx, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// InfraCreateProviderSecret is the resolver for the infra_createProviderSecret field.
func (r *mutationResolver) InfraCreateProviderSecret(ctx context.Context, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.CreateProviderSecret(ictx, secret)
}

// InfraUpdateProviderSecret is the resolver for the infra_updateProviderSecret field.
func (r *mutationResolver) InfraUpdateProviderSecret(ctx context.Context, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.UpdateProviderSecret(ictx, secret)
}

// InfraDeleteProviderSecret is the resolver for the infra_deleteProviderSecret field.
func (r *mutationResolver) InfraDeleteProviderSecret(ctx context.Context, secretName string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.DeleteProviderSecret(ictx, secretName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// InfraCreateDomainEntry is the resolver for the infra_createDomainEntry field.
func (r *mutationResolver) InfraCreateDomainEntry(ctx context.Context, domainEntry entities.DomainEntry) (*entities.DomainEntry, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateDomainEntry(ictx, domainEntry)
}

// InfraUpdateDomainEntry is the resolver for the infra_updateDomainEntry field.
func (r *mutationResolver) InfraUpdateDomainEntry(ctx context.Context, domainEntry entities.DomainEntry) (*entities.DomainEntry, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateDomainEntry(ictx, domainEntry)
}

// InfraDeleteDomainEntry is the resolver for the infra_deleteDomainEntry field.
func (r *mutationResolver) InfraDeleteDomainEntry(ctx context.Context, domainName string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteDomainEntry(ictx, domainName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// InfraCreateNodePool is the resolver for the infra_createNodePool field.
func (r *mutationResolver) InfraCreateNodePool(ctx context.Context, clusterName string, pool entities.NodePool) (*entities.NodePool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.CreateNodePool(ictx, clusterName, pool)
}

// InfraUpdateNodePool is the resolver for the infra_updateNodePool field.
func (r *mutationResolver) InfraUpdateNodePool(ctx context.Context, clusterName string, pool entities.NodePool) (*entities.NodePool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.UpdateNodePool(ictx, clusterName, pool)
}

// InfraDeleteNodePool is the resolver for the infra_deleteNodePool field.
func (r *mutationResolver) InfraDeleteNodePool(ctx context.Context, clusterName string, poolName string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.DeleteNodePool(ictx, clusterName, poolName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// InfraCreateVPNDevice is the resolver for the infra_createVPNDevice field.
func (r *mutationResolver) InfraCreateVPNDevice(ctx context.Context, clusterName string, vpnDevice entities.VPNDevice) (*entities.VPNDevice, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.CreateVPNDevice(cc, clusterName, vpnDevice)
}

// InfraUpdateVPNDevice is the resolver for the infra_updateVPNDevice field.
func (r *mutationResolver) InfraUpdateVPNDevice(ctx context.Context, clusterName string, vpnDevice entities.VPNDevice) (*entities.VPNDevice, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateVPNDevice(cc, clusterName, vpnDevice)
}

// InfraUpdateVPNDevicePorts is the resolver for the infra_updateVPNDevicePorts field.
func (r *mutationResolver) InfraUpdateVPNDevicePorts(ctx context.Context, clusterName string, deviceName string, ports []*v1.Port) (bool, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.UpdateVpnDevicePorts(cc, clusterName, deviceName, ports); err != nil {
		return false, err
	}

	return true, nil
}

// InfraUpdateVPNDeviceNs is the resolver for the infra_updateVPNDeviceNs field.
func (r *mutationResolver) InfraUpdateVPNDeviceNs(ctx context.Context, clusterName string, deviceName string, namespace string) (bool, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.UpdateVpnDeviceNs(cc, clusterName, deviceName, namespace); err != nil {
		return false, err
	}

	return true, nil
}

// InfraDeleteVPNDevice is the resolver for the infra_deleteVPNDevice field.
func (r *mutationResolver) InfraDeleteVPNDevice(ctx context.Context, clusterName string, deviceName string) (bool, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteVPNDevice(cc, clusterName, deviceName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// InfraCreateClusterManagedService is the resolver for the infra_createClusterManagedService field.
func (r *mutationResolver) InfraCreateClusterManagedService(ctx context.Context, clusterName string, service entities.ClusterManagedService) (*entities.ClusterManagedService, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateClusterManagedService(ictx, clusterName, service)
}

// InfraUpdateClusterManagedService is the resolver for the infra_updateClusterManagedService field.
func (r *mutationResolver) InfraUpdateClusterManagedService(ctx context.Context, clusterName string, service entities.ClusterManagedService) (*entities.ClusterManagedService, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.UpdateClusterManagedService(ictx, clusterName, service)
}

// InfraDeleteClusterManagedService is the resolver for the infra_deleteClusterManagedService field.
func (r *mutationResolver) InfraDeleteClusterManagedService(ctx context.Context, clusterName string, serviceName string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteClusterManagedService(ictx, clusterName, serviceName); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// InfraCreateHelmRelease is the resolver for the infra_createHelmRelease field.
func (r *mutationResolver) InfraCreateHelmRelease(ctx context.Context, clusterName string, release entities.HelmRelease) (*entities.HelmRelease, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateHelmRelease(ictx, clusterName, release)
}

// InfraUpdateHelmRelease is the resolver for the infra_updateHelmRelease field.
func (r *mutationResolver) InfraUpdateHelmRelease(ctx context.Context, clusterName string, release entities.HelmRelease) (*entities.HelmRelease, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateHelmRelease(ictx, clusterName, release)
}

// InfraDeleteHelmRelease is the resolver for the infra_deleteHelmRelease field.
func (r *mutationResolver) InfraDeleteHelmRelease(ctx context.Context, clusterName string, releaseName string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteHelmRelease(ictx, clusterName, releaseName); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCheckNameAvailability is the resolver for the infra_checkNameAvailability field.
func (r *queryResolver) InfraCheckNameAvailability(ctx context.Context, resType domain.ResType, clusterName *string, name string) (*domain.CheckNameAvailabilityOutput, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.CheckNameAvailability(ictx, resType, clusterName, name)
}

// InfraListClusters is the resolver for the infra_listClusters field.
func (r *queryResolver) InfraListClusters(ctx context.Context, search *model.SearchCluster, pagination *repos.CursorPagination) (*model.ClusterPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}

		if search.CloudProviderName != nil {
			filter["spec.cloudProvider"] = *search.CloudProviderName
		}

		if search.Region != nil {
			filter["spec.region"] = *search.Region
		}

		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pClusters, err := r.Domain.ListClusters(ictx, filter, *pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	ce := make([]*model.ClusterEdge, len(pClusters.Edges))
	for i := range pClusters.Edges {
		ce[i] = &model.ClusterEdge{
			Node:   pClusters.Edges[i].Node,
			Cursor: pClusters.Edges[i].Cursor,
		}
	}

	m := model.ClusterPaginatedRecords{
		Edges: ce,
		PageInfo: &model.PageInfo{
			EndCursor:       &pClusters.PageInfo.EndCursor,
			HasNextPage:     pClusters.PageInfo.HasNextPage,
			HasPreviousPage: pClusters.PageInfo.HasPrevPage,
			StartCursor:     &pClusters.PageInfo.StartCursor,
		},
		TotalCount: int(pClusters.TotalCount),
	}

	return &m, nil
}

// InfraGetCluster is the resolver for the infra_getCluster field.
func (r *queryResolver) InfraGetCluster(ctx context.Context, name string) (*entities.Cluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetCluster(ictx, name)
}

// InfraListNodePools is the resolver for the infra_listNodePools field.
func (r *queryResolver) InfraListNodePools(ctx context.Context, clusterName string, search *model.SearchNodepool, pagination *repos.CursorPagination) (*model.NodePoolPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pNodePools, err := r.Domain.ListNodePools(ictx, clusterName, filter, *pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pe := make([]*model.NodePoolEdge, len(pNodePools.Edges))
	for i := range pNodePools.Edges {
		pe[i] = &model.NodePoolEdge{
			Node:   pNodePools.Edges[i].Node,
			Cursor: pNodePools.Edges[i].Cursor,
		}
	}

	m := model.NodePoolPaginatedRecords{
		Edges: pe,
		PageInfo: &model.PageInfo{
			EndCursor:       &pNodePools.PageInfo.EndCursor,
			HasNextPage:     pNodePools.PageInfo.HasNextPage,
			HasPreviousPage: pNodePools.PageInfo.HasPrevPage,
			StartCursor:     &pNodePools.PageInfo.StartCursor,
		},
		TotalCount: int(pNodePools.TotalCount),
	}

	return &m, nil
}

// InfraGetNodePool is the resolver for the infra_getNodePool field.
func (r *queryResolver) InfraGetNodePool(ctx context.Context, clusterName string, poolName string) (*entities.NodePool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetNodePool(ictx, clusterName, poolName)
}

// InfraListProviderSecrets is the resolver for the infra_listProviderSecrets field.
func (r *queryResolver) InfraListProviderSecrets(ctx context.Context, search *model.SearchProviderSecret, pagination *repos.CursorPagination) (*model.CloudProviderSecretPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}

		if search.CloudProviderName != nil {
			filter["cloudProviderName"] = *search.CloudProviderName
		}
	}

	pSecrets, err := r.Domain.ListProviderSecrets(ictx, filter, *pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pe := make([]*model.CloudProviderSecretEdge, len(pSecrets.Edges))
	for i := range pSecrets.Edges {
		pe[i] = &model.CloudProviderSecretEdge{
			Node:   pSecrets.Edges[i].Node,
			Cursor: pSecrets.Edges[i].Cursor,
		}
	}

	m := model.CloudProviderSecretPaginatedRecords{
		Edges: pe,
		PageInfo: &model.PageInfo{
			EndCursor:       &pSecrets.PageInfo.EndCursor,
			HasNextPage:     pSecrets.PageInfo.HasNextPage,
			HasPreviousPage: pSecrets.PageInfo.HasPrevPage,
			StartCursor:     &pSecrets.PageInfo.StartCursor,
		},
		TotalCount: int(pSecrets.TotalCount),
	}

	return &m, nil
}

// InfraGetProviderSecret is the resolver for the infra_getProviderSecret field.
func (r *queryResolver) InfraGetProviderSecret(ctx context.Context, name string) (*entities.CloudProviderSecret, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetProviderSecret(ictx, name)
}

// InfraListDomainEntries is the resolver for the infra_listDomainEntries field.
func (r *queryResolver) InfraListDomainEntries(ctx context.Context, search *model.SearchDomainEntry, pagination *repos.CursorPagination) (*model.DomainEntryPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}

		if search.ClusterName != nil {
			filter["spec.clusterName"] = *search.ClusterName
		}
	}

	dEntries, err := r.Domain.ListDomainEntries(ictx, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	edges := make([]*model.DomainEntryEdge, len(dEntries.Edges))
	for i := range dEntries.Edges {
		edges[i] = &model.DomainEntryEdge{
			Node:   dEntries.Edges[i].Node,
			Cursor: dEntries.Edges[i].Cursor,
		}
	}

	m := model.DomainEntryPaginatedRecords{
		Edges: edges,
		PageInfo: &model.PageInfo{
			EndCursor:       &dEntries.PageInfo.EndCursor,
			HasNextPage:     dEntries.PageInfo.HasNextPage,
			HasPreviousPage: dEntries.PageInfo.HasPrevPage,
			StartCursor:     &dEntries.PageInfo.StartCursor,
		},
		TotalCount: int(dEntries.TotalCount),
	}

	return &m, nil
}

// InfraGetDomainEntry is the resolver for the infra_getDomainEntry field.
func (r *queryResolver) InfraGetDomainEntry(ctx context.Context, domainName string) (*entities.DomainEntry, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetDomainEntry(ictx, domainName)
}

// InfraCheckAwsAccess is the resolver for the infra_checkAwsAccess field.
func (r *queryResolver) InfraCheckAwsAccess(ctx context.Context, cloudproviderName string) (*model.CheckAwsAccessOutput, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	output, err := r.Domain.ValidateProviderSecretAWSAccess(ictx, cloudproviderName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return &model.CheckAwsAccessOutput{
		Result:          output.Result,
		InstallationURL: output.InstallationURL,
	}, nil
}

// InfraListVPNDevices is the resolver for the infra_listVPNDevices field.
func (r *queryResolver) InfraListVPNDevices(ctx context.Context, clusterName *string, search *model.SearchVPNDevices, pq *repos.CursorPagination) (*model.VPNDevicePaginatedRecords, error) {
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	cc, err := toInfraContext(ctx)
	if err != nil {
		if cc.AccountName == "" {
			return nil, errors.NewE(err)
		}
	}

	devices, err := r.Domain.ListVPNDevices(cc, cc.AccountName, clusterName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ve := make([]*model.VPNDeviceEdge, len(devices.Edges))
	for i := range devices.Edges {
		ve[i] = &model.VPNDeviceEdge{
			Node:   devices.Edges[i].Node,
			Cursor: devices.Edges[i].Cursor,
		}
	}

	m := model.VPNDevicePaginatedRecords{
		Edges: ve,
		PageInfo: &model.PageInfo{
			EndCursor:       &devices.PageInfo.EndCursor,
			HasNextPage:     devices.PageInfo.HasNextPage,
			HasPreviousPage: devices.PageInfo.HasPrevPage,
			StartCursor:     &devices.PageInfo.StartCursor,
		},
		TotalCount: int(devices.TotalCount),
	}

	return &m, nil
}

// InfraGetVPNDevice is the resolver for the infra_getVPNDevice field.
func (r *queryResolver) InfraGetVPNDevice(ctx context.Context, clusterName string, name string) (*entities.VPNDevice, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetVPNDevice(cc, clusterName, name)
}

// InfraListClusterManagedServices is the resolver for the infra_listClusterManagedServices field.
func (r *queryResolver) InfraListClusterManagedServices(ctx context.Context, clusterName string, search *model.SearchClusterManagedService, pagination *repos.CursorPagination) (*model.ClusterManagedServicePaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}

		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pClusters, err := r.Domain.ListClusterManagedServices(ictx, clusterName, filter, *pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	ce := make([]*model.ClusterManagedServiceEdge, len(pClusters.Edges))
	for i := range pClusters.Edges {
		ce[i] = &model.ClusterManagedServiceEdge{
			Node:   pClusters.Edges[i].Node,
			Cursor: pClusters.Edges[i].Cursor,
		}
	}

	m := model.ClusterManagedServicePaginatedRecords{
		Edges: ce,
		PageInfo: &model.PageInfo{
			EndCursor:       &pClusters.PageInfo.EndCursor,
			HasNextPage:     pClusters.PageInfo.HasNextPage,
			HasPreviousPage: pClusters.PageInfo.HasPrevPage,
			StartCursor:     &pClusters.PageInfo.StartCursor,
		},
		TotalCount: int(pClusters.TotalCount),
	}

	return &m, nil
}

// InfraGetClusterManagedService is the resolver for the infra_getClusterManagedService field.
func (r *queryResolver) InfraGetClusterManagedService(ctx context.Context, clusterName string, name string) (*entities.ClusterManagedService, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetClusterManagedService(ictx, clusterName, name)
}

// InfraListHelmReleases is the resolver for the infra_listHelmReleases field.
func (r *queryResolver) InfraListHelmReleases(ctx context.Context, clusterName string, search *model.SearchHelmRelease, pagination *repos.CursorPagination) (*model.HelmReleasePaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}

		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pRelease, err := r.Domain.ListHelmReleases(ictx, clusterName, filter, *pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	ce := make([]*model.HelmReleaseEdge, len(pRelease.Edges))
	for i := range pRelease.Edges {
		ce[i] = &model.HelmReleaseEdge{
			Node:   pRelease.Edges[i].Node,
			Cursor: pRelease.Edges[i].Cursor,
		}
	}

	m := model.HelmReleasePaginatedRecords{
		Edges: ce,
		PageInfo: &model.PageInfo{
			EndCursor:       &pRelease.PageInfo.EndCursor,
			HasNextPage:     pRelease.PageInfo.HasNextPage,
			HasPreviousPage: pRelease.PageInfo.HasPrevPage,
			StartCursor:     &pRelease.PageInfo.StartCursor,
		},
		TotalCount: int(pRelease.TotalCount),
	}

	return &m, nil
}

// InfraGetHelmRelease is the resolver for the infra_getHelmRelease field.
func (r *queryResolver) InfraGetHelmRelease(ctx context.Context, clusterName string, name string) (*entities.HelmRelease, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetHelmRelease(ictx, clusterName, name)
}

// InfraListManagedServiceTemplates is the resolver for the infra_listManagedServiceTemplates field.
func (r *queryResolver) InfraListManagedServiceTemplates(ctx context.Context) ([]*entities.MsvcTemplate, error) {
	return r.Domain.ListManagedSvcTemplates()
}

// InfraGetManagedServiceTemplate is the resolver for the infra_getManagedServiceTemplate field.
func (r *queryResolver) InfraGetManagedServiceTemplate(ctx context.Context, category string, name string) (*entities.MsvcTemplateEntry, error) {
	return r.Domain.GetManagedSvcTemplate(category, name)
}

// InfraListPVCs is the resolver for the infra_listPVCs field.
func (r *queryResolver) InfraListPVCs(ctx context.Context, clusterName string, search *model.SearchPersistentVolumeClaims, pq *repos.CursorPagination) (*model.PersistentVolumeClaimPaginatedRecords, error) {
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pvcs, err := r.Domain.ListPVCs(cc, clusterName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ve := make([]*model.PersistentVolumeClaimEdge, len(pvcs.Edges))
	for i := range pvcs.Edges {
		ve[i] = &model.PersistentVolumeClaimEdge{
			Node:   pvcs.Edges[i].Node,
			Cursor: pvcs.Edges[i].Cursor,
		}
	}

	m := model.PersistentVolumeClaimPaginatedRecords{
		Edges: ve,
		PageInfo: &model.PageInfo{
			EndCursor:       &pvcs.PageInfo.EndCursor,
			HasNextPage:     pvcs.PageInfo.HasNextPage,
			HasPreviousPage: pvcs.PageInfo.HasPrevPage,
			StartCursor:     &pvcs.PageInfo.StartCursor,
		},
		TotalCount: int(pvcs.TotalCount),
	}

	return &m, nil
}

// InfraGetPvc is the resolver for the infra_getPVC field.
func (r *queryResolver) InfraGetPvc(ctx context.Context, clusterName string, name string) (*entities.PersistentVolumeClaim, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetPVC(cc, clusterName, name)
}

// InfraListNamespaces is the resolver for the infra_listNamespaces field.
func (r *queryResolver) InfraListNamespaces(ctx context.Context, clusterName string, search *model.SearchNamespaces, pq *repos.CursorPagination) (*model.NamespacePaginatedRecords, error) {
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	namespaces, err := r.Domain.ListNamespaces(cc, clusterName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ve := make([]*model.NamespaceEdge, len(namespaces.Edges))
	for i := range namespaces.Edges {
		ve[i] = &model.NamespaceEdge{
			Node:   namespaces.Edges[i].Node,
			Cursor: namespaces.Edges[i].Cursor,
		}
	}

	m := model.NamespacePaginatedRecords{
		Edges: ve,
		PageInfo: &model.PageInfo{
			EndCursor:       &namespaces.PageInfo.EndCursor,
			HasNextPage:     namespaces.PageInfo.HasNextPage,
			HasPreviousPage: namespaces.PageInfo.HasPrevPage,
			StartCursor:     &namespaces.PageInfo.StartCursor,
		},
		TotalCount: int(namespaces.TotalCount),
	}

	return &m, nil
}

// InfraGetNamespace is the resolver for the infra_getNamespace field.
func (r *queryResolver) InfraGetNamespace(ctx context.Context, clusterName string, name string) (*entities.Namespace, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetNamespace(cc, clusterName, name)
}

// InfraListPVs is the resolver for the infra_listPVs field.
func (r *queryResolver) InfraListPVs(ctx context.Context, clusterName string, search *model.SearchPersistentVolumes, pq *repos.CursorPagination) (*model.PersistentVolumePaginatedRecords, error) {
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pvs, err := r.Domain.ListPVs(cc, clusterName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ve := make([]*model.PersistentVolumeEdge, len(pvs.Edges))
	for i := range pvs.Edges {
		ve[i] = &model.PersistentVolumeEdge{
			Node:   pvs.Edges[i].Node,
			Cursor: pvs.Edges[i].Cursor,
		}
	}

	m := model.PersistentVolumePaginatedRecords{
		Edges: ve,
		PageInfo: &model.PageInfo{
			EndCursor:       &pvs.PageInfo.EndCursor,
			HasNextPage:     pvs.PageInfo.HasNextPage,
			HasPreviousPage: pvs.PageInfo.HasPrevPage,
			StartCursor:     &pvs.PageInfo.StartCursor,
		},
		TotalCount: int(pvs.TotalCount),
	}

	return &m, nil
}

// InfraGetPv is the resolver for the infra_getPV field.
func (r *queryResolver) InfraGetPv(ctx context.Context, clusterName string, name string) (*entities.PersistentVolume, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetPV(cc, clusterName, name)
}

// InfraListVolumeAttachments is the resolver for the infra_listVolumeAttachments field.
func (r *queryResolver) InfraListVolumeAttachments(ctx context.Context, clusterName string, search *model.SearchVolumeAttachments, pq *repos.CursorPagination) (*model.VolumeAttachmentPaginatedRecords, error) {
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	volatt, err := r.Domain.ListVolumeAttachments(cc, clusterName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ve := make([]*model.VolumeAttachmentEdge, len(volatt.Edges))
	for i := range volatt.Edges {
		ve[i] = &model.VolumeAttachmentEdge{
			Node:   volatt.Edges[i].Node,
			Cursor: volatt.Edges[i].Cursor,
		}
	}

	m := model.VolumeAttachmentPaginatedRecords{
		Edges: ve,
		PageInfo: &model.PageInfo{
			EndCursor:       &volatt.PageInfo.EndCursor,
			HasNextPage:     volatt.PageInfo.HasNextPage,
			HasPreviousPage: volatt.PageInfo.HasPrevPage,
			StartCursor:     &volatt.PageInfo.StartCursor,
		},
		TotalCount: int(volatt.TotalCount),
	}

	return &m, nil
}

// InfraGetVolumeAttachment is the resolver for the infra_getVolumeAttachment field.
func (r *queryResolver) InfraGetVolumeAttachment(ctx context.Context, clusterName string, name string) (*entities.VolumeAttachment, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetVolumeAttachment(cc, clusterName, name)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
