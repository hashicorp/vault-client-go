// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KVv2WriteConfigRequest struct for KVv2WriteConfigRequest
type KVv2WriteConfigRequest struct {
	// If true, the backend will require the cas parameter to be set for each write
	CasRequired bool `json:"cas_required"`

	// If set, the length of time before a version is deleted. A negative duration disables the use of delete_version_after on all keys. A zero duration clears the current setting. Accepts a Go duration format string.
	DeleteVersionAfter int32 `json:"delete_version_after"`

	// The number of versions to keep for each key. Defaults to 10
	MaxVersions int32 `json:"max_versions"`
}

// NewKVv2WriteConfigRequestWithDefaults instantiates a new KVv2WriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKVv2WriteConfigRequestWithDefaults() *KVv2WriteConfigRequest {
	var this KVv2WriteConfigRequest

	return &this
}

func (o KVv2WriteConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
