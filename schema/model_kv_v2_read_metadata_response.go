// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// KvV2ReadMetadataResponse struct for KvV2ReadMetadataResponse
type KvV2ReadMetadataResponse struct {
	CasRequired bool `json:"cas_required"`

	CreatedTime time.Time `json:"created_time"`

	CurrentVersion int64 `json:"current_version"`

	// User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret.
	CustomMetadata map[string]interface{} `json:"custom_metadata"`

	// The length of time before a version is deleted.
	DeleteVersionAfter int32 `json:"delete_version_after"`

	// The number of versions to keep
	MaxVersions int64 `json:"max_versions"`

	OldestVersion int64 `json:"oldest_version"`

	UpdatedTime time.Time `json:"updated_time"`

	Versions map[string]interface{} `json:"versions"`
}

// NewKvV2ReadMetadataResponseWithDefaults instantiates a new KvV2ReadMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2ReadMetadataResponseWithDefaults() *KvV2ReadMetadataResponse {
	var this KvV2ReadMetadataResponse

	return &this
}

func (o KvV2ReadMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cas_required"] = o.CasRequired
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["current_version"] = o.CurrentVersion
	toSerialize["custom_metadata"] = o.CustomMetadata
	toSerialize["delete_version_after"] = o.DeleteVersionAfter
	toSerialize["max_versions"] = o.MaxVersions
	toSerialize["oldest_version"] = o.OldestVersion
	toSerialize["updated_time"] = o.UpdatedTime
	toSerialize["versions"] = o.Versions

	return json.Marshal(toSerialize)
}
