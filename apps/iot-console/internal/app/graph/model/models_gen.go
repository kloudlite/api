// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/kloudlite/api/apps/iot-console/internal/entities"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/repos"
)

type GithubComKloudliteAPIAppsIotConsoleInternalEntitiesExposedService struct {
	IP   string `json:"ip"`
	Name string `json:"name"`
}

type GithubComKloudliteAPIAppsIotConsoleInternalEntitiesExposedServiceIn struct {
	IP   string `json:"ip"`
	Name string `json:"name"`
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

type GithubComKloudliteOperatorApisCrdsV1AppInterceptPortMappings struct {
	AppPort    int `json:"appPort"`
	DevicePort int `json:"devicePort"`
}

type GithubComKloudliteOperatorApisCrdsV1AppInterceptPortMappingsIn struct {
	AppPort    int `json:"appPort"`
	DevicePort int `json:"devicePort"`
}

type GithubComKloudliteOperatorApisCrdsV1AppRouter struct {
	BackendProtocol *string                                        `json:"backendProtocol,omitempty"`
	BasicAuth       *GithubComKloudliteOperatorApisCrdsV1BasicAuth `json:"basicAuth,omitempty"`
	Cors            *GithubComKloudliteOperatorApisCrdsV1Cors      `json:"cors,omitempty"`
	Domains         []string                                       `json:"domains"`
	HTTPS           *GithubComKloudliteOperatorApisCrdsV1HTTPS     `json:"https,omitempty"`
	IngressClass    *string                                        `json:"ingressClass,omitempty"`
	MaxBodySizeInMb *int                                           `json:"maxBodySizeInMB,omitempty"`
	RateLimit       *GithubComKloudliteOperatorApisCrdsV1RateLimit `json:"rateLimit,omitempty"`
	Routes          []*GithubComKloudliteOperatorApisCrdsV1Route   `json:"routes,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppRouterIn struct {
	BackendProtocol *string                                          `json:"backendProtocol,omitempty"`
	BasicAuth       *GithubComKloudliteOperatorApisCrdsV1BasicAuthIn `json:"basicAuth,omitempty"`
	Cors            *GithubComKloudliteOperatorApisCrdsV1CorsIn      `json:"cors,omitempty"`
	Domains         []string                                         `json:"domains"`
	HTTPS           *GithubComKloudliteOperatorApisCrdsV1HTTPSIn     `json:"https,omitempty"`
	IngressClass    *string                                          `json:"ingressClass,omitempty"`
	MaxBodySizeInMb *int                                             `json:"maxBodySizeInMB,omitempty"`
	RateLimit       *GithubComKloudliteOperatorApisCrdsV1RateLimitIn `json:"rateLimit,omitempty"`
	Routes          []*GithubComKloudliteOperatorApisCrdsV1RouteIn   `json:"routes,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppSpec struct {
	Containers                []*GithubComKloudliteOperatorApisCrdsV1AppContainer `json:"containers"`
	DisplayName               *string                                             `json:"displayName,omitempty"`
	Freeze                    *bool                                               `json:"freeze,omitempty"`
	Hpa                       *GithubComKloudliteOperatorApisCrdsV1Hpa            `json:"hpa,omitempty"`
	Intercept                 *GithubComKloudliteOperatorApisCrdsV1Intercept      `json:"intercept,omitempty"`
	NodeSelector              map[string]interface{}                              `json:"nodeSelector,omitempty"`
	Region                    *string                                             `json:"region,omitempty"`
	Replicas                  *int                                                `json:"replicas,omitempty"`
	Router                    *GithubComKloudliteOperatorApisCrdsV1AppRouter      `json:"router,omitempty"`
	ServiceAccount            *string                                             `json:"serviceAccount,omitempty"`
	Services                  []*GithubComKloudliteOperatorApisCrdsV1AppSvc       `json:"services,omitempty"`
	Tolerations               []*K8sIoAPICoreV1Toleration                         `json:"tolerations,omitempty"`
	TopologySpreadConstraints []*K8sIoAPICoreV1TopologySpreadConstraint           `json:"topologySpreadConstraints,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppSpecIn struct {
	Containers                []*GithubComKloudliteOperatorApisCrdsV1AppContainerIn `json:"containers"`
	DisplayName               *string                                               `json:"displayName,omitempty"`
	Freeze                    *bool                                                 `json:"freeze,omitempty"`
	Hpa                       *GithubComKloudliteOperatorApisCrdsV1HPAIn            `json:"hpa,omitempty"`
	Intercept                 *GithubComKloudliteOperatorApisCrdsV1InterceptIn      `json:"intercept,omitempty"`
	NodeSelector              map[string]interface{}                                `json:"nodeSelector,omitempty"`
	Region                    *string                                               `json:"region,omitempty"`
	Replicas                  *int                                                  `json:"replicas,omitempty"`
	Router                    *GithubComKloudliteOperatorApisCrdsV1AppRouterIn      `json:"router,omitempty"`
	ServiceAccount            *string                                               `json:"serviceAccount,omitempty"`
	Services                  []*GithubComKloudliteOperatorApisCrdsV1AppSvcIn       `json:"services,omitempty"`
	Tolerations               []*K8sIoAPICoreV1TolerationIn                         `json:"tolerations,omitempty"`
	TopologySpreadConstraints []*K8sIoAPICoreV1TopologySpreadConstraintIn           `json:"topologySpreadConstraints,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppSvc struct {
	Port     int     `json:"port"`
	Protocol *string `json:"protocol,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1AppSvcIn struct {
	Port     int     `json:"port"`
	Protocol *string `json:"protocol,omitempty"`
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

type GithubComKloudliteOperatorApisCrdsV1Hpa struct {
	Enabled         bool `json:"enabled"`
	MaxReplicas     *int `json:"maxReplicas,omitempty"`
	MinReplicas     *int `json:"minReplicas,omitempty"`
	ThresholdCPU    *int `json:"thresholdCpu,omitempty"`
	ThresholdMemory *int `json:"thresholdMemory,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1HPAIn struct {
	Enabled         bool `json:"enabled"`
	MaxReplicas     *int `json:"maxReplicas,omitempty"`
	MinReplicas     *int `json:"minReplicas,omitempty"`
	ThresholdCPU    *int `json:"thresholdCpu,omitempty"`
	ThresholdMemory *int `json:"thresholdMemory,omitempty"`
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
	Enabled      bool                                                            `json:"enabled"`
	PortMappings []*GithubComKloudliteOperatorApisCrdsV1AppInterceptPortMappings `json:"portMappings,omitempty"`
	ToDevice     string                                                          `json:"toDevice"`
}

type GithubComKloudliteOperatorApisCrdsV1InterceptIn struct {
	Enabled      bool                                                              `json:"enabled"`
	PortMappings []*GithubComKloudliteOperatorApisCrdsV1AppInterceptPortMappingsIn `json:"portMappings,omitempty"`
	ToDevice     string                                                            `json:"toDevice"`
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
	App     string `json:"app"`
	Path    string `json:"path"`
	Port    int    `json:"port"`
	Rewrite *bool  `json:"rewrite,omitempty"`
}

type GithubComKloudliteOperatorApisCrdsV1RouteIn struct {
	App     string `json:"app"`
	Path    string `json:"path"`
	Port    int    `json:"port"`
	Rewrite *bool  `json:"rewrite,omitempty"`
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
	Debug      *string                                     `json:"debug,omitempty"`
	Error      *string                                     `json:"error,omitempty"`
	Generation *int                                        `json:"generation,omitempty"`
	Info       *string                                     `json:"info,omitempty"`
	Message    *string                                     `json:"message,omitempty"`
	StartedAt  *string                                     `json:"startedAt,omitempty"`
	State      *GithubComKloudliteOperatorPkgOperatorState `json:"state,omitempty"`
	Status     bool                                        `json:"status"`
}

type GithubComKloudliteOperatorPkgOperatorCheckMeta struct {
	Debug       *bool   `json:"debug,omitempty"`
	Description *string `json:"description,omitempty"`
	Hide        *bool   `json:"hide,omitempty"`
	Name        string  `json:"name"`
	Title       string  `json:"title"`
}

type GithubComKloudliteOperatorPkgOperatorResourceRef struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
}

