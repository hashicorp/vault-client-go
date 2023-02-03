// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteKeyRequest struct for OIDCWriteKeyRequest
type OIDCWriteKeyRequest struct {
	// Signing algorithm to use. This will default to RS256.
	Algorithm string `json:"algorithm"`

	// Comma separated string or array of role client ids allowed to use this key for signing. If empty no roles are allowed. If \"*\" all roles are allowed.
	AllowedClientIds []string `json:"allowed_client_ids"`

	// How often to generate a new keypair.
	RotationPeriod int32 `json:"rotation_period"`

	// Controls how long the public portion of a key will be available for verification after being rotated.
	VerificationTtl int32 `json:"verification_ttl"`
}

// NewOIDCWriteKeyRequestWithDefaults instantiates a new OIDCWriteKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteKeyRequestWithDefaults() *OIDCWriteKeyRequest {
	var this OIDCWriteKeyRequest

	this.Algorithm = "RS256"

	return &this
}

func (o OIDCWriteKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
