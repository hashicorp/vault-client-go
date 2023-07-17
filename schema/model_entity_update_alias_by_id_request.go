// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// EntityUpdateAliasByIdRequest struct for EntityUpdateAliasByIdRequest
type EntityUpdateAliasByIdRequest struct {
	// Entity ID to which this alias should be tied to
	CanonicalId string `json:"canonical_id,omitempty"`

	// User provided key-value pairs
	CustomMetadata map[string]interface{} `json:"custom_metadata,omitempty"`

	// Entity ID to which this alias belongs to. This field is deprecated, use canonical_id.
	EntityId string `json:"entity_id,omitempty"`

	// (Unused)
	MountAccessor string `json:"mount_accessor,omitempty"`

	// (Unused)
	Name string `json:"name,omitempty"`
}
