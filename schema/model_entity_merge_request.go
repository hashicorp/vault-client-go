// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// EntityMergeRequest struct for EntityMergeRequest
type EntityMergeRequest struct {
	// Alias IDs to keep in case of conflicting aliases. Ignored if no conflicting aliases found
	ConflictingAliasIdsToKeep []string `json:"conflicting_alias_ids_to_keep,omitempty"`

	// Setting this will follow the 'mine' strategy for merging MFA secrets. If there are secrets of the same type both in entities that are merged from and in entity into which all others are getting merged, secrets in the destination will be unaltered. If not set, this API will throw an error containing all the conflicts.
	Force bool `json:"force,omitempty"`

	// Entity IDs which need to get merged
	FromEntityIds []string `json:"from_entity_ids,omitempty"`

	// Entity ID into which all the other entities need to get merged
	ToEntityId string `json:"to_entity_id,omitempty"`
}
