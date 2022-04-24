// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"kloudlite.io/pkg/repos"
)

type Account struct {
	ID       repos.ID   `json:"id"`
	Projects []*Project `json:"projects"`
	Clusters []*Cluster `json:"clusters"`
}

func (Account) IsEntity() {}

type App struct {
	ID                repos.ID        `json:"id"`
	Name              string          `json:"name"`
	Namespace         string          `json:"namespace"`
	Description       *string         `json:"description"`
	ReadableID        repos.ID        `json:"readableId"`
	Replicas          *int            `json:"replicas"`
	Services          []*AppService   `json:"services"`
	AttachedResources []*ManagedRes   `json:"attachedResources"`
	Containers        []*AppContainer `json:"containers"`
	Project           *Project        `json:"project"`
	Version           *int            `json:"version"`
}

type AppContainer struct {
	Name             string        `json:"name"`
	Image            string        `json:"image"`
	ImagePullPolicy  string        `json:"imagePullPolicy"`
	Env              []*AppEnv     `json:"env"`
	ResourceCPU      *ContainerRes `json:"resourceCpu"`
	ResourceMemory   *ContainerRes `json:"resourceMemory"`
	Volumes          []*AppVolume  `json:"volumes"`
	ManagedResources []*ManagedRes `json:"managedResources"`
}

type AppContainerIn struct {
	Name             string             `json:"name"`
	Image            *string            `json:"image"`
	ImagePullPolicy  *string            `json:"imagePullPolicy"`
	Env              []*AppEnvInput     `json:"env"`
	ResourceCPU      *ContainerResInput `json:"resourceCpu"`
	ResourceMemory   *ContainerResInput `json:"resourceMemory"`
	Volumes          []*IAppVolume      `json:"volumes"`
	ManagedResources []string           `json:"managedResources"`
	PullSecret       *string            `json:"pullSecret"`
}

type AppContainerInput struct {
	Name              string         `json:"name"`
	Image             string         `json:"image"`
	PullSecret        *string        `json:"pull_secret"`
	EnvVars           []*EnvVar      `json:"env_vars"`
	CPUMin            string         `json:"cpu_min"`
	CPUMax            string         `json:"cpu_max"`
	MemMin            string         `json:"mem_min"`
	MemMax            string         `json:"mem_max"`
	AttachedResources []*AttachedRes `json:"attached_resources"`
}

type AppContainerUpdateInput struct {
	Name             *string            `json:"name"`
	Image            *string            `json:"image"`
	ImagePullPolicy  *string            `json:"imagePullPolicy"`
	Env              []*AppEnvInput     `json:"env"`
	ResourceCPU      *ContainerResInput `json:"resourceCpu"`
	ResourceMemory   *ContainerResInput `json:"resourceMemory"`
	Volumes          []*IAppVolume      `json:"volumes"`
	ManagedResources []string           `json:"managedResources"`
}

type AppEnv struct {
	Key    string  `json:"key"`
	Type   string  `json:"type"`
	Value  *string `json:"value"`
	RefKey *string `json:"refKey"`
	RefID  *string `json:"refId"`
}

type AppEnvInput struct {
	Key    string  `json:"key"`
	Type   string  `json:"type"`
	Value  *string `json:"value"`
	RefKey *string `json:"refKey"`
	RefID  *string `json:"refId"`
}

