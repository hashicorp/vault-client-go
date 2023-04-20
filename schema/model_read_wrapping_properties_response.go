// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// ReadWrappingPropertiesResponse struct for ReadWrappingPropertiesResponse
type ReadWrappingPropertiesResponse struct {
	CreationPath string `json:"creation_path"`

	CreationTime time.Time `json:"creation_time"`

	CreationTtl int32 `json:"creation_ttl"`
}

// NewReadWrappingPropertiesResponseWithDefaults instantiates a new ReadWrappingPropertiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadWrappingPropertiesResponseWithDefaults() *ReadWrappingPropertiesResponse {
	var this ReadWrappingPropertiesResponse

	return &this
}

func (o ReadWrappingPropertiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["creation_path"] = o.CreationPath
	toSerialize["creation_time"] = o.CreationTime
	toSerialize["creation_ttl"] = o.CreationTtl

	return json.Marshal(toSerialize)
}
