// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SshIssueCertificateRequest struct for SshIssueCertificateRequest
type SshIssueCertificateRequest struct {
	// Type of certificate to be created; either \"user\" or \"host\".
	CertType string `json:"cert_type,omitempty"`

	// Critical options that the certificate should be signed for.
	CriticalOptions map[string]interface{} `json:"critical_options,omitempty"`

	// Extensions that the certificate should be signed for.
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	// Specifies the number of bits to use for the generated keys.
	KeyBits int32 `json:"key_bits,omitempty"`

	// Key id that the created certificate should have. If not specified, the display name of the token will be used.
	KeyId string `json:"key_id,omitempty"`

	// Specifies the desired key type; must be `rsa`, `ed25519` or `ec`
	KeyType string `json:"key_type,omitempty"`

	// The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL.
	Ttl string `json:"ttl,omitempty"`

	// Valid principals, either usernames or hostnames, that the certificate should be signed for.
	ValidPrincipals string `json:"valid_principals,omitempty"`
}
