package domain

import (
	networkingv1 "k8s.io/api/networking/v1"

	name_check "github.com/kloudlite/api/apps/infra/internal/domain/name-check"
	provider_secret "github.com/kloudlite/api/apps/infra/internal/domain/provider-secret"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

// type InfraContext struct {
// 	context.Context
// 	UserId      repos.ID
// 	UserEmail   string
// 	UserName    string
// 	AccountName string
// }
//
// func (i InfraContext) GetUserId() repos.ID {
// 	return i.UserId
// }
//
// func (i InfraContext) GetUserEmail() string {
// 	return i.UserEmail
// }
//
// func (i InfraContext) GetUserName() string {
// 	return i.UserName
// }
//
// type UpdateAndDeleteOpts struct {
// 	MessageTimestamp time.Time
// }

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

type Domain interface {
	NameCheckDomain

	ClusterDomain
	HelmReleaseDomain
	BYOKClusterDomain
	NodepoolDomain

	ProviderSecretDomain

	GlobalVPNDomain
	GlobalVPNConnectionDomain

	PersistentVolumeDomain
	PersistentVolumeClaimDomain
	NamespaceDomain
	VolumeAttachmentDomain

	IngressHostsDomain
	ClusterManagedServiceDomain

	ManagedSvcTemplatesDomain

	// ListNodes(ctx InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Node], error)
	// GetNode(ctx InfraContext, clusterName string, nodeName string) (*entities.Node, error)

	// OnNodeUpdateMessage(ctx InfraContext, clusterName string, node entities.Node) error
	// OnNodeDeleteMessage(ctx InfraContext, clusterName string, node entities.Node) error
}

type NameCheckDomain interface {
	CheckNameAvailability(ctx domainT.InfraContext, typeArg name_check.ResType, clusterName *string, name string) (*name_check.CheckNameAvailabilityOutput, error)
}

type ProviderSecretDomain interface {
	CreateProviderSecret(ctx domainT.InfraContext, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error)
	UpdateProviderSecret(ctx domainT.InfraContext, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error)
	DeleteProviderSecret(ctx domainT.InfraContext, secretName string) error

	ListProviderSecrets(ctx domainT.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.CloudProviderSecret], error)
	GetProviderSecret(ctx domainT.InfraContext, name string) (*entities.CloudProviderSecret, error)

	ValidateProviderSecretAWSAccess(ctx domainT.InfraContext, name string) (*provider_secret.AWSAccessValidationOutput, error)
}

type NodepoolDomain interface {
	CreateNodePool(ctx domainT.InfraContext, clusterName string, nodePool entities.NodePool) (*entities.NodePool, error)
	UpdateNodePool(ctx domainT.InfraContext, clusterName string, nodePool entities.NodePool) (*entities.NodePool, error)
	DeleteNodePool(ctx domainT.InfraContext, clusterName string, poolName string) error

	ListNodePools(ctx domainT.InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.NodePool], error)
	GetNodePool(ctx domainT.InfraContext, clusterName string, poolName string) (*entities.NodePool, error)

	OnNodePoolDeleteMessage(ctx domainT.InfraContext, clusterName string, nodePool entities.NodePool) error
	OnNodePoolUpdateMessage(ctx domainT.InfraContext, clusterName string, nodePool entities.NodePool, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
	OnNodepoolApplyError(ctx domainT.InfraContext, clusterName string, name string, errMsg string, opts domainT.UpdateAndDeleteOpts) error
}

type GlobalVPNDomain interface {
	CreateGlobalVPN(ctx domainT.InfraContext, cluster entities.GlobalVPN) (*entities.GlobalVPN, error)
	UpdateGlobalVPN(ctx domainT.InfraContext, cluster entities.GlobalVPN) (*entities.GlobalVPN, error)
	DeleteGlobalVPN(ctx domainT.InfraContext, name string) error

	ListGlobalVPN(ctx domainT.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.GlobalVPN], error)
	GetGlobalVPN(ctx domainT.InfraContext, name string) (*entities.GlobalVPN, error)

	CreateGlobalVPNDevice(ctx domainT.InfraContext, device entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error)
	UpdateGlobalVPNDevice(ctx domainT.InfraContext, device entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error)
	DeleteGlobalVPNDevice(ctx domainT.InfraContext, gvpn string, device string) error

	ListGlobalVPNDevice(ctx domainT.InfraContext, gvpn string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.GlobalVPNDevice], error)
	GetGlobalVPNDevice(ctx domainT.InfraContext, gvpn string, device string) (*entities.GlobalVPNDevice, error)
	GetGlobalVPNDeviceWgConfig(ctx domainT.InfraContext, gvpn string, device string) (string, error)
}

