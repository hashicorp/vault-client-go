// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV2ListResponse struct for KvV2ListResponse
type KvV2ListResponse struct {
	Keys []string `json:"keys,omitempty"`
}

// NewKvV2ListResponseWithDefaults instantiates a new KvV2ListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2ListResponseWithDefaults() *KvV2ListResponse {
	var this KvV2ListResponse

	return &this
}
