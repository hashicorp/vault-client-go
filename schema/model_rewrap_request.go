// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RewrapRequest struct for RewrapRequest
type RewrapRequest struct {
	Token string `json:"token"`
}

// NewRewrapRequestWithDefaults instantiates a new RewrapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewrapRequestWithDefaults() *RewrapRequest {
	var this RewrapRequest

	return &this
}

func (o RewrapRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
