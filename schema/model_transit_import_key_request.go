// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitImportKeyRequest struct for TransitImportKeyRequest
type TransitImportKeyRequest struct {
	// Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled.
	AllowPlaintextBackup bool `json:"allow_plaintext_backup,omitempty"`

	// True if the imported key may be rotated within Vault; false otherwise.
	AllowRotation bool `json:"allow_rotation,omitempty"`

	// Amount of time the key should live before being automatically rotated. A value of 0 (default) disables automatic rotation for the key.
	AutoRotatePeriod string `json:"auto_rotate_period,omitempty"`

	// The base64-encoded ciphertext of the keys. The AES key should be encrypted using OAEP with the wrapping key and then concatenated with the import key, wrapped by the AES key.
	Ciphertext string `json:"ciphertext,omitempty"`

	// Base64 encoded context for key derivation. When reading a key with key derivation enabled, if the key type supports public keys, this will return the public key for the given context.
	Context string `json:"context,omitempty"`

	// Enables key derivation mode. This allows for per-transaction unique keys for encryption operations.
	Derived bool `json:"derived,omitempty"`

	// Enables keys to be exportable. This allows for all the valid keys in the key ring to be exported.
	Exportable bool `json:"exportable,omitempty"`

	// The hash function used as a random oracle in the OAEP wrapping of the user-generated, ephemeral AES key. Can be one of \"SHA1\", \"SHA224\", \"SHA256\" (default), \"SHA384\", or \"SHA512\"
	HashFunction string `json:"hash_function,omitempty"`

	// The plaintext PEM public key to be imported. If \"ciphertext\" is set, this field is ignored.
	PublicKey string `json:"public_key,omitempty"`

	// The type of key being imported. Currently, \"aes128-gcm96\" (symmetric), \"aes256-gcm96\" (symmetric), \"ecdsa-p256\" (asymmetric), \"ecdsa-p384\" (asymmetric), \"ecdsa-p521\" (asymmetric), \"ed25519\" (asymmetric), \"rsa-2048\" (asymmetric), \"rsa-3072\" (asymmetric), \"rsa-4096\" (asymmetric) are supported. Defaults to \"aes256-gcm96\".
	Type string `json:"type,omitempty"`
}
