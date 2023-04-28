// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AppSpec struct {
	Tolerations    []*AppSpecTolerations  `json:"tolerations,omitempty"`
	AccountName    string                 `json:"accountName"`
	Containers     []*AppSpecContainers   `json:"containers"`
	Region         string                 `json:"region"`
	NodeSelector   map[string]interface{} `json:"nodeSelector,omitempty"`
	Replicas       *int                   `json:"replicas,omitempty"`
	ServiceAccount *string                `json:"serviceAccount,omitempty"`
	Services       []*AppSpecServices     `json:"services,omitempty"`
	Freeze         *bool                  `json:"freeze,omitempty"`
	Hpa            *AppSpecHpa            `json:"hpa,omitempty"`
	Intercept      *AppSpecIntercept      `json:"intercept,omitempty"`
}

type AppSpecContainers struct {
	Args            []*string                        `json:"args,omitempty"`
	Command         []*string                        `json:"command,omitempty"`
	Env             []*AppSpecContainersEnv          `json:"env,omitempty"`
	ImagePullPolicy *string                          `json:"imagePullPolicy,omitempty"`
	LivenessProbe   *AppSpecContainersLivenessProbe  `json:"livenessProbe,omitempty"`
	ReadinessProbe  *AppSpecContainersReadinessProbe `json:"readinessProbe,omitempty"`
	ResourceCPU     *AppSpecContainersResourceCPU    `json:"resourceCpu,omitempty"`
	EnvFrom         []*AppSpecContainersEnvFrom      `json:"envFrom,omitempty"`
	Image           string                           `json:"image"`
	Name            string                           `json:"name"`
	ResourceMemory  *AppSpecContainersResourceMemory `json:"resourceMemory,omitempty"`
	Volumes         []*AppSpecContainersVolumes      `json:"volumes,omitempty"`
}

type AppSpecContainersEnv struct {
	RefName *string `json:"refName,omitempty"`
	Type    *string `json:"type,omitempty"`
	Value   *string `json:"value,omitempty"`
	Key     string  `json:"key"`
	RefKey  *string `json:"refKey,omitempty"`
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
	RefName *string `json:"refName,omitempty"`
	Type    *string `json:"type,omitempty"`
	Value   *string `json:"value,omitempty"`
	Key     string  `json:"key"`
	RefKey  *string `json:"refKey,omitempty"`
}

type AppSpecContainersIn struct {
	Args            []*string                          `json:"args,omitempty"`
	Command         []*string                          `json:"command,omitempty"`
	Env             []*AppSpecContainersEnvIn          `json:"env,omitempty"`
	ImagePullPolicy *string                            `json:"imagePullPolicy,omitempty"`
	LivenessProbe   *AppSpecContainersLivenessProbeIn  `json:"livenessProbe,omitempty"`
	ReadinessProbe  *AppSpecContainersReadinessProbeIn `json:"readinessProbe,omitempty"`
	ResourceCPU     *AppSpecContainersResourceCPUIn    `json:"resourceCpu,omitempty"`
	EnvFrom         []*AppSpecContainersEnvFromIn      `json:"envFrom,omitempty"`
	Image           string                             `json:"image"`
	Name            string                             `json:"name"`
	ResourceMemory  *AppSpecContainersResourceMemoryIn `json:"resourceMemory,omitempty"`
	Volumes         []*AppSpecContainersVolumesIn      `json:"volumes,omitempty"`
}

type AppSpecContainersLivenessProbe struct {
	Type             string                                 `json:"type"`
	FailureThreshold *int                                   `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGet `json:"httpGet,omitempty"`
	InitialDelay     *int                                   `json:"initialDelay,omitempty"`
	Interval         *int                                   `json:"interval,omitempty"`
	Shell            *AppSpecContainersLivenessProbeShell   `json:"shell,omitempty"`
	TCP              *AppSpecContainersLivenessProbeTCP     `json:"tcp,omitempty"`
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
	Type             string                                   `json:"type"`
	FailureThreshold *int                                     `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGetIn `json:"httpGet,omitempty"`
	InitialDelay     *int                                     `json:"initialDelay,omitempty"`
	Interval         *int                                     `json:"interval,omitempty"`
	Shell            *AppSpecContainersLivenessProbeShellIn   `json:"shell,omitempty"`
	TCP              *AppSpecContainersLivenessProbeTCPIn     `json:"tcp,omitempty"`
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
	Type             string                                  `json:"type"`
	FailureThreshold *int                                    `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGet `json:"httpGet,omitempty"`
	InitialDelay     *int                                    `json:"initialDelay,omitempty"`
	Interval         *int                                    `json:"interval,omitempty"`
	Shell            *AppSpecContainersReadinessProbeShell   `json:"shell,omitempty"`
	TCP              *AppSpecContainersReadinessProbeTCP     `json:"tcp,omitempty"`
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
	Type             string                                    `json:"type"`
	FailureThreshold *int                                      `json:"failureThreshold,omitempty"`
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGetIn `json:"httpGet,omitempty"`
	InitialDelay     *int                                      `json:"initialDelay,omitempty"`
	Interval         *int                                      `json:"interval,omitempty"`
	Shell            *AppSpecContainersReadinessProbeShellIn   `json:"shell,omitempty"`
	TCP              *AppSpecContainersReadinessProbeTCPIn     `json:"tcp,omitempty"`
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
	Tolerations    []*AppSpecTolerationsIn `json:"tolerations,omitempty"`
	AccountName    string                  `json:"accountName"`
	Containers     []*AppSpecContainersIn  `json:"containers"`
	Region         string                  `json:"region"`
	NodeSelector   map[string]interface{}  `json:"nodeSelector,omitempty"`
	Replicas       *int                    `json:"replicas,omitempty"`
	ServiceAccount *string                 `json:"serviceAccount,omitempty"`
	Services       []*AppSpecServicesIn    `json:"services,omitempty"`
	Freeze         *bool                   `json:"freeze,omitempty"`
	Hpa            *AppSpecHpaIn           `json:"hpa,omitempty"`
	Intercept      *AppSpecInterceptIn     `json:"intercept,omitempty"`
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
	Type       *string `json:"type,omitempty"`
	Name       *string `json:"name,omitempty"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort,omitempty"`
}

