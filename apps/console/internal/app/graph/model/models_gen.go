// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AppSpec struct {
	Services       []*AppSpecServices     `json:"services,omitempty"`
	DisplayName    *string                `json:"displayName,omitempty"`
	Freeze         *bool                  `json:"freeze,omitempty"`
	Intercept      *AppSpecIntercept      `json:"intercept,omitempty"`
	NodeSelector   map[string]interface{} `json:"nodeSelector,omitempty"`
	Region         *string                `json:"region,omitempty"`
	Replicas       *int                   `json:"replicas,omitempty"`
	ServiceAccount *string                `json:"serviceAccount,omitempty"`
	Tolerations    []*AppSpecTolerations  `json:"tolerations,omitempty"`
	Containers     []*AppSpecContainers   `json:"containers"`
	Hpa            *AppSpecHpa            `json:"hpa,omitempty"`
}

type AppSpecContainers struct {
	Args            []*string                        `json:"args,omitempty"`
	Env             []*AppSpecContainersEnv          `json:"env,omitempty"`
	EnvFrom         []*AppSpecContainersEnvFrom      `json:"envFrom,omitempty"`
	Name            string                           `json:"name"`
	ReadinessProbe  *AppSpecContainersReadinessProbe `json:"readinessProbe,omitempty"`
	ResourceMemory  *AppSpecContainersResourceMemory `json:"resourceMemory,omitempty"`
	Volumes         []*AppSpecContainersVolumes      `json:"volumes,omitempty"`
	Command         []*string                        `json:"command,omitempty"`
	Image           string                           `json:"image"`
	ImagePullPolicy *string                          `json:"imagePullPolicy,omitempty"`
	LivenessProbe   *AppSpecContainersLivenessProbe  `json:"livenessProbe,omitempty"`
	ResourceCPU     *AppSpecContainersResourceCPU    `json:"resourceCpu,omitempty"`
}

type AppSpecContainersEnv struct {
	Key      string  `json:"key"`
	Optional *bool   `json:"optional,omitempty"`
	RefKey   *string `json:"refKey,omitempty"`
	RefName  *string `json:"refName,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type AppSpecContainersEnvFrom struct {
	RefName string `json:"refName"`
	Type    string `json:"type"`
}

type AppSpecContainersEnvFromIn struct {
	RefName string `json:"refName"`
	Type    string `json:"type"`
}

type AppSpecContainersEnvIn struct {
	Key      string  `json:"key"`
	Optional *bool   `json:"optional,omitempty"`
	RefKey   *string `json:"refKey,omitempty"`
	RefName  *string `json:"refName,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type AppSpecContainersIn struct {
	Args            []*string                          `json:"args,omitempty"`
	Env             []*AppSpecContainersEnvIn          `json:"env,omitempty"`
	EnvFrom         []*AppSpecContainersEnvFromIn      `json:"envFrom,omitempty"`
	Name            string                             `json:"name"`
	ReadinessProbe  *AppSpecContainersReadinessProbeIn `json:"readinessProbe,omitempty"`
	ResourceMemory  *AppSpecContainersResourceMemoryIn `json:"resourceMemory,omitempty"`
	Volumes         []*AppSpecContainersVolumesIn      `json:"volumes,omitempty"`
	Command         []*string                          `json:"command,omitempty"`
	Image           string                             `json:"image"`
	ImagePullPolicy *string                            `json:"imagePullPolicy,omitempty"`
	LivenessProbe   *AppSpecContainersLivenessProbeIn  `json:"livenessProbe,omitempty"`
	ResourceCPU     *AppSpecContainersResourceCPUIn    `json:"resourceCpu,omitempty"`
}

type AppSpecContainersLivenessProbe struct {
	TCP              *AppSpecContainersLivenessProbeTCP     `json:"tcp,omitempty"`
	Type             string                                 `json:"type"`
	FailureThreshold *int                                   `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGet `json:"httpGet,omitempty"`
	InitialDelay     *int                                   `json:"initialDelay,omitempty"`
	Interval         *int                                   `json:"interval,omitempty"`
	Shell            *AppSpecContainersLivenessProbeShell   `json:"shell,omitempty"`
}

type AppSpecContainersLivenessProbeHTTPGet struct {
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
}

type AppSpecContainersLivenessProbeHTTPGetIn struct {
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
}

type AppSpecContainersLivenessProbeIn struct {
	TCP              *AppSpecContainersLivenessProbeTCPIn     `json:"tcp,omitempty"`
	Type             string                                   `json:"type"`
	FailureThreshold *int                                     `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGetIn `json:"httpGet,omitempty"`
	InitialDelay     *int                                     `json:"initialDelay,omitempty"`
	Interval         *int                                     `json:"interval,omitempty"`
	Shell            *AppSpecContainersLivenessProbeShellIn   `json:"shell,omitempty"`
}

