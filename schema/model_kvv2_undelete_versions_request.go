// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KVv2UndeleteVersionsRequest struct for KVv2UndeleteVersionsRequest
type KVv2UndeleteVersionsRequest struct {
	// The versions to unarchive. The versions will be restored and their data will be returned on normal get requests.
	Versions []int32 `json:"versions"`
}

// NewKVv2UndeleteVersionsRequestWithDefaults instantiates a new KVv2UndeleteVersionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKVv2UndeleteVersionsRequestWithDefaults() *KVv2UndeleteVersionsRequest {
	var this KVv2UndeleteVersionsRequest

	return &this
}
