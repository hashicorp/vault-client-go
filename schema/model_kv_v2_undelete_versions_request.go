// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KvV2UndeleteVersionsRequest struct for KvV2UndeleteVersionsRequest
type KvV2UndeleteVersionsRequest struct {
	// The versions to unarchive. The versions will be restored and their data will be returned on normal get requests.
	Versions []int32 `json:"versions,omitempty"`
}
