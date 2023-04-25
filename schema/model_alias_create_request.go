// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AliasCreateRequest struct for AliasCreateRequest
type AliasCreateRequest struct {
	// Entity ID to which this alias belongs to
	CanonicalId string `json:"canonical_id"`

	// Entity ID to which this alias belongs to. This field is deprecated in favor of 'canonical_id'.
	EntityId string `json:"entity_id"`

	// ID of the alias
	Id string `json:"id"`

	// Mount accessor to which this alias belongs to
	MountAccessor string `json:"mount_accessor"`

	// Name of the alias
	Name string `json:"name"`
}

// NewAliasCreateRequestWithDefaults instantiates a new AliasCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasCreateRequestWithDefaults() *AliasCreateRequest {
	var this AliasCreateRequest

	return &this
}

func (o AliasCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["canonical_id"] = o.CanonicalId
	toSerialize["entity_id"] = o.EntityId
	toSerialize["id"] = o.Id
	toSerialize["mount_accessor"] = o.MountAccessor
	toSerialize["name"] = o.Name

	return json.Marshal(toSerialize)
}