type AppSpecContainersLivenessProbeShell struct {
	Command []*string `json:"command,omitempty"`
}

type AppSpecContainersLivenessProbeShellIn struct {
	Command []*string `json:"command,omitempty"`
}

type AppSpecContainersLivenessProbeTCP struct {
	Port int `json:"port"`
}

type AppSpecContainersLivenessProbeTCPIn struct {
	Port int `json:"port"`
}

type AppSpecContainersReadinessProbe struct {
	FailureThreshold *int                                    `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGet `json:"httpGet,omitempty"`
	InitialDelay     *int                                    `json:"initialDelay,omitempty"`
	Interval         *int                                    `json:"interval,omitempty"`
	Shell            *AppSpecContainersReadinessProbeShell   `json:"shell,omitempty"`
	TCP              *AppSpecContainersReadinessProbeTCP     `json:"tcp,omitempty"`
	Type             string                                  `json:"type"`
}

type AppSpecContainersReadinessProbeHTTPGet struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersReadinessProbeHTTPGetIn struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersReadinessProbeIn struct {
	FailureThreshold *int                                      `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGetIn `json:"httpGet,omitempty"`
	InitialDelay     *int                                      `json:"initialDelay,omitempty"`
	Interval         *int                                      `json:"interval,omitempty"`
	Shell            *AppSpecContainersReadinessProbeShellIn   `json:"shell,omitempty"`
	TCP              *AppSpecContainersReadinessProbeTCPIn     `json:"tcp,omitempty"`
	Type             string                                    `json:"type"`
}

type AppSpecContainersReadinessProbeShell struct {
	Command []*string `json:"command,omitempty"`
}

type AppSpecContainersReadinessProbeShellIn struct {
	Command []*string `json:"command,omitempty"`
}

type AppSpecContainersReadinessProbeTCP struct {
	Port int `json:"port"`
}

type AppSpecContainersReadinessProbeTCPIn struct {
	Port int `json:"port"`
}

type AppSpecContainersResourceCPU struct {
	Max *string `json:"max,omitempty"`
	Min *string `json:"min,omitempty"`
}

type AppSpecContainersResourceCPUIn struct {
	Max *string `json:"max,omitempty"`
	Min *string `json:"min,omitempty"`
}

type AppSpecContainersResourceMemory struct {
	Max *string `json:"max,omitempty"`
	Min *string `json:"min,omitempty"`
}

type AppSpecContainersResourceMemoryIn struct {
	Max *string `json:"max,omitempty"`
	Min *string `json:"min,omitempty"`
}

type AppSpecContainersVolumes struct {
	Items     []*AppSpecContainersVolumesItems `json:"items,omitempty"`
	MountPath string                           `json:"mountPath"`
	RefName   string                           `json:"refName"`
	Type      string                           `json:"type"`
}

type AppSpecContainersVolumesIn struct {
	Items     []*AppSpecContainersVolumesItemsIn `json:"items,omitempty"`
	MountPath string                             `json:"mountPath"`
	RefName   string                             `json:"refName"`
	Type      string                             `json:"type"`
}

type AppSpecContainersVolumesItems struct {
	FileName *string `json:"fileName,omitempty"`
	Key      string  `json:"key"`
}

type AppSpecContainersVolumesItemsIn struct {
	FileName *string `json:"fileName,omitempty"`
	Key      string  `json:"key"`
}

type AppSpecHpa struct {
	MaxReplicas     *int  `json:"maxReplicas,omitempty"`
	MinReplicas     *int  `json:"minReplicas,omitempty"`
	ThresholdCPU    *int  `json:"thresholdCpu,omitempty"`
	ThresholdMemory *int  `json:"thresholdMemory,omitempty"`
	Enabled         *bool `json:"enabled,omitempty"`
}

type AppSpecHpaIn struct {
	MaxReplicas     *int  `json:"maxReplicas,omitempty"`
	MinReplicas     *int  `json:"minReplicas,omitempty"`
	ThresholdCPU    *int  `json:"thresholdCpu,omitempty"`
	ThresholdMemory *int  `json:"thresholdMemory,omitempty"`
	Enabled         *bool `json:"enabled,omitempty"`
}

