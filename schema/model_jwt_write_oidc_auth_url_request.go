// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// JWTWriteOIDCAuthURLRequest struct for JWTWriteOIDCAuthURLRequest
type JWTWriteOIDCAuthURLRequest struct {
	// Optional client-provided nonce that must match during callback, if present.
	ClientNonce string `json:"client_nonce"`

	// The OAuth redirect_uri to use in the authorization URL.
	RedirectUri string `json:"redirect_uri"`

	// The role to issue an OIDC authorization URL against.
	Role string `json:"role"`
}

// NewJWTWriteOIDCAuthURLRequestWithDefaults instantiates a new JWTWriteOIDCAuthURLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWTWriteOIDCAuthURLRequestWithDefaults() *JWTWriteOIDCAuthURLRequest {
	var this JWTWriteOIDCAuthURLRequest

	return &this
}

func (o JWTWriteOIDCAuthURLRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
