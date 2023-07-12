// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SshSignCertificateRequest struct for SshSignCertificateRequest
type SshSignCertificateRequest struct {
	// Type of certificate to be created; either \"user\" or \"host\".
	CertType string `json:"cert_type,omitempty"`

	// Critical options that the certificate should be signed for.
	CriticalOptions map[string]interface{} `json:"critical_options,omitempty"`

	// Extensions that the certificate should be signed for.
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	// Key id that the created certificate should have. If not specified, the display name of the token will be used.
	KeyId string `json:"key_id,omitempty"`

	// SSH public key that should be signed.
	PublicKey string `json:"public_key,omitempty"`

	// The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL.
	Ttl string `json:"ttl,omitempty"`

	// Valid principals, either usernames or hostnames, that the certificate should be signed for.
	ValidPrincipals string `json:"valid_principals,omitempty"`
}
