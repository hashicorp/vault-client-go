// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KubernetesWriteCredentialsRequest struct for KubernetesWriteCredentialsRequest
type KubernetesWriteCredentialsRequest struct {
	// If true, generate a ClusterRoleBinding to grant permissions across the whole cluster instead of within a namespace. Requires the Vault role to have kubernetes_role_type set to ClusterRole.
	ClusterRoleBinding bool `json:"cluster_role_binding"`

	// The name of the Kubernetes namespace in which to generate the credentials
	KubernetesNamespace string `json:"kubernetes_namespace"`

	// The TTL of the generated credentials
	Ttl int32 `json:"ttl"`
}

// NewKubernetesWriteCredentialsRequestWithDefaults instantiates a new KubernetesWriteCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesWriteCredentialsRequestWithDefaults() *KubernetesWriteCredentialsRequest {
	var this KubernetesWriteCredentialsRequest

	return &this
}

func (o KubernetesWriteCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
