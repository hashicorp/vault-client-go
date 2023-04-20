// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitConfigureCacheRequest struct for TransitConfigureCacheRequest
type TransitConfigureCacheRequest struct {
	// Size of cache, use 0 for an unlimited cache size, defaults to 0
	Size int32 `json:"size"`
}

// NewTransitConfigureCacheRequestWithDefaults instantiates a new TransitConfigureCacheRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitConfigureCacheRequestWithDefaults() *TransitConfigureCacheRequest {
	var this TransitConfigureCacheRequest

	this.Size = 0

	return &this
}

func (o TransitConfigureCacheRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["size"] = o.Size

	return json.Marshal(toSerialize)
}