type GlobalVPNConnectionDomain interface {
	OnGlobalVPNConnectionDeleteMessage(ctx domainT.InfraContext, clusterName string, clusterConn entities.GlobalVPNConnection) error
	OnGlobalVPNConnectionUpdateMessage(ctx domainT.InfraContext, clusterName string, clusterConn entities.GlobalVPNConnection, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
	OnGlobalVPNConnectionApplyError(ctx domainT.InfraContext, clusterName string, name string, errMsg string, opts domainT.UpdateAndDeleteOpts) error
}

type HelmReleaseDomain interface {
	ListHelmReleases(ctx domainT.InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.HelmRelease], error)
	GetHelmRelease(ctx domainT.InfraContext, clusterName string, serviceName string) (*entities.HelmRelease, error)

	CreateHelmRelease(ctx domainT.InfraContext, clusterName string, service entities.HelmRelease) (*entities.HelmRelease, error)
	UpdateHelmRelease(ctx domainT.InfraContext, clusterName string, service entities.HelmRelease) (*entities.HelmRelease, error)
	DeleteHelmRelease(ctx domainT.InfraContext, clusterName string, name string) error

	OnHelmReleaseApplyError(ctx domainT.InfraContext, clusterName string, name string, errMsg string, opts domainT.UpdateAndDeleteOpts) error
	OnHelmReleaseDeleteMessage(ctx domainT.InfraContext, clusterName string, service entities.HelmRelease) error
	OnHelmReleaseUpdateMessage(ctx domainT.InfraContext, clusterName string, service entities.HelmRelease, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
}

type ClusterDomain interface {
	CreateCluster(ctx domainT.InfraContext, cluster entities.Cluster) (*entities.Cluster, error)
	UpdateCluster(ctx domainT.InfraContext, cluster entities.Cluster) (*entities.Cluster, error)
	DeleteCluster(ctx domainT.InfraContext, name string) error

	ListClusters(ctx domainT.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Cluster], error)
	GetCluster(ctx domainT.InfraContext, name string) (*entities.Cluster, error)

	GetClusterAdminKubeconfig(ctx domainT.InfraContext, clusterName string) (*string, error)

	OnClusterDeleteMessage(ctx domainT.InfraContext, cluster entities.Cluster) error
	OnClusterUpdateMessage(ctx domainT.InfraContext, cluster entities.Cluster, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error

	UpgradeHelmKloudliteAgent(ctx domainT.InfraContext, clusterName string) error
}

type BYOKClusterDomain interface {
	CreateBYOKCluster(ctx domainT.InfraContext, cluster entities.BYOKCluster) (*entities.BYOKCluster, error)
	UpdateBYOKCluster(ctx domainT.InfraContext, clusterName string, displayName string) (*entities.BYOKCluster, error)
	ListBYOKCluster(ctx domainT.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.BYOKCluster], error)
	GetBYOKCluster(ctx domainT.InfraContext, name string) (*entities.BYOKCluster, error)
	GetBYOKClusterSetupInstructions(ctx domainT.InfraContext, name string) (*string, error)

	DeleteBYOKCluster(ctx domainT.InfraContext, name string) error
	UpsertBYOKClusterKubeconfig(ctx domainT.InfraContext, clusterName string, kubeconfig []byte) error
}

