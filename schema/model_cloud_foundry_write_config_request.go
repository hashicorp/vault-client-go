// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CloudFoundryWriteConfigRequest struct for CloudFoundryWriteConfigRequest
type CloudFoundryWriteConfigRequest struct {
	// CF’s API address.
	CfApiAddr string `json:"cf_api_addr"`

	// The PEM-format certificates that are presented for mutual TLS with the CloudFoundry API. If not set, mutual TLS is not used
	CfApiMutualTlsCertificate string `json:"cf_api_mutual_tls_certificate"`

	// The PEM-format private key that are used for mutual TLS with the CloudFoundry API. If not set, mutual TLS is not used
	CfApiMutualTlsKey string `json:"cf_api_mutual_tls_key"`

	// The PEM-format CA certificates that are acceptable for the CF API to present.
	CfApiTrustedCertificates []string `json:"cf_api_trusted_certificates"`

	// The client id for CF’s API.
	CfClientId string `json:"cf_client_id"`

	// The client secret for CF’s API.
	CfClientSecret string `json:"cf_client_secret"`

	// The password for CF’s API.
	CfPassword string `json:"cf_password"`

	// The username for CF’s API.
	CfUsername string `json:"cf_username"`

	// The PEM-format CA certificates that are required to have issued the instance certificates presented for logging in.
	IdentityCaCertificates []string `json:"identity_ca_certificates"`

	// Duration in seconds for the maximum acceptable length in the future a \"signing_time\" can be. Useful for clock drift. Set low to reduce the opportunity for replay attacks.
	LoginMaxSecondsNotAfter int32 `json:"login_max_seconds_not_after"`

	// Duration in seconds for the maximum acceptable age of a \"signing_time\". Useful for clock drift. Set low to reduce the opportunity for replay attacks.
	LoginMaxSecondsNotBefore int32 `json:"login_max_seconds_not_before"`

	// Deprecated. Please use \"cf_api_addr\".
	// Deprecated
	PcfApiAddr string `json:"pcf_api_addr"`

	// Deprecated. Please use \"cf_api_trusted_certificates\".
	// Deprecated
	PcfApiTrustedCertificates []string `json:"pcf_api_trusted_certificates"`

	// Deprecated. Please use \"cf_password\".
	// Deprecated
	PcfPassword string `json:"pcf_password"`

	// Deprecated. Please use \"cf_username\".
	// Deprecated
	PcfUsername string `json:"pcf_username"`
}

// NewCloudFoundryWriteConfigRequestWithDefaults instantiates a new CloudFoundryWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFoundryWriteConfigRequestWithDefaults() *CloudFoundryWriteConfigRequest {
	var this CloudFoundryWriteConfigRequest

	this.LoginMaxSecondsNotAfter = 60
	this.LoginMaxSecondsNotBefore = 300

	return &this
}

func (o CloudFoundryWriteConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
