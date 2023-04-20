// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// UiHeadersConfigureRequest struct for UiHeadersConfigureRequest
type UiHeadersConfigureRequest struct {
	// Returns multiple values if true
	Multivalue bool `json:"multivalue"`

	// The values to set the header.
	Values []string `json:"values"`
}

// NewUiHeadersConfigureRequestWithDefaults instantiates a new UiHeadersConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiHeadersConfigureRequestWithDefaults() *UiHeadersConfigureRequest {
	var this UiHeadersConfigureRequest

	return &this
}

func (o UiHeadersConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["multivalue"] = o.Multivalue
	toSerialize["values"] = o.Values

	return json.Marshal(toSerialize)
}
