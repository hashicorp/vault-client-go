// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GroupUpdateAliasByIdRequest struct for GroupUpdateAliasByIdRequest
type GroupUpdateAliasByIdRequest struct {
	// ID of the group to which this is an alias.
	CanonicalId string `json:"canonical_id,omitempty"`

	// Mount accessor to which this alias belongs to.
	MountAccessor string `json:"mount_accessor,omitempty"`

	// Alias of the group.
	Name string `json:"name,omitempty"`
}
