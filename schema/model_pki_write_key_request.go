// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PKIWriteKeyRequest struct for PKIWriteKeyRequest
type PKIWriteKeyRequest struct {
	// Human-readable name for this key.
	KeyName string `json:"key_name"`
}

// NewPKIWriteKeyRequestWithDefaults instantiates a new PKIWriteKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteKeyRequestWithDefaults() *PKIWriteKeyRequest {
	var this PKIWriteKeyRequest

	return &this
}