type GithubComKloudliteOperatorPkgOperatorStatus struct {
	CheckList           []*GithubComKloudliteOperatorPkgOperatorCheckMeta   `json:"checkList,omitempty"`
	Checks              map[string]interface{}                              `json:"checks,omitempty"`
	IsReady             bool                                                `json:"isReady"`
	LastReadyGeneration *int                                                `json:"lastReadyGeneration,omitempty"`
	LastReconcileTime   *string                                             `json:"lastReconcileTime,omitempty"`
	Message             *GithubComKloudliteOperatorPkgRawJSONRawJSON        `json:"message,omitempty"`
	Resources           []*GithubComKloudliteOperatorPkgOperatorResourceRef `json:"resources,omitempty"`
}

type GithubComKloudliteOperatorPkgRawJSONRawJSON struct {
	RawMessage interface{} `json:"RawMessage,omitempty"`
}

type IOTAppEdge struct {
	Cursor string           `json:"cursor"`
	Node   *entities.IOTApp `json:"node"`
}

type IOTAppPaginatedRecords struct {
	Edges      []*IOTAppEdge `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

type IOTDeploymentEdge struct {
	Cursor string                  `json:"cursor"`
	Node   *entities.IOTDeployment `json:"node"`
}

type IOTDeploymentPaginatedRecords struct {
	Edges      []*IOTDeploymentEdge `json:"edges"`
	PageInfo   *PageInfo            `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}

