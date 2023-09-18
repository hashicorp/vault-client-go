// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RekeyAttemptInitializeRequest struct for RekeyAttemptInitializeRequest
type RekeyAttemptInitializeRequest struct {
	// Specifies if using PGP-encrypted keys, whether Vault should also store a plaintext backup of the PGP-encrypted keys.
	Backup bool `json:"backup,omitempty"`

	// Specifies an array of PGP public keys used to encrypt the output unseal keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as secret_shares.
	PgpKeys []string `json:"pgp_keys,omitempty"`

	// Turns on verification functionality
	RequireVerification bool `json:"require_verification,omitempty"`

	// Specifies the number of shares to split the unseal key into.
	SecretShares int32 `json:"secret_shares,omitempty"`

	// Specifies the number of shares required to reconstruct the unseal key. This must be less than or equal secret_shares. If using Vault HSM with auto-unsealing, this value must be the same as secret_shares.
	SecretThreshold int32 `json:"secret_threshold,omitempty"`
}
