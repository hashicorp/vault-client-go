// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PersonaCreateRequest struct for PersonaCreateRequest
type PersonaCreateRequest struct {
	// Entity ID to which this persona belongs to
	EntityId string `json:"entity_id,omitempty"`

	// ID of the persona
	Id string `json:"id,omitempty"`

	// Metadata to be associated with the persona. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Mount accessor to which this persona belongs to
	MountAccessor string `json:"mount_accessor,omitempty"`

	// Name of the persona
	Name string `json:"name,omitempty"`
}
