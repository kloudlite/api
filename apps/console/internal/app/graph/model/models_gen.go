// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/kloudlite/api/apps/console/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type AppEdge struct {
	Cursor string        `json:"cursor"`
	Node   *entities.App `json:"node"`
}

type AppPaginatedRecords struct {
	Edges      []*AppEdge `json:"edges"`
	PageInfo   *PageInfo  `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

type ConfigEdge struct {
	Cursor string           `json:"cursor"`
	Node   *entities.Config `json:"node"`
}

type ConfigKeyRef struct {
	ConfigName string `json:"configName"`
	Key        string `json:"key"`
}

type ConfigKeyValueRefIn struct {
	ConfigName string `json:"configName"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}

type ConfigPaginatedRecords struct {
	Edges      []*ConfigEdge `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

type EnvironmentEdge struct {
	Cursor string                `json:"cursor"`
	Node   *entities.Environment `json:"node"`
}

type EnvironmentPaginatedRecords struct {
	Edges      []*EnvironmentEdge `json:"edges"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

type GithubComKloudliteOperatorApisCrdsV1AppContainer struct {
	Args            []string                                               `json:"args,omitempty"`
	Command         []string                                               `json:"command,omitempty"`
	Env             []*GithubComKloudliteOperatorApisCrdsV1ContainerEnv    `json:"env,omitempty"`
	EnvFrom         []*GithubComKloudliteOperatorApisCrdsV1EnvFrom         `json:"envFrom,omitempty"`
	Image           string                                                 `json:"image"`
	ImagePullPolicy *string                                                `json:"imagePullPolicy,omitempty"`
	LivenessProbe   *GithubComKloudliteOperatorApisCrdsV1Probe             `json:"livenessProbe,omitempty"`
	Name            string                                                 `json:"name"`
	ReadinessProbe  *GithubComKloudliteOperatorApisCrdsV1Probe             `json:"readinessProbe,omitempty"`
	ResourceCPU     *GithubComKloudliteOperatorApisCrdsV1ContainerResource `json:"resourceCpu,omitempty"`
	ResourceMemory  *GithubComKloudliteOperatorApisCrdsV1ContainerResource `json:"resourceMemory,omitempty"`
	Volumes         []*GithubComKloudliteOperatorApisCrdsV1ContainerVolume `json:"volumes,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppContainerIn struct {
	Args            []string                                                 `json:"args,omitempty"`
	Command         []string                                                 `json:"command,omitempty"`
	Env             []*GithubComKloudliteOperatorApisCrdsV1ContainerEnvIn    `json:"env,omitempty"`
	EnvFrom         []*GithubComKloudliteOperatorApisCrdsV1EnvFromIn         `json:"envFrom,omitempty"`
	Image           string                                                   `json:"image"`
	ImagePullPolicy *string                                                  `json:"imagePullPolicy,omitempty"`
	LivenessProbe   *GithubComKloudliteOperatorApisCrdsV1ProbeIn             `json:"livenessProbe,omitempty"`
	Name            string                                                   `json:"name"`
	ReadinessProbe  *GithubComKloudliteOperatorApisCrdsV1ProbeIn             `json:"readinessProbe,omitempty"`
	ResourceCPU     *GithubComKloudliteOperatorApisCrdsV1ContainerResourceIn `json:"resourceCpu,omitempty"`
	ResourceMemory  *GithubComKloudliteOperatorApisCrdsV1ContainerResourceIn `json:"resourceMemory,omitempty"`
	Volumes         []*GithubComKloudliteOperatorApisCrdsV1ContainerVolumeIn `json:"volumes,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppSpec struct {
	Containers     []*GithubComKloudliteOperatorApisCrdsV1AppContainer `json:"containers"`
	DisplayName    *string                                             `json:"displayName,omitempty"`
	Freeze         *bool                                               `json:"freeze,omitempty"`
	Hpa            *GithubComKloudliteOperatorApisCrdsV1Hpa            `json:"hpa,omitempty"`
	Intercept      *GithubComKloudliteOperatorApisCrdsV1Intercept      `json:"intercept,omitempty"`
	NodeSelector   map[string]interface{}                              `json:"nodeSelector,omitempty"`
	Region         *string                                             `json:"region,omitempty"`
	Replicas       *int                                                `json:"replicas,omitempty"`
	ServiceAccount *string                                             `json:"serviceAccount,omitempty"`
	Services       []*GithubComKloudliteOperatorApisCrdsV1AppSvc       `json:"services,omitempty"`
	Tolerations    []*K8sIoAPICoreV1Toleration                         `json:"tolerations,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppSpecIn struct {
	Containers     []*GithubComKloudliteOperatorApisCrdsV1AppContainerIn `json:"containers"`
	DisplayName    *string                                               `json:"displayName,omitempty"`
	Freeze         *bool                                                 `json:"freeze,omitempty"`
	Hpa            *GithubComKloudliteOperatorApisCrdsV1HPAIn            `json:"hpa,omitempty"`
	Intercept      *GithubComKloudliteOperatorApisCrdsV1InterceptIn      `json:"intercept,omitempty"`
	NodeSelector   map[string]interface{}                                `json:"nodeSelector,omitempty"`
	Region         *string                                               `json:"region,omitempty"`
	Replicas       *int                                                  `json:"replicas,omitempty"`
	ServiceAccount *string                                               `json:"serviceAccount,omitempty"`
	Services       []*GithubComKloudliteOperatorApisCrdsV1AppSvcIn       `json:"services,omitempty"`
	Tolerations    []*K8sIoAPICoreV1TolerationIn                         `json:"tolerations,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppSvc struct {
	Name       *string `json:"name,omitempty"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppSvcIn struct {
	Name       *string `json:"name,omitempty"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1BasicAuth struct {
	Enabled    bool    `json:"enabled"`
	SecretName *string `json:"secretName,omitempty"`
	Username   *string `json:"username,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1BasicAuthIn struct {
	Enabled    bool    `json:"enabled"`
	SecretName *string `json:"secretName,omitempty"`
	Username   *string `json:"username,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1ContainerEnv struct {
	Key      string                                              `json:"key"`
	Optional *bool                                               `json:"optional,omitempty"`
	RefKey   *string                                             `json:"refKey,omitempty"`
	RefName  *string                                             `json:"refName,omitempty"`
	Type     *GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret `json:"type,omitempty"`
	Value    *string                                             `json:"value,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1ContainerEnvIn struct {
	Key      string                                              `json:"key"`
	Optional *bool                                               `json:"optional,omitempty"`
	RefKey   *string                                             `json:"refKey,omitempty"`
	RefName  *string                                             `json:"refName,omitempty"`
	Type     *GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret `json:"type,omitempty"`
	Value    *string                                             `json:"value,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1ContainerResource struct {
	Max *string `json:"max,omitempty"`
	Min *string `json:"min,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1ContainerResourceIn struct {
	Max *string `json:"max,omitempty"`
	Min *string `json:"min,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1ContainerVolume struct {
	Items     []*GithubComKloudliteOperatorApisCrdsV1ContainerVolumeItem `json:"items,omitempty"`
	MountPath string                                                     `json:"mountPath"`
	RefName   string                                                     `json:"refName"`
	Type      GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret         `json:"type"`
}

type GithubComKloudliteOperatorApisCrdsV1ContainerVolumeIn struct {
	Items     []*GithubComKloudliteOperatorApisCrdsV1ContainerVolumeItemIn `json:"items,omitempty"`
	MountPath string                                                       `json:"mountPath"`
	RefName   string                                                       `json:"refName"`
	Type      GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret           `json:"type"`
}

type GithubComKloudliteOperatorApisCrdsV1ContainerVolumeItem struct {
	FileName *string `json:"fileName,omitempty"`
	Key      string  `json:"key"`
}

type GithubComKloudliteOperatorApisCrdsV1ContainerVolumeItemIn struct {
	FileName *string `json:"fileName,omitempty"`
	Key      string  `json:"key"`
}

type GithubComKloudliteOperatorApisCrdsV1Cors struct {
	AllowCredentials *bool    `json:"allowCredentials,omitempty"`
	Enabled          *bool    `json:"enabled,omitempty"`
	Origins          []string `json:"origins,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1CorsIn struct {
	AllowCredentials *bool    `json:"allowCredentials,omitempty"`
	Enabled          *bool    `json:"enabled,omitempty"`
	Origins          []string `json:"origins,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1EnvFrom struct {
	RefName string                                             `json:"refName"`
	Type    GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret `json:"type"`
}

type GithubComKloudliteOperatorApisCrdsV1EnvFromIn struct {
	RefName string                                             `json:"refName"`
	Type    GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret `json:"type"`
}

type GithubComKloudliteOperatorApisCrdsV1EnvironmentRouting struct {
	Mode                *GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode `json:"mode,omitempty"`
	PrivateIngressClass *string                                                     `json:"privateIngressClass,omitempty"`
	PublicIngressClass  *string                                                     `json:"publicIngressClass,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingIn struct {
	Mode *GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode `json:"mode,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1EnvironmentSpec struct {
	ProjectName     string                                                  `json:"projectName"`
	Routing         *GithubComKloudliteOperatorApisCrdsV1EnvironmentRouting `json:"routing,omitempty"`
	TargetNamespace *string                                                 `json:"targetNamespace,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1EnvironmentSpecIn struct {
	ProjectName     string                                                    `json:"projectName"`
	Routing         *GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingIn `json:"routing,omitempty"`
	TargetNamespace *string                                                   `json:"targetNamespace,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1Hpa struct {
	Enabled         *bool `json:"enabled,omitempty"`
	MaxReplicas     *int  `json:"maxReplicas,omitempty"`
	MinReplicas     *int  `json:"minReplicas,omitempty"`
	ThresholdCPU    *int  `json:"thresholdCpu,omitempty"`
	ThresholdMemory *int  `json:"thresholdMemory,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1HPAIn struct {
	Enabled         *bool `json:"enabled,omitempty"`
	MaxReplicas     *int  `json:"maxReplicas,omitempty"`
	MinReplicas     *int  `json:"minReplicas,omitempty"`
	ThresholdCPU    *int  `json:"thresholdCpu,omitempty"`
	ThresholdMemory *int  `json:"thresholdMemory,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1HTTPGetProbe struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type GithubComKloudliteOperatorApisCrdsV1HTTPGetProbeIn struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type GithubComKloudliteOperatorApisCrdsV1HTTPS struct {
	ClusterIssuer *string `json:"clusterIssuer,omitempty"`
	Enabled       bool    `json:"enabled"`
	ForceRedirect *bool   `json:"forceRedirect,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1HTTPSIn struct {
	ClusterIssuer *string `json:"clusterIssuer,omitempty"`
	Enabled       bool    `json:"enabled"`
	ForceRedirect *bool   `json:"forceRedirect,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1Intercept struct {
	Enabled  bool   `json:"enabled"`
	ToDevice string `json:"toDevice"`
}

type GithubComKloudliteOperatorApisCrdsV1InterceptIn struct {
	Enabled  bool   `json:"enabled"`
	ToDevice string `json:"toDevice"`
}

type GithubComKloudliteOperatorApisCrdsV1ManagedResourceSpec struct {
	ResourceTemplate *GithubComKloudliteOperatorApisCrdsV1MresResourceTemplate `json:"resourceTemplate"`
}

type GithubComKloudliteOperatorApisCrdsV1ManagedResourceSpecIn struct {
	ResourceTemplate *GithubComKloudliteOperatorApisCrdsV1MresResourceTemplateIn `json:"resourceTemplate"`
}

type GithubComKloudliteOperatorApisCrdsV1ManagedServiceSpec struct {
	ServiceTemplate *GithubComKloudliteOperatorApisCrdsV1ServiceTemplate `json:"serviceTemplate"`
}

type GithubComKloudliteOperatorApisCrdsV1ManagedServiceSpecIn struct {
	ServiceTemplate *GithubComKloudliteOperatorApisCrdsV1ServiceTemplateIn `json:"serviceTemplate"`
}

type GithubComKloudliteOperatorApisCrdsV1MresResourceTemplate struct {
	APIVersion string                                            `json:"apiVersion"`
	Kind       string                                            `json:"kind"`
	MsvcRef    *GithubComKloudliteOperatorApisCrdsV1MsvcNamedRef `json:"msvcRef"`
	Spec       map[string]interface{}                            `json:"spec"`
}

type GithubComKloudliteOperatorApisCrdsV1MresResourceTemplateIn struct {
	MsvcRef *GithubComKloudliteOperatorApisCrdsV1MsvcNamedRefIn `json:"msvcRef"`
	Spec    map[string]interface{}                              `json:"spec"`
}

type GithubComKloudliteOperatorApisCrdsV1MsvcNamedRef struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
}

type GithubComKloudliteOperatorApisCrdsV1MsvcNamedRefIn struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
}

