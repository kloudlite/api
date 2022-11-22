// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"kloudlite.io/pkg/repos"
)

type Account struct {
	ID       repos.ID   `json:"id"`
	Projects []*Project `json:"projects"`
	Devices  []*Device  `json:"devices"`
}

func (Account) IsEntity() {}

type App struct {
	ID               repos.ID          `json:"id"`
	IsLambda         bool              `json:"isLambda"`
	Name             string            `json:"name"`
	Namespace        string            `json:"namespace"`
	CreatedAt        string            `json:"createdAt"`
	UpdatedAt        *string           `json:"updatedAt"`
	Description      *string           `json:"description"`
	ReadableID       repos.ID          `json:"readableId"`
	Replicas         *int              `json:"replicas"`
	Services         []*ExposedService `json:"services"`
	Containers       []*AppContainer   `json:"containers"`
	Project          *Project          `json:"project"`
	Status           string            `json:"status"`
	AutoScale        *AutoScale        `json:"autoScale"`
	Conditions       []*MetaCondition  `json:"conditions"`
	Restart          bool              `json:"restart"`
	DoFreeze         bool              `json:"doFreeze"`
	DoUnfreeze       bool              `json:"doUnfreeze"`
	IsFrozen         bool              `json:"isFrozen"`
	CurrentIntercept *repos.ID         `json:"currentIntercept"`
	Intercept        bool              `json:"intercept"`
	CloseIntercept   bool              `json:"closeIntercept"`
}

func (App) IsEntity() {}

type AppContainer struct {
	Name              string         `json:"name"`
	Image             *string        `json:"image"`
	PullSecret        *string        `json:"pullSecret"`
	EnvVars           []*EnvVar      `json:"envVars"`
	AttachedResources []*AttachedRes `json:"attachedResources"`
	ComputePlan       string         `json:"computePlan"`
	Quantity          float64        `json:"quantity"`
	IsShared          *bool          `json:"isShared"`
	Mounts            []*Mount       `json:"mounts"`
}

type AppContainerIn struct {
	Name              string              `json:"name"`
	Image             *string             `json:"image"`
	PullSecret        *string             `json:"pullSecret"`
	EnvVars           []*EnvVarInput      `json:"envVars"`
	Mounts            []*MountInput       `json:"mounts"`
	ComputePlan       string              `json:"computePlan"`
	Quantity          float64             `json:"quantity"`
	AttachedResources []*AttachedResInput `json:"attachedResources"`
	IsShared          *bool               `json:"isShared"`
}

