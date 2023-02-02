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

// TransitGenerateDataKeyRequest struct for TransitGenerateDataKeyRequest
type TransitGenerateDataKeyRequest struct {
	// Number of bits for the key; currently 128, 256, and 512 bits are supported. Defaults to 256.

	Bits int32 `json:"bits"`

	// Context for key derivation. Required for derived keys.

	Context string `json:"context"`

	// The version of the Vault key to use for encryption of the data key. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.

	KeyVersion int32 `json:"key_version"`

	// Nonce for when convergent encryption v1 is used (only in Vault 0.6.1)

	Nonce string `json:"nonce"`
}

// NewTransitGenerateDataKeyRequestWithDefaults instantiates a new TransitGenerateDataKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitGenerateDataKeyRequestWithDefaults() *TransitGenerateDataKeyRequest {
	var this TransitGenerateDataKeyRequest

	this.Bits = 256

	return &this
}

func (o TransitGenerateDataKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bits"] = o.Bits

	toSerialize["context"] = o.Context

	toSerialize["key_version"] = o.KeyVersion

	toSerialize["nonce"] = o.Nonce

	return json.Marshal(toSerialize)
}