type GithubComKloudliteOperatorApisCrdsV1Probe struct {
	FailureThreshold *int                                              `json:"failureThreshold,omitempty"`
	HTTPGet          *GithubComKloudliteOperatorApisCrdsV1HTTPGetProbe `json:"httpGet,omitempty"`
	InitialDelay     *int                                              `json:"initialDelay,omitempty"`
	Interval         *int                                              `json:"interval,omitempty"`
	Shell            *GithubComKloudliteOperatorApisCrdsV1ShellProbe   `json:"shell,omitempty"`
	TCP              *GithubComKloudliteOperatorApisCrdsV1TCPProbe     `json:"tcp,omitempty"`
	Type             string                                            `json:"type"`
}

type GithubComKloudliteOperatorApisCrdsV1ProbeIn struct {
	FailureThreshold *int                                                `json:"failureThreshold,omitempty"`
	HTTPGet          *GithubComKloudliteOperatorApisCrdsV1HTTPGetProbeIn `json:"httpGet,omitempty"`
	InitialDelay     *int                                                `json:"initialDelay,omitempty"`
	Interval         *int                                                `json:"interval,omitempty"`
	Shell            *GithubComKloudliteOperatorApisCrdsV1ShellProbeIn   `json:"shell,omitempty"`
	TCP              *GithubComKloudliteOperatorApisCrdsV1TCPProbeIn     `json:"tcp,omitempty"`
	Type             string                                              `json:"type"`
}