type AppSpecServicesIn struct {
	Type       *string `json:"type,omitempty"`
	Name       *string `json:"name,omitempty"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort,omitempty"`
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
	Kind       *string `json:"kind,omitempty"`
	Name       string  `json:"name"`
	APIVersion string  `json:"apiVersion"`
}

type ManagedResourceSpecMsvcRefIn struct {
	Kind       *string `json:"kind,omitempty"`
	Name       string  `json:"name"`
	APIVersion string  `json:"apiVersion"`
}

type ManagedServiceSpec struct {
	Region       string                           `json:"region"`
	Tolerations  []*ManagedServiceSpecTolerations `json:"tolerations,omitempty"`
	Inputs       map[string]interface{}           `json:"inputs,omitempty"`
	MsvcKind     *ManagedServiceSpecMsvcKind      `json:"msvcKind"`
	NodeSelector map[string]interface{}           `json:"nodeSelector,omitempty"`
}

type ManagedServiceSpecIn struct {
	Region       string                             `json:"region"`
	Tolerations  []*ManagedServiceSpecTolerationsIn `json:"tolerations,omitempty"`
	Inputs       map[string]interface{}             `json:"inputs,omitempty"`
	MsvcKind     *ManagedServiceSpecMsvcKindIn      `json:"msvcKind"`
	NodeSelector map[string]interface{}             `json:"nodeSelector,omitempty"`
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
	Key               *string `json:"key,omitempty"`
	Operator          *string `json:"operator,omitempty"`
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
}

type ManagedServiceSpecTolerationsIn struct {
	Key               *string `json:"key,omitempty"`
	Operator          *string `json:"operator,omitempty"`
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`
	Value             *string `json:"value,omitempty"`
	Effect            *string `json:"effect,omitempty"`
}

type ProjectSpec struct {
	Logo            *string `json:"logo,omitempty"`
	TargetNamespace string  `json:"targetNamespace"`
	AccountName     string  `json:"accountName"`
	ClusterName     string  `json:"clusterName"`
	DisplayName     *string `json:"displayName,omitempty"`
}

type ProjectSpecIn struct {
	Logo            *string `json:"logo,omitempty"`
	TargetNamespace string  `json:"targetNamespace"`
	AccountName     string  `json:"accountName"`
	ClusterName     string  `json:"clusterName"`
	DisplayName     *string `json:"displayName,omitempty"`
}

type RouterSpec struct {
	HTTPS           *RouterSpecHTTPS     `json:"https,omitempty"`
	MaxBodySizeInMb *int                 `json:"maxBodySizeInMB,omitempty"`
	RateLimit       *RouterSpecRateLimit `json:"rateLimit,omitempty"`
	Region          *string              `json:"region,omitempty"`
	Routes          []*RouterSpecRoutes  `json:"routes,omitempty"`
	BasicAuth       *RouterSpecBasicAuth `json:"basicAuth,omitempty"`
	Cors            *RouterSpecCors      `json:"cors,omitempty"`
	Domains         []*string            `json:"domains"`
}

type RouterSpecBasicAuth struct {
	SecretName *string `json:"secretName,omitempty"`
	Username   *string `json:"username,omitempty"`
	Enabled    bool    `json:"enabled"`
}

type RouterSpecBasicAuthIn struct {
	SecretName *string `json:"secretName,omitempty"`
	Username   *string `json:"username,omitempty"`
	Enabled    bool    `json:"enabled"`
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
	Enabled       bool  `json:"enabled"`
	ForceRedirect *bool `json:"forceRedirect,omitempty"`
}

type RouterSpecHTTPSIn struct {
	Enabled       bool  `json:"enabled"`
	ForceRedirect *bool `json:"forceRedirect,omitempty"`
}

type RouterSpecIn struct {
	HTTPS           *RouterSpecHTTPSIn     `json:"https,omitempty"`
	MaxBodySizeInMb *int                   `json:"maxBodySizeInMB,omitempty"`
	RateLimit       *RouterSpecRateLimitIn `json:"rateLimit,omitempty"`
	Region          *string                `json:"region,omitempty"`
	Routes          []*RouterSpecRoutesIn  `json:"routes,omitempty"`
	BasicAuth       *RouterSpecBasicAuthIn `json:"basicAuth,omitempty"`
	Cors            *RouterSpecCorsIn      `json:"cors,omitempty"`
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
	Lambda  *string `json:"lambda,omitempty"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite,omitempty"`
	App     *string `json:"app,omitempty"`
}

type RouterSpecRoutesIn struct {
	Lambda  *string `json:"lambda,omitempty"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite,omitempty"`
	App     *string `json:"app,omitempty"`
}
