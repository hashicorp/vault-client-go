// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KubernetesWriteRoleRequest struct for KubernetesWriteRoleRequest
type KubernetesWriteRoleRequest struct {
	// A label selector for Kubernetes namespaces in which credentials can be generated. Accepts either a JSON or YAML object. If set with allowed_kubernetes_namespaces, the conditions are conjuncted.
	AllowedKubernetesNamespaceSelector string `json:"allowed_kubernetes_namespace_selector"`

	// A list of the Kubernetes namespaces in which credentials can be generated. If set to \"*\" all namespaces are allowed.
	AllowedKubernetesNamespaces []string `json:"allowed_kubernetes_namespaces"`

	// Additional annotations to apply to all generated Kubernetes objects.
	ExtraAnnotations map[string]interface{} `json:"extra_annotations"`

	// Additional labels to apply to all generated Kubernetes objects.
	ExtraLabels map[string]interface{} `json:"extra_labels"`

	// The Role or ClusterRole rules to use when generating a role. Accepts either a JSON or YAML object. If set, the entire chain of Kubernetes objects will be generated.
	GeneratedRoleRules string `json:"generated_role_rules"`

	// The pre-existing Role or ClusterRole to bind a generated service account to. If set, Kubernetes token, service account, and role binding objects will be created.
	KubernetesRoleName string `json:"kubernetes_role_name"`

	// Specifies whether the Kubernetes role is a Role or ClusterRole.
	KubernetesRoleType string `json:"kubernetes_role_type"`

	// The name template to use when generating service accounts, roles and role bindings. If unset, a default template is used.
	NameTemplate string `json:"name_template"`

	// The pre-existing service account to generate tokens for. Mutually exclusive with all role parameters. If set, only a Kubernetes service account token will be created.
	ServiceAccountName string `json:"service_account_name"`

	// The default ttl for generated Kubernetes service account tokens. If not set or set to 0, will use system default.
	TokenDefaultTtl int32 `json:"token_default_ttl"`

	// The maximum ttl for generated Kubernetes service account tokens. If not set or set to 0, will use system default.
	TokenMaxTtl int32 `json:"token_max_ttl"`
}

// NewKubernetesWriteRoleRequestWithDefaults instantiates a new KubernetesWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesWriteRoleRequestWithDefaults() *KubernetesWriteRoleRequest {
	var this KubernetesWriteRoleRequest

	this.KubernetesRoleType = "Role"

	return &this
}

func (o KubernetesWriteRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
