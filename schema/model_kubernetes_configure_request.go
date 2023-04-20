// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KubernetesConfigureRequest struct for KubernetesConfigureRequest
type KubernetesConfigureRequest struct {
	// Disable defaulting to the local CA certificate and service account JWT when running in a Kubernetes pod.
	DisableLocalCaJwt bool `json:"disable_local_ca_jwt"`

	// PEM encoded CA certificate to use to verify the Kubernetes API server certificate. Defaults to the local pod's CA if found.
	KubernetesCaCert string `json:"kubernetes_ca_cert"`

	// Kubernetes API URL to connect to. Defaults to https://$KUBERNETES_SERVICE_HOST:KUBERNETES_SERVICE_PORT if those environment variables are set.
	KubernetesHost string `json:"kubernetes_host"`

	// The JSON web token of the service account used by the secret engine to manage Kubernetes credentials. Defaults to the local pod's JWT if found.
	ServiceAccountJwt string `json:"service_account_jwt"`
}

// NewKubernetesConfigureRequestWithDefaults instantiates a new KubernetesConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesConfigureRequestWithDefaults() *KubernetesConfigureRequest {
	var this KubernetesConfigureRequest

	this.DisableLocalCaJwt = false

	return &this
}

func (o KubernetesConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["disable_local_ca_jwt"] = o.DisableLocalCaJwt
	toSerialize["kubernetes_ca_cert"] = o.KubernetesCaCert
	toSerialize["kubernetes_host"] = o.KubernetesHost
	toSerialize["service_account_jwt"] = o.ServiceAccountJwt

	return json.Marshal(toSerialize)
}
