package domain

import (
	"context"
	"kloudlite.io/apps/infra/internal/domain/entities"
)

type InfraContext struct {
	context.Context
	AccountName string
}

type Domain interface {
	CreateCluster(ctx InfraContext, cluster entities.Cluster) (*entities.Cluster, error)
	ListClusters(ctx InfraContext) ([]*entities.Cluster, error)
	GetCluster(ctx InfraContext, name string) (*entities.Cluster, error)
	UpdateCluster(ctx InfraContext, cluster entities.Cluster) (*entities.Cluster, error)
	DeleteCluster(ctx InfraContext, name string) error
	OnDeleteClusterMessage(ctx InfraContext, cluster entities.Cluster) error
	OnUpdateClusterMessage(ctx InfraContext, cluster entities.Cluster) error

	GetProviderSecret(ctx InfraContext, name string) (*entities.Secret, error)

	CreateCloudProvider(ctx InfraContext, cloudProvider entities.CloudProvider, providerSecret entities.Secret) (*entities.CloudProvider, error)
	ListCloudProviders(ctx InfraContext) ([]*entities.CloudProvider, error)
	GetCloudProvider(ctx InfraContext, name string) (*entities.CloudProvider, error)
	UpdateCloudProvider(ctx InfraContext, cloudProvider entities.CloudProvider, providerSecret *entities.Secret) (*entities.CloudProvider, error)
	DeleteCloudProvider(ctx InfraContext, name string) error
	OnDeleteCloudProviderMessage(ctx InfraContext, cloudProvider entities.CloudProvider) error
	OnUpdateCloudProviderMessage(ctx InfraContext, cloudProvider entities.CloudProvider) error

	CreateEdge(ctx InfraContext, edge entities.Edge) (*entities.Edge, error)
	ListEdges(ctx InfraContext, clusterName string, providerName *string) ([]*entities.Edge, error)
	GetEdge(ctx InfraContext, clusterName string, name string) (*entities.Edge, error)
	UpdateEdge(ctx InfraContext, edge entities.Edge) (*entities.Edge, error)
	DeleteEdge(ctx InfraContext, clusterName string, name string) error
	OnDeleteEdgeMessage(ctx InfraContext, edge entities.Edge) error
	OnUpdateEdgeMessage(ctx InfraContext, edge entities.Edge) error

	GetNodePools(ctx InfraContext, clusterName string, edgeName string) ([]*entities.NodePool, error)
	GetMasterNodes(ctx InfraContext, clusterName string) ([]*entities.MasterNode, error)
	GetWorkerNodes(ctx InfraContext, clusterName string, edgeName string) ([]*entities.WorkerNode, error)
	DeleteWorkerNode(ctx InfraContext, clusterName string, edgeName string, name string) (bool, error)
	OnDeleteWorkerNodeMessage(ctx InfraContext, workerNode entities.WorkerNode) error
	OnUpdateWorkerNodeMessage(ctx InfraContext, workerNode entities.WorkerNode) error
}