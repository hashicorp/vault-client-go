// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AliasUpdateByIdRequest struct for AliasUpdateByIdRequest
type AliasUpdateByIdRequest struct {
	// Entity ID to which this alias should be tied to
	CanonicalId string `json:"canonical_id,omitempty"`

	// Entity ID to which this alias should be tied to. This field is deprecated in favor of 'canonical_id'.
	EntityId string `json:"entity_id,omitempty"`

	// Mount accessor to which this alias belongs to
	MountAccessor string `json:"mount_accessor,omitempty"`

	// Name of the alias
	Name string `json:"name,omitempty"`
}
