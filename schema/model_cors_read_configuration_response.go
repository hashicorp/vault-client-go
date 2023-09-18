// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// CorsReadConfigurationResponse struct for CorsReadConfigurationResponse
type CorsReadConfigurationResponse struct {
	AllowedHeaders []string `json:"allowed_headers,omitempty"`

	AllowedOrigins []string `json:"allowed_origins,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
}
