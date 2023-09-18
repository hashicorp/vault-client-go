// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV2DeleteVersionsRequest struct for KvV2DeleteVersionsRequest
type KvV2DeleteVersionsRequest struct {
	// The versions to be archived. The versioned data will not be deleted, but it will no longer be returned in normal get requests.
	Versions []int32 `json:"versions,omitempty"`
}
