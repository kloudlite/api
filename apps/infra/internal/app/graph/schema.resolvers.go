package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"

	"kloudlite.io/apps/infra/internal/app/graph/generated"
	"kloudlite.io/apps/infra/internal/app/graph/model"
	"kloudlite.io/apps/infra/internal/domain"
	"kloudlite.io/apps/infra/internal/domain/entities"
)

// InfraCreateBYOCCluster is the resolver for the infra_createBYOCCluster field.
func (r *mutationResolver) InfraCreateBYOCCluster(ctx context.Context, cluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	return r.Domain.CreateBYOCCluster(toInfraContext(ctx), cluster)
}

// InfraUpdateBYOCCluster is the resolver for the infra_updateBYOCCluster field.
func (r *mutationResolver) InfraUpdateBYOCCluster(ctx context.Context, cluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	return r.Domain.UpdateBYOCCluster(toInfraContext(ctx), cluster)
}

// InfraDeleteBYOCCluster is the resolver for the infra_deleteBYOCCluster field.
func (r *mutationResolver) InfraDeleteBYOCCluster(ctx context.Context, name string) (bool, error) {
	err := r.Domain.DeleteBYOCCluster(toInfraContext(ctx), name)
	if err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateCluster is the resolver for the infra_createCluster field.
func (r *mutationResolver) InfraCreateCluster(ctx context.Context, cluster entities.Cluster) (*entities.Cluster, error) {
	return r.Domain.CreateCluster(toInfraContext(ctx), cluster)
}

// InfraUpdateCluster is the resolver for the infra_updateCluster field.
func (r *mutationResolver) InfraUpdateCluster(ctx context.Context, cluster entities.Cluster) (*entities.Cluster, error) {
	return r.Domain.UpdateCluster(toInfraContext(ctx), cluster)
}

// InfraDeleteCluster is the resolver for the infra_deleteCluster field.
func (r *mutationResolver) InfraDeleteCluster(ctx context.Context, name string) (bool, error) {
	if err := r.Domain.DeleteCluster(toInfraContext(ctx), name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateCloudProvider is the resolver for the infra_createCloudProvider field.
func (r *mutationResolver) InfraCreateCloudProvider(ctx context.Context, cloudProvider entities.CloudProvider, providerSecret entities.Secret) (*entities.CloudProvider, error) {
	return r.Domain.CreateCloudProvider(toInfraContext(ctx), cloudProvider, providerSecret)
}

// InfraUpdateCloudProvider is the resolver for the infra_updateCloudProvider field.
func (r *mutationResolver) InfraUpdateCloudProvider(ctx context.Context, cloudProvider entities.CloudProvider, providerSecret *entities.Secret) (*entities.CloudProvider, error) {
	return r.Domain.UpdateCloudProvider(toInfraContext(ctx), cloudProvider, providerSecret)
}

// InfraDeleteCloudProvider is the resolver for the infra_deleteCloudProvider field.
func (r *mutationResolver) InfraDeleteCloudProvider(ctx context.Context, name string) (bool, error) {
	if err := r.Domain.DeleteCloudProvider(toInfraContext(ctx), name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateEdge is the resolver for the infra_createEdge field.
func (r *mutationResolver) InfraCreateEdge(ctx context.Context, edge entities.Edge) (*entities.Edge, error) {
	return r.Domain.CreateEdge(toInfraContext(ctx), edge)
}

// InfraUpdateEdge is the resolver for the infra_updateEdge field.
func (r *mutationResolver) InfraUpdateEdge(ctx context.Context, edge entities.Edge) (*entities.Edge, error) {
	return r.Domain.UpdateEdge(toInfraContext(ctx), edge)
}

// InfraDeleteEdge is the resolver for the infra_deleteEdge field.
func (r *mutationResolver) InfraDeleteEdge(ctx context.Context, clusterName string, name string) (bool, error) {
	if err := r.Domain.DeleteEdge(toInfraContext(ctx), clusterName, name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraDeleteWorkerNode is the resolver for the infra_deleteWorkerNode field.
func (r *mutationResolver) InfraDeleteWorkerNode(ctx context.Context, clusterName string, edgeName string, name string) (bool, error) {
	return r.Domain.DeleteWorkerNode(toInfraContext(ctx), clusterName, edgeName, name)
}

// InfraCheckNameAvailability is the resolver for the infra_checkNameAvailability field.
func (r *queryResolver) InfraCheckNameAvailability(ctx context.Context, resType domain.ResType, name string) (*domain.CheckNameAvailabilityOutput, error) {
	return r.Domain.CheckNameAvailability(toInfraContext(ctx), resType, name)
}

// InfraListBYOCClusters is the resolver for the infra_listBYOCClusters field.
func (r *queryResolver) InfraListBYOCClusters(ctx context.Context) (*model.BYOCClusterPaginatedRecords, error) {
	pClusters, err := r.Domain.ListBYOCClusters(toInfraContext(ctx))
	if err != nil {
		return nil, err
	}

	ae := make([]*model.BYOCClusterEdge, len(pClusters.Edges))
	for i := range pClusters.Edges {
		ae[i] = &model.BYOCClusterEdge{
			Node:   pClusters.Edges[i].Node,
			Cursor: pClusters.Edges[i].Cursor,
		}
	}

	m := model.BYOCClusterPaginatedRecords{
		Edges: ae,
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

// InfraGetBYOCCluster is the resolver for the infra_getBYOCCluster field.
func (r *queryResolver) InfraGetBYOCCluster(ctx context.Context, name string) (*entities.BYOCCluster, error) {
	return r.Domain.GetBYOCCluster(toInfraContext(ctx), name)
}

// InfraListClusters is the resolver for the infra_listClusters field.
func (r *queryResolver) InfraListClusters(ctx context.Context) (*model.ClusterPaginatedRecords, error) {
	pClusters, err := r.Domain.ListClusters(toInfraContext(ctx))
	if err != nil {
		return nil, err
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
	return r.Domain.GetCluster(toInfraContext(ctx), name)
}

// InfraListCloudProviders is the resolver for the infra_listCloudProviders field.
func (r *queryResolver) InfraListCloudProviders(ctx context.Context) (*model.CloudProviderPaginatedRecords, error) {
	pCloudProviders, err := r.Domain.ListCloudProviders(toInfraContext(ctx))
	if err != nil {
		return nil, err
	}

	cpe := make([]*model.CloudProviderEdge, len(pCloudProviders.Edges))
	for i := range pCloudProviders.Edges {
		cpe[i] = &model.CloudProviderEdge{
			Node:   pCloudProviders.Edges[i].Node,
			Cursor: pCloudProviders.Edges[i].Cursor,
		}
	}

	m := model.CloudProviderPaginatedRecords{
		Edges: cpe,
		PageInfo: &model.PageInfo{
			EndCursor:       &pCloudProviders.PageInfo.EndCursor,
			HasNextPage:     pCloudProviders.PageInfo.HasNextPage,
			HasPreviousPage: pCloudProviders.PageInfo.HasPrevPage,
			StartCursor:     &pCloudProviders.PageInfo.StartCursor,
		},
		TotalCount: int(pCloudProviders.TotalCount),
	}

	return &m, nil
}

// InfraGetCloudProvider is the resolver for the infra_getCloudProvider field.
func (r *queryResolver) InfraGetCloudProvider(ctx context.Context, name string) (*entities.CloudProvider, error) {
	return r.Domain.GetCloudProvider(toInfraContext(ctx), name)
}

// InfraListEdges is the resolver for the infra_listEdges field.
func (r *queryResolver) InfraListEdges(ctx context.Context, clusterName string, providerName *string) (*model.EdgePaginatedRecords, error) {
	pEdges, err := r.Domain.ListEdges(toInfraContext(ctx), clusterName, providerName)
	if err != nil {
		return nil, err
	}

	pe := make([]*model.EdgeEdge, len(pEdges.Edges))
	for i := range pEdges.Edges {
		pe[i] = &model.EdgeEdge{
			Node:   pEdges.Edges[i].Node,
			Cursor: pEdges.Edges[i].Cursor,
		}
	}

	m := model.EdgePaginatedRecords{
		Edges: pe,
		PageInfo: &model.PageInfo{
			EndCursor:       &pEdges.PageInfo.EndCursor,
			HasNextPage:     pEdges.PageInfo.HasNextPage,
			HasPreviousPage: pEdges.PageInfo.HasPrevPage,
			StartCursor:     &pEdges.PageInfo.StartCursor,
		},
		TotalCount: int(pEdges.TotalCount),
	}

	return &m, nil

}

// InfraGetEdge is the resolver for the infra_getEdge field.
func (r *queryResolver) InfraGetEdge(ctx context.Context, clusterName string, name string) (*entities.Edge, error) {
	return r.Domain.GetEdge(toInfraContext(ctx), clusterName, name)
}

// InfraListMasterNodes is the resolver for the infra_listMasterNodes field.
func (r *queryResolver) InfraListMasterNodes(ctx context.Context, clusterName string) (*model.MasterNodePaginatedRecords, error) {
	pMasterNodes, err := r.Domain.ListMasterNodes(toInfraContext(ctx), clusterName)
	if err != nil {
		return nil, err
	}

	mne := make([]*model.MasterNodeEdge, len(pMasterNodes.Edges))
	for i := range pMasterNodes.Edges {
		mne[i] = &model.MasterNodeEdge{
			Node:   pMasterNodes.Edges[i].Node,
			Cursor: pMasterNodes.Edges[i].Cursor,
		}
	}

	m := model.MasterNodePaginatedRecords{
		Edges: mne,
		PageInfo: &model.PageInfo{
			EndCursor:       &pMasterNodes.PageInfo.EndCursor,
			HasNextPage:     pMasterNodes.PageInfo.HasNextPage,
			HasPreviousPage: pMasterNodes.PageInfo.HasPrevPage,
			StartCursor:     &pMasterNodes.PageInfo.StartCursor,
		},
		TotalCount: int(pMasterNodes.TotalCount),
	}

	return &m, nil
}

// InfraListWorkerNodes is the resolver for the infra_listWorkerNodes field.
func (r *queryResolver) InfraListWorkerNodes(ctx context.Context, clusterName string, edgeName string) (*model.WorkerNodePaginatedRecords, error) {
	pWorkerNodes, err := r.Domain.ListWorkerNodes(toInfraContext(ctx), clusterName, edgeName)
	if err != nil {
		return nil, err
	}

	wne := make([]*model.WorkerNodeEdge, len(pWorkerNodes.Edges))
	for i := range pWorkerNodes.Edges {
		wne[i] = &model.WorkerNodeEdge{
			Node:   pWorkerNodes.Edges[i].Node,
			Cursor: pWorkerNodes.Edges[i].Cursor,
		}
	}

	m := model.WorkerNodePaginatedRecords{
		Edges: wne,
		PageInfo: &model.PageInfo{
			EndCursor:       &pWorkerNodes.PageInfo.EndCursor,
			HasNextPage:     pWorkerNodes.PageInfo.HasNextPage,
			HasPreviousPage: pWorkerNodes.PageInfo.HasPrevPage,
			StartCursor:     &pWorkerNodes.PageInfo.StartCursor,
		},
		TotalCount: int(pWorkerNodes.TotalCount),
	}

	return &m, nil

}

// InfraListNodePools is the resolver for the infra_listNodePools field.
func (r *queryResolver) InfraListNodePools(ctx context.Context, clusterName string, edgeName string) (*model.NodePoolPaginatedRecords, error) {
	pNodePools, err := r.Domain.ListNodePools(toInfraContext(ctx), clusterName, edgeName)
	if err != nil {
		return nil, err
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
func (r *queryResolver) InfraGetNodePool(ctx context.Context, clusterName string, edgeName string, poolName string) (*entities.NodePool, error) {
	return r.Domain.GetNodePool(toInfraContext(ctx), clusterName, edgeName, poolName)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
