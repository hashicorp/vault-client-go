// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// UnwrapRequest struct for UnwrapRequest
type UnwrapRequest struct {
	Token string `json:"token"`
}

// NewUnwrapRequestWithDefaults instantiates a new UnwrapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnwrapRequestWithDefaults() *UnwrapRequest {
	var this UnwrapRequest

	return &this
}

func (o UnwrapRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
