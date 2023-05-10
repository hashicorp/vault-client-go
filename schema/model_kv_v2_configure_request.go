// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KvV2ConfigureRequest struct for KvV2ConfigureRequest
type KvV2ConfigureRequest struct {
	// If true, the backend will require the cas parameter to be set for each write
	CasRequired bool `json:"cas_required"`

	// If set, the length of time before a version is deleted. A negative duration disables the use of delete_version_after on all keys. A zero duration clears the current setting. Accepts a Go duration format string.
	DeleteVersionAfter int32 `json:"delete_version_after"`

	// The number of versions to keep for each key. Defaults to 10
	MaxVersions int32 `json:"max_versions"`
}

// NewKvV2ConfigureRequestWithDefaults instantiates a new KvV2ConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2ConfigureRequestWithDefaults() *KvV2ConfigureRequest {
	var this KvV2ConfigureRequest

	return &this
}

func (o KvV2ConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cas_required"] = o.CasRequired
	toSerialize["delete_version_after"] = o.DeleteVersionAfter
	toSerialize["max_versions"] = o.MaxVersions

	return json.Marshal(toSerialize)
}
