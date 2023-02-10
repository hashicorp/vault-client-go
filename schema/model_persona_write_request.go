// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PersonaWriteRequest struct for PersonaWriteRequest
type PersonaWriteRequest struct {
	// Entity ID to which this persona belongs to
	EntityId string `json:"entity_id"`

	// ID of the persona
	Id string `json:"id"`

	// Metadata to be associated with the persona. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata"`

	// Mount accessor to which this persona belongs to
	MountAccessor string `json:"mount_accessor"`

	// Name of the persona
	Name string `json:"name"`
}

// NewPersonaWriteRequestWithDefaults instantiates a new PersonaWriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonaWriteRequestWithDefaults() *PersonaWriteRequest {
	var this PersonaWriteRequest

	return &this
}

func (o PersonaWriteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["entity_id"] = o.EntityId
	toSerialize["id"] = o.Id
	toSerialize["metadata"] = o.Metadata
	toSerialize["mount_accessor"] = o.MountAccessor
	toSerialize["name"] = o.Name

	return json.Marshal(toSerialize)
}
