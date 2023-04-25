// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KvV2DeleteVersionsRequest struct for KvV2DeleteVersionsRequest
type KvV2DeleteVersionsRequest struct {
	// The versions to be archived. The versioned data will not be deleted, but it will no longer be returned in normal get requests.
	Versions []int32 `json:"versions"`
}

// NewKvV2DeleteVersionsRequestWithDefaults instantiates a new KvV2DeleteVersionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2DeleteVersionsRequestWithDefaults() *KvV2DeleteVersionsRequest {
	var this KvV2DeleteVersionsRequest

	return &this
}

func (o KvV2DeleteVersionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["versions"] = o.Versions

	return json.Marshal(toSerialize)
}
