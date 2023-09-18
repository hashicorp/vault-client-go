// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AliasCreateRequest struct for AliasCreateRequest
type AliasCreateRequest struct {
	// Entity ID to which this alias belongs to
	CanonicalId string `json:"canonical_id,omitempty"`

	// Entity ID to which this alias belongs to. This field is deprecated in favor of 'canonical_id'.
	EntityId string `json:"entity_id,omitempty"`

	// ID of the alias
	Id string `json:"id,omitempty"`

	// Mount accessor to which this alias belongs to
	MountAccessor string `json:"mount_accessor,omitempty"`

	// Name of the alias
	Name string `json:"name,omitempty"`
}
