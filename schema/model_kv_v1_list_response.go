// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV1ListResponse struct for KvV1ListResponse
type KvV1ListResponse struct {
	Keys []string `json:"keys,omitempty"`
}

// NewKvV1ListResponseWithDefaults instantiates a new KvV1ListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV1ListResponseWithDefaults() *KvV1ListResponse {
	var this KvV1ListResponse

	return &this
}
