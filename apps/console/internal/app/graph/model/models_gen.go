// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AppSpec struct {
	Containers     []*AppSpecContainers   `json:"containers"`
	Hpa            *AppSpecHpa            `json:"hpa,omitempty"`
	Services       []*AppSpecServices     `json:"services,omitempty"`
	ServiceAccount *string                `json:"serviceAccount,omitempty"`
	Tolerations    []*AppSpecTolerations  `json:"tolerations,omitempty"`
	DisplayName    *string                `json:"displayName,omitempty"`
	Freeze         *bool                  `json:"freeze,omitempty"`
	Intercept      *AppSpecIntercept      `json:"intercept,omitempty"`
	NodeSelector   map[string]interface{} `json:"nodeSelector,omitempty"`
	Region         *string                `json:"region,omitempty"`
	Replicas       *int                   `json:"replicas,omitempty"`
}

type AppSpecContainers struct {
	Command         []*string                        `json:"command,omitempty"`
	ReadinessProbe  *AppSpecContainersReadinessProbe `json:"readinessProbe,omitempty"`
	Volumes         []*AppSpecContainersVolumes      `json:"volumes,omitempty"`
	Args            []*string                        `json:"args,omitempty"`
	Env             []*AppSpecContainersEnv          `json:"env,omitempty"`
	EnvFrom         []*AppSpecContainersEnvFrom      `json:"envFrom,omitempty"`
	Image           string                           `json:"image"`
	ImagePullPolicy *string                          `json:"imagePullPolicy,omitempty"`
	LivenessProbe   *AppSpecContainersLivenessProbe  `json:"livenessProbe,omitempty"`
	Name            string                           `json:"name"`
	ResourceCPU     *AppSpecContainersResourceCPU    `json:"resourceCpu,omitempty"`
	ResourceMemory  *AppSpecContainersResourceMemory `json:"resourceMemory,omitempty"`
}

type AppSpecContainersEnv struct {
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
	Key      string  `json:"key"`
	Optional *bool   `json:"optional,omitempty"`
	RefKey   *string `json:"refKey,omitempty"`
	RefName  *string `json:"refName,omitempty"`
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
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
	Key      string  `json:"key"`
	Optional *bool   `json:"optional,omitempty"`
	RefKey   *string `json:"refKey,omitempty"`
	RefName  *string `json:"refName,omitempty"`
}

type AppSpecContainersIn struct {
	Command         []*string                          `json:"command,omitempty"`
	ReadinessProbe  *AppSpecContainersReadinessProbeIn `json:"readinessProbe,omitempty"`
	Volumes         []*AppSpecContainersVolumesIn      `json:"volumes,omitempty"`
	Args            []*string                          `json:"args,omitempty"`
	Env             []*AppSpecContainersEnvIn          `json:"env,omitempty"`
	EnvFrom         []*AppSpecContainersEnvFromIn      `json:"envFrom,omitempty"`
	Image           string                             `json:"image"`
	ImagePullPolicy *string                            `json:"imagePullPolicy,omitempty"`
	LivenessProbe   *AppSpecContainersLivenessProbeIn  `json:"livenessProbe,omitempty"`
	Name            string                             `json:"name"`
	ResourceCPU     *AppSpecContainersResourceCPUIn    `json:"resourceCpu,omitempty"`
	ResourceMemory  *AppSpecContainersResourceMemoryIn `json:"resourceMemory,omitempty"`
}

type AppSpecContainersLivenessProbe struct {
	FailureThreshold *int                                   `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGet `json:"httpGet,omitempty"`
	InitialDelay     *int                                   `json:"initialDelay,omitempty"`
	Interval         *int                                   `json:"interval,omitempty"`
	Shell            *AppSpecContainersLivenessProbeShell   `json:"shell,omitempty"`
	TCP              *AppSpecContainersLivenessProbeTCP     `json:"tcp,omitempty"`
	Type             string                                 `json:"type"`
}

type AppSpecContainersLivenessProbeHTTPGet struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersLivenessProbeHTTPGetIn struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersLivenessProbeIn struct {
	FailureThreshold *int                                     `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGetIn `json:"httpGet,omitempty"`
	InitialDelay     *int                                     `json:"initialDelay,omitempty"`
	Interval         *int                                     `json:"interval,omitempty"`
	Shell            *AppSpecContainersLivenessProbeShellIn   `json:"shell,omitempty"`
	TCP              *AppSpecContainersLivenessProbeTCPIn     `json:"tcp,omitempty"`
	Type             string                                   `json:"type"`
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
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGet `json:"httpGet,omitempty"`
	InitialDelay     *int                                    `json:"initialDelay,omitempty"`
	Interval         *int                                    `json:"interval,omitempty"`
	Shell            *AppSpecContainersReadinessProbeShell   `json:"shell,omitempty"`
	TCP              *AppSpecContainersReadinessProbeTCP     `json:"tcp,omitempty"`
	Type             string                                  `json:"type"`
	FailureThreshold *int                                    `json:"failureThreshold,omitempty"`
}

