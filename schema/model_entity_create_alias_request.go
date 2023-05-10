// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EntityCreateAliasRequest struct for EntityCreateAliasRequest
type EntityCreateAliasRequest struct {
	// Entity ID to which this alias belongs
	CanonicalId string `json:"canonical_id"`

	// User provided key-value pairs
	CustomMetadata map[string]interface{} `json:"custom_metadata"`

	// Entity ID to which this alias belongs. This field is deprecated, use canonical_id.
	EntityId string `json:"entity_id"`

	// ID of the entity alias. If set, updates the corresponding entity alias.
	Id string `json:"id"`

	// Mount accessor to which this alias belongs to; unused for a modify
	MountAccessor string `json:"mount_accessor"`

	// Name of the alias; unused for a modify
	Name string `json:"name"`
}

// NewEntityCreateAliasRequestWithDefaults instantiates a new EntityCreateAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityCreateAliasRequestWithDefaults() *EntityCreateAliasRequest {
	var this EntityCreateAliasRequest

	return &this
}

func (o EntityCreateAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["canonical_id"] = o.CanonicalId
	toSerialize["custom_metadata"] = o.CustomMetadata
	toSerialize["entity_id"] = o.EntityId
	toSerialize["id"] = o.Id
	toSerialize["mount_accessor"] = o.MountAccessor
	toSerialize["name"] = o.Name

	return json.Marshal(toSerialize)
}
