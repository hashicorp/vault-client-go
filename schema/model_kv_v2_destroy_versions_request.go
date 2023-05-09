// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV2DestroyVersionsRequest struct for KvV2DestroyVersionsRequest
type KvV2DestroyVersionsRequest struct {
	// The versions to destroy. Their data will be permanently deleted.
	Versions []int32 `json:"versions,omitempty"`
}

// NewKvV2DestroyVersionsRequestWithDefaults instantiates a new KvV2DestroyVersionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2DestroyVersionsRequestWithDefaults() *KvV2DestroyVersionsRequest {
	var this KvV2DestroyVersionsRequest

	return &this
}
