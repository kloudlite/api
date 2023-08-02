package domain

import (
	"context"
	entities "kloudlite.io/apps/infra/internal/entities"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

type InfraContext struct {
	context.Context
	UserId      repos.ID
	AccountName string
}

type Domain interface {
	CheckNameAvailability(ctx InfraContext, typeArg ResType, name string) (*CheckNameAvailabilityOutput, error)

	// ListBYOCClusters(ctx InfraContext, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities2.BYOCCluster], error)
	// GetBYOCCluster(ctx InfraContext, name string) (*entities2.BYOCCluster, error)

	// CreateBYOCCluster(ctx InfraContext, cluster entities2.BYOCCluster) (*entities2.BYOCCluster, error)
	// UpdateBYOCCluster(ctx InfraContext, cluster entities2.BYOCCluster) (*entities2.BYOCCluster, error)
	// DeleteBYOCCluster(ctx InfraContext, name string) error
	// ResyncBYOCCluster(ctx InfraContext, name string) error

	// OnDeleteBYOCClusterMessage(ctx InfraContext, cluster entities2.BYOCCluster) error
	// OnBYOCClusterHelmUpdates(ctx InfraContext, cluster entities2.BYOCCluster) error

	CreateCluster(ctx InfraContext, cluster entities.Cluster) (*entities.Cluster, error)
	ListClusters(ctx InfraContext, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.Cluster], error)
	GetCluster(ctx InfraContext, name string) (*entities.Cluster, error)

	UpdateCluster(ctx InfraContext, cluster entities.Cluster) (*entities.Cluster, error)
	DeleteCluster(ctx InfraContext, name string) error

	OnDeleteClusterMessage(ctx InfraContext, cluster entities.Cluster) error
	OnUpdateClusterMessage(ctx InfraContext, cluster entities.Cluster) error

	CreateProviderSecret(ctx InfraContext, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error)
	UpdateProviderSecret(ctx InfraContext, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error)
	DeleteProviderSecret(ctx InfraContext, secretName string) error

	ListProviderSecrets(ctx InfraContext, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.CloudProviderSecret], error)
	GetProviderSecret(ctx InfraContext, name string) (*entities.CloudProviderSecret, error)

	// CreateCloudProvider(ctx InfraContext, cloudProvider entities.CloudProvider, providerSecret entities2.CloudProviderSecret) (*entities.CloudProvider, error)
	// ListCloudProviders(ctx InfraContext, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.CloudProvider], error)
	// GetCloudProvider(ctx InfraContext, name string) (*entities.CloudProvider, error)
	// UpdateCloudProvider(ctx InfraContext, cloudProvider entities.CloudProvider, providerSecret *entities2.CloudProviderSecret) (*entities.CloudProvider, error)
	// DeleteCloudProvider(ctx InfraContext, name string) error
	// OnDeleteCloudProviderMessage(ctx InfraContext, cloudProvider entities.CloudProvider) error
	// OnUpdateCloudProviderMessage(ctx InfraContext, cloudProvider entities.CloudProvider) error

	// ListEdges(ctx InfraContext, clusterName string, providerName *string, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.Edge], error)
	// GetEdge(ctx InfraContext, clusterName string, name string) (*entities.Edge, error)
	//
	// CreateEdge(ctx InfraContext, edge entities.Edge) (*entities.Edge, error)
	// UpdateEdge(ctx InfraContext, edge entities.Edge) (*entities.Edge, error)
	// DeleteEdge(ctx InfraContext, clusterName string, name string) error
	//
	// OnDeleteEdgeMessage(ctx InfraContext, edge entities.Edge) error
	// OnUpdateEdgeMessage(ctx InfraContext, edge entities.Edge) error

	CreateNodePool(ctx InfraContext, clusterName string, nodePool entities.NodePool) (*entities.NodePool, error)
	UpdateNodePool(ctx InfraContext, clusterName string, nodePool entities.NodePool) (*entities.NodePool, error)

	DeleteNodePool(ctx InfraContext, clusterName string, poolName string) error
	ListNodePools(ctx InfraContext, clusterName string, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.NodePool], error)
	GetNodePool(ctx InfraContext, clusterName string, poolName string) (*entities.NodePool, error)

	OnDeleteNodePoolMessage(ctx InfraContext, clusterName string, nodePool entities.NodePool) error
	OnUpdateNodePoolMessage(ctx InfraContext, clusterName string, nodePool entities.NodePool) error

	// ListMasterNodes(ctx InfraContext, clusterName string) ([]*entities.MasterNode, error)
	// ListWorkerNodes(ctx InfraContext, clusterName string, edgeName string) ([]*entities.WorkerNode, error)
	// DeleteWorkerNode(ctx InfraContext, clusterName string, edgeName string, name string) (bool, error)
	//
	// OnDeleteWorkerNodeMessage(ctx InfraContext, workerNode entities.WorkerNode) error
	// OnUpdateWorkerNodeMessage(ctx InfraContext, workerNode entities.WorkerNode) error
}
