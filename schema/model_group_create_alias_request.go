// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GroupCreateAliasRequest struct for GroupCreateAliasRequest
type GroupCreateAliasRequest struct {
	// ID of the group to which this is an alias.
	CanonicalId string `json:"canonical_id"`

	// ID of the group alias.
	Id string `json:"id"`

	// Mount accessor to which this alias belongs to.
	MountAccessor string `json:"mount_accessor"`

	// Alias of the group.
	Name string `json:"name"`
}

// NewGroupCreateAliasRequestWithDefaults instantiates a new GroupCreateAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCreateAliasRequestWithDefaults() *GroupCreateAliasRequest {
	var this GroupCreateAliasRequest

	return &this
}

func (o GroupCreateAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["canonical_id"] = o.CanonicalId
	toSerialize["id"] = o.Id
	toSerialize["mount_accessor"] = o.MountAccessor
	toSerialize["name"] = o.Name

	return json.Marshal(toSerialize)
}
