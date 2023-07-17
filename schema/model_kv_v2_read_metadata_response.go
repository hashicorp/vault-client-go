// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// KvV2ReadMetadataResponse struct for KvV2ReadMetadataResponse
type KvV2ReadMetadataResponse struct {
	CasRequired bool `json:"cas_required,omitempty"`

	CreatedTime time.Time `json:"created_time,omitempty"`

	CurrentVersion int64 `json:"current_version,omitempty"`

	// User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret.
	CustomMetadata map[string]interface{} `json:"custom_metadata,omitempty"`

	// The length of time before a version is deleted.
	DeleteVersionAfter string `json:"delete_version_after,omitempty"`

	// The number of versions to keep
	MaxVersions int64 `json:"max_versions,omitempty"`

	OldestVersion int64 `json:"oldest_version,omitempty"`

	UpdatedTime time.Time `json:"updated_time,omitempty"`

	Versions map[string]interface{} `json:"versions,omitempty"`
}
