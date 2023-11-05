// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/infra/internal/entities"
	"kloudlite.io/pkg/repos"
)

type BYOCClusterEdge struct {
	Cursor string                `json:"cursor"`
	Node   *entities.BYOCCluster `json:"node"`
}

type BYOCClusterPaginatedRecords struct {
	Edges      []*BYOCClusterEdge `json:"edges"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

type CheckAwsAccessOutput struct {
	Result          bool        `json:"result"`
	InstallationURL interface{} `json:"installationUrl,omitempty"`
}

type CloudProviderSecretEdge struct {
	Cursor string                        `json:"cursor"`
	Node   *entities.CloudProviderSecret `json:"node"`
}

type CloudProviderSecretPaginatedRecords struct {
	Edges      []*CloudProviderSecretEdge `json:"edges"`
	PageInfo   *PageInfo                  `json:"pageInfo"`
	TotalCount int                        `json:"totalCount"`
}

type ClusterEdge struct {
	Cursor string            `json:"cursor"`
	Node   *entities.Cluster `json:"node"`
}

type ClusterPaginatedRecords struct {
	Edges      []*ClusterEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

type DomainEntryEdge struct {
	Cursor string                `json:"cursor"`
	Node   *entities.DomainEntry `json:"node"`
}

type DomainEntryPaginatedRecords struct {
	Edges      []*DomainEntryEdge `json:"edges"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

type GithubComKloudliteOperatorApisClustersV1AWSClusterConfig struct {
	K3sMasters    *GithubComKloudliteOperatorApisClustersV1AWSK3sMastersConfig `json:"k3sMasters,omitempty"`
	NodePools     map[string]interface{}                                       `json:"nodePools,omitempty"`
	Region        string                                                       `json:"region"`
	SpotNodePools map[string]interface{}                                       `json:"spotNodePools,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1AWSClusterConfigIn struct {
	K3sMasters *GithubComKloudliteOperatorApisClustersV1AWSK3sMastersConfigIn `json:"k3sMasters,omitempty"`
	Region     string                                                         `json:"region"`
}

type GithubComKloudliteOperatorApisClustersV1AWSK3sMastersConfig struct {
	IamInstanceProfileRole *string                `json:"iamInstanceProfileRole,omitempty"`
	ImageID                string                 `json:"imageId"`
	ImageSSHUsername       string                 `json:"imageSSHUsername"`
	InstanceType           string                 `json:"instanceType"`
	Nodes                  map[string]interface{} `json:"nodes,omitempty"`
	NvidiaGpuEnabled       bool                   `json:"nvidiaGpuEnabled"`
	RootVolumeSize         int                    `json:"rootVolumeSize"`
	RootVolumeType         string                 `json:"rootVolumeType"`
}

type GithubComKloudliteOperatorApisClustersV1AWSK3sMastersConfigIn struct {
	IamInstanceProfileRole *string `json:"iamInstanceProfileRole,omitempty"`
	InstanceType           string  `json:"instanceType"`
}

type GithubComKloudliteOperatorApisClustersV1AWSNodePoolConfig struct {
	NormalPool *GithubComKloudliteOperatorApisClustersV1AwsNodePool     `json:"normalPool,omitempty"`
	PoolType   string                                                   `json:"poolType"`
	SpotPool   *GithubComKloudliteOperatorApisClustersV1AwsSpotNodePool `json:"spotPool,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1AWSNodePoolConfigIn struct {
	NormalPool *GithubComKloudliteOperatorApisClustersV1AwsNodePoolIn     `json:"normalPool,omitempty"`
	PoolType   string                                                     `json:"poolType"`
	SpotPool   *GithubComKloudliteOperatorApisClustersV1AwsSpotNodePoolIn `json:"spotPool,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1AwsNodePool struct {
	Ami                    string                 `json:"ami"`
	AmiSSHUsername         string                 `json:"amiSSHUsername"`
	AvailabilityZone       *string                `json:"availabilityZone,omitempty"`
	IamInstanceProfileRole *string                `json:"iamInstanceProfileRole,omitempty"`
	InstanceType           string                 `json:"instanceType"`
	Nodes                  map[string]interface{} `json:"nodes,omitempty"`
	NvidiaGpuEnabled       bool                   `json:"nvidiaGpuEnabled"`
	RootVolumeSize         int                    `json:"rootVolumeSize"`
	RootVolumeType         string                 `json:"rootVolumeType"`
}

type GithubComKloudliteOperatorApisClustersV1AwsNodePoolIn struct {
	Ami                    string                 `json:"ami"`
	AmiSSHUsername         string                 `json:"amiSSHUsername"`
	AvailabilityZone       *string                `json:"availabilityZone,omitempty"`
	IamInstanceProfileRole *string                `json:"iamInstanceProfileRole,omitempty"`
	InstanceType           string                 `json:"instanceType"`
	Nodes                  map[string]interface{} `json:"nodes,omitempty"`
	NvidiaGpuEnabled       bool                   `json:"nvidiaGpuEnabled"`
	RootVolumeSize         int                    `json:"rootVolumeSize"`
	RootVolumeType         string                 `json:"rootVolumeType"`
}

type GithubComKloudliteOperatorApisClustersV1AwsSpotCPUNode struct {
	MemoryPerVcpu *GithubComKloudliteOperatorApisCommonTypesMinMaxFloat `json:"memoryPerVcpu,omitempty"`
	Vcpu          *GithubComKloudliteOperatorApisCommonTypesMinMaxFloat `json:"vcpu"`
}

type GithubComKloudliteOperatorApisClustersV1AwsSpotCPUNodeIn struct {
	MemoryPerVcpu *GithubComKloudliteOperatorApisCommonTypesMinMaxFloatIn `json:"memoryPerVcpu,omitempty"`
	Vcpu          *GithubComKloudliteOperatorApisCommonTypesMinMaxFloatIn `json:"vcpu"`
}

type GithubComKloudliteOperatorApisClustersV1AwsSpotGpuNode struct {
	InstanceTypes []string `json:"instanceTypes"`
}

type GithubComKloudliteOperatorApisClustersV1AwsSpotGpuNodeIn struct {
	InstanceTypes []string `json:"instanceTypes"`
}

type GithubComKloudliteOperatorApisClustersV1AwsSpotNodePool struct {
	Ami                      string                                                  `json:"ami"`
	AmiSSHUsername           string                                                  `json:"amiSSHUsername"`
	AvailabilityZone         *string                                                 `json:"availabilityZone,omitempty"`
	CPUNode                  *GithubComKloudliteOperatorApisClustersV1AwsSpotCPUNode `json:"cpuNode,omitempty"`
	GpuNode                  *GithubComKloudliteOperatorApisClustersV1AwsSpotGpuNode `json:"gpuNode,omitempty"`
	IamInstanceProfileRole   *string                                                 `json:"iamInstanceProfileRole,omitempty"`
	Nodes                    map[string]interface{}                                  `json:"nodes,omitempty"`
	NvidiaGpuEnabled         bool                                                    `json:"nvidiaGpuEnabled"`
	RootVolumeSize           int                                                     `json:"rootVolumeSize"`
	RootVolumeType           string                                                  `json:"rootVolumeType"`
	SpotFleetTaggingRoleName string                                                  `json:"spotFleetTaggingRoleName"`
}

type GithubComKloudliteOperatorApisClustersV1AwsSpotNodePoolIn struct {
	Ami                      string                                                    `json:"ami"`
	AmiSSHUsername           string                                                    `json:"amiSSHUsername"`
	AvailabilityZone         *string                                                   `json:"availabilityZone,omitempty"`
	CPUNode                  *GithubComKloudliteOperatorApisClustersV1AwsSpotCPUNodeIn `json:"cpuNode,omitempty"`
	GpuNode                  *GithubComKloudliteOperatorApisClustersV1AwsSpotGpuNodeIn `json:"gpuNode,omitempty"`
	IamInstanceProfileRole   *string                                                   `json:"iamInstanceProfileRole,omitempty"`
	Nodes                    map[string]interface{}                                    `json:"nodes,omitempty"`
	NvidiaGpuEnabled         bool                                                      `json:"nvidiaGpuEnabled"`
	RootVolumeSize           int                                                       `json:"rootVolumeSize"`
	RootVolumeType           string                                                    `json:"rootVolumeType"`
	SpotFleetTaggingRoleName string                                                    `json:"spotFleetTaggingRoleName"`
}

type GithubComKloudliteOperatorApisClustersV1BYOCSpec struct {
	AccountName        string   `json:"accountName"`
	DisplayName        *string  `json:"displayName,omitempty"`
	IncomingKafkaTopic string   `json:"incomingKafkaTopic"`
	IngressClasses     []string `json:"ingressClasses,omitempty"`
	Provider           string   `json:"provider"`
	PublicIps          []string `json:"publicIps,omitempty"`
	Region             string   `json:"region"`
	StorageClasses     []string `json:"storageClasses,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1BYOCSpecIn struct {
	AccountName        string   `json:"accountName"`
	DisplayName        *string  `json:"displayName,omitempty"`
	IncomingKafkaTopic string   `json:"incomingKafkaTopic"`
	IngressClasses     []string `json:"ingressClasses,omitempty"`
	Provider           string   `json:"provider"`
	PublicIps          []string `json:"publicIps,omitempty"`
	Region             string   `json:"region"`
	StorageClasses     []string `json:"storageClasses,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1CloudProviderCredentialKeys struct {
	KeyAccessKey               string `json:"keyAccessKey"`
	KeyAWSAccountID            string `json:"keyAWSAccountId"`
	KeyAWSAssumeRoleExternalID string `json:"keyAWSAssumeRoleExternalID"`
	KeyAWSAssumeRoleRoleArn    string `json:"keyAWSAssumeRoleRoleARN"`
	KeySecretKey               string `json:"keySecretKey"`
}

type GithubComKloudliteOperatorApisClustersV1ClusterOutput struct {
	KeyK3sAgentJoinToken  string `json:"keyK3sAgentJoinToken"`
	KeyK3sServerJoinToken string `json:"keyK3sServerJoinToken"`
	KeyKubeconfig         string `json:"keyKubeconfig"`
	SecretName            string `json:"secretName"`
}

type GithubComKloudliteOperatorApisClustersV1ClusterSpec struct {
	AccountID              string                                                               `json:"accountId"`
	AccountName            string                                                               `json:"accountName"`
	AvailabilityMode       GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode  `json:"availabilityMode"`
	Aws                    *GithubComKloudliteOperatorApisClustersV1AWSClusterConfig            `json:"aws,omitempty"`
	BackupToS3Enabled      bool                                                                 `json:"backupToS3Enabled"`
	CloudflareEnabled      *bool                                                                `json:"cloudflareEnabled,omitempty"`
	CloudProvider          GithubComKloudliteOperatorApisCommonTypesCloudProvider               `json:"cloudProvider"`
	ClusterInternalDNSHost *string                                                              `json:"clusterInternalDnsHost,omitempty"`
	ClusterTokenRef        *GithubComKloudliteOperatorApisCommonTypesSecretKeyRef               `json:"clusterTokenRef,omitempty"`
	CredentialKeys         *GithubComKloudliteOperatorApisClustersV1CloudProviderCredentialKeys `json:"credentialKeys,omitempty"`
	CredentialsRef         *GithubComKloudliteOperatorApisCommonTypesSecretRef                  `json:"credentialsRef"`
	KloudliteRelease       string                                                               `json:"kloudliteRelease"`
	MessageQueueTopicName  string                                                               `json:"messageQueueTopicName"`
	Output                 *GithubComKloudliteOperatorApisClustersV1ClusterOutput               `json:"output,omitempty"`
	PublicDNSHost          string                                                               `json:"publicDNSHost"`
	TaintMasterNodes       bool                                                                 `json:"taintMasterNodes"`
}

type GithubComKloudliteOperatorApisClustersV1ClusterSpecIn struct {
	AvailabilityMode  GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode `json:"availabilityMode"`
	Aws               *GithubComKloudliteOperatorApisClustersV1AWSClusterConfigIn         `json:"aws,omitempty"`
	CloudflareEnabled *bool                                                               `json:"cloudflareEnabled,omitempty"`
	CloudProvider     GithubComKloudliteOperatorApisCommonTypesCloudProvider              `json:"cloudProvider"`
	CredentialsRef    *GithubComKloudliteOperatorApisCommonTypesSecretRefIn               `json:"credentialsRef"`
}

type GithubComKloudliteOperatorApisClustersV1MasterNodeProps struct {
	AvailabilityZone string  `json:"availabilityZone"`
	LastRecreatedAt  *string `json:"lastRecreatedAt,omitempty"`
	Role             string  `json:"role"`
}

type GithubComKloudliteOperatorApisClustersV1NodePoolSpec struct {
	Aws           *GithubComKloudliteOperatorApisClustersV1AWSNodePoolConfig `json:"aws,omitempty"`
	CloudProvider string                                                     `json:"cloudProvider"`
	MaxCount      int                                                        `json:"maxCount"`
	MinCount      int                                                        `json:"minCount"`
	TargetCount   int                                                        `json:"targetCount"`
}

type GithubComKloudliteOperatorApisClustersV1NodePoolSpecIn struct {
	Aws           *GithubComKloudliteOperatorApisClustersV1AWSNodePoolConfigIn `json:"aws,omitempty"`
	CloudProvider string                                                       `json:"cloudProvider"`
	MaxCount      int                                                          `json:"maxCount"`
	MinCount      int                                                          `json:"minCount"`
	TargetCount   int                                                          `json:"targetCount"`
}

type GithubComKloudliteOperatorApisClustersV1NodeProps struct {
	LastRecreatedAt *string `json:"lastRecreatedAt,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1NodePropsIn struct {
	LastRecreatedAt *string `json:"lastRecreatedAt,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1NodeSpec struct {
	NodepoolName string `json:"nodepoolName"`
}

type GithubComKloudliteOperatorApisClustersV1NodeSpecIn struct {
	NodepoolName string `json:"nodepoolName"`
}

type GithubComKloudliteOperatorApisCommonTypesMinMaxFloat struct {
	Max string `json:"max"`
	Min string `json:"min"`
}

type GithubComKloudliteOperatorApisCommonTypesMinMaxFloatIn struct {
	Max string `json:"max"`
	Min string `json:"min"`
}

type GithubComKloudliteOperatorApisCommonTypesSecretKeyRef struct {
	Key       string  `json:"key"`
	Name      string  `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
}

type GithubComKloudliteOperatorApisCommonTypesSecretRef struct {
	Name      string  `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
}

type GithubComKloudliteOperatorApisCommonTypesSecretRefIn struct {
	Name      string  `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
}

type GithubComKloudliteOperatorApisWireguardV1DeviceSpec struct {
	Offset     *int                                             `json:"offset,omitempty"`
	Ports      []*GithubComKloudliteOperatorApisWireguardV1Port `json:"ports,omitempty"`
	ServerName string                                           `json:"serverName"`
}

type GithubComKloudliteOperatorApisWireguardV1DeviceSpecIn struct {
	Offset     *int                                               `json:"offset,omitempty"`
	Ports      []*GithubComKloudliteOperatorApisWireguardV1PortIn `json:"ports,omitempty"`
	ServerName string                                             `json:"serverName"`
}

type GithubComKloudliteOperatorApisWireguardV1Port struct {
	Port       *int `json:"port,omitempty"`
	TargetPort *int `json:"targetPort,omitempty"`
}

type GithubComKloudliteOperatorApisWireguardV1PortIn struct {
	Port       *int `json:"port,omitempty"`
	TargetPort *int `json:"targetPort,omitempty"`
}

type GithubComKloudliteOperatorPkgRawJSONRawJSON struct {
	RawMessage interface{} `json:"RawMessage,omitempty"`
}

type KloudliteIoAppsInfraInternalEntitiesAWSSecretCredentials struct {
	AccessKey               *string `json:"accessKey,omitempty"`
	AwsAccountID            *string `json:"awsAccountId,omitempty"`
	AwsAssumeRoleExternalID *string `json:"awsAssumeRoleExternalId,omitempty"`
	AwsAssumeRoleRoleArn    *string `json:"awsAssumeRoleRoleARN,omitempty"`
	SecretKey               *string `json:"secretKey,omitempty"`
}

type KloudliteIoAppsInfraInternalEntitiesAWSSecretCredentialsIn struct {
	AccessKey    *string `json:"accessKey,omitempty"`
	AwsAccountID *string `json:"awsAccountId,omitempty"`
	SecretKey    *string `json:"secretKey,omitempty"`
}

type KloudliteIoAppsInfraInternalEntitiesHelmStatusVal struct {
	IsReady *bool  `json:"isReady,omitempty"`
	Message string `json:"message"`
}

type NodeEdge struct {
	Cursor string         `json:"cursor"`
	Node   *entities.Node `json:"node"`
}

type NodeIn struct {
	Metadata *v1.ObjectMeta                                      `json:"metadata,omitempty"`
	Spec     *GithubComKloudliteOperatorApisClustersV1NodeSpecIn `json:"spec"`
}

type NodePaginatedRecords struct {
	Edges      []*NodeEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type NodePoolEdge struct {
	Cursor string             `json:"cursor"`
	Node   *entities.NodePool `json:"node"`
}

type NodePoolPaginatedRecords struct {
	Edges      []*NodePoolEdge `json:"edges"`
	PageInfo   *PageInfo       `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor,omitempty"`
	HasNextPage     *bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage *bool   `json:"hasPreviousPage,omitempty"`
	StartCursor     *string `json:"startCursor,omitempty"`
}

type SearchCluster struct {
	CloudProviderName *repos.MatchFilter `json:"cloudProviderName,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	Region            *repos.MatchFilter `json:"region,omitempty"`
	Text              *repos.MatchFilter `json:"text,omitempty"`
}

type SearchDomainEntry struct {
	ClusterName *repos.MatchFilter `json:"clusterName,omitempty"`
	Text        *repos.MatchFilter `json:"text,omitempty"`
}

type SearchNodepool struct {
	Text *repos.MatchFilter `json:"text,omitempty"`
}

type SearchProviderSecret struct {
	CloudProviderName *repos.MatchFilter `json:"cloudProviderName,omitempty"`
	Text              *repos.MatchFilter `json:"text,omitempty"`
}

type SearchVPNDevices struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type VPNDeviceEdge struct {
	Cursor string              `json:"cursor"`
	Node   *entities.VPNDevice `json:"node"`
}

type VPNDevicePaginatedRecords struct {
	Edges      []*VPNDeviceEdge `json:"edges"`
	PageInfo   *PageInfo        `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

type GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode string

const (
	GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityModeDev GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode = "dev"
	GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityModeHa  GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode = "HA"
)

var AllGithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode = []GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode{
	GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityModeDev,
	GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityModeHa,
}

func (e GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode) IsValid() bool {
	switch e {
	case GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityModeDev, GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityModeHa:
		return true
	}
	return false
}

func (e GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode) String() string {
	return string(e)
}

func (e *GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___operator___apis___clusters___v1__ClusterSpecAvailabilityMode", str)
	}
	return nil
}

func (e GithubComKloudliteOperatorApisClustersV1ClusterSpecAvailabilityMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GithubComKloudliteOperatorApisCommonTypesCloudProvider string

const (
	GithubComKloudliteOperatorApisCommonTypesCloudProviderAws   GithubComKloudliteOperatorApisCommonTypesCloudProvider = "aws"
	GithubComKloudliteOperatorApisCommonTypesCloudProviderAzure GithubComKloudliteOperatorApisCommonTypesCloudProvider = "azure"
	GithubComKloudliteOperatorApisCommonTypesCloudProviderDo    GithubComKloudliteOperatorApisCommonTypesCloudProvider = "do"
	GithubComKloudliteOperatorApisCommonTypesCloudProviderGcp   GithubComKloudliteOperatorApisCommonTypesCloudProvider = "gcp"
)

var AllGithubComKloudliteOperatorApisCommonTypesCloudProvider = []GithubComKloudliteOperatorApisCommonTypesCloudProvider{
	GithubComKloudliteOperatorApisCommonTypesCloudProviderAws,
	GithubComKloudliteOperatorApisCommonTypesCloudProviderAzure,
	GithubComKloudliteOperatorApisCommonTypesCloudProviderDo,
	GithubComKloudliteOperatorApisCommonTypesCloudProviderGcp,
}

func (e GithubComKloudliteOperatorApisCommonTypesCloudProvider) IsValid() bool {
	switch e {
	case GithubComKloudliteOperatorApisCommonTypesCloudProviderAws, GithubComKloudliteOperatorApisCommonTypesCloudProviderAzure, GithubComKloudliteOperatorApisCommonTypesCloudProviderDo, GithubComKloudliteOperatorApisCommonTypesCloudProviderGcp:
		return true
	}
	return false
}

func (e GithubComKloudliteOperatorApisCommonTypesCloudProvider) String() string {
	return string(e)
}

func (e *GithubComKloudliteOperatorApisCommonTypesCloudProvider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteOperatorApisCommonTypesCloudProvider(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___operator___apis___common____types__CloudProvider", str)
	}
	return nil
}

func (e GithubComKloudliteOperatorApisCommonTypesCloudProvider) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
