// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KubernetesConfigureAuthRequest struct for KubernetesConfigureAuthRequest
type KubernetesConfigureAuthRequest struct {
	// Disable JWT issuer validation (Deprecated, will be removed in a future release)
	// Deprecated
	DisableIssValidation bool `json:"disable_iss_validation,omitempty"`

	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod
	DisableLocalCaJwt bool `json:"disable_local_ca_jwt,omitempty"`

	// Optional JWT issuer. If no issuer is specified, then this plugin will use kubernetes.io/serviceaccount as the default issuer. (Deprecated, will be removed in a future release)
	// Deprecated
	Issuer string `json:"issuer,omitempty"`

	// PEM encoded CA cert for use by the TLS client used to talk with the API.
	KubernetesCaCert string `json:"kubernetes_ca_cert,omitempty"`

	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost string `json:"kubernetes_host,omitempty"`

	// Optional list of PEM-formated public keys or certificates used to verify the signatures of kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `json:"pem_keys,omitempty"`

	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt string `json:"token_reviewer_jwt,omitempty"`
}
