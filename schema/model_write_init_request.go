// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteInitRequest struct for WriteInitRequest
type WriteInitRequest struct {
	// Specifies an array of PGP public keys used to encrypt the output unseal keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as `secret_shares`.
	PgpKeys []string `json:"pgp_keys"`

	// Specifies an array of PGP public keys used to encrypt the output recovery keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as `recovery_shares`.
	RecoveryPgpKeys []string `json:"recovery_pgp_keys"`

	// Specifies the number of shares to split the recovery key into.
	RecoveryShares int32 `json:"recovery_shares"`

	// Specifies the number of shares required to reconstruct the recovery key. This must be less than or equal to `recovery_shares`.
	RecoveryThreshold int32 `json:"recovery_threshold"`

	// Specifies a PGP public key used to encrypt the initial root token. The key must be base64-encoded from its original binary representation.
	RootTokenPgpKey string `json:"root_token_pgp_key"`

	// Specifies the number of shares to split the unseal key into.
	SecretShares int32 `json:"secret_shares"`

	// Specifies the number of shares required to reconstruct the unseal key. This must be less than or equal secret_shares. If using Vault HSM with auto-unsealing, this value must be the same as `secret_shares`.
	SecretThreshold int32 `json:"secret_threshold"`

	// Specifies the number of shares that should be encrypted by the HSM and stored for auto-unsealing. Currently must be the same as `secret_shares`.
	StoredShares int32 `json:"stored_shares"`
}

// NewWriteInitRequestWithDefaults instantiates a new WriteInitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteInitRequestWithDefaults() *WriteInitRequest {
	var this WriteInitRequest

	return &this
}
