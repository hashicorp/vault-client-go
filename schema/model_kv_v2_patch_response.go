// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// KvV2PatchResponse struct for KvV2PatchResponse
type KvV2PatchResponse struct {
	CreatedTime time.Time `json:"created_time,omitempty"`

	CustomMetadata map[string]interface{} `json:"custom_metadata,omitempty"`

	DeletionTime string `json:"deletion_time,omitempty"`

	Destroyed bool `json:"destroyed,omitempty"`

	Version int64 `json:"version,omitempty"`
}

// NewKvV2PatchResponseWithDefaults instantiates a new KvV2PatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2PatchResponseWithDefaults() *KvV2PatchResponse {
	var this KvV2PatchResponse

	return &this
}
