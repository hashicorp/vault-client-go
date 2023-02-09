// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SSHWriteIssueRequest struct for SSHWriteIssueRequest
type SSHWriteIssueRequest struct {
	// Type of certificate to be created; either \"user\" or \"host\".
	CertType string `json:"cert_type"`

	// Critical options that the certificate should be signed for.
	CriticalOptions map[string]interface{} `json:"critical_options"`

	// Extensions that the certificate should be signed for.
	Extensions map[string]interface{} `json:"extensions"`

	// Specifies the number of bits to use for the generated keys.
	KeyBits int32 `json:"key_bits"`

	// Key id that the created certificate should have. If not specified, the display name of the token will be used.
	KeyId string `json:"key_id"`

	// Specifies the desired key type; must be `rsa`, `ed25519` or `ec`
	KeyType string `json:"key_type"`

	// The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL.
	Ttl int32 `json:"ttl"`

	// Valid principals, either usernames or hostnames, that the certificate should be signed for.
	ValidPrincipals string `json:"valid_principals"`
}

// NewSSHWriteIssueRequestWithDefaults instantiates a new SSHWriteIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHWriteIssueRequestWithDefaults() *SSHWriteIssueRequest {
	var this SSHWriteIssueRequest

	this.CertType = "user"
	this.KeyBits = 0
	this.KeyType = "rsa"

	return &this
}

func (o SSHWriteIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
