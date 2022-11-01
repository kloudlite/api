package op_crds

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type SecretMetadata struct {
	Name            string                  `json:"name,omitempty"`
	Namespace       string                  `json:"namespace,omitempty"`
	Annotations     map[string]string       `json:"annotations,omitempty"`
	Labels          map[string]string       `json:"labels,omitempty"`
	OwnerReferences []metav1.OwnerReference `json:"ownerReference,omitempty"`
}

const SecretAPIVersion = "v1"
const SecretKind = "Secret"

type Secret struct {
	APIVersion string         `json:"apiVersion,omitempty"`
	Kind       string         `json:"kind,omitempty"`
	Metadata   SecretMetadata `json:"metadata,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
	StringData map[string]any `json:"stringData,omitempty"`
}
