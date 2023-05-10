// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// UiHeadersReadConfigurationResponse struct for UiHeadersReadConfigurationResponse
type UiHeadersReadConfigurationResponse struct {
	// returns the first header value when `multivalue` request parameter is false
	Value string `json:"value,omitempty"`

	// returns all header values when `multivalue` request parameter is true
	Values []string `json:"values,omitempty"`
}

// NewUiHeadersReadConfigurationResponseWithDefaults instantiates a new UiHeadersReadConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiHeadersReadConfigurationResponseWithDefaults() *UiHeadersReadConfigurationResponse {
	var this UiHeadersReadConfigurationResponse

	return &this
}
