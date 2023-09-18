// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV2ReadConfigurationResponse struct for KvV2ReadConfigurationResponse
type KvV2ReadConfigurationResponse struct {
	// If true, the backend will require the cas parameter to be set for each write
	CasRequired bool `json:"cas_required,omitempty"`

	// The length of time before a version is deleted.
	DeleteVersionAfter string `json:"delete_version_after,omitempty"`

	// The number of versions to keep for each key.
	MaxVersions int32 `json:"max_versions,omitempty"`
}
