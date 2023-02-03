// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCRotateKeyRequest struct for OIDCRotateKeyRequest
type OIDCRotateKeyRequest struct {
	// Controls how long the public portion of a key will be available for verification after being rotated. Setting verification_ttl here will override the verification_ttl set on the key.
	VerificationTtl int32 `json:"verification_ttl"`
}

// NewOIDCRotateKeyRequestWithDefaults instantiates a new OIDCRotateKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCRotateKeyRequestWithDefaults() *OIDCRotateKeyRequest {
	var this OIDCRotateKeyRequest

	return &this
}

func (o OIDCRotateKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
