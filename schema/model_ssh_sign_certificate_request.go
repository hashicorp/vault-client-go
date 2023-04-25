// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SshSignCertificateRequest struct for SshSignCertificateRequest
type SshSignCertificateRequest struct {
	// Type of certificate to be created; either \"user\" or \"host\".
	CertType string `json:"cert_type"`

	// Critical options that the certificate should be signed for.
	CriticalOptions map[string]interface{} `json:"critical_options"`

	// Extensions that the certificate should be signed for.
	Extensions map[string]interface{} `json:"extensions"`

	// Key id that the created certificate should have. If not specified, the display name of the token will be used.
	KeyId string `json:"key_id"`

	// SSH public key that should be signed.
	PublicKey string `json:"public_key"`

	// The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL.
	Ttl int32 `json:"ttl"`

	// Valid principals, either usernames or hostnames, that the certificate should be signed for.
	ValidPrincipals string `json:"valid_principals"`
}

// NewSshSignCertificateRequestWithDefaults instantiates a new SshSignCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshSignCertificateRequestWithDefaults() *SshSignCertificateRequest {
	var this SshSignCertificateRequest

	this.CertType = "user"

	return &this
}

func (o SshSignCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cert_type"] = o.CertType
	toSerialize["critical_options"] = o.CriticalOptions
	toSerialize["extensions"] = o.Extensions
	toSerialize["key_id"] = o.KeyId
	toSerialize["public_key"] = o.PublicKey
	toSerialize["ttl"] = o.Ttl
	toSerialize["valid_principals"] = o.ValidPrincipals

	return json.Marshal(toSerialize)
}
