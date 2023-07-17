// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KubernetesConfigureRequest struct for KubernetesConfigureRequest
type KubernetesConfigureRequest struct {
	// Disable defaulting to the local CA certificate and service account JWT when running in a Kubernetes pod.
	DisableLocalCaJwt bool `json:"disable_local_ca_jwt,omitempty"`

	// PEM encoded CA certificate to use to verify the Kubernetes API server certificate. Defaults to the local pod's CA if found.
	KubernetesCaCert string `json:"kubernetes_ca_cert,omitempty"`

	// Kubernetes API URL to connect to. Defaults to https://$KUBERNETES_SERVICE_HOST:KUBERNETES_SERVICE_PORT if those environment variables are set.
	KubernetesHost string `json:"kubernetes_host,omitempty"`

	// The JSON web token of the service account used by the secret engine to manage Kubernetes credentials. Defaults to the local pod's JWT if found.
	ServiceAccountJwt string `json:"service_account_jwt,omitempty"`
}
