// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// EntityLookUpRequest struct for EntityLookUpRequest
type EntityLookUpRequest struct {
	// ID of the alias.
	AliasId string `json:"alias_id,omitempty"`

	// Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with 'alias_name'.
	AliasMountAccessor string `json:"alias_mount_accessor,omitempty"`

	// Name of the alias. This should be supplied in conjunction with 'alias_mount_accessor'.
	AliasName string `json:"alias_name,omitempty"`

	// ID of the entity.
	Id string `json:"id,omitempty"`

	// Name of the entity.
	Name string `json:"name,omitempty"`
}
