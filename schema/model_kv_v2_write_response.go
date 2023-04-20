// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// KvV2WriteResponse struct for KvV2WriteResponse
type KvV2WriteResponse struct {
	CreatedTime time.Time `json:"created_time"`

	CustomMetadata map[string]interface{} `json:"custom_metadata"`

	DeletionTime string `json:"deletion_time"`

	Destroyed bool `json:"destroyed"`

	Version int64 `json:"version"`
}

// NewKvV2WriteResponseWithDefaults instantiates a new KvV2WriteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2WriteResponseWithDefaults() *KvV2WriteResponse {
	var this KvV2WriteResponse

	return &this
}

func (o KvV2WriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["created_time"] = o.CreatedTime
	toSerialize["custom_metadata"] = o.CustomMetadata
	toSerialize["deletion_time"] = o.DeletionTime
	toSerialize["destroyed"] = o.Destroyed
	toSerialize["version"] = o.Version

	return json.Marshal(toSerialize)
}
