// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EntityWriteRequest struct for EntityWriteRequest
type EntityWriteRequest struct {
	// If set true, tokens tied to this identity will not be able to be used (but will not be revoked).
	Disabled bool `json:"disabled"`

	// ID of the entity. If set, updates the corresponding existing entity.
	Id string `json:"id"`

	// Metadata to be associated with the entity. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata"`

	// Name of the entity
	Name string `json:"name"`

	// Policies to be tied to the entity.
	Policies []string `json:"policies"`
}

// NewEntityWriteRequestWithDefaults instantiates a new EntityWriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWriteRequestWithDefaults() *EntityWriteRequest {
	var this EntityWriteRequest

	return &this
}

func (o EntityWriteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
