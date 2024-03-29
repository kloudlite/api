// DO NOT EDIT. generated by "github.com/kloudlite/api/cmd/struct-json-path"

package field_constants

// constant vars generated for struct App
const (
	AppEnabled                = "enabled"
	AppSpec                   = "spec"
	AppSpecContainers         = "spec.containers"
	AppSpecDisplayName        = "spec.displayName"
	AppSpecFreeze             = "spec.freeze"
	AppSpecHpa                = "spec.hpa"
	AppSpecHpaEnabled         = "spec.hpa.enabled"
	AppSpecHpaMaxReplicas     = "spec.hpa.maxReplicas"
	AppSpecHpaMinReplicas     = "spec.hpa.minReplicas"
	AppSpecHpaThresholdCpu    = "spec.hpa.thresholdCpu"
	AppSpecHpaThresholdMemory = "spec.hpa.thresholdMemory"
	AppSpecIntercept          = "spec.intercept"
	AppSpecInterceptEnabled   = "spec.intercept.enabled"
	AppSpecInterceptToDevice  = "spec.intercept.toDevice"
	AppSpecNodeSelector       = "spec.nodeSelector"
	AppSpecRegion             = "spec.region"
	AppSpecReplicas           = "spec.replicas"
	AppSpecServiceAccount     = "spec.serviceAccount"
	AppSpecServices           = "spec.services"
	AppSpecTolerations        = "spec.tolerations"
)

// constant vars generated for struct Config
const (
	ConfigBinaryData = "binaryData"
	ConfigData       = "data"
	ConfigImmutable  = "immutable"
)

// constant vars generated for struct ConsoleVPNDevice
const (
	ConsoleVPNDeviceLinkedClusters        = "linkedClusters"
	ConsoleVPNDeviceSpec                  = "spec"
	ConsoleVPNDeviceSpecActiveNamespace   = "spec.activeNamespace"
	ConsoleVPNDeviceSpecCnameRecords      = "spec.cnameRecords"
	ConsoleVPNDeviceSpecDisabled          = "spec.disabled"
	ConsoleVPNDeviceSpecNoExternalService = "spec.noExternalService"
	ConsoleVPNDeviceSpecNodeSelector      = "spec.nodeSelector"
	ConsoleVPNDeviceSpecPorts             = "spec.ports"
	ConsoleVPNDeviceWireguardConfig       = "wireguardConfig"
	ConsoleVPNDeviceWireguardConfigs      = "wireguardConfigs"
)

// constant vars generated for struct Environment
const (
	EnvironmentSpec                           = "spec"
	EnvironmentSpecProjectName                = "spec.projectName"
	EnvironmentSpecRouting                    = "spec.routing"
	EnvironmentSpecRoutingMode                = "spec.routing.mode"
	EnvironmentSpecRoutingPrivateIngressClass = "spec.routing.privateIngressClass"
	EnvironmentSpecRoutingPublicIngressClass  = "spec.routing.publicIngressClass"
	EnvironmentSpecTargetNamespace            = "spec.targetNamespace"
)

// constant vars generated for struct ImagePullSecret
const (
	ImagePullSecretDockerConfigJson   = "dockerConfigJson"
	ImagePullSecretFormat             = "format"
	ImagePullSecretGeneratedK8sSecret = "generatedK8sSecret"
	ImagePullSecretRegistryPassword   = "registryPassword"
	ImagePullSecretRegistryURL        = "registryURL"
	ImagePullSecretRegistryUsername   = "registryUsername"
)

// constant vars generated for struct ManagedResource
const (
	ManagedResourceEnabled                               = "enabled"
	ManagedResourceSpec                                  = "spec"
	ManagedResourceSpecResourceName                      = "spec.resourceName"
	ManagedResourceSpecResourceTemplate                  = "spec.resourceTemplate"
	ManagedResourceSpecResourceTemplateApiVersion        = "spec.resourceTemplate.apiVersion"
	ManagedResourceSpecResourceTemplateKind              = "spec.resourceTemplate.kind"
	ManagedResourceSpecResourceTemplateMsvcRef           = "spec.resourceTemplate.msvcRef"
	ManagedResourceSpecResourceTemplateMsvcRefApiVersion = "spec.resourceTemplate.msvcRef.apiVersion"
	ManagedResourceSpecResourceTemplateMsvcRefKind       = "spec.resourceTemplate.msvcRef.kind"
	ManagedResourceSpecResourceTemplateMsvcRefName       = "spec.resourceTemplate.msvcRef.name"
	ManagedResourceSpecResourceTemplateMsvcRefNamespace  = "spec.resourceTemplate.msvcRef.namespace"
	ManagedResourceSpecResourceTemplateSpec              = "spec.resourceTemplate.spec"
	ManagedResourceSyncedOutputSecretRef                 = "syncedOutputSecretRef"
)

// constant vars generated for struct Project
const (
	ProjectSpec                = "spec"
	ProjectSpecTargetNamespace = "spec.targetNamespace"
)

// constant vars generated for struct ProjectManagedService
const (
	ProjectManagedServiceSpec                                  = "spec"
	ProjectManagedServiceSpecMsvcSpec                          = "spec.msvcSpec"
	ProjectManagedServiceSpecMsvcSpecServiceTemplate           = "spec.msvcSpec.serviceTemplate"
	ProjectManagedServiceSpecMsvcSpecServiceTemplateApiVersion = "spec.msvcSpec.serviceTemplate.apiVersion"
	ProjectManagedServiceSpecMsvcSpecServiceTemplateKind       = "spec.msvcSpec.serviceTemplate.kind"
	ProjectManagedServiceSpecMsvcSpecServiceTemplateSpec       = "spec.msvcSpec.serviceTemplate.spec"
	ProjectManagedServiceSpecTargetNamespace                   = "spec.targetNamespace"
	ProjectManagedServiceSyncedOutputSecretRef                 = "syncedOutputSecretRef"
)

