/*

HashiCorp Vault API



HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.



API version: 1.13.0


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenRevokeOrphanRequest struct for TokenRevokeOrphanRequest
type TokenRevokeOrphanRequest struct {
	// Token to revoke (request body)

	Token string `json:"token"`
}

// NewTokenRevokeOrphanRequestWithDefaults instantiates a new TokenRevokeOrphanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRevokeOrphanRequestWithDefaults() *TokenRevokeOrphanRequest {
	var this TokenRevokeOrphanRequest

	return &this
}

func (o TokenRevokeOrphanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
