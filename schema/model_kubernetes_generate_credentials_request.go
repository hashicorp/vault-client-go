// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KubernetesGenerateCredentialsRequest struct for KubernetesGenerateCredentialsRequest
type KubernetesGenerateCredentialsRequest struct {
	// The intended audiences of the generated credentials
	Audiences []string `json:"audiences,omitempty"`

	// If true, generate a ClusterRoleBinding to grant permissions across the whole cluster instead of within a namespace. Requires the Vault role to have kubernetes_role_type set to ClusterRole.
	ClusterRoleBinding bool `json:"cluster_role_binding,omitempty"`

	// The name of the Kubernetes namespace in which to generate the credentials
	KubernetesNamespace string `json:"kubernetes_namespace"`

	// The TTL of the generated credentials
	Ttl string `json:"ttl,omitempty"`
}
