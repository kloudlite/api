package ports

import "github.com/kloudlite/api/apps/infra/internal/domain/types"

type PublishMsg string

const (
	PublishAdd    PublishMsg = "added"
	PublishDelete PublishMsg = "deleted"
	PublishUpdate PublishMsg = "updated"
)

type ResourceType string

const (
	ResourceTypeClusterManagedService ResourceType = "cluster_managed_service"
	ResourceTypeCluster               ResourceType = "cluster"
	ResourceTypeClusterGroup          ResourceType = "cluster_group"
	ResourceTypeBYOKCluster           ResourceType = "byok_cluster"
	ResourceTypeDomainEntries         ResourceType = "domain_entries"
	ResourceTypeHelmRelease           ResourceType = "helm_release"
	ResourceTypeNodePool              ResourceType = "nodepool"
	ResourceTypeClusterConnection     ResourceType = "cluster_connection"
	ResourceTypePVC                   ResourceType = "persistance_volume_claim"
	ResourceTypePV                    ResourceType = "persistance_volume"
	ResourceTypeVolumeAttachment      ResourceType = "volume_attachment"
)

type ResourceEventPublisher interface {
	PublishInfraEvent(ctx types.InfraContext, resourceType ResourceType, resName string, update PublishMsg)
	PublishResourceEvent(ctx types.InfraContext, clusterName string, resourceType ResourceType, resName string, update PublishMsg)
}
