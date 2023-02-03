// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TokenRevokeRequest struct for TokenRevokeRequest
type TokenRevokeRequest struct {
	// Token to revoke (request body)
	Token string `json:"token"`
}

// NewTokenRevokeRequestWithDefaults instantiates a new TokenRevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRevokeRequestWithDefaults() *TokenRevokeRequest {
	var this TokenRevokeRequest

	return &this
}
