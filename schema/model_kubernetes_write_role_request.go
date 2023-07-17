// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KubernetesWriteRoleRequest struct for KubernetesWriteRoleRequest
type KubernetesWriteRoleRequest struct {
	// A label selector for Kubernetes namespaces in which credentials can be generated. Accepts either a JSON or YAML object. If set with allowed_kubernetes_namespaces, the conditions are conjuncted.
	AllowedKubernetesNamespaceSelector string `json:"allowed_kubernetes_namespace_selector,omitempty"`

	// A list of the Kubernetes namespaces in which credentials can be generated. If set to \"*\" all namespaces are allowed.
	AllowedKubernetesNamespaces []string `json:"allowed_kubernetes_namespaces,omitempty"`

	// Additional annotations to apply to all generated Kubernetes objects.
	ExtraAnnotations map[string]interface{} `json:"extra_annotations,omitempty"`

	// Additional labels to apply to all generated Kubernetes objects.
	ExtraLabels map[string]interface{} `json:"extra_labels,omitempty"`

	// The Role or ClusterRole rules to use when generating a role. Accepts either a JSON or YAML object. If set, the entire chain of Kubernetes objects will be generated.
	GeneratedRoleRules string `json:"generated_role_rules,omitempty"`

	// The pre-existing Role or ClusterRole to bind a generated service account to. If set, Kubernetes token, service account, and role binding objects will be created.
	KubernetesRoleName string `json:"kubernetes_role_name,omitempty"`

	// Specifies whether the Kubernetes role is a Role or ClusterRole.
	KubernetesRoleType string `json:"kubernetes_role_type,omitempty"`

	// The name template to use when generating service accounts, roles and role bindings. If unset, a default template is used.
	NameTemplate string `json:"name_template,omitempty"`

	// The pre-existing service account to generate tokens for. Mutually exclusive with all role parameters. If set, only a Kubernetes service account token will be created.
	ServiceAccountName string `json:"service_account_name,omitempty"`

	// The default audiences for generated Kubernetes service account tokens. If not set or set to \"\", will use k8s cluster default.
	TokenDefaultAudiences []string `json:"token_default_audiences,omitempty"`

	// The default ttl for generated Kubernetes service account tokens. If not set or set to 0, will use system default.
	TokenDefaultTtl string `json:"token_default_ttl,omitempty"`

	// The maximum ttl for generated Kubernetes service account tokens. If not set or set to 0, will use system default.
	TokenMaxTtl string `json:"token_max_ttl,omitempty"`
}