type AppSpecIn struct {
	Services       []*AppSpecServicesIn    `json:"services,omitempty"`
	DisplayName    *string                 `json:"displayName,omitempty"`
	Freeze         *bool                   `json:"freeze,omitempty"`
	Intercept      *AppSpecInterceptIn     `json:"intercept,omitempty"`
	NodeSelector   map[string]interface{}  `json:"nodeSelector,omitempty"`
	Region         *string                 `json:"region,omitempty"`
	Replicas       *int                    `json:"replicas,omitempty"`
	ServiceAccount *string                 `json:"serviceAccount,omitempty"`
	Tolerations    []*AppSpecTolerationsIn `json:"tolerations,omitempty"`
	Containers     []*AppSpecContainersIn  `json:"containers"`
	Hpa            *AppSpecHpaIn           `json:"hpa,omitempty"`
}

type AppSpecIntercept struct {
	Enabled  bool   `json:"enabled"`
	ToDevice string `json:"toDevice"`
}

type AppSpecInterceptIn struct {
	Enabled  bool   `json:"enabled"`
	ToDevice string `json:"toDevice"`
}

type AppSpecServices struct {
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort,omitempty"`
	Type       *string `json:"type,omitempty"`
	Name       *string `json:"name,omitempty"`
}

type AppSpecServicesIn struct {
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort,omitempty"`
	Type       *string `json:"type,omitempty"`
	Name       *string `json:"name,omitempty"`
}

type AppSpecTolerations struct {
	Key               *string `json:"key,omitempty"`
	Operator          *string `json:"operator,omitempty"`
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
}

type AppSpecTolerationsIn struct {
	Key               *string `json:"key,omitempty"`
	Operator          *string `json:"operator,omitempty"`
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
}

type EnvironmentSpec struct {
	ProjectName     string  `json:"projectName"`
	TargetNamespace *string `json:"targetNamespace,omitempty"`
}

type EnvironmentSpecIn struct {
	ProjectName     string  `json:"projectName"`
	TargetNamespace *string `json:"targetNamespace,omitempty"`
}

type ManagedResourceSpec struct {
	Inputs   map[string]interface{}       `json:"inputs,omitempty"`
	MresKind *ManagedResourceSpecMresKind `json:"mresKind"`
	MsvcRef  *ManagedResourceSpecMsvcRef  `json:"msvcRef"`
}

type ManagedResourceSpecIn struct {
	Inputs   map[string]interface{}         `json:"inputs,omitempty"`
	MresKind *ManagedResourceSpecMresKindIn `json:"mresKind"`
	MsvcRef  *ManagedResourceSpecMsvcRefIn  `json:"msvcRef"`
}

type ManagedResourceSpecMresKind struct {
	Kind string `json:"kind"`
}

type ManagedResourceSpecMresKindIn struct {
	Kind string `json:"kind"`
}

type ManagedResourceSpecMsvcRef struct {
	Name       string  `json:"name"`
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind,omitempty"`
}

type ManagedResourceSpecMsvcRefIn struct {
	Name       string  `json:"name"`
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind,omitempty"`
}

type ManagedServiceSpec struct {
	Inputs       map[string]interface{}           `json:"inputs,omitempty"`
	MsvcKind     *ManagedServiceSpecMsvcKind      `json:"msvcKind"`
	NodeSelector map[string]interface{}           `json:"nodeSelector,omitempty"`
	Region       *string                          `json:"region,omitempty"`
	Tolerations  []*ManagedServiceSpecTolerations `json:"tolerations,omitempty"`
}

type ManagedServiceSpecIn struct {
	Inputs       map[string]interface{}             `json:"inputs,omitempty"`
	MsvcKind     *ManagedServiceSpecMsvcKindIn      `json:"msvcKind"`
	NodeSelector map[string]interface{}             `json:"nodeSelector,omitempty"`
	Region       *string                            `json:"region,omitempty"`
	Tolerations  []*ManagedServiceSpecTolerationsIn `json:"tolerations,omitempty"`
}

type ManagedServiceSpecMsvcKind struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind,omitempty"`
}

type ManagedServiceSpecMsvcKindIn struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind,omitempty"`
}

type ManagedServiceSpecTolerations struct {
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
	Key               *string `json:"key,omitempty"`
	Operator          *string `json:"operator,omitempty"`
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`
}

type ManagedServiceSpecTolerationsIn struct {
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
	Key               *string `json:"key,omitempty"`
	Operator          *string `json:"operator,omitempty"`
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`
}

