// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV2ConfigureRequest struct for KvV2ConfigureRequest
type KvV2ConfigureRequest struct {
	// If true, the backend will require the cas parameter to be set for each write
	CasRequired bool `json:"cas_required,omitempty"`

	// If set, the length of time before a version is deleted. A negative duration disables the use of delete_version_after on all keys. A zero duration clears the current setting. Accepts a Go duration format string.
	DeleteVersionAfter string `json:"delete_version_after,omitempty"`

	// The number of versions to keep for each key. Defaults to 10
	MaxVersions int32 `json:"max_versions,omitempty"`
}