type AppSpecContainersReadinessProbeHTTPGet struct {
	Port        int                    `json:"port"`
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
	Path        string                 `json:"path"`
}

type AppSpecContainersReadinessProbeHTTPGetIn struct {
	Port        int                    `json:"port"`
	HTTPHeaders map[string]interface{} `json:"httpHeaders,omitempty"`
	Path        string                 `json:"path"`
}

type AppSpecContainersReadinessProbeIn struct {
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGetIn `json:"httpGet,omitempty"`
	InitialDelay     *int                                      `json:"initialDelay,omitempty"`
	Interval         *int                                      `json:"interval,omitempty"`
	Shell            *AppSpecContainersReadinessProbeShellIn   `json:"shell,omitempty"`
	TCP              *AppSpecContainersReadinessProbeTCPIn     `json:"tcp,omitempty"`
	Type             string                                    `json:"type"`
	FailureThreshold *int                                      `json:"failureThreshold,omitempty"`
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
	Enabled         *bool `json:"enabled,omitempty"`
	MaxReplicas     *int  `json:"maxReplicas,omitempty"`
	MinReplicas     *int  `json:"minReplicas,omitempty"`
	ThresholdCPU    *int  `json:"thresholdCpu,omitempty"`
	ThresholdMemory *int  `json:"thresholdMemory,omitempty"`
}

type AppSpecHpaIn struct {
	Enabled         *bool `json:"enabled,omitempty"`
	MaxReplicas     *int  `json:"maxReplicas,omitempty"`
	MinReplicas     *int  `json:"minReplicas,omitempty"`
	ThresholdCPU    *int  `json:"thresholdCpu,omitempty"`
	ThresholdMemory *int  `json:"thresholdMemory,omitempty"`
}

type AppSpecIn struct {
	Containers     []*AppSpecContainersIn  `json:"containers"`
	Hpa            *AppSpecHpaIn           `json:"hpa,omitempty"`
	Services       []*AppSpecServicesIn    `json:"services,omitempty"`
	ServiceAccount *string                 `json:"serviceAccount,omitempty"`
	Tolerations    []*AppSpecTolerationsIn `json:"tolerations,omitempty"`
	DisplayName    *string                 `json:"displayName,omitempty"`
	Freeze         *bool                   `json:"freeze,omitempty"`
	Intercept      *AppSpecInterceptIn     `json:"intercept,omitempty"`
	NodeSelector   map[string]interface{}  `json:"nodeSelector,omitempty"`
	Region         *string                 `json:"region,omitempty"`
	Replicas       *int                    `json:"replicas,omitempty"`
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
	Name       *string `json:"name,omitempty"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type AppSpecServicesIn struct {
	Name       *string `json:"name,omitempty"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type AppSpecTolerations struct {
	Operator          *string `json:"operator,omitempty"`
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
	Key               *string `json:"key,omitempty"`
}

type AppSpecTolerationsIn struct {
	Operator          *string `json:"operator,omitempty"`
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
	Key               *string `json:"key,omitempty"`
}

type EnvironmentSpec struct {
	ProjectName     string `json:"projectName"`
	TargetNamespace string `json:"targetNamespace"`
}

type EnvironmentSpecIn struct {
	ProjectName     string `json:"projectName"`
	TargetNamespace string `json:"targetNamespace"`
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
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind,omitempty"`
	Name       string  `json:"name"`
}

type ManagedResourceSpecMsvcRefIn struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind,omitempty"`
	Name       string  `json:"name"`
}

type ManagedServiceSpec struct {
	NodeSelector map[string]interface{}           `json:"nodeSelector,omitempty"`
	Region       *string                          `json:"region,omitempty"`
	Tolerations  []*ManagedServiceSpecTolerations `json:"tolerations,omitempty"`
	Inputs       map[string]interface{}           `json:"inputs,omitempty"`
	MsvcKind     *ManagedServiceSpecMsvcKind      `json:"msvcKind"`
}