type GithubComKloudliteOperatorApisCrdsV1ProjectManagedServiceSpec struct {
	MsvcSpec        *GithubComKloudliteOperatorApisCrdsV1ManagedServiceSpec `json:"msvcSpec"`
	TargetNamespace string                                                  `json:"targetNamespace"`
}

type GithubComKloudliteOperatorApisCrdsV1ProjectManagedServiceSpecIn struct {
	MsvcSpec        *GithubComKloudliteOperatorApisCrdsV1ManagedServiceSpecIn `json:"msvcSpec"`
	TargetNamespace string                                                    `json:"targetNamespace"`
}

type GithubComKloudliteOperatorApisCrdsV1ProjectSpec struct {
	AccountName     string  `json:"accountName"`
	ClusterName     string  `json:"clusterName"`
	DisplayName     *string `json:"displayName,omitempty"`
	Logo            *string `json:"logo,omitempty"`
	TargetNamespace string  `json:"targetNamespace"`
}

type GithubComKloudliteOperatorApisCrdsV1ProjectSpecIn struct {
	DisplayName     *string `json:"displayName,omitempty"`
	Logo            *string `json:"logo,omitempty"`
	TargetNamespace string  `json:"targetNamespace"`
}

type GithubComKloudliteOperatorApisCrdsV1RateLimit struct {
	Connections *int  `json:"connections,omitempty"`
	Enabled     *bool `json:"enabled,omitempty"`
	Rpm         *int  `json:"rpm,omitempty"`
	Rps         *int  `json:"rps,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1RateLimitIn struct {
	Connections *int  `json:"connections,omitempty"`
	Enabled     *bool `json:"enabled,omitempty"`
	Rpm         *int  `json:"rpm,omitempty"`
	Rps         *int  `json:"rps,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1Route struct {
	App     *string `json:"app,omitempty"`
	Lambda  *string `json:"lambda,omitempty"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1RouteIn struct {
	App     *string `json:"app,omitempty"`
	Lambda  *string `json:"lambda,omitempty"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1RouterSpec struct {
	BackendProtocol *string                                        `json:"backendProtocol,omitempty"`
	BasicAuth       *GithubComKloudliteOperatorApisCrdsV1BasicAuth `json:"basicAuth,omitempty"`
	Cors            *GithubComKloudliteOperatorApisCrdsV1Cors      `json:"cors,omitempty"`
	Domains         []string                                       `json:"domains"`
	HTTPS           *GithubComKloudliteOperatorApisCrdsV1HTTPS     `json:"https,omitempty"`
	IngressClass    *string                                        `json:"ingressClass,omitempty"`
	MaxBodySizeInMb *int                                           `json:"maxBodySizeInMB,omitempty"`
	RateLimit       *GithubComKloudliteOperatorApisCrdsV1RateLimit `json:"rateLimit,omitempty"`
	Region          *string                                        `json:"region,omitempty"`
	Routes          []*GithubComKloudliteOperatorApisCrdsV1Route   `json:"routes,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1RouterSpecIn struct {
	BackendProtocol *string                                          `json:"backendProtocol,omitempty"`
	BasicAuth       *GithubComKloudliteOperatorApisCrdsV1BasicAuthIn `json:"basicAuth,omitempty"`
	Cors            *GithubComKloudliteOperatorApisCrdsV1CorsIn      `json:"cors,omitempty"`
	Domains         []string                                         `json:"domains"`
	HTTPS           *GithubComKloudliteOperatorApisCrdsV1HTTPSIn     `json:"https,omitempty"`
	IngressClass    *string                                          `json:"ingressClass,omitempty"`
	MaxBodySizeInMb *int                                             `json:"maxBodySizeInMB,omitempty"`
	RateLimit       *GithubComKloudliteOperatorApisCrdsV1RateLimitIn `json:"rateLimit,omitempty"`
	Region          *string                                          `json:"region,omitempty"`
	Routes          []*GithubComKloudliteOperatorApisCrdsV1RouteIn   `json:"routes,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1ServiceTemplate struct {
	APIVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Spec       map[string]interface{} `json:"spec"`
}

type GithubComKloudliteOperatorApisCrdsV1ServiceTemplateIn struct {
	APIVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Spec       map[string]interface{} `json:"spec"`
}

type GithubComKloudliteOperatorApisCrdsV1ShellProbe struct {
	Command []string `json:"command,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1ShellProbeIn struct {
	Command []string `json:"command,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1TCPProbe struct {
	Port int `json:"port"`
}

type GithubComKloudliteOperatorApisCrdsV1TCPProbeIn struct {
	Port int `json:"port"`
}

type GithubComKloudliteOperatorPkgOperatorCheck struct {
	Generation *int    `json:"generation,omitempty"`
	Message    *string `json:"message,omitempty"`
	Status     bool    `json:"status"`
}

type GithubComKloudliteOperatorPkgOperatorResourceRef struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
}

