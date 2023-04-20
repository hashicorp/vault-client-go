// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// UiHeadersReadConfigurationResponse struct for UiHeadersReadConfigurationResponse
type UiHeadersReadConfigurationResponse struct {
	// returns the first header value when `multivalue` request parameter is false
	Value string `json:"value"`

	// returns all header values when `multivalue` request parameter is true
	Values []string `json:"values"`
}

// NewUiHeadersReadConfigurationResponseWithDefaults instantiates a new UiHeadersReadConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiHeadersReadConfigurationResponseWithDefaults() *UiHeadersReadConfigurationResponse {
	var this UiHeadersReadConfigurationResponse

	return &this
}

func (o UiHeadersReadConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["value"] = o.Value
	toSerialize["values"] = o.Values

	return json.Marshal(toSerialize)
}