type IOTDeviceBlueprintEdge struct {
	Cursor string                       `json:"cursor"`
	Node   *entities.IOTDeviceBlueprint `json:"node"`
}

type IOTDeviceBlueprintPaginatedRecords struct {
	Edges      []*IOTDeviceBlueprintEdge `json:"edges"`
	PageInfo   *PageInfo                 `json:"pageInfo"`
	TotalCount int                       `json:"totalCount"`
}

type IOTDeviceEdge struct {
	Cursor string              `json:"cursor"`
	Node   *entities.IOTDevice `json:"node"`
}

type IOTDevicePaginatedRecords struct {
	Edges      []*IOTDeviceEdge `json:"edges"`
	PageInfo   *PageInfo        `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

type IOTEnvironment struct {
	AccountName       string                     `json:"accountName"`
	CreatedBy         *common.CreatedOrUpdatedBy `json:"createdBy"`
	CreationTime      string                     `json:"creationTime"`
	DisplayName       string                     `json:"displayName"`
	ID                repos.ID                   `json:"id"`
	LastUpdatedBy     *common.CreatedOrUpdatedBy `json:"lastUpdatedBy"`
	MarkedForDeletion *bool                      `json:"markedForDeletion,omitempty"`
	Name              string                     `json:"name"`
	ProjectName       string                     `json:"projectName"`
	RecordVersion     int                        `json:"recordVersion"`
	UpdateTime        string                     `json:"updateTime"`
}

type IOTEnvironmentEdge struct {
	Cursor string          `json:"cursor"`
	Node   *IOTEnvironment `json:"node"`
}

type IOTEnvironmentIn struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
}

type IOTEnvironmentPaginatedRecords struct {
	Edges      []*IOTEnvironmentEdge `json:"edges"`
	PageInfo   *PageInfo             `json:"pageInfo"`
	TotalCount int                   `json:"totalCount"`
}

type IOTProjectEdge struct {
	Cursor string               `json:"cursor"`
	Node   *entities.IOTProject `json:"node"`
}

type IOTProjectPaginatedRecords struct {
	Edges      []*IOTProjectEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
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

type K8sIoAPICoreV1TopologySpreadConstraint struct {
	LabelSelector      *K8sIoApimachineryPkgApisMetaV1LabelSelector `json:"labelSelector,omitempty"`
	MatchLabelKeys     []string                                     `json:"matchLabelKeys,omitempty"`
	MaxSkew            int                                          `json:"maxSkew"`
	MinDomains         *int                                         `json:"minDomains,omitempty"`
	NodeAffinityPolicy *string                                      `json:"nodeAffinityPolicy,omitempty"`
	NodeTaintsPolicy   *string                                      `json:"nodeTaintsPolicy,omitempty"`
	TopologyKey        string                                       `json:"topologyKey"`
	WhenUnsatisfiable  K8sIoAPICoreV1UnsatisfiableConstraintAction  `json:"whenUnsatisfiable"`
}

type K8sIoAPICoreV1TopologySpreadConstraintIn struct {
	LabelSelector      *K8sIoApimachineryPkgApisMetaV1LabelSelectorIn `json:"labelSelector,omitempty"`
	MatchLabelKeys     []string                                       `json:"matchLabelKeys,omitempty"`
	MaxSkew            int                                            `json:"maxSkew"`
	MinDomains         *int                                           `json:"minDomains,omitempty"`
	NodeAffinityPolicy *string                                        `json:"nodeAffinityPolicy,omitempty"`
	NodeTaintsPolicy   *string                                        `json:"nodeTaintsPolicy,omitempty"`
	TopologyKey        string                                         `json:"topologyKey"`
	WhenUnsatisfiable  K8sIoAPICoreV1UnsatisfiableConstraintAction    `json:"whenUnsatisfiable"`
}

type K8sIoApimachineryPkgApisMetaV1LabelSelector struct {
	MatchExpressions []*K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement `json:"matchExpressions,omitempty"`
	MatchLabels      map[string]interface{}                                    `json:"matchLabels,omitempty"`
}

type K8sIoApimachineryPkgApisMetaV1LabelSelectorIn struct {
	MatchExpressions []*K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirementIn `json:"matchExpressions,omitempty"`
	MatchLabels      map[string]interface{}                                      `json:"matchLabels,omitempty"`
}

type K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement struct {
	Key      string                                              `json:"key"`
	Operator K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator `json:"operator"`
	Values   []string                                            `json:"values,omitempty"`
}

type K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirementIn struct {
	Key      string                                              `json:"key"`
	Operator K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator `json:"operator"`
	Values   []string                                            `json:"values,omitempty"`
}

type Mutation struct {
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor,omitempty"`
	HasNextPage *bool   `json:"hasNextPage,omitempty"`
	HasPrevPage *bool   `json:"hasPrevPage,omitempty"`
	StartCursor *string `json:"startCursor,omitempty"`
}