type GithubComKloudliteOperatorPkgRawJSONRawJSON struct {
	RawMessage interface{} `json:"RawMessage,omitempty"`
}

type ImagePullSecretEdge struct {
	Cursor string                    `json:"cursor"`
	Node   *entities.ImagePullSecret `json:"node"`
}

type ImagePullSecretPaginatedRecords struct {
	Edges      []*ImagePullSecretEdge `json:"edges"`
	PageInfo   *PageInfo              `json:"pageInfo"`
	TotalCount int                    `json:"totalCount"`
}

type K8sIoAPICoreV1Toleration struct {
	Effect            *K8sIoAPICoreV1TaintEffect        `json:"effect,omitempty"`
	Key               *string                           `json:"key,omitempty"`
	Operator          *K8sIoAPICoreV1TolerationOperator `json:"operator,omitempty"`
	TolerationSeconds *int                              `json:"tolerationSeconds,omitempty"`
	Value             *string                           `json:"value,omitempty"`
}

type K8sIoAPICoreV1TolerationIn struct {
	Effect            *K8sIoAPICoreV1TaintEffect        `json:"effect,omitempty"`
	Key               *string                           `json:"key,omitempty"`
	Operator          *K8sIoAPICoreV1TolerationOperator `json:"operator,omitempty"`
	TolerationSeconds *int                              `json:"tolerationSeconds,omitempty"`
	Value             *string                           `json:"value,omitempty"`
}