type ProjectSpec struct {
	DisplayName     *string `json:"displayName,omitempty"`
	Logo            *string `json:"logo,omitempty"`
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	AccountName     string  `json:"accountName"`
	ClusterName     string  `json:"clusterName"`
}

type ProjectSpecIn struct {
	DisplayName     *string `json:"displayName,omitempty"`
	Logo            *string `json:"logo,omitempty"`
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	AccountName     string  `json:"accountName"`
	ClusterName     string  `json:"clusterName"`
}

type RouterSpec struct {
	Cors            *RouterSpecCors      `json:"cors,omitempty"`
	Region          *string              `json:"region,omitempty"`
	Routes          []*RouterSpecRoutes  `json:"routes,omitempty"`
	HTTPS           *RouterSpecHTTPS     `json:"https,omitempty"`
	IngressClass    *string              `json:"ingressClass,omitempty"`
	MaxBodySizeInMb *int                 `json:"maxBodySizeInMB,omitempty"`
	RateLimit       *RouterSpecRateLimit `json:"rateLimit,omitempty"`
	BackendProtocol *string              `json:"backendProtocol,omitempty"`
	BasicAuth       *RouterSpecBasicAuth `json:"basicAuth,omitempty"`
	Domains         []*string            `json:"domains"`
}

type RouterSpecBasicAuth struct {
	Enabled    bool    `json:"enabled"`
	SecretName *string `json:"secretName,omitempty"`
	Username   *string `json:"username,omitempty"`
}

type RouterSpecBasicAuthIn struct {
	Enabled    bool    `json:"enabled"`
	SecretName *string `json:"secretName,omitempty"`
	Username   *string `json:"username,omitempty"`
}

type RouterSpecCors struct {
	Enabled          *bool     `json:"enabled,omitempty"`
	Origins          []*string `json:"origins,omitempty"`
	AllowCredentials *bool     `json:"allowCredentials,omitempty"`
}

type RouterSpecCorsIn struct {
	Enabled          *bool     `json:"enabled,omitempty"`
	Origins          []*string `json:"origins,omitempty"`
	AllowCredentials *bool     `json:"allowCredentials,omitempty"`
}

type RouterSpecHTTPS struct {
	ForceRedirect *bool   `json:"forceRedirect,omitempty"`
	ClusterIssuer *string `json:"clusterIssuer,omitempty"`
	Enabled       bool    `json:"enabled"`
}

type RouterSpecHTTPSIn struct {
	ForceRedirect *bool   `json:"forceRedirect,omitempty"`
	ClusterIssuer *string `json:"clusterIssuer,omitempty"`
	Enabled       bool    `json:"enabled"`
}

type RouterSpecIn struct {
	Cors            *RouterSpecCorsIn      `json:"cors,omitempty"`
	Region          *string                `json:"region,omitempty"`
	Routes          []*RouterSpecRoutesIn  `json:"routes,omitempty"`
	HTTPS           *RouterSpecHTTPSIn     `json:"https,omitempty"`
	IngressClass    *string                `json:"ingressClass,omitempty"`
	MaxBodySizeInMb *int                   `json:"maxBodySizeInMB,omitempty"`
	RateLimit       *RouterSpecRateLimitIn `json:"rateLimit,omitempty"`
	BackendProtocol *string                `json:"backendProtocol,omitempty"`
	BasicAuth       *RouterSpecBasicAuthIn `json:"basicAuth,omitempty"`
	Domains         []*string              `json:"domains"`
}

type RouterSpecRateLimit struct {
	Connections *int  `json:"connections,omitempty"`
	Enabled     *bool `json:"enabled,omitempty"`
	Rpm         *int  `json:"rpm,omitempty"`
	Rps         *int  `json:"rps,omitempty"`
}

type RouterSpecRateLimitIn struct {
	Connections *int  `json:"connections,omitempty"`
	Enabled     *bool `json:"enabled,omitempty"`
	Rpm         *int  `json:"rpm,omitempty"`
	Rps         *int  `json:"rps,omitempty"`
}

type RouterSpecRoutes struct {
	App     *string `json:"app,omitempty"`
	Lambda  *string `json:"lambda,omitempty"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite,omitempty"`
}

type RouterSpecRoutesIn struct {
	App     *string `json:"app,omitempty"`
	Lambda  *string `json:"lambda,omitempty"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite,omitempty"`
}