// constant vars generated for struct ResourceMapping
const (
	ResourceMappingBaseEntity                  = "BaseEntity"
	ResourceMappingBaseEntityCreationTime      = "BaseEntity.creationTime"
	ResourceMappingBaseEntityId                = "BaseEntity.id"
	ResourceMappingBaseEntityMarkedForDeletion = "BaseEntity.markedForDeletion"
	ResourceMappingBaseEntityRecordVersion     = "BaseEntity.recordVersion"
	ResourceMappingBaseEntityUpdateTime        = "BaseEntity.updateTime"
	ResourceMappingResourceHeirarchy           = "resourceHeirarchy"
	ResourceMappingResourceName                = "resourceName"
	ResourceMappingResourceNamespace           = "resourceNamespace"
	ResourceMappingResourceType                = "resourceType"
)

// constant vars generated for struct Router
const (
	RouterEnabled                  = "enabled"
	RouterSpec                     = "spec"
	RouterSpecBackendProtocol      = "spec.backendProtocol"
	RouterSpecBasicAuth            = "spec.basicAuth"
	RouterSpecBasicAuthEnabled     = "spec.basicAuth.enabled"
	RouterSpecBasicAuthSecretName  = "spec.basicAuth.secretName"
	RouterSpecBasicAuthUsername    = "spec.basicAuth.username"
	RouterSpecCors                 = "spec.cors"
	RouterSpecCorsAllowCredentials = "spec.cors.allowCredentials"
	RouterSpecCorsEnabled          = "spec.cors.enabled"
	RouterSpecCorsOrigins          = "spec.cors.origins"
	RouterSpecDomains              = "spec.domains"
	RouterSpecHttps                = "spec.https"
	RouterSpecHttpsClusterIssuer   = "spec.https.clusterIssuer"
	RouterSpecHttpsEnabled         = "spec.https.enabled"
	RouterSpecHttpsForceRedirect   = "spec.https.forceRedirect"
	RouterSpecIngressClass         = "spec.ingressClass"
	RouterSpecMaxBodySizeInMB      = "spec.maxBodySizeInMB"
	RouterSpecRateLimit            = "spec.rateLimit"
	RouterSpecRateLimitConnections = "spec.rateLimit.connections"
	RouterSpecRateLimitEnabled     = "spec.rateLimit.enabled"
	RouterSpecRateLimitRpm         = "spec.rateLimit.rpm"
	RouterSpecRateLimitRps         = "spec.rateLimit.rps"
	RouterSpecRoutes               = "spec.routes"
)

// constant vars generated for struct Secret
const (
	SecretData       = "data"
	SecretImmutable  = "immutable"
	SecretStringData = "stringData"
	SecretType       = "type"
)

// constant vars generated for struct
const (
	AccountName                        = "accountName"
	ApiVersion                         = "apiVersion"
	ClusterName                        = "clusterName"
	CreatedBy                          = "createdBy"
	CreatedByUserEmail                 = "createdBy.userEmail"
	CreatedByUserId                    = "createdBy.userId"
	CreatedByUserName                  = "createdBy.userName"
	CreationTime                       = "creationTime"
	DisplayName                        = "displayName"
	EnvironmentName                    = "environmentName"
	Id                                 = "id"
	Kind                               = "kind"
	LastUpdatedBy                      = "lastUpdatedBy"
	LastUpdatedByUserEmail             = "lastUpdatedBy.userEmail"
	LastUpdatedByUserId                = "lastUpdatedBy.userId"
	LastUpdatedByUserName              = "lastUpdatedBy.userName"
	MarkedForDeletion                  = "markedForDeletion"
	Metadata                           = "metadata"
	MetadataAnnotations                = "metadata.annotations"
	MetadataCreationTimestamp          = "metadata.creationTimestamp"
	MetadataDeletionGracePeriodSeconds = "metadata.deletionGracePeriodSeconds"
	MetadataDeletionTimestamp          = "metadata.deletionTimestamp"
	MetadataFinalizers                 = "metadata.finalizers"
	MetadataGenerateName               = "metadata.generateName"
	MetadataGeneration                 = "metadata.generation"
	MetadataLabels                     = "metadata.labels"
	MetadataManagedFields              = "metadata.managedFields"
	MetadataName                       = "metadata.name"
	MetadataNamespace                  = "metadata.namespace"
	MetadataOwnerReferences            = "metadata.ownerReferences"
	MetadataResourceVersion            = "metadata.resourceVersion"
	MetadataSelfLink                   = "metadata.selfLink"
	MetadataUid                        = "metadata.uid"
	ProjectName                        = "projectName"
	RecordVersion                      = "recordVersion"
	Status                             = "status"
	StatusChecks                       = "status.checks"
	StatusIsReady                      = "status.isReady"
	StatusLastReadyGeneration          = "status.lastReadyGeneration"
	StatusLastReconcileTime            = "status.lastReconcileTime"
	StatusMessage                      = "status.message"
	StatusMessageItems                 = "status.message.items"
	StatusResources                    = "status.resources"
	SyncStatus                         = "syncStatus"
	SyncStatusAction                   = "syncStatus.action"
	SyncStatusError                    = "syncStatus.error"
	SyncStatusLastSyncedAt             = "syncStatus.lastSyncedAt"
	SyncStatusRecordVersion            = "syncStatus.recordVersion"
	SyncStatusState                    = "syncStatus.state"
	SyncStatusSyncScheduledAt          = "syncStatus.syncScheduledAt"
	UpdateTime                         = "updateTime"
)
