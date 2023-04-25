// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// InternalUiListNamespacesResponse struct for InternalUiListNamespacesResponse
type InternalUiListNamespacesResponse struct {
	// field is only returned if there are one or more namespaces
	Keys []string `json:"keys"`
}

// NewInternalUiListNamespacesResponseWithDefaults instantiates a new InternalUiListNamespacesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalUiListNamespacesResponseWithDefaults() *InternalUiListNamespacesResponse {
	var this InternalUiListNamespacesResponse

	return &this
}

func (o InternalUiListNamespacesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