type AppInput struct {
	Name        string                 `json:"name"`
	IsLambda    bool                   `json:"isLambda"`
	ProjectID   string                 `json:"projectId"`
	Description *string                `json:"description"`
	AutoScale   *AutoScaleIn           `json:"autoScale"`
	ReadableID  repos.ID               `json:"readableId"`
	Replicas    *int                   `json:"replicas"`
	Services    []*ExposedServiceIn    `json:"services"`
	Containers  []*AppContainerIn      `json:"containers"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type AppService struct {
	Type       string `json:"type"`
	Port       int    `json:"port"`
	TargetPort *int   `json:"targetPort"`
}

type AppServiceInput struct {
	Type       string `json:"type"`
	Port       int    `json:"port"`
	TargetPort *int   `json:"targetPort"`
}

type AttachedRes struct {
	ResID repos.ID `json:"res_id"`
}

type AttachedResInput struct {
	ResID repos.ID `json:"res_id"`
}

type AutoScale struct {
	MinReplicas     int `json:"minReplicas"`
	MaxReplicas     int `json:"maxReplicas"`
	UsagePercentage int `json:"usage_percentage"`
}

type AutoScaleIn struct {
	MinReplicas     int `json:"minReplicas"`
	MaxReplicas     int `json:"maxReplicas"`
	UsagePercentage int `json:"usage_percentage"`
}

type CCMData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CSEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CSEntryIn struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CloudProvider struct {
	ID       repos.ID      `json:"id"`
	Name     string        `json:"name"`
	Provider string        `json:"provider"`
	Edges    []*EdgeRegion `json:"edges"`
	Status   string        `json:"status"`
	IsShared bool          `json:"isShared"`
}

type CloudProviderIn struct {
	Name        string                 `json:"name"`
	Provider    string                 `json:"provider"`
	Credentials map[string]interface{} `json:"credentials"`
}

type CloudProviderUpdateIn struct {
	Name        *string                `json:"name"`
	Credentials map[string]interface{} `json:"credentials"`
}

type ClusterIn struct {
	Name       string `json:"name"`
	SubDomain  string `json:"subDomain"`
	KubeConfig string `json:"kubeConfig"`
}

type ClusterOut struct {
	ID        repos.ID `json:"id"`
	Name      string   `json:"name"`
	SubDomain string   `json:"subDomain"`
}

type ComputePlan struct {
	Name                  string `json:"name"`
	Desc                  string `json:"desc"`
	SharingEnabled        bool   `json:"sharingEnabled"`
	DedicatedEnabled      bool   `json:"dedicatedEnabled"`
	MemoryPerVCPUCpu      int    `json:"memoryPerVCPUCpu"`
	MaxDedicatedCPUPerPod int    `json:"maxDedicatedCPUPerPod"`
	MaxSharedCPUPerPod    int    `json:"maxSharedCPUPerPod"`
}

func (ComputePlan) IsEntity() {}

type Config struct {
	ID          repos.ID   `json:"id"`
	Name        string     `json:"name"`
	Project     *Project   `json:"project"`
	Description *string    `json:"description"`
	Namespace   string     `json:"namespace"`
	Entries     []*CSEntry `json:"entries"`
	Status      string     `json:"status"`
}

type Device struct {
	ID                   repos.ID               `json:"id"`
	User                 *User                  `json:"user"`
	Name                 string                 `json:"name"`
	Configuration        map[string]interface{} `json:"configuration"`
	Account              *Account               `json:"account"`
	Ports                []*Port                `json:"ports"`
	Region               *string                `json:"region"`
	InterceptingServices []*App                 `json:"interceptingServices"`
}

func (Device) IsEntity() {}

type DeviceIn struct {
	ID     repos.ID  `json:"id"`
	Name   string    `json:"name"`
	Region string    `json:"region"`
	Ports  []*PortIn `json:"ports"`
}

type DockerCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EdgeNode struct {
	NodeIndex    int                    `json:"nodeIndex"`
	Status       map[string]interface{} `json:"status"`
	Name         string                 `json:"name"`
	Config       string                 `json:"config"`
	CreationTime string                 `json:"creationTime"`
}

type EdgeRegion struct {
	ID        repos.ID       `json:"id"`
	Name      string         `json:"name"`
	Region    string         `json:"region"`
	Provider  *CloudProvider `json:"provider"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt *string        `json:"updatedAt"`
	Pools     []*NodePool    `json:"pools"`
	Status    string         `json:"status"`
}

type EdgeRegionIn struct {
	Name   string        `json:"name"`
	Region string        `json:"region"`
	Pools  []*NodePoolIn `json:"pools"`
}

type EdgeRegionUpdateIn struct {
	Name  *string       `json:"name"`
	Pools []*NodePoolIn `json:"pools"`
}

type EnvVal struct {
	Type  string  `json:"type"`
	Value *string `json:"value"`
	Ref   *string `json:"ref"`
	Key   *string `json:"key"`
}

type EnvValInput struct {
	Type  string  `json:"type"`
	Value *string `json:"value"`
	Ref   *string `json:"ref"`
	Key   *string `json:"key"`
}

type EnvVar struct {
	Key   string  `json:"key"`
	Value *EnvVal `json:"value"`
}

type EnvVarInput struct {
	Key   string       `json:"key"`
	Value *EnvValInput `json:"value"`
}

type ExposedService struct {
	Type    string `json:"type"`
	Target  int    `json:"target"`
	Exposed int    `json:"exposed"`
}

type ExposedServiceIn struct {
	Type    string `json:"type"`
	Target  int    `json:"target"`
	Exposed int    `json:"exposed"`
}

type Kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type KVInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type LambdaPlan struct {
	Name string `json:"name"`
}

func (LambdaPlan) IsEntity() {}

type LoadEnv struct {
	EnvVars    map[string]interface{} `json:"envVars"`
	MountFiles map[string]interface{} `json:"mountFiles"`
}

