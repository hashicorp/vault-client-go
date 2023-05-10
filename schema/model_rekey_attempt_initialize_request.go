// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RekeyAttemptInitializeRequest struct for RekeyAttemptInitializeRequest
type RekeyAttemptInitializeRequest struct {
	// Specifies if using PGP-encrypted keys, whether Vault should also store a plaintext backup of the PGP-encrypted keys.
	Backup bool `json:"backup"`

	// Specifies an array of PGP public keys used to encrypt the output unseal keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as secret_shares.
	PgpKeys []string `json:"pgp_keys"`

	// Turns on verification functionality
	RequireVerification bool `json:"require_verification"`

	// Specifies the number of shares to split the unseal key into.
	SecretShares int32 `json:"secret_shares"`

	// Specifies the number of shares required to reconstruct the unseal key. This must be less than or equal secret_shares. If using Vault HSM with auto-unsealing, this value must be the same as secret_shares.
	SecretThreshold int32 `json:"secret_threshold"`
}

// NewRekeyAttemptInitializeRequestWithDefaults instantiates a new RekeyAttemptInitializeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyAttemptInitializeRequestWithDefaults() *RekeyAttemptInitializeRequest {
	var this RekeyAttemptInitializeRequest

	return &this
}

func (o RekeyAttemptInitializeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["backup"] = o.Backup
	toSerialize["pgp_keys"] = o.PgpKeys
	toSerialize["require_verification"] = o.RequireVerification
	toSerialize["secret_shares"] = o.SecretShares
	toSerialize["secret_threshold"] = o.SecretThreshold

	return json.Marshal(toSerialize)
}
