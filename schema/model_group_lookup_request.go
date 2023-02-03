// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GroupLookupRequest struct for GroupLookupRequest
type GroupLookupRequest struct {
	// ID of the alias.
	AliasId string `json:"alias_id"`

	// Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with 'alias_name'.
	AliasMountAccessor string `json:"alias_mount_accessor"`

	// Name of the alias. This should be supplied in conjunction with 'alias_mount_accessor'.
	AliasName string `json:"alias_name"`

	// ID of the group.
	Id string `json:"id"`

	// Name of the group.
	Name string `json:"name"`
}

// NewGroupLookupRequestWithDefaults instantiates a new GroupLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupLookupRequestWithDefaults() *GroupLookupRequest {
	var this GroupLookupRequest

	return &this
}
