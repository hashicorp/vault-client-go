// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EntityUpdateAliasByIdRequest struct for EntityUpdateAliasByIdRequest
type EntityUpdateAliasByIdRequest struct {
	// Entity ID to which this alias should be tied to
	CanonicalId string `json:"canonical_id"`

	// User provided key-value pairs
	CustomMetadata map[string]interface{} `json:"custom_metadata"`

	// Entity ID to which this alias belongs to. This field is deprecated, use canonical_id.
	EntityId string `json:"entity_id"`

	// (Unused)
	MountAccessor string `json:"mount_accessor"`

	// (Unused)
	Name string `json:"name"`
}

// NewEntityUpdateAliasByIdRequestWithDefaults instantiates a new EntityUpdateAliasByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityUpdateAliasByIdRequestWithDefaults() *EntityUpdateAliasByIdRequest {
	var this EntityUpdateAliasByIdRequest

	return &this
}

func (o EntityUpdateAliasByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["canonical_id"] = o.CanonicalId
	toSerialize["custom_metadata"] = o.CustomMetadata
	toSerialize["entity_id"] = o.EntityId
	toSerialize["mount_accessor"] = o.MountAccessor
	toSerialize["name"] = o.Name

	return json.Marshal(toSerialize)
}
