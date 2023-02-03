// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenWriteCreateRequest struct for TokenWriteCreateRequest
type TokenWriteCreateRequest struct {
	// Name to associate with this token
	DisplayName string `json:"display_name"`

	// Name of the entity alias to associate with this token
	EntityAlias string `json:"entity_alias"`

	// Explicit Max TTL of this token
	ExplicitMaxTtl string `json:"explicit_max_ttl"`

	// Value for the token
	Id string `json:"id"`

	// Arbitrary key=value metadata to associate with the token
	Metadata map[string]interface{} `json:"metadata"`

	// Do not include default policy for this token
	NoDefaultPolicy bool `json:"no_default_policy"`

	// Create the token with no parent
	NoParent bool `json:"no_parent"`

	// Max number of uses for this token
	NumUses int32 `json:"num_uses"`

	// Renew period
	Period string `json:"period"`

	// List of policies for the token
	Policies []string `json:"policies"`

	// Allow token to be renewed past its initial TTL up to system/mount maximum TTL
	Renewable bool `json:"renewable"`

	// Time to live for this token
	Ttl string `json:"ttl"`

	// Token type
	Type string `json:"type"`
}

// NewTokenWriteCreateRequestWithDefaults instantiates a new TokenWriteCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWriteCreateRequestWithDefaults() *TokenWriteCreateRequest {
	var this TokenWriteCreateRequest

	return &this
}
