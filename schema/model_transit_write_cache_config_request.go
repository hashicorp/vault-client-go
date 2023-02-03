// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitWriteCacheConfigRequest struct for TransitWriteCacheConfigRequest
type TransitWriteCacheConfigRequest struct {
	// Size of cache, use 0 for an unlimited cache size, defaults to 0

	Size int32 `json:"size"`
}

// NewTransitWriteCacheConfigRequestWithDefaults instantiates a new TransitWriteCacheConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitWriteCacheConfigRequestWithDefaults() *TransitWriteCacheConfigRequest {
	var this TransitWriteCacheConfigRequest

	this.Size = 0

	return &this
}

func (o TransitWriteCacheConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["size"] = o.Size

	return json.Marshal(toSerialize)
}
