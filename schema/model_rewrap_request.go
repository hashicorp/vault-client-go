// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RewrapRequest struct for RewrapRequest
type RewrapRequest struct {
	Token string `json:"token,omitempty"`
}

// NewRewrapRequestWithDefaults instantiates a new RewrapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewrapRequestWithDefaults() *RewrapRequest {
	var this RewrapRequest

	return &this
}
