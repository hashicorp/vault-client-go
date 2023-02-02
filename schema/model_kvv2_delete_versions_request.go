/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KVv2DeleteVersionsRequest struct for KVv2DeleteVersionsRequest
type KVv2DeleteVersionsRequest struct {
	// The versions to be archived. The versioned data will not be deleted, but it will no longer be returned in normal get requests.
	Versions []int32 `json:"versions"`
}

// NewKVv2DeleteVersionsRequestWithDefaults instantiates a new KVv2DeleteVersionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKVv2DeleteVersionsRequestWithDefaults() *KVv2DeleteVersionsRequest {
	var this KVv2DeleteVersionsRequest

	return &this
}

func (o KVv2DeleteVersionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["versions"] = o.Versions

	return json.Marshal(toSerialize)
}