type ManagedResourceEdge struct {
	Cursor string                    `json:"cursor"`
	Node   *entities.ManagedResource `json:"node"`
}

type ManagedResourcePaginatedRecords struct {
	Edges      []*ManagedResourceEdge `json:"edges"`
	PageInfo   *PageInfo              `json:"pageInfo"`
	TotalCount int                    `json:"totalCount"`
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor,omitempty"`
	HasNextPage     *bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage *bool   `json:"hasPreviousPage,omitempty"`
	StartCursor     *string `json:"startCursor,omitempty"`
}

type ProjectEdge struct {
	Cursor string            `json:"cursor"`
	Node   *entities.Project `json:"node"`
}

type ProjectManagedServiceEdge struct {
	Cursor string                          `json:"cursor"`
	Node   *entities.ProjectManagedService `json:"node"`
}

type ProjectManagedServicePaginatedRecords struct {
	Edges      []*ProjectManagedServiceEdge `json:"edges"`
	PageInfo   *PageInfo                    `json:"pageInfo"`
	TotalCount int                          `json:"totalCount"`
}

type ProjectPaginatedRecords struct {
	Edges      []*ProjectEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

type RouterEdge struct {
	Cursor string           `json:"cursor"`
	Node   *entities.Router `json:"node"`
}

type RouterPaginatedRecords struct {
	Edges      []*RouterEdge `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

type SearchApps struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchConfigs struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchEnvironments struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	ProjectName       *repos.MatchFilter `json:"projectName,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchImagePullSecrets struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchManagedResources struct {
	Text               *repos.MatchFilter `json:"text,omitempty"`
	ManagedServiceName *repos.MatchFilter `json:"managedServiceName,omitempty"`
	IsReady            *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion  *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchProjectManagedService struct {
	Text               *repos.MatchFilter `json:"text,omitempty"`
	ManagedServiceName *repos.MatchFilter `json:"managedServiceName,omitempty"`
	IsReady            *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion  *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchProjects struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchRouters struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchSecrets struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SecretEdge struct {
	Cursor string           `json:"cursor"`
	Node   *entities.Secret `json:"node"`
}

type SecretKeyRef struct {
	Key         string `json:"key"`
	SeceretName string `json:"seceretName"`
}

type SecretKeyValueRefIn struct {
	Key         string `json:"key"`
	SeceretName string `json:"seceretName"`
	Value       string `json:"value"`
}

type SecretPaginatedRecords struct {
	Edges      []*SecretEdge `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

type GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat string

const (
	GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormatDockerConfigJSON GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat = "dockerConfigJson"
	GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormatParams           GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat = "params"
)

var AllGithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat = []GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat{
	GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormatDockerConfigJSON,
	GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormatParams,
}

func (e GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat) IsValid() bool {
	switch e {
	case GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormatDockerConfigJSON, GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormatParams:
		return true
	}
	return false
}

func (e GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat) String() string {
	return string(e)
}

func (e *GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___api___apps___console___internal___entities__ImagePullSecretFormat", str)
	}
	return nil
}

func (e GithubComKloudliteAPIAppsConsoleInternalEntitiesImagePullSecretFormat) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret string

const (
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretConfig GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret = "config"
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretSecret GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret = "secret"
)

var AllGithubComKloudliteOperatorApisCrdsV1ConfigOrSecret = []GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret{
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretConfig,
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretSecret,
}

func (e GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret) IsValid() bool {
	switch e {
	case GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretConfig, GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretSecret:
		return true
	}
	return false
}

func (e GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret) String() string {
	return string(e)
}

func (e *GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___operator___apis___crds___v1__ConfigOrSecret", str)
	}
	return nil
}

func (e GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode string

const (
	GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingModePrivate GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode = "private"
	GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingModePublic  GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode = "public"
)

var AllGithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode = []GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode{
	GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingModePrivate,
	GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingModePublic,
}

func (e GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode) IsValid() bool {
	switch e {
	case GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingModePrivate, GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingModePublic:
		return true
	}
	return false
}

func (e GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode) String() string {
	return string(e)
}

func (e *GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___operator___apis___crds___v1__EnvironmentRoutingMode", str)
	}
	return nil
}

func (e GithubComKloudliteOperatorApisCrdsV1EnvironmentRoutingMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type K8sIoAPICoreV1SecretType string

const (
	K8sIoAPICoreV1SecretTypeBootstrapKubernetesIoToken      K8sIoAPICoreV1SecretType = "bootstrap__kubernetes__io___token"
	K8sIoAPICoreV1SecretTypeKubernetesIoBasicAuth           K8sIoAPICoreV1SecretType = "kubernetes__io___basic____auth"
	K8sIoAPICoreV1SecretTypeKubernetesIoDockercfg           K8sIoAPICoreV1SecretType = "kubernetes__io___dockercfg"
	K8sIoAPICoreV1SecretTypeKubernetesIoDockerconfigjson    K8sIoAPICoreV1SecretType = "kubernetes__io___dockerconfigjson"
	K8sIoAPICoreV1SecretTypeKubernetesIoServiceAccountToken K8sIoAPICoreV1SecretType = "kubernetes__io___service____account____token"
	K8sIoAPICoreV1SecretTypeKubernetesIoSSHAuth             K8sIoAPICoreV1SecretType = "kubernetes__io___ssh____auth"
	K8sIoAPICoreV1SecretTypeKubernetesIoTLS                 K8sIoAPICoreV1SecretType = "kubernetes__io___tls"
	K8sIoAPICoreV1SecretTypeOpaque                          K8sIoAPICoreV1SecretType = "Opaque"
)

var AllK8sIoAPICoreV1SecretType = []K8sIoAPICoreV1SecretType{
	K8sIoAPICoreV1SecretTypeBootstrapKubernetesIoToken,
	K8sIoAPICoreV1SecretTypeKubernetesIoBasicAuth,
	K8sIoAPICoreV1SecretTypeKubernetesIoDockercfg,
	K8sIoAPICoreV1SecretTypeKubernetesIoDockerconfigjson,
	K8sIoAPICoreV1SecretTypeKubernetesIoServiceAccountToken,
	K8sIoAPICoreV1SecretTypeKubernetesIoSSHAuth,
	K8sIoAPICoreV1SecretTypeKubernetesIoTLS,
	K8sIoAPICoreV1SecretTypeOpaque,
}

func (e K8sIoAPICoreV1SecretType) IsValid() bool {
	switch e {
	case K8sIoAPICoreV1SecretTypeBootstrapKubernetesIoToken, K8sIoAPICoreV1SecretTypeKubernetesIoBasicAuth, K8sIoAPICoreV1SecretTypeKubernetesIoDockercfg, K8sIoAPICoreV1SecretTypeKubernetesIoDockerconfigjson, K8sIoAPICoreV1SecretTypeKubernetesIoServiceAccountToken, K8sIoAPICoreV1SecretTypeKubernetesIoSSHAuth, K8sIoAPICoreV1SecretTypeKubernetesIoTLS, K8sIoAPICoreV1SecretTypeOpaque:
		return true
	}
	return false
}

func (e K8sIoAPICoreV1SecretType) String() string {
	return string(e)
}

func (e *K8sIoAPICoreV1SecretType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = K8sIoAPICoreV1SecretType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid K8s__io___api___core___v1__SecretType", str)
	}
	return nil
}

func (e K8sIoAPICoreV1SecretType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type K8sIoAPICoreV1TaintEffect string

const (
	K8sIoAPICoreV1TaintEffectNoExecute        K8sIoAPICoreV1TaintEffect = "NoExecute"
	K8sIoAPICoreV1TaintEffectNoSchedule       K8sIoAPICoreV1TaintEffect = "NoSchedule"
	K8sIoAPICoreV1TaintEffectPreferNoSchedule K8sIoAPICoreV1TaintEffect = "PreferNoSchedule"
)

var AllK8sIoAPICoreV1TaintEffect = []K8sIoAPICoreV1TaintEffect{
	K8sIoAPICoreV1TaintEffectNoExecute,
	K8sIoAPICoreV1TaintEffectNoSchedule,
	K8sIoAPICoreV1TaintEffectPreferNoSchedule,
}

func (e K8sIoAPICoreV1TaintEffect) IsValid() bool {
	switch e {
	case K8sIoAPICoreV1TaintEffectNoExecute, K8sIoAPICoreV1TaintEffectNoSchedule, K8sIoAPICoreV1TaintEffectPreferNoSchedule:
		return true
	}
	return false
}

func (e K8sIoAPICoreV1TaintEffect) String() string {
	return string(e)
}

func (e *K8sIoAPICoreV1TaintEffect) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = K8sIoAPICoreV1TaintEffect(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid K8s__io___api___core___v1__TaintEffect", str)
	}
	return nil
}

func (e K8sIoAPICoreV1TaintEffect) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type K8sIoAPICoreV1TolerationOperator string

const (
	K8sIoAPICoreV1TolerationOperatorEqual  K8sIoAPICoreV1TolerationOperator = "Equal"
	K8sIoAPICoreV1TolerationOperatorExists K8sIoAPICoreV1TolerationOperator = "Exists"
)

var AllK8sIoAPICoreV1TolerationOperator = []K8sIoAPICoreV1TolerationOperator{
	K8sIoAPICoreV1TolerationOperatorEqual,
	K8sIoAPICoreV1TolerationOperatorExists,
}

func (e K8sIoAPICoreV1TolerationOperator) IsValid() bool {
	switch e {
	case K8sIoAPICoreV1TolerationOperatorEqual, K8sIoAPICoreV1TolerationOperatorExists:
		return true
	}
	return false
}

func (e K8sIoAPICoreV1TolerationOperator) String() string {
	return string(e)
}

func (e *K8sIoAPICoreV1TolerationOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = K8sIoAPICoreV1TolerationOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid K8s__io___api___core___v1__TolerationOperator", str)
	}
	return nil
}

func (e K8sIoAPICoreV1TolerationOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
