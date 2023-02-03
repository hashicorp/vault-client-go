// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EntityWriteAliasByIDRequest struct for EntityWriteAliasByIDRequest
type EntityWriteAliasByIDRequest struct {
	// Entity ID to which this alias should be tied to
	CanonicalId string `json:"canonical_id"`

	// User provided key-value pairs
	CustomMetadata map[string]interface{} `json:"custom_metadata"`

	// Entity ID to which this alias belongs to. This field is deprecated, use canonical_id.
	EntityId string `json:"entity_id"`

	// (Unused)
	MountAccessor string `json:"mount_accessor"`

	// (Unused)
	Name string `json:"name"`
}

// NewEntityWriteAliasByIDRequestWithDefaults instantiates a new EntityWriteAliasByIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWriteAliasByIDRequestWithDefaults() *EntityWriteAliasByIDRequest {
	var this EntityWriteAliasByIDRequest

	return &this
}
