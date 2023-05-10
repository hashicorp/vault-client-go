// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OidcWriteKeyRequest struct for OidcWriteKeyRequest
type OidcWriteKeyRequest struct {
	// Signing algorithm to use. This will default to RS256.
	Algorithm string `json:"algorithm"`

	// Comma separated string or array of role client ids allowed to use this key for signing. If empty no roles are allowed. If \"*\" all roles are allowed.
	AllowedClientIds []string `json:"allowed_client_ids"`

	// How often to generate a new keypair.
	RotationPeriod int32 `json:"rotation_period"`

	// Controls how long the public portion of a key will be available for verification after being rotated.
	VerificationTtl int32 `json:"verification_ttl"`
}

// NewOidcWriteKeyRequestWithDefaults instantiates a new OidcWriteKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcWriteKeyRequestWithDefaults() *OidcWriteKeyRequest {
	var this OidcWriteKeyRequest

	this.Algorithm = "RS256"

	return &this
}

func (o OidcWriteKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["algorithm"] = o.Algorithm
	toSerialize["allowed_client_ids"] = o.AllowedClientIds
	toSerialize["rotation_period"] = o.RotationPeriod
	toSerialize["verification_ttl"] = o.VerificationTtl

	return json.Marshal(toSerialize)
}
