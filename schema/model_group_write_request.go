// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GroupWriteRequest struct for GroupWriteRequest
type GroupWriteRequest struct {
	// ID of the group. If set, updates the corresponding existing group.
	Id string `json:"id"`

	// Entity IDs to be assigned as group members.
	MemberEntityIds []string `json:"member_entity_ids"`

	// Group IDs to be assigned as group members.
	MemberGroupIds []string `json:"member_group_ids"`

	// Metadata to be associated with the group. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata"`

	// Name of the group.
	Name string `json:"name"`

	// Policies to be tied to the group.
	Policies []string `json:"policies"`

	// Type of the group, 'internal' or 'external'. Defaults to 'internal'
	Type string `json:"type"`
}

// NewGroupWriteRequestWithDefaults instantiates a new GroupWriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWriteRequestWithDefaults() *GroupWriteRequest {
	var this GroupWriteRequest

	return &this
}

func (o GroupWriteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["id"] = o.Id
	toSerialize["member_entity_ids"] = o.MemberEntityIds
	toSerialize["member_group_ids"] = o.MemberGroupIds
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["policies"] = o.Policies
	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
