// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// EntityCreateAliasRequest struct for EntityCreateAliasRequest
type EntityCreateAliasRequest struct {
	// Entity ID to which this alias belongs
	CanonicalId string `json:"canonical_id,omitempty"`

	// User provided key-value pairs
	CustomMetadata map[string]interface{} `json:"custom_metadata,omitempty"`

	// Entity ID to which this alias belongs. This field is deprecated, use canonical_id.
	EntityId string `json:"entity_id,omitempty"`

	// ID of the entity alias. If set, updates the corresponding entity alias.
	Id string `json:"id,omitempty"`

	// Mount accessor to which this alias belongs to; unused for a modify
	MountAccessor string `json:"mount_accessor,omitempty"`

	// Name of the alias; unused for a modify
	Name string `json:"name,omitempty"`
}
