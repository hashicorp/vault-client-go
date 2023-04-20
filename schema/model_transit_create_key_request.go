// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitCreateKeyRequest struct for TransitCreateKeyRequest
type TransitCreateKeyRequest struct {
	// Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled.
	AllowPlaintextBackup bool `json:"allow_plaintext_backup"`

	// Amount of time the key should live before being automatically rotated. A value of 0 (default) disables automatic rotation for the key.
	AutoRotatePeriod int32 `json:"auto_rotate_period"`

	// Base64 encoded context for key derivation. When reading a key with key derivation enabled, if the key type supports public keys, this will return the public key for the given context.
	Context string `json:"context"`

	// Whether to support convergent encryption. This is only supported when using a key with key derivation enabled and will require all requests to carry both a context and 96-bit (12-byte) nonce. The given nonce will be used in place of a randomly generated nonce. As a result, when the same context and nonce are supplied, the same ciphertext is generated. It is *very important* when using this mode that you ensure that all nonces are unique for a given context. Failing to do so will severely impact the ciphertext's security.
	ConvergentEncryption bool `json:"convergent_encryption"`

	// Enables key derivation mode. This allows for per-transaction unique keys for encryption operations.
	Derived bool `json:"derived"`

	// Enables keys to be exportable. This allows for all the valid keys in the key ring to be exported.
	Exportable bool `json:"exportable"`

	// The key size in bytes for the algorithm. Only applies to HMAC and must be no fewer than 32 bytes and no more than 512
	KeySize int32 `json:"key_size"`

	// The UUID of the managed key to use for this transit key
	ManagedKeyId string `json:"managed_key_id"`

	// The name of the managed key to use for this transit key
	ManagedKeyName string `json:"managed_key_name"`

	// The type of key to create. Currently, \"aes128-gcm96\" (symmetric), \"aes256-gcm96\" (symmetric), \"ecdsa-p256\" (asymmetric), \"ecdsa-p384\" (asymmetric), \"ecdsa-p521\" (asymmetric), \"ed25519\" (asymmetric), \"rsa-2048\" (asymmetric), \"rsa-3072\" (asymmetric), \"rsa-4096\" (asymmetric) are supported. Defaults to \"aes256-gcm96\".
	Type string `json:"type"`
}

// NewTransitCreateKeyRequestWithDefaults instantiates a new TransitCreateKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitCreateKeyRequestWithDefaults() *TransitCreateKeyRequest {
	var this TransitCreateKeyRequest

	this.AutoRotatePeriod = 0
	this.KeySize = 0
	this.Type = "aes256-gcm96"

	return &this
}

func (o TransitCreateKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allow_plaintext_backup"] = o.AllowPlaintextBackup
	toSerialize["auto_rotate_period"] = o.AutoRotatePeriod
	toSerialize["context"] = o.Context
	toSerialize["convergent_encryption"] = o.ConvergentEncryption
	toSerialize["derived"] = o.Derived
	toSerialize["exportable"] = o.Exportable
	toSerialize["key_size"] = o.KeySize
	toSerialize["managed_key_id"] = o.ManagedKeyId
	toSerialize["managed_key_name"] = o.ManagedKeyName
	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
