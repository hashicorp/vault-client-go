// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CorsConfigureRequest struct for CorsConfigureRequest
type CorsConfigureRequest struct {
	// A comma-separated string or array of strings indicating headers that are allowed on cross-origin requests.
	AllowedHeaders []string `json:"allowed_headers"`

	// A comma-separated string or array of strings indicating origins that may make cross-origin requests.
	AllowedOrigins []string `json:"allowed_origins"`

	// Enables or disables CORS headers on requests.
	Enable bool `json:"enable"`
}

// NewCorsConfigureRequestWithDefaults instantiates a new CorsConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorsConfigureRequestWithDefaults() *CorsConfigureRequest {
	var this CorsConfigureRequest

	return &this
}

func (o CorsConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allowed_headers"] = o.AllowedHeaders
	toSerialize["allowed_origins"] = o.AllowedOrigins
	toSerialize["enable"] = o.Enable

	return json.Marshal(toSerialize)
}
