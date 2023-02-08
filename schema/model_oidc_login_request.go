// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCLoginRequest struct for OIDCLoginRequest
type OIDCLoginRequest struct {
	// The signed JWT to validate.
	Jwt string `json:"jwt"`

	// The role to log in against.
	Role string `json:"role"`
}

// NewOIDCLoginRequestWithDefaults instantiates a new OIDCLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCLoginRequestWithDefaults() *OIDCLoginRequest {
	var this OIDCLoginRequest

	return &this
}

func (o OIDCLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
