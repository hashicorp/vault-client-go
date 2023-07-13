// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV2ListMetadataResponse struct for KvV2ListMetadataResponse
type KvV2ListMetadataResponse struct {
	Keys []string `json:"keys,omitempty"`
}

// NewKvV2ListMetadataResponseWithDefaults instantiates a new KvV2ListMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2ListMetadataResponseWithDefaults() *KvV2ListMetadataResponse {
	var this KvV2ListMetadataResponse

	return &this
}
