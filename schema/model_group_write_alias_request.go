// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GroupWriteAliasRequest struct for GroupWriteAliasRequest
type GroupWriteAliasRequest struct {
	// ID of the group to which this is an alias.
	CanonicalId string `json:"canonical_id"`

	// ID of the group alias.
	Id string `json:"id"`

	// Mount accessor to which this alias belongs to.
	MountAccessor string `json:"mount_accessor"`

	// Alias of the group.
	Name string `json:"name"`
}

// NewGroupWriteAliasRequestWithDefaults instantiates a new GroupWriteAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWriteAliasRequestWithDefaults() *GroupWriteAliasRequest {
	var this GroupWriteAliasRequest

	return &this
}

func (o GroupWriteAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["canonical_id"] = o.CanonicalId
	toSerialize["id"] = o.Id
	toSerialize["mount_accessor"] = o.MountAccessor
	toSerialize["name"] = o.Name

	return json.Marshal(toSerialize)
}