type Query struct {
}

type SearchIOTApps struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchIOTDeployments struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchIOTDeviceBlueprints struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchIOTDevices struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type SearchIOTProjects struct {
	Text              *repos.MatchFilter `json:"text,omitempty"`
	IsReady           *repos.MatchFilter `json:"isReady,omitempty"`
	MarkedForDeletion *repos.MatchFilter `json:"markedForDeletion,omitempty"`
}

type GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType string

const (
	GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintTypeGroupBlueprint     GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType = "group_blueprint"
	GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintTypeSingletonBlueprint GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType = "singleton_blueprint"
)

var AllGithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType = []GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType{
	GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintTypeGroupBlueprint,
	GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintTypeSingletonBlueprint,
}

func (e GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType) IsValid() bool {
	switch e {
	case GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintTypeGroupBlueprint, GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintTypeSingletonBlueprint:
		return true
	}
	return false
}

func (e GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType) String() string {
	return string(e)
}

func (e *GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___api___apps___iot____console___internal___entities__BluePrintType", str)
	}
	return nil
}

func (e GithubComKloudliteAPIAppsIotConsoleInternalEntitiesBluePrintType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GithubComKloudliteAPIPkgReposMatchType string

const (
	GithubComKloudliteAPIPkgReposMatchTypeArray      GithubComKloudliteAPIPkgReposMatchType = "array"
	GithubComKloudliteAPIPkgReposMatchTypeExact      GithubComKloudliteAPIPkgReposMatchType = "exact"
	GithubComKloudliteAPIPkgReposMatchTypeNotInArray GithubComKloudliteAPIPkgReposMatchType = "not_in_array"
	GithubComKloudliteAPIPkgReposMatchTypeRegex      GithubComKloudliteAPIPkgReposMatchType = "regex"
)

var AllGithubComKloudliteAPIPkgReposMatchType = []GithubComKloudliteAPIPkgReposMatchType{
	GithubComKloudliteAPIPkgReposMatchTypeArray,
	GithubComKloudliteAPIPkgReposMatchTypeExact,
	GithubComKloudliteAPIPkgReposMatchTypeNotInArray,
	GithubComKloudliteAPIPkgReposMatchTypeRegex,
}

func (e GithubComKloudliteAPIPkgReposMatchType) IsValid() bool {
	switch e {
	case GithubComKloudliteAPIPkgReposMatchTypeArray, GithubComKloudliteAPIPkgReposMatchTypeExact, GithubComKloudliteAPIPkgReposMatchTypeNotInArray, GithubComKloudliteAPIPkgReposMatchTypeRegex:
		return true
	}
	return false
}

func (e GithubComKloudliteAPIPkgReposMatchType) String() string {
	return string(e)
}

func (e *GithubComKloudliteAPIPkgReposMatchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteAPIPkgReposMatchType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___api___pkg___repos__MatchType", str)
	}
	return nil
}

