// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV2WriteMetadataRequest struct for KvV2WriteMetadataRequest
type KvV2WriteMetadataRequest struct {
	// If true the key will require the cas parameter to be set on all write requests. If false, the backend’s configuration will be used.
	CasRequired bool `json:"cas_required,omitempty"`

	// User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret.
	CustomMetadata map[string]interface{} `json:"custom_metadata,omitempty"`

	// The length of time before a version is deleted. If not set, the backend's configured delete_version_after is used. Cannot be greater than the backend's delete_version_after. A zero duration clears the current setting. A negative duration will cause an error.
	DeleteVersionAfter string `json:"delete_version_after,omitempty"`

	// The number of versions to keep. If not set, the backend’s configured max version is used.
	MaxVersions int32 `json:"max_versions,omitempty"`
}
