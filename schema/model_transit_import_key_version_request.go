// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitImportKeyVersionRequest struct for TransitImportKeyVersionRequest
type TransitImportKeyVersionRequest struct {
	// The base64-encoded ciphertext of the keys. The AES key should be encrypted using OAEP with the wrapping key and then concatenated with the import key, wrapped by the AES key.
	Ciphertext string `json:"ciphertext,omitempty"`

	// The hash function used as a random oracle in the OAEP wrapping of the user-generated, ephemeral AES key. Can be one of \"SHA1\", \"SHA224\", \"SHA256\" (default), \"SHA384\", or \"SHA512\"
	HashFunction string `json:"hash_function,omitempty"`

	// The plaintext public key to be imported. If \"ciphertext\" is set, this field is ignored.
	PublicKey string `json:"public_key,omitempty"`

	// Key version to be updated, if left empty, a new version will be created unless a private key is specified and the 'Latest' key is missing a private key.
	Version int32 `json:"version,omitempty"`
}
