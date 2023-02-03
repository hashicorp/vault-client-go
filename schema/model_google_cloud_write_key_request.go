// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudWriteKeyRequest struct for GoogleCloudWriteKeyRequest
type GoogleCloudWriteKeyRequest struct {
	// Private key algorithm for service account key - defaults to KEY_ALG_RSA_2048\"

	KeyAlgorithm string `json:"key_algorithm"`

	// Private key type for service account key - defaults to TYPE_GOOGLE_CREDENTIALS_FILE\"

	KeyType string `json:"key_type"`

	// Lifetime of the service account key

	Ttl int32 `json:"ttl"`
}

// NewGoogleCloudWriteKeyRequestWithDefaults instantiates a new GoogleCloudWriteKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteKeyRequestWithDefaults() *GoogleCloudWriteKeyRequest {
	var this GoogleCloudWriteKeyRequest

	this.KeyAlgorithm = "KEY_ALG_RSA_2048"

	this.KeyType = "TYPE_GOOGLE_CREDENTIALS_FILE"

	return &this
}

func (o GoogleCloudWriteKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key_algorithm"] = o.KeyAlgorithm

	toSerialize["key_type"] = o.KeyType

	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