type ClusterManagedServiceDomain interface {
	ListClusterManagedServices(ctx domainT.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.ClusterManagedService], error)
	GetClusterManagedService(ctx domainT.InfraContext, serviceName string) (*entities.ClusterManagedService, error)
	CreateClusterManagedService(ctx domainT.InfraContext, cmsvc entities.ClusterManagedService) (*entities.ClusterManagedService, error)
	UpdateClusterManagedService(ctx domainT.InfraContext, cmsvc entities.ClusterManagedService) (*entities.ClusterManagedService, error)
	DeleteClusterManagedService(ctx domainT.InfraContext, name string) error

	OnClusterManagedServiceDeleteMessage(ctx domainT.InfraContext, clusterName string, service entities.ClusterManagedService) error
	OnClusterManagedServiceApplyError(ctx domainT.InfraContext, clusterName, name, errMsg string, opts domainT.UpdateAndDeleteOpts) error
	OnClusterManagedServiceUpdateMessage(ctx domainT.InfraContext, clusterName string, service entities.ClusterManagedService, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
}

type PersistentVolumeDomain interface {
	ListPVs(ctx domainT.InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.PersistentVolume], error)
	GetPV(ctx domainT.InfraContext, clusterName string, pvName string) (*entities.PersistentVolume, error)
	OnPVUpdateMessage(ctx domainT.InfraContext, clusterName string, pv entities.PersistentVolume, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
	OnPVDeleteMessage(ctx domainT.InfraContext, clusterName string, pv entities.PersistentVolume) error
	DeletePV(ctx domainT.InfraContext, clusterName string, pvName string) error
}

type PersistentVolumeClaimDomain interface {
	ListPVCs(ctx domainT.InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.PersistentVolumeClaim], error)
	GetPVC(ctx domainT.InfraContext, clusterName string, pvcName string) (*entities.PersistentVolumeClaim, error)
	OnPVCUpdateMessage(ctx domainT.InfraContext, clusterName string, pvc entities.PersistentVolumeClaim, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
	OnPVCDeleteMessage(ctx domainT.InfraContext, clusterName string, pvc entities.PersistentVolumeClaim) error
}

type VolumeAttachmentDomain interface {
	ListVolumeAttachments(ctx domainT.InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.VolumeAttachment], error)
	GetVolumeAttachment(ctx domainT.InfraContext, clusterName string, volAttachmentName string) (*entities.VolumeAttachment, error)
	OnVolumeAttachmentUpdateMessage(ctx domainT.InfraContext, clusterName string, volumeAttachment entities.VolumeAttachment, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
	OnVolumeAttachmentDeleteMessage(ctx domainT.InfraContext, clusterName string, volumeAttachment entities.VolumeAttachment) error
}

type NamespaceDomain interface {
	ListNamespaces(ctx domainT.InfraContext, clusterName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Namespace], error)
	GetNamespace(ctx domainT.InfraContext, clusterName string, namespace string) (*entities.Namespace, error)
	OnNamespaceUpdateMessage(ctx domainT.InfraContext, clusterName string, namespace entities.Namespace, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
	OnNamespaceDeleteMessage(ctx domainT.InfraContext, clusterName string, namespace entities.Namespace) error
}

type IngressHostsDomain interface {
	ListDomainEntries(ctx domainT.InfraContext, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.DomainEntry], error)
	GetDomainEntry(ctx domainT.InfraContext, name string) (*entities.DomainEntry, error)

	CreateDomainEntry(ctx domainT.InfraContext, domainName entities.DomainEntry) (*entities.DomainEntry, error)
	UpdateDomainEntry(ctx domainT.InfraContext, domainName entities.DomainEntry) (*entities.DomainEntry, error)
	DeleteDomainEntry(ctx domainT.InfraContext, name string) error

	OnIngressUpdateMessage(ctx domainT.InfraContext, clusterName string, ingress networkingv1.Ingress, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error
	OnIngressDeleteMessage(ctx domainT.InfraContext, clusterName string, ingress networkingv1.Ingress) error
}

type ManagedSvcTemplatesDomain interface {
	ListManagedSvcTemplates() ([]*entities.MsvcTemplate, error)
	GetManagedSvcTemplate(category string, name string) (*entities.MsvcTemplateEntry, error)
}
