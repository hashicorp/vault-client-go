// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// ReadWrappingPropertiesResponse struct for ReadWrappingPropertiesResponse
type ReadWrappingPropertiesResponse struct {
	CreationPath string `json:"creation_path,omitempty"`

	CreationTime time.Time `json:"creation_time,omitempty"`

	CreationTtl string `json:"creation_ttl,omitempty"`
}

// NewReadWrappingPropertiesResponseWithDefaults instantiates a new ReadWrappingPropertiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadWrappingPropertiesResponseWithDefaults() *ReadWrappingPropertiesResponse {
	var this ReadWrappingPropertiesResponse

	return &this
}
