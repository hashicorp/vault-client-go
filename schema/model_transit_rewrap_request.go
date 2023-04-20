// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitRewrapRequest struct for TransitRewrapRequest
type TransitRewrapRequest struct {
	// Specifies a list of items to be re-encrypted in a single batch. When this parameter is set, if the parameters 'ciphertext', 'context' and 'nonce' are also set, they will be ignored. Any batch output will preserve the order of the batch input.
	BatchInput []map[string]interface{} `json:"batch_input"`

	// Ciphertext value to rewrap
	Ciphertext string `json:"ciphertext"`

	// Base64 encoded context for key derivation. Required for derived keys.
	Context string `json:"context"`

	// The version of the key to use for encryption. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.
	KeyVersion int32 `json:"key_version"`

	// Nonce for when convergent encryption is used
	Nonce string `json:"nonce"`
}

// NewTransitRewrapRequestWithDefaults instantiates a new TransitRewrapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitRewrapRequestWithDefaults() *TransitRewrapRequest {
	var this TransitRewrapRequest

	return &this
}

func (o TransitRewrapRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["batch_input"] = o.BatchInput
	toSerialize["ciphertext"] = o.Ciphertext
	toSerialize["context"] = o.Context
	toSerialize["key_version"] = o.KeyVersion
	toSerialize["nonce"] = o.Nonce

	return json.Marshal(toSerialize)
}
