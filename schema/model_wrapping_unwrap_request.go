// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WrappingUnwrapRequest struct for WrappingUnwrapRequest
type WrappingUnwrapRequest struct {
	Token string `json:"token"`
}

// NewWrappingUnwrapRequestWithDefaults instantiates a new WrappingUnwrapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWrappingUnwrapRequestWithDefaults() *WrappingUnwrapRequest {
	var this WrappingUnwrapRequest

	return &this
}

func (o WrappingUnwrapRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
