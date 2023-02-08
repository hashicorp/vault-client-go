// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PersonaIDWriteByIDRequest struct for PersonaIDWriteByIDRequest
type PersonaIDWriteByIDRequest struct {
	// Entity ID to which this persona should be tied to
	EntityId string `json:"entity_id"`

	// Metadata to be associated with the persona. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata"`

	// Mount accessor to which this persona belongs to
	MountAccessor string `json:"mount_accessor"`

	// Name of the persona
	Name string `json:"name"`
}

// NewPersonaIDWriteByIDRequestWithDefaults instantiates a new PersonaIDWriteByIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonaIDWriteByIDRequestWithDefaults() *PersonaIDWriteByIDRequest {
	var this PersonaIDWriteByIDRequest

	return &this
}

func (o PersonaIDWriteByIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
