// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalUiReadResultantAclResponse struct for InternalUiReadResultantAclResponse
type InternalUiReadResultantAclResponse struct {
	ExactPaths map[string]interface{} `json:"exact_paths,omitempty"`

	GlobPaths map[string]interface{} `json:"glob_paths,omitempty"`

	Root bool `json:"root,omitempty"`
}