type AppFlowInput struct {
	Name            string                 `json:"name"`
	Description     *string                `json:"description"`
	ExposedServices []*ExposedServiceInput `json:"exposed_services"`
	Containers      []*AppContainerInput   `json:"containers"`
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

type AppVolume struct {
	Name      string           `json:"name"`
	MountPath string           `json:"mountPath"`
	Type      string           `json:"type"`
	RefID     repos.ID         `json:"refId"`
	Items     []*AppVolumeItem `json:"items"`
}

type AppVolumeItem struct {
	Key      string `json:"key"`
	FileName string `json:"fileName"`
}

type AttachedRes struct {
	ResID repos.ID `json:"res_id"`
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

type Cluster struct {
	ID          repos.ID  `json:"id"`
	Name        string    `json:"name"`
	Provider    string    `json:"provider"`
	Region      string    `json:"region"`
	IP          *string   `json:"ip"`
	Devices     []*Device `json:"devices"`
	UserDevices []*Device `json:"userDevices"`
	NodesCount  int       `json:"nodesCount"`
	Status      string    `json:"status"`
	Account     *Account  `json:"account"`
}

func (Cluster) IsEntity() {}

type Config struct {
	ID          repos.ID   `json:"id"`
	Name        string     `json:"name"`
	Project     *Project   `json:"project"`
	Description *string    `json:"description"`
	Namespace   string     `json:"namespace"`
	Entries     []*CSEntry `json:"entries"`
}

type ContainerRes struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type ContainerResInput struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type Device struct {
	ID            repos.ID `json:"id"`
	User          *User    `json:"user"`
	Name          string   `json:"name"`
	Cluster       *Cluster `json:"cluster"`
	Configuration string   `json:"configuration"`
	IP            string   `json:"ip"`
}

func (Device) IsEntity() {}

type EnvVal struct {
	Type  string  `json:"type"`
	Value *string `json:"value"`
	Ref   *string `json:"ref"`
	Key   *string `json:"key"`
}

type EnvVar struct {
	Key   string  `json:"key"`
	Value *EnvVal `json:"value"`
}

type ExposedServiceInput struct {
	Type    string `json:"type"`
	Target  int    `json:"target"`
	Exposed int    `json:"exposed"`
}

type GitPipeline struct {
	ID          repos.ID               `json:"id"`
	PipelineEnv string                 `json:"pipelineEnv"`
	GitProvider *string                `json:"gitProvider"`
	GitRepoURL  *string                `json:"gitRepoUrl"`
	BuildArgs   []*Kv                  `json:"buildArgs"`
	PullSecret  *string                `json:"pullSecret"`
	Name        string                 `json:"name"`
	ImageName   string                 `json:"imageName"`
	DockerFile  *string                `json:"dockerFile"`
	ContextDir  *string                `json:"contextDir"`
	Github      map[string]interface{} `json:"github"`
	Gitlab      map[string]interface{} `json:"gitlab"`
	Project     *Project               `json:"project"`
}

type GitPipelineInput struct {
	GitRepoURL  string     `json:"gitRepoUrl"`
	GitProvider string     `json:"gitProvider"`
	DockerFile  *string    `json:"dockerFile"`
	ContextDir  *string    `json:"contextDir"`
	BuildArgs   []*KVInput `json:"buildArgs"`
	PipelineEnv string     `json:"pipelineEnv"`
	PullSecret  *string    `json:"pullSecret"`
}

type IAppVolume struct {
	Name      string            `json:"name"`
	MountPath string            `json:"mountPath"`
	Type      string            `json:"type"`
	RefID     repos.ID          `json:"refId"`
	Items     []*IAppVolumeItem `json:"items"`
}

type IAppVolumeItem struct {
	Key      string `json:"key"`
	FileName string `json:"fileName"`
}

type IAppVolumeUpdate struct {
	MountPath *string           `json:"mountPath"`
	Type      *string           `json:"type"`
	RefID     *repos.ID         `json:"refId"`
	Items     []*IAppVolumeItem `json:"items"`
}

type Kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type KVInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ManagedRes struct {
	ID           repos.ID               `json:"id"`
	Name         string                 `json:"name"`
	ResourceType string                 `json:"resourceType"`
	Installation *ManagedSvc            `json:"installation"`
	Values       map[string]interface{} `json:"values"`
}

type ManagedSvc struct {
	ID        repos.ID               `json:"id"`
	Name      string                 `json:"name"`
	Project   *Project               `json:"project"`
	Source    string                 `json:"source"`
	Values    map[string]interface{} `json:"values"`
	Resources []*ManagedRes          `json:"resources"`
}

type NewResourcesIn struct {
	Configs    []map[string]interface{} `json:"configs"`
	Secrets    []map[string]interface{} `json:"secrets"`
	MServices  []map[string]interface{} `json:"mServices"`
	MResources []map[string]interface{} `json:"mResources"`
}

type Project struct {
	ID          repos.ID             `json:"id"`
	Name        string               `json:"name"`
	DisplayName string               `json:"displayName"`
	ReadableID  repos.ID             `json:"readableId"`
	Logo        *string              `json:"logo"`
	Description *string              `json:"description"`
	Account     *Account             `json:"account"`
	Memberships []*ProjectMembership `json:"memberships"`
}

type ProjectMembership struct {
	User    *User    `json:"user"`
	Role    string   `json:"role"`
	Project *Project `json:"project"`
}

type Route struct {
	Path    string `json:"path"`
	AppName string `json:"appName"`
	Port    int    `json:"port"`
}

type RouteInput struct {
	Path    string `json:"path"`
	AppName string `json:"appName"`
	Port    int    `json:"port"`
}

type Router struct {
	ID      repos.ID `json:"id"`
	Name    string   `json:"name"`
	Project *Project `json:"project"`
	Domains []string `json:"domains"`
	Routes  []*Route `json:"routes"`
}

type Secret struct {
	ID          repos.ID   `json:"id"`
	Name        string     `json:"name"`
	Project     *Project   `json:"project"`
	Description *string    `json:"description"`
	Namespace   string     `json:"namespace"`
	Entries     []*CSEntry `json:"entries"`
}

type User struct {
	ID      repos.ID  `json:"id"`
	Devices []*Device `json:"devices"`
}

func (User) IsEntity() {}