type ManagedRes struct {
	ID           repos.ID               `json:"id"`
	Name         string                 `json:"name"`
	ResourceType string                 `json:"resourceType"`
	Installation *ManagedSvc            `json:"installation"`
	Values       map[string]interface{} `json:"values"`
	Outputs      map[string]interface{} `json:"outputs"`
	Status       string                 `json:"status"`
	CreatedAt    string                 `json:"createdAt"`
	UpdatedAt    *string                `json:"updatedAt"`
}

type ManagedSvc struct {
	ID         repos.ID               `json:"id"`
	Name       string                 `json:"name"`
	Project    *Project               `json:"project"`
	Source     string                 `json:"source"`
	Values     map[string]interface{} `json:"values"`
	Resources  []*ManagedRes          `json:"resources"`
	Status     string                 `json:"status"`
	Conditions []*MetaCondition       `json:"conditions"`
	Outputs    map[string]interface{} `json:"outputs"`
	CreatedAt  string                 `json:"createdAt"`
	UpdatedAt  *string                `json:"updatedAt"`
}

type MetaCondition struct {
	Status        string `json:"status"`
	ConditionType string `json:"conditionType"`
	LastTimeStamp string `json:"lastTimeStamp"`
	Reason        string `json:"reason"`
	Message       string `json:"message"`
}

type Mount struct {
	Type string `json:"type"`
	Ref  string `json:"ref"`
	Path string `json:"path"`
}

type MountInput struct {
	Type string `json:"type"`
	Ref  string `json:"ref"`
	Path string `json:"path"`
}

type NewResourcesIn struct {
	Configs    []map[string]interface{} `json:"configs"`
	Secrets    []map[string]interface{} `json:"secrets"`
	MServices  []map[string]interface{} `json:"mServices"`
	MResources []map[string]interface{} `json:"mResources"`
}

type NodePool struct {
	Name   string `json:"name"`
	Config string `json:"config"`
	Min    int    `json:"min"`
	Max    int    `json:"max"`
}

type NodePoolIn struct {
	Name   string `json:"name"`
	Config string `json:"config"`
	Min    int    `json:"min"`
	Max    int    `json:"max"`
}

type Port struct {
	Port       int  `json:"port"`
	TargetPort *int `json:"targetPort"`
}

type PortIn struct {
	Port       int  `json:"port"`
	TargetPort *int `json:"targetPort"`
}

type Project struct {
	ID                repos.ID             `json:"id"`
	Name              string               `json:"name"`
	DisplayName       string               `json:"displayName"`
	ReadableID        repos.ID             `json:"readableId"`
	Logo              *string              `json:"logo"`
	Description       *string              `json:"description"`
	Account           *Account             `json:"account"`
	Memberships       []*ProjectMembership `json:"memberships"`
	Status            string               `json:"status"`
	Cluster           *string              `json:"cluster"`
	DockerCredentials *DockerCredentials   `json:"dockerCredentials"`
	RegionID          repos.ID             `json:"regionId"`
	Region            *EdgeRegion          `json:"region"`
}

type ProjectMembership struct {
	User    *User    `json:"user"`
	Role    string   `json:"role"`
	Project *Project `json:"project"`
}

type Route struct {
	Path    string `json:"path"`
	AppName string `json:"appName"`
	Port    *int   `json:"port"`
}

type RouteInput struct {
	Path    string `json:"path"`
	AppName string `json:"appName"`
	Port    *int   `json:"port"`
}

type Router struct {
	ID      repos.ID `json:"id"`
	Name    string   `json:"name"`
	Project *Project `json:"project"`
	Domains []string `json:"domains"`
	Routes  []*Route `json:"routes"`
	Status  string   `json:"status"`
}

type Secret struct {
	ID          repos.ID   `json:"id"`
	Name        string     `json:"name"`
	Project     *Project   `json:"project"`
	Description *string    `json:"description"`
	Namespace   string     `json:"namespace"`
	Entries     []*CSEntry `json:"entries"`
	Status      string     `json:"status"`
}

type StoragePlan struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (StoragePlan) IsEntity() {}

type User struct {
	ID      repos.ID  `json:"id"`
	Devices []*Device `json:"devices"`
}

func (User) IsEntity() {}
