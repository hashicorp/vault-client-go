// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// VersionHistoryResponse struct for VersionHistoryResponse
type VersionHistoryResponse struct {
	KeyInfo map[string]interface{} `json:"key_info,omitempty"`

	Keys []string `json:"keys,omitempty"`
}

// NewVersionHistoryResponseWithDefaults instantiates a new VersionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionHistoryResponseWithDefaults() *VersionHistoryResponse {
	var this VersionHistoryResponse

	return &this
}
