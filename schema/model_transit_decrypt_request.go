// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitDecryptRequest struct for TransitDecryptRequest
type TransitDecryptRequest struct {
	// When using an AEAD cipher mode, such as AES-GCM, this parameter allows passing associated data (AD/AAD) into the encryption function; this data must be passed on subsequent decryption requests but can be transited in plaintext. On successful decryption, both the ciphertext and the associated data are attested not to have been tampered with.
	AssociatedData string `json:"associated_data,omitempty"`

	// Specifies a list of items to be decrypted in a single batch. When this parameter is set, if the parameters 'ciphertext', 'context' and 'nonce' are also set, they will be ignored. Any batch output will preserve the order of the batch input.
	BatchInput []map[string]interface{} `json:"batch_input,omitempty"`

	// The ciphertext to decrypt, provided as returned by encrypt.
	Ciphertext string `json:"ciphertext,omitempty"`

	// Base64 encoded context for key derivation. Required if key derivation is enabled.
	Context string `json:"context,omitempty"`

	// Base64 encoded nonce value used during encryption. Must be provided if convergent encryption is enabled for this key and the key was generated with Vault 0.6.1. Not required for keys created in 0.6.2+.
	Nonce string `json:"nonce,omitempty"`

	// Ordinarily, if a batch item fails to decrypt due to a bad input, but other batch items succeed, the HTTP response code is 400 (Bad Request). Some applications may want to treat partial failures differently. Providing the parameter returns the given response code integer instead of a 400 in this case. If all values fail HTTP 400 is still returned.
	PartialFailureResponseCode int32 `json:"partial_failure_response_code,omitempty"`
}
