// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CorsReadConfigurationResponse struct for CorsReadConfigurationResponse
type CorsReadConfigurationResponse struct {
	AllowedHeaders []string `json:"allowed_headers"`

	AllowedOrigins []string `json:"allowed_origins"`

	Enabled bool `json:"enabled"`
}

// NewCorsReadConfigurationResponseWithDefaults instantiates a new CorsReadConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorsReadConfigurationResponseWithDefaults() *CorsReadConfigurationResponse {
	var this CorsReadConfigurationResponse

	return &this
}

func (o CorsReadConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allowed_headers"] = o.AllowedHeaders
	toSerialize["allowed_origins"] = o.AllowedOrigins
	toSerialize["enabled"] = o.Enabled

	return json.Marshal(toSerialize)
}
