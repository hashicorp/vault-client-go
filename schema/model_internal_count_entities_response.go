// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalCountEntitiesResponse struct for InternalCountEntitiesResponse
type InternalCountEntitiesResponse struct {
	Counters map[string]interface{} `json:"counters,omitempty"`
}

// NewInternalCountEntitiesResponseWithDefaults instantiates a new InternalCountEntitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalCountEntitiesResponseWithDefaults() *InternalCountEntitiesResponse {
	var this InternalCountEntitiesResponse

	return &this
}
