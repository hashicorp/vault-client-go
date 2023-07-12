// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// NomadConfigureAccessRequest struct for NomadConfigureAccessRequest
type NomadConfigureAccessRequest struct {
	// Nomad server address
	Address string `json:"address,omitempty"`

	// CA certificate to use when verifying Nomad server certificate, must be x509 PEM encoded.
	CaCert string `json:"ca_cert,omitempty"`

	// Client certificate used for Nomad's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert string `json:"client_cert,omitempty"`

	// Client key used for Nomad's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey string `json:"client_key,omitempty"`

	// Max length for name of generated Nomad tokens
	MaxTokenNameLength int32 `json:"max_token_name_length,omitempty"`

	// Token for API calls
	Token string `json:"token,omitempty"`
}
