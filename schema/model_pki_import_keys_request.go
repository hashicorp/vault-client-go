// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIImportKeysRequest struct for PKIImportKeysRequest
type PKIImportKeysRequest struct {
	// Optional name to be used for this key
	KeyName string `json:"key_name"`

	// PEM-format, unencrypted secret key
	PemBundle string `json:"pem_bundle"`
}

// NewPKIImportKeysRequestWithDefaults instantiates a new PKIImportKeysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIImportKeysRequestWithDefaults() *PKIImportKeysRequest {
	var this PKIImportKeysRequest

	return &this
}

func (o PKIImportKeysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
