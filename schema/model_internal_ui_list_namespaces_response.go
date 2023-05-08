// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalUiListNamespacesResponse struct for InternalUiListNamespacesResponse
type InternalUiListNamespacesResponse struct {
	// field is only returned if there are one or more namespaces
	Keys []string `json:"keys,omitempty"`
}

// NewInternalUiListNamespacesResponseWithDefaults instantiates a new InternalUiListNamespacesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalUiListNamespacesResponseWithDefaults() *InternalUiListNamespacesResponse {
	var this InternalUiListNamespacesResponse

	return &this
}
