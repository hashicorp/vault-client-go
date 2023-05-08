// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// ReadWrappingPropertiesRequest struct for ReadWrappingPropertiesRequest
type ReadWrappingPropertiesRequest struct {
	Token string `json:"token,omitempty"`
}

// NewReadWrappingPropertiesRequestWithDefaults instantiates a new ReadWrappingPropertiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadWrappingPropertiesRequestWithDefaults() *ReadWrappingPropertiesRequest {
	var this ReadWrappingPropertiesRequest

	return &this
}
