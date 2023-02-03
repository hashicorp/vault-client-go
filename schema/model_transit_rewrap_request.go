// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitRewrapRequest struct for TransitRewrapRequest
type TransitRewrapRequest struct {
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
