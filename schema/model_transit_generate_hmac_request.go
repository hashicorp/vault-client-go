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

// TransitGenerateHMACRequest struct for TransitGenerateHMACRequest
type TransitGenerateHMACRequest struct {
	// Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \"sha2-256\".

	Algorithm string `json:"algorithm"`

	// The base64-encoded input data

	Input string `json:"input"`

	// The version of the key to use for generating the HMAC. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.

	KeyVersion int32 `json:"key_version"`

	// Algorithm to use (POST URL parameter)

	Urlalgorithm string `json:"urlalgorithm"`
}

// NewTransitGenerateHMACRequestWithDefaults instantiates a new TransitGenerateHMACRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitGenerateHMACRequestWithDefaults() *TransitGenerateHMACRequest {
	var this TransitGenerateHMACRequest

	this.Algorithm = "sha2-256"

	return &this
}

func (o TransitGenerateHMACRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["algorithm"] = o.Algorithm

	toSerialize["input"] = o.Input

	toSerialize["key_version"] = o.KeyVersion

	toSerialize["urlalgorithm"] = o.Urlalgorithm

	return json.Marshal(toSerialize)
}
