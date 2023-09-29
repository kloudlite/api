package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"

	"kloudlite.io/apps/infra/internal/app/graph/generated"
	"kloudlite.io/apps/infra/internal/app/graph/model"
	"kloudlite.io/apps/infra/internal/domain"
	"kloudlite.io/apps/infra/internal/entities"
	"kloudlite.io/pkg/repos"
)

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

// InfraCreateBYOCCluster is the resolver for the infra_createBYOCCluster field.
func (r *mutationResolver) InfraCreateBYOCCluster(ctx context.Context, byocCluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	return r.Domain.CreateBYOCCluster(toInfraContext(ctx), byocCluster)
}

// InfraUpdateBYOCCluster is the resolver for the infra_updateBYOCCluster field.
func (r *mutationResolver) InfraUpdateBYOCCluster(ctx context.Context, byocCluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	return r.Domain.UpdateBYOCCluster(toInfraContext(ctx), byocCluster)
}

// InfraDeleteBYOCCluster is the resolver for the infra_deleteBYOCCluster field.
func (r *mutationResolver) InfraDeleteBYOCCluster(ctx context.Context, name string) (bool, error) {
	if err := r.Domain.DeleteBYOCCluster(toInfraContext(ctx), name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateProviderSecret is the resolver for the infra_createProviderSecret field.
func (r *mutationResolver) InfraCreateProviderSecret(ctx context.Context, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error) {
	return r.Domain.CreateProviderSecret(toInfraContext(ctx), secret)
}

// InfraUpdateProviderSecret is the resolver for the infra_updateProviderSecret field.
func (r *mutationResolver) InfraUpdateProviderSecret(ctx context.Context, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error) {
	return r.Domain.UpdateProviderSecret(toInfraContext(ctx), secret)
}

// InfraDeleteProviderSecret is the resolver for the infra_deleteProviderSecret field.
func (r *mutationResolver) InfraDeleteProviderSecret(ctx context.Context, secretName string) (bool, error) {
	if err := r.Domain.DeleteProviderSecret(toInfraContext(ctx), secretName); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateNodePool is the resolver for the infra_createNodePool field.
func (r *mutationResolver) InfraCreateNodePool(ctx context.Context, clusterName string, pool entities.NodePool) (*entities.NodePool, error) {
	return r.Domain.CreateNodePool(toInfraContext(ctx), clusterName, pool)
}

// InfraUpdateNodePool is the resolver for the infra_updateNodePool field.
func (r *mutationResolver) InfraUpdateNodePool(ctx context.Context, clusterName string, pool entities.NodePool) (*entities.NodePool, error) {
	return r.Domain.UpdateNodePool(toInfraContext(ctx), clusterName, pool)
}

// InfraDeleteNodePool is the resolver for the infra_deleteNodePool field.
func (r *mutationResolver) InfraDeleteNodePool(ctx context.Context, clusterName string, poolName string) (bool, error) {
	if err := r.Domain.DeleteNodePool(toInfraContext(ctx), clusterName, poolName); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCheckNameAvailability is the resolver for the infra_checkNameAvailability field.
func (r *queryResolver) InfraCheckNameAvailability(ctx context.Context, resType domain.ResType, clusterName *string, name string) (*domain.CheckNameAvailabilityOutput, error) {
	return r.Domain.CheckNameAvailability(toInfraContext(ctx), resType, clusterName, name)
}

// InfraListClusters is the resolver for the infra_listClusters field.
func (r *queryResolver) InfraListClusters(ctx context.Context, search *model.SearchCluster, pagination *repos.CursorPagination) (*model.ClusterPaginatedRecords, error) {
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

	pClusters, err := r.Domain.ListClusters(toInfraContext(ctx), filter, *pagination)
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

// InfraListBYOCClusters is the resolver for the infra_listBYOCClusters field.
func (r *queryResolver) InfraListBYOCClusters(ctx context.Context, search *model.SearchCluster, pagination *repos.CursorPagination) (*model.BYOCClusterPaginatedRecords, error) {
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

	pClusters, err := r.Domain.ListBYOCClusters(toInfraContext(ctx), filter, *pagination)
	if err != nil {
		return nil, err
	}

	bce := make([]*model.BYOCClusterEdge, len(pClusters.Edges))
	for i := range pClusters.Edges {
		bce[i] = &model.BYOCClusterEdge{
			Node:   pClusters.Edges[i].Node,
			Cursor: pClusters.Edges[i].Cursor,
		}
	}

	m := model.BYOCClusterPaginatedRecords{
		Edges: bce,
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

// InfraListNodePools is the resolver for the infra_listNodePools field.
func (r *queryResolver) InfraListNodePools(ctx context.Context, clusterName string, search *model.SearchNodepool, pagination *repos.CursorPagination) (*model.NodePoolPaginatedRecords, error) {
	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pNodePools, err := r.Domain.ListNodePools(toInfraContext(ctx), clusterName, filter, *pagination)
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
func (r *queryResolver) InfraGetNodePool(ctx context.Context, clusterName string, poolName string) (*entities.NodePool, error) {
	return r.Domain.GetNodePool(toInfraContext(ctx), clusterName, poolName)
}

// InfraListProviderSecrets is the resolver for the infra_listProviderSecrets field.
func (r *queryResolver) InfraListProviderSecrets(ctx context.Context, search *model.SearchProviderSecret, pagination *repos.CursorPagination) (*model.CloudProviderSecretPaginatedRecords, error) {
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

	pSecrets, err := r.Domain.ListProviderSecrets(toInfraContext(ctx), filter, *pagination)
	if err != nil {
		return nil, err
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
	return r.Domain.GetProviderSecret(toInfraContext(ctx), name)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
