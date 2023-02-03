// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KVv2WriteRequest struct for KVv2WriteRequest
type KVv2WriteRequest struct {
	// The contents of the data map will be stored and returned on read.
	Data map[string]interface{} `json:"data"`

	// Options for writing a KV entry. Set the \"cas\" value to use a Check-And-Set operation. If not set the write will be allowed. If set to 0 a write will only be allowed if the key doesn’t exist. If the index is non-zero the write will only be allowed if the key’s current version matches the version specified in the cas parameter.
	Options map[string]interface{} `json:"options"`

	// If provided during a read, the value at the version number will be returned
	Version int32 `json:"version"`
}

// NewKVv2WriteRequestWithDefaults instantiates a new KVv2WriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKVv2WriteRequestWithDefaults() *KVv2WriteRequest {
	var this KVv2WriteRequest

	return &this
}

func (o KVv2WriteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
