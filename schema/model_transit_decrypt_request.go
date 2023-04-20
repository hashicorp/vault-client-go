// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitDecryptRequest struct for TransitDecryptRequest
type TransitDecryptRequest struct {
	// When using an AEAD cipher mode, such as AES-GCM, this parameter allows passing associated data (AD/AAD) into the encryption function; this data must be passed on subsequent decryption requests but can be transited in plaintext. On successful decryption, both the ciphertext and the associated data are attested not to have been tampered with.
	AssociatedData string `json:"associated_data"`

	// Specifies a list of items to be decrypted in a single batch. When this parameter is set, if the parameters 'ciphertext', 'context' and 'nonce' are also set, they will be ignored. Any batch output will preserve the order of the batch input.
	BatchInput []map[string]interface{} `json:"batch_input"`

	// The ciphertext to decrypt, provided as returned by encrypt.
	Ciphertext string `json:"ciphertext"`

	// Base64 encoded context for key derivation. Required if key derivation is enabled.
	Context string `json:"context"`

	// Base64 encoded nonce value used during encryption. Must be provided if convergent encryption is enabled for this key and the key was generated with Vault 0.6.1. Not required for keys created in 0.6.2+.
	Nonce string `json:"nonce"`

	// Ordinarily, if a batch item fails to decrypt due to a bad input, but other batch items succeed, the HTTP response code is 400 (Bad Request). Some applications may want to treat partial failures differently. Providing the parameter returns the given response code integer instead of a 400 in this case. If all values fail HTTP 400 is still returned.
	PartialFailureResponseCode int32 `json:"partial_failure_response_code"`
}

// NewTransitDecryptRequestWithDefaults instantiates a new TransitDecryptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitDecryptRequestWithDefaults() *TransitDecryptRequest {
	var this TransitDecryptRequest

	return &this
}

func (o TransitDecryptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["associated_data"] = o.AssociatedData
	toSerialize["batch_input"] = o.BatchInput
	toSerialize["ciphertext"] = o.Ciphertext
	toSerialize["context"] = o.Context
	toSerialize["nonce"] = o.Nonce
	toSerialize["partial_failure_response_code"] = o.PartialFailureResponseCode

	return json.Marshal(toSerialize)
}
