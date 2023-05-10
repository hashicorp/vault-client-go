// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// UnwrapRequest struct for UnwrapRequest
type UnwrapRequest struct {
	Token string `json:"token,omitempty"`
}

// NewUnwrapRequestWithDefaults instantiates a new UnwrapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnwrapRequestWithDefaults() *UnwrapRequest {
	var this UnwrapRequest

	return &this
}