func (e GithubComKloudliteAPIPkgReposMatchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret string

const (
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretConfig GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret = "config"
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretPvc    GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret = "pvc"
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretSecret GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret = "secret"
)

var AllGithubComKloudliteOperatorApisCrdsV1ConfigOrSecret = []GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret{
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretConfig,
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretPvc,
	GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretSecret,
}

func (e GithubComKloudliteOperatorApisCrdsV1ConfigOrSecret) IsValid() bool {
	switch e {
	case GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretConfig, GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretPvc, GithubComKloudliteOperatorApisCrdsV1ConfigOrSecretSecret:
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

type GithubComKloudliteOperatorPkgOperatorState string

const (
	GithubComKloudliteOperatorPkgOperatorStateErroredDuringReconcilation GithubComKloudliteOperatorPkgOperatorState = "errored____during____reconcilation"
	GithubComKloudliteOperatorPkgOperatorStateFinishedReconcilation      GithubComKloudliteOperatorPkgOperatorState = "finished____reconcilation"
	GithubComKloudliteOperatorPkgOperatorStateUnderReconcilation         GithubComKloudliteOperatorPkgOperatorState = "under____reconcilation"
	GithubComKloudliteOperatorPkgOperatorStateYetToBeReconciled          GithubComKloudliteOperatorPkgOperatorState = "yet____to____be____reconciled"
)

var AllGithubComKloudliteOperatorPkgOperatorState = []GithubComKloudliteOperatorPkgOperatorState{
	GithubComKloudliteOperatorPkgOperatorStateErroredDuringReconcilation,
	GithubComKloudliteOperatorPkgOperatorStateFinishedReconcilation,
	GithubComKloudliteOperatorPkgOperatorStateUnderReconcilation,
	GithubComKloudliteOperatorPkgOperatorStateYetToBeReconciled,
}

func (e GithubComKloudliteOperatorPkgOperatorState) IsValid() bool {
	switch e {
	case GithubComKloudliteOperatorPkgOperatorStateErroredDuringReconcilation, GithubComKloudliteOperatorPkgOperatorStateFinishedReconcilation, GithubComKloudliteOperatorPkgOperatorStateUnderReconcilation, GithubComKloudliteOperatorPkgOperatorStateYetToBeReconciled:
		return true
	}
	return false
}

func (e GithubComKloudliteOperatorPkgOperatorState) String() string {
	return string(e)
}

func (e *GithubComKloudliteOperatorPkgOperatorState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GithubComKloudliteOperatorPkgOperatorState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Github__com___kloudlite___operator___pkg___operator__State", str)
	}
	return nil
}

func (e GithubComKloudliteOperatorPkgOperatorState) MarshalGQL(w io.Writer) {
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

type K8sIoAPICoreV1UnsatisfiableConstraintAction string

const (
	K8sIoAPICoreV1UnsatisfiableConstraintActionDoNotSchedule  K8sIoAPICoreV1UnsatisfiableConstraintAction = "DoNotSchedule"
	K8sIoAPICoreV1UnsatisfiableConstraintActionScheduleAnyway K8sIoAPICoreV1UnsatisfiableConstraintAction = "ScheduleAnyway"
)

var AllK8sIoAPICoreV1UnsatisfiableConstraintAction = []K8sIoAPICoreV1UnsatisfiableConstraintAction{
	K8sIoAPICoreV1UnsatisfiableConstraintActionDoNotSchedule,
	K8sIoAPICoreV1UnsatisfiableConstraintActionScheduleAnyway,
}

func (e K8sIoAPICoreV1UnsatisfiableConstraintAction) IsValid() bool {
	switch e {
	case K8sIoAPICoreV1UnsatisfiableConstraintActionDoNotSchedule, K8sIoAPICoreV1UnsatisfiableConstraintActionScheduleAnyway:
		return true
	}
	return false
}

func (e K8sIoAPICoreV1UnsatisfiableConstraintAction) String() string {
	return string(e)
}

func (e *K8sIoAPICoreV1UnsatisfiableConstraintAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = K8sIoAPICoreV1UnsatisfiableConstraintAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid K8s__io___api___core___v1__UnsatisfiableConstraintAction", str)
	}
	return nil
}

func (e K8sIoAPICoreV1UnsatisfiableConstraintAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator string

const (
	K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorDoesNotExist K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator = "DoesNotExist"
	K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorExists       K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator = "Exists"
	K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorIn           K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator = "In"
	K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorNotIn        K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator = "NotIn"
)

var AllK8sIoApimachineryPkgApisMetaV1LabelSelectorOperator = []K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator{
	K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorDoesNotExist,
	K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorExists,
	K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorIn,
	K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorNotIn,
}

func (e K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator) IsValid() bool {
	switch e {
	case K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorDoesNotExist, K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorExists, K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorIn, K8sIoApimachineryPkgApisMetaV1LabelSelectorOperatorNotIn:
		return true
	}
	return false
}

func (e K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator) String() string {
	return string(e)
}

func (e *K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid K8s__io___apimachinery___pkg___apis___meta___v1__LabelSelectorOperator", str)
	}
	return nil
}

func (e K8sIoApimachineryPkgApisMetaV1LabelSelectorOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
