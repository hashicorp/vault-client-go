// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// ConsulConfigureAccessRequest struct for ConsulConfigureAccessRequest
type ConsulConfigureAccessRequest struct {
	// Consul server address
	Address string `json:"address,omitempty"`

	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert string `json:"ca_cert,omitempty"`

	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert string `json:"client_cert,omitempty"`

	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey string `json:"client_key,omitempty"`

	// URI scheme for the Consul address
	Scheme string `json:"scheme,omitempty"`

	// Token for API calls
	Token string `json:"token,omitempty"`
}
