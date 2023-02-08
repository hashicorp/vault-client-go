// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GroupWriteAliasByIDRequest struct for GroupWriteAliasByIDRequest
type GroupWriteAliasByIDRequest struct {
	// ID of the group to which this is an alias.
	CanonicalId string `json:"canonical_id"`

	// Mount accessor to which this alias belongs to.
	MountAccessor string `json:"mount_accessor"`

	// Alias of the group.
	Name string `json:"name"`
}

// NewGroupWriteAliasByIDRequestWithDefaults instantiates a new GroupWriteAliasByIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWriteAliasByIDRequestWithDefaults() *GroupWriteAliasByIDRequest {
	var this GroupWriteAliasByIDRequest

	return &this
}

func (o GroupWriteAliasByIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
