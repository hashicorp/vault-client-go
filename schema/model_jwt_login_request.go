// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// JwtLoginRequest struct for JwtLoginRequest
type JwtLoginRequest struct {
	// The signed JWT to validate.
	Jwt string `json:"jwt"`

	// The role to log in against.
	Role string `json:"role"`
}

// NewJwtLoginRequestWithDefaults instantiates a new JwtLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJwtLoginRequestWithDefaults() *JwtLoginRequest {
	var this JwtLoginRequest

	return &this
}

func (o JwtLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["jwt"] = o.Jwt
	toSerialize["role"] = o.Role

	return json.Marshal(toSerialize)
}
