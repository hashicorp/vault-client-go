// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteGenerateRootRequest struct for WriteGenerateRootRequest
type WriteGenerateRootRequest struct {
	// Specifies a base64-encoded PGP public key.

	PgpKey string `json:"pgp_key"`
}

// NewWriteGenerateRootRequestWithDefaults instantiates a new WriteGenerateRootRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteGenerateRootRequestWithDefaults() *WriteGenerateRootRequest {
	var this WriteGenerateRootRequest

	return &this
}

func (o WriteGenerateRootRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["pgp_key"] = o.PgpKey

	return json.Marshal(toSerialize)
}
