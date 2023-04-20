// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KubernetesGenerateCredentialsRequest struct for KubernetesGenerateCredentialsRequest
type KubernetesGenerateCredentialsRequest struct {
	// The intended audiences of the generated credentials
	Audiences []string `json:"audiences"`

	// If true, generate a ClusterRoleBinding to grant permissions across the whole cluster instead of within a namespace. Requires the Vault role to have kubernetes_role_type set to ClusterRole.
	ClusterRoleBinding bool `json:"cluster_role_binding"`

	// The name of the Kubernetes namespace in which to generate the credentials
	KubernetesNamespace string `json:"kubernetes_namespace"`

	// The TTL of the generated credentials
	Ttl int32 `json:"ttl"`
}

// NewKubernetesGenerateCredentialsRequestWithDefaults instantiates a new KubernetesGenerateCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesGenerateCredentialsRequestWithDefaults() *KubernetesGenerateCredentialsRequest {
	var this KubernetesGenerateCredentialsRequest

	return &this
}

func (o KubernetesGenerateCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["audiences"] = o.Audiences
	toSerialize["cluster_role_binding"] = o.ClusterRoleBinding
	toSerialize["kubernetes_namespace"] = o.KubernetesNamespace
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
