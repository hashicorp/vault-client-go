// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GroupUpdateByIdRequest struct for GroupUpdateByIdRequest
type GroupUpdateByIdRequest struct {
	// Entity IDs to be assigned as group members.
	MemberEntityIds []string `json:"member_entity_ids,omitempty"`

	// Group IDs to be assigned as group members.
	MemberGroupIds []string `json:"member_group_ids,omitempty"`

	// Metadata to be associated with the group. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Name of the group.
	Name string `json:"name,omitempty"`

	// Policies to be tied to the group.
	Policies []string `json:"policies,omitempty"`

	// Type of the group, 'internal' or 'external'. Defaults to 'internal'
	Type string `json:"type,omitempty"`
}
