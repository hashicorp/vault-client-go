// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteRoleRequest struct for OIDCWriteRoleRequest
type OIDCWriteRoleRequest struct {
	// Optional client_id
	ClientId string `json:"client_id"`

	// The OIDC key to use for generating tokens. The specified key must already exist.
	Key string `json:"key"`

	// The template string to use for generating tokens. This may be in string-ified JSON or base64 format.
	Template string `json:"template"`

	// TTL of the tokens generated against the role.
	Ttl int32 `json:"ttl"`
}

// NewOIDCWriteRoleRequestWithDefaults instantiates a new OIDCWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteRoleRequestWithDefaults() *OIDCWriteRoleRequest {
	var this OIDCWriteRoleRequest

	return &this
}

func (o OIDCWriteRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
