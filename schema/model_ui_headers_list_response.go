// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// UiHeadersListResponse struct for UiHeadersListResponse
type UiHeadersListResponse struct {
	// Lists of configured UI headers. Omitted if list is empty
	Keys []string `json:"keys"`
}

// NewUiHeadersListResponseWithDefaults instantiates a new UiHeadersListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiHeadersListResponseWithDefaults() *UiHeadersListResponse {
	var this UiHeadersListResponse

	return &this
}

func (o UiHeadersListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
