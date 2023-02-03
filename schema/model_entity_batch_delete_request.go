// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EntityBatchDeleteRequest struct for EntityBatchDeleteRequest
type EntityBatchDeleteRequest struct {
	// Entity IDs to delete
	EntityIds []string `json:"entity_ids"`
}

// NewEntityBatchDeleteRequestWithDefaults instantiates a new EntityBatchDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityBatchDeleteRequestWithDefaults() *EntityBatchDeleteRequest {
	var this EntityBatchDeleteRequest

	return &this
}

func (o EntityBatchDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