type ManagedServiceSpecIn struct {
	NodeSelector map[string]interface{}             `json:"nodeSelector,omitempty"`
	Region       *string                            `json:"region,omitempty"`
	Tolerations  []*ManagedServiceSpecTolerationsIn `json:"tolerations,omitempty"`
	Inputs       map[string]interface{}             `json:"inputs,omitempty"`
	MsvcKind     *ManagedServiceSpecMsvcKindIn      `json:"msvcKind"`
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
	TargetNamespace string  `json:"targetNamespace"`
	AccountName     string  `json:"accountName"`
	ClusterName     string  `json:"clusterName"`
}

type ProjectSpecIn struct {
	DisplayName     *string `json:"displayName,omitempty"`
	Logo            *string `json:"logo,omitempty"`
	TargetNamespace string  `json:"targetNamespace"`
	AccountName     string  `json:"accountName"`
	ClusterName     string  `json:"clusterName"`
}

type RouterSpec struct {
	Region          *string              `json:"region,omitempty"`
	Routes          []*RouterSpecRoutes  `json:"routes,omitempty"`
	BackendProtocol *string              `json:"backendProtocol,omitempty"`
	BasicAuth       *RouterSpecBasicAuth `json:"basicAuth,omitempty"`
	Domains         []*string            `json:"domains"`
	IngressClass    *string              `json:"ingressClass,omitempty"`
	RateLimit       *RouterSpecRateLimit `json:"rateLimit,omitempty"`
	Cors            *RouterSpecCors      `json:"cors,omitempty"`
	HTTPS           *RouterSpecHTTPS     `json:"https,omitempty"`
	MaxBodySizeInMb *int                 `json:"maxBodySizeInMB,omitempty"`
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
	AllowCredentials *bool     `json:"allowCredentials,omitempty"`
	Enabled          *bool     `json:"enabled,omitempty"`
	Origins          []*string `json:"origins,omitempty"`
}

type RouterSpecCorsIn struct {
	AllowCredentials *bool     `json:"allowCredentials,omitempty"`
	Enabled          *bool     `json:"enabled,omitempty"`
	Origins          []*string `json:"origins,omitempty"`
}

type RouterSpecHTTPS struct {
	ClusterIssuer *string `json:"clusterIssuer,omitempty"`
	Enabled       bool    `json:"enabled"`
	ForceRedirect *bool   `json:"forceRedirect,omitempty"`
}

type RouterSpecHTTPSIn struct {
	ClusterIssuer *string `json:"clusterIssuer,omitempty"`
	Enabled       bool    `json:"enabled"`
	ForceRedirect *bool   `json:"forceRedirect,omitempty"`
}

type RouterSpecIn struct {
	Region          *string                `json:"region,omitempty"`
	Routes          []*RouterSpecRoutesIn  `json:"routes,omitempty"`
	BackendProtocol *string                `json:"backendProtocol,omitempty"`
	BasicAuth       *RouterSpecBasicAuthIn `json:"basicAuth,omitempty"`
	Domains         []*string              `json:"domains"`
	IngressClass    *string                `json:"ingressClass,omitempty"`
	RateLimit       *RouterSpecRateLimitIn `json:"rateLimit,omitempty"`
	Cors            *RouterSpecCorsIn      `json:"cors,omitempty"`
	HTTPS           *RouterSpecHTTPSIn     `json:"https,omitempty"`
	MaxBodySizeInMb *int                   `json:"maxBodySizeInMB,omitempty"`
}

type RouterSpecRateLimit struct {
	Rps         *int  `json:"rps,omitempty"`
	Connections *int  `json:"connections,omitempty"`
	Enabled     *bool `json:"enabled,omitempty"`
	Rpm         *int  `json:"rpm,omitempty"`
}

type RouterSpecRateLimitIn struct {
	Rps         *int  `json:"rps,omitempty"`
	Connections *int  `json:"connections,omitempty"`
	Enabled     *bool `json:"enabled,omitempty"`
	Rpm         *int  `json:"rpm,omitempty"`
}

type RouterSpecRoutes struct {
	Rewrite *bool   `json:"rewrite,omitempty"`
	App     *string `json:"app,omitempty"`
	Lambda  *string `json:"lambda,omitempty"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
}

type RouterSpecRoutesIn struct {
	Rewrite *bool   `json:"rewrite,omitempty"`
	App     *string `json:"app,omitempty"`
	Lambda  *string `json:"lambda,omitempty"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
}
