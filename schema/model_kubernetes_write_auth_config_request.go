// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KubernetesWriteAuthConfigRequest struct for KubernetesWriteAuthConfigRequest
type KubernetesWriteAuthConfigRequest struct {
	// Disable JWT issuer validation (Deprecated, will be removed in a future release)
	// Deprecated
	DisableIssValidation bool `json:"disable_iss_validation"`

	// Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod
	DisableLocalCaJwt bool `json:"disable_local_ca_jwt"`

	// Optional JWT issuer. If no issuer is specified, then this plugin will use kubernetes.io/serviceaccount as the default issuer. (Deprecated, will be removed in a future release)
	// Deprecated
	Issuer string `json:"issuer"`

	// PEM encoded CA cert for use by the TLS client used to talk with the API.
	KubernetesCaCert string `json:"kubernetes_ca_cert"`

	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost string `json:"kubernetes_host"`

	// Optional list of PEM-formated public keys or certificates used to verify the signatures of kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `json:"pem_keys"`

	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	TokenReviewerJwt string `json:"token_reviewer_jwt"`
}

// NewKubernetesWriteAuthConfigRequestWithDefaults instantiates a new KubernetesWriteAuthConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesWriteAuthConfigRequestWithDefaults() *KubernetesWriteAuthConfigRequest {
	var this KubernetesWriteAuthConfigRequest

	this.DisableIssValidation = true
	this.DisableLocalCaJwt = false

	return &this
}
