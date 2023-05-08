// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TokenCreateRequest struct for TokenCreateRequest
type TokenCreateRequest struct {
	// Name to associate with this token
	DisplayName string `json:"display_name,omitempty"`

	// Name of the entity alias to associate with this token
	EntityAlias string `json:"entity_alias,omitempty"`

	// Explicit Max TTL of this token
	ExplicitMaxTtl string `json:"explicit_max_ttl,omitempty"`

	// Value for the token
	Id string `json:"id,omitempty"`

	// Arbitrary key=value metadata to associate with the token
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Do not include default policy for this token
	NoDefaultPolicy bool `json:"no_default_policy,omitempty"`

	// Create the token with no parent
	NoParent bool `json:"no_parent,omitempty"`

	// Max number of uses for this token
	NumUses int32 `json:"num_uses,omitempty"`

	// Renew period
	Period string `json:"period,omitempty"`

	// List of policies for the token
	Policies []string `json:"policies,omitempty"`

	// Allow token to be renewed past its initial TTL up to system/mount maximum TTL
	Renewable bool `json:"renewable,omitempty"`

	// Time to live for this token
	Ttl string `json:"ttl,omitempty"`

	// Token type
	Type string `json:"type,omitempty"`
}

// NewTokenCreateRequestWithDefaults instantiates a new TokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCreateRequestWithDefaults() *TokenCreateRequest {
	var this TokenCreateRequest

	return &this
}
