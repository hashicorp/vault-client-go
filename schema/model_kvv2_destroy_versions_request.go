// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KVv2DestroyVersionsRequest struct for KVv2DestroyVersionsRequest
type KVv2DestroyVersionsRequest struct {
	// The versions to destroy. Their data will be permanently deleted.
	Versions []int32 `json:"versions"`
}

// NewKVv2DestroyVersionsRequestWithDefaults instantiates a new KVv2DestroyVersionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKVv2DestroyVersionsRequestWithDefaults() *KVv2DestroyVersionsRequest {
	var this KVv2DestroyVersionsRequest

	return &this
}

func (o KVv2DestroyVersionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
