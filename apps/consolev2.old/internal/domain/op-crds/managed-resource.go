package op_crds

type MsvcRef struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
}

type MresKind struct {
	Kind string `json:"kind"`
}

type ManagedResourceSpec struct {
	MsvcRef  MsvcRef           `json:"msvcRef"`
	MresKind MresKind          `json:"mresKind"`
	Inputs   map[string]string `json:"inputs,omitempty"`
}

type ManagedResourceMetadata struct {
	Name        string            `json:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

const ManagedResourceAPIVersion = "crds.kloudlite.io/v1"
const ManagedResourceKind = "ManagedResource"

type ManagedResource struct {
	APIVersion string                  `json:"apiVersion,omitempty"`
	Kind       string                  `json:"kind,omitempty"`
	Metadata   ManagedResourceMetadata `json:"metadata"`
	Spec       ManagedResourceSpec     `json:"spec,omitempty"`
}