// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InitializeRequest struct for InitializeRequest
type InitializeRequest struct {
	// Specifies an array of PGP public keys used to encrypt the output unseal keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as `secret_shares`.
	PgpKeys []string `json:"pgp_keys,omitempty"`

	// Specifies an array of PGP public keys used to encrypt the output recovery keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as `recovery_shares`.
	RecoveryPgpKeys []string `json:"recovery_pgp_keys,omitempty"`

	// Specifies the number of shares to split the recovery key into.
	RecoveryShares int32 `json:"recovery_shares,omitempty"`

	// Specifies the number of shares required to reconstruct the recovery key. This must be less than or equal to `recovery_shares`.
	RecoveryThreshold int32 `json:"recovery_threshold,omitempty"`

	// Specifies a PGP public key used to encrypt the initial root token. The key must be base64-encoded from its original binary representation.
	RootTokenPgpKey string `json:"root_token_pgp_key,omitempty"`

	// Specifies the number of shares to split the unseal key into.
	SecretShares int32 `json:"secret_shares,omitempty"`

	// Specifies the number of shares required to reconstruct the unseal key. This must be less than or equal secret_shares. If using Vault HSM with auto-unsealing, this value must be the same as `secret_shares`.
	SecretThreshold int32 `json:"secret_threshold,omitempty"`

	// Specifies the number of shares that should be encrypted by the HSM and stored for auto-unsealing. Currently must be the same as `secret_shares`.
	StoredShares int32 `json:"stored_shares,omitempty"`
}
