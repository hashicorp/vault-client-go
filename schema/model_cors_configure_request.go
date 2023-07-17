// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// CorsConfigureRequest struct for CorsConfigureRequest
type CorsConfigureRequest struct {
	// A comma-separated string or array of strings indicating headers that are allowed on cross-origin requests.
	AllowedHeaders []string `json:"allowed_headers,omitempty"`

	// A comma-separated string or array of strings indicating origins that may make cross-origin requests.
	AllowedOrigins []string `json:"allowed_origins,omitempty"`

	// Enables or disables CORS headers on requests.
	Enable bool `json:"enable,omitempty"`
}
