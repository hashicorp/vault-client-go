// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EntityLookUpRequest struct for EntityLookUpRequest
type EntityLookUpRequest struct {
	// ID of the alias.
	AliasId string `json:"alias_id"`

	// Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with 'alias_name'.
	AliasMountAccessor string `json:"alias_mount_accessor"`

	// Name of the alias. This should be supplied in conjunction with 'alias_mount_accessor'.
	AliasName string `json:"alias_name"`

	// ID of the entity.
	Id string `json:"id"`

	// Name of the entity.
	Name string `json:"name"`
}

// NewEntityLookUpRequestWithDefaults instantiates a new EntityLookUpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityLookUpRequestWithDefaults() *EntityLookUpRequest {
	var this EntityLookUpRequest

	return &this
}

func (o EntityLookUpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["alias_id"] = o.AliasId
	toSerialize["alias_mount_accessor"] = o.AliasMountAccessor
	toSerialize["alias_name"] = o.AliasName
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	return json.Marshal(toSerialize)
}
