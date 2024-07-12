package types

import (
	"fmt"
	"strings"
)

type AccountType string

const (
	AccountTypeFree    AccountType = "free"
	AccountTypePremium AccountType = "premium"
)

type ResourceType string

const (
	ResourceAccount ResourceType = "account"

	ResourceEnvironment      ResourceType = "environment"
	ResourceConsoleVPNDevice ResourceType = "console_vpn_device"
	ResourceInfraVPNDevice   ResourceType = "infra_vpn_device"
)

type Role string

const (
	RoleResourceOwner Role = "resource_owner"

	RoleEnvironmentOwner        Role = "environment_owner"
	RoleEnvironmentCollaborator Role = "environment_collaborator"

	RoleAccountOwner  Role = "account_owner"
	RoleAccountAdmin  Role = "account_admin"
	RoleAccountMember Role = "account_member"
)

type Action string

const (
	CreateAccount Action = "create-account"
	ListAccounts  Action = "list-accounts"
	GetAccount    Action = "get-account"
	UpdateAccount Action = "update-account"
	DeleteAccount Action = "delete-account"

	CreateSecretsInAccount Action = "create-secrets-in-account"
	ReadSecretsFromAccount Action = "read-secrets-from-account"

	InviteAccountMember Action = "invite-account-member"
	InviteAccountAdmin  Action = "invite-account-admin"

	ListAccountInvitations Action = "list-account-invitations"
	GetAccountInvitation   Action = "get-account-invitation"

	DeleteAccountInvitation Action = "delete-account-invitation"

	ListMembershipsForAccount Action = "list-memberships-for-account"

	RemoveAccountMembership Action = "remove-account-membership"
	UpdateAccountMembership Action = "update-account-membership"

	CreateEnvironmentMembership Action = "create-environment-membership"
	RemoveEnvironmentMembership Action = "remove-environment-membership"
	UpdateEnvironmentMembership Action = "update-environment-membership"

	ActivateAccount   Action = "activate-account"
	DeactivateAccount Action = "deactivate-account"

	// clusters
	CreateCluster Action = "create-cluster"
	DeleteCluster Action = "delete-cluster"
	ListClusters  Action = "list-clusters"
	GetCluster    Action = "get-cluster"
	UpdateCluster Action = "update-cluster"

	// cluster managed services
	CreateClusterManagedService Action = "create-cluster-managed-service"
	CloneClusterManagedService  Action = "clone-cluster-managed-service"
	DeleteClusterManagedService Action = "delete-cluster-managed-service"
	ListClusterManagedServices  Action = "list-cluster-managed-services"
	GetClusterManagedService    Action = "get-cluster-managed-service"
	UpdateClusterManagedService Action = "update-cluster-managed-service"

	// helm releases
	CreateHelmRelease Action = "create-helm-release"
	DeleteHelmRelease Action = "delete-helm-release"
	ListHelmReleases  Action = "list-helm-releases"
	GetHelmRelease    Action = "get-helm-release"
	UpdateHelmRelease Action = "update-helm-release"

	// nodepools
	CreateNodepool Action = "create-nodepool"
	DeleteNodepool Action = "delete-nodepool"
	ListNodepools  Action = "list-nodepools"
	GetNodepool    Action = "get-nodepool"
	UpdateNodepool Action = "update-nodepool"

	// managed resource
	CreateManagedResource Action = "create-managed-resource"
	DeleteManagedResource Action = "delete-managed-resource"
	ListManagedResources  Action = "list-managed-resources"
	GetManagedResource    Action = "get-managed-resource"
	UpdateManagedResource Action = "update-managed-resource"

	CreateCloudProviderSecret Action = "create-cloud-provider-secret"
	UpdateCloudProviderSecret Action = "update-cloud-provider-secret"
	DeleteCloudProviderSecret Action = "delete-cloud-provider-secret"

	ListCloudProviderSecrets Action = "list-cloud-provider-secrets"
	GetCloudProviderSecret   Action = "get-cloud-provider-secret"

	CreateEnvironment Action = "create-environment"
	CloneEnvironment  Action = "clone-environment"
	UpdateEnvironment Action = "update-environment"
	DeleteEnvironment Action = "delete-environment"
	GetEnvironment    Action = "get-environment"
	ListEnvironments  Action = "list-environments"

	MutateResourcesInEnvironment Action = "mutate-resources-in-environment"
	ReadResourcesInEnvironment   Action = "read-resources-in-environment"

	ListVPNDevices Action = "list-vpn-devices"
	GetVPNDevice   Action = "get-vpn-device"

	GetVPNDeviceConnectConfig Action = "get-vpn-device-connect-config"

	CreateVPNDevice Action = "create-vpn-device"
	UpdateVPNDevice Action = "update-vpn-device"
	DeleteVPNDevice Action = "delete-vpn-device"

	CreateDomainEntry Action = "create-domain-entry"
	UpdateDomainEntry Action = "update-domain-entry"
	DeleteDomainEntry Action = "delete-domain-entry"

	ListDomainEntries Action = "list-domain-entries"
	GetDomainEntry    Action = "get-domain-entry"

	ReadLogs    Action = "read-logs"
	ReadMetrics Action = "read-metrics"

	// build runs
	ListBuildRuns  Action = "list-build-runs"
	GetBuildRun    Action = "get-build-run"
	CreateBuildRun Action = "create-build-run"
	UpdateBuildRun Action = "update-build-run"
	DeleteBuildRun Action = "delete-build-run"

	// build integrations
	ListBuildIntegrations  Action = "list-build-integrations"
	GetBuildIntegration    Action = "get-build-integration"
	UpdateBuildIntegration Action = "update-build-integration"
	CreateBuildIntegration Action = "create-build-integration"
	DeleteBuildIntegration Action = "delete-build-integration"

	// image pull secrets
	ListImagePullSecrets  Action = "list-image-pull-secrets"
	GetImagePullSecret    Action = "get-image-pull-secret"
	UpdateImagePullSecret Action = "update-image-pull-secret"
	CreateImagePullSecret Action = "create-image-pull-secret"
	DeleteImagePullSecret Action = "delete-image-pull-secret"
)

func NewResourceRef(accountName string, resourceType ResourceType, resourceName string) string {
	return fmt.Sprintf("%s/%s/%s", accountName, resourceType, resourceName)
}

func ParseResourceRef(rref string) (accountName, resourceType, resourceName string, err error) {
	sp := strings.SplitN(rref, "/", 3)
	if len(sp) != 3 {
		return "", "", "", fmt.Errorf("invalid resource ref %s", rref)
	}

	return sp[0], sp[1], sp[2], nil
}
