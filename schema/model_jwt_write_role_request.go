// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// JWTWriteRoleRequest struct for JWTWriteRoleRequest
type JWTWriteRoleRequest struct {
	// Comma-separated list of allowed values for redirect_uri
	AllowedRedirectUris []string `json:"allowed_redirect_uris"`

	// Comma-separated list of 'aud' claims that are valid for login; any match is sufficient
	BoundAudiences []string `json:"bound_audiences"`

	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs"`

	// Map of claims/values which must match for login
	BoundClaims map[string]interface{} `json:"bound_claims"`

	// How to interpret values in the map of claims/values (which must match for login): allowed values are 'string' or 'glob'
	BoundClaimsType string `json:"bound_claims_type"`

	// The 'sub' claim that is valid for login. Optional.
	BoundSubject string `json:"bound_subject"`

	// Mappings of claims (key) that will be copied to a metadata field (value)
	ClaimMappings map[string]interface{} `json:"claim_mappings"`

	// Duration in seconds of leeway when validating all claims to account for clock skew. Defaults to 60 (1 minute) if set to 0 and can be disabled if set to -1.
	ClockSkewLeeway int32 `json:"clock_skew_leeway"`

	// Duration in seconds of leeway when validating expiration of a token to account for clock skew. Defaults to 150 (2.5 minutes) if set to 0 and can be disabled if set to -1.
	ExpirationLeeway int32 `json:"expiration_leeway"`

	// The claim to use for the Identity group alias names
	GroupsClaim string `json:"groups_claim"`

	// Specifies the allowable elapsed time in seconds since the last time the user was actively authenticated.
	MaxAge int32 `json:"max_age"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl int32 `json:"max_ttl"`

	// Duration in seconds of leeway when validating not before values of a token to account for clock skew. Defaults to 150 (2.5 minutes) if set to 0 and can be disabled if set to -1.
	NotBeforeLeeway int32 `json:"not_before_leeway"`

	// Use \"token_num_uses\" instead. If this and \"token_num_uses\" are both specified, only \"token_num_uses\" will be used.
	// Deprecated
	NumUses int32 `json:"num_uses"`

	// Comma-separated list of OIDC scopes
	OidcScopes []string `json:"oidc_scopes"`

	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period int32 `json:"period"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`

	// Type of the role, either 'jwt' or 'oidc'.
	RoleType string `json:"role_type"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`

	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl int32 `json:"token_explicit_max_ttl"`

	// The maximum lifetime of the generated token
	TokenMaxTtl int32 `json:"token_max_ttl"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod int32 `json:"token_period"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`

	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type"`

	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl int32 `json:"ttl"`

	// The claim to use for the Identity entity alias name
	UserClaim string `json:"user_claim"`

	// If true, the user_claim value will use JSON pointer syntax for referencing claims.
	UserClaimJsonPointer bool `json:"user_claim_json_pointer"`

	// Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive information may be present in OIDC responses.
	VerboseOidcLogging bool `json:"verbose_oidc_logging"`
}

// NewJWTWriteRoleRequestWithDefaults instantiates a new JWTWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWTWriteRoleRequestWithDefaults() *JWTWriteRoleRequest {
	var this JWTWriteRoleRequest

	this.BoundClaimsType = "string"
	this.ExpirationLeeway = 150
	this.NotBeforeLeeway = 150
	this.TokenType = "default-service"

	return &this
}

func (o JWTWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allowed_redirect_uris"] = o.AllowedRedirectUris
	toSerialize["bound_audiences"] = o.BoundAudiences
	toSerialize["bound_cidrs"] = o.BoundCidrs
	toSerialize["bound_claims"] = o.BoundClaims
	toSerialize["bound_claims_type"] = o.BoundClaimsType
	toSerialize["bound_subject"] = o.BoundSubject
	toSerialize["claim_mappings"] = o.ClaimMappings
	toSerialize["clock_skew_leeway"] = o.ClockSkewLeeway
	toSerialize["expiration_leeway"] = o.ExpirationLeeway
	toSerialize["groups_claim"] = o.GroupsClaim
	toSerialize["max_age"] = o.MaxAge
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["not_before_leeway"] = o.NotBeforeLeeway
	toSerialize["num_uses"] = o.NumUses
	toSerialize["oidc_scopes"] = o.OidcScopes
	toSerialize["period"] = o.Period
	toSerialize["policies"] = o.Policies
	toSerialize["role_type"] = o.RoleType
	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	toSerialize["token_explicit_max_ttl"] = o.TokenExplicitMaxTtl
	toSerialize["token_max_ttl"] = o.TokenMaxTtl
	toSerialize["token_no_default_policy"] = o.TokenNoDefaultPolicy
	toSerialize["token_num_uses"] = o.TokenNumUses
	toSerialize["token_period"] = o.TokenPeriod
	toSerialize["token_policies"] = o.TokenPolicies
	toSerialize["token_ttl"] = o.TokenTtl
	toSerialize["token_type"] = o.TokenType
	toSerialize["ttl"] = o.Ttl
	toSerialize["user_claim"] = o.UserClaim
	toSerialize["user_claim_json_pointer"] = o.UserClaimJsonPointer
	toSerialize["verbose_oidc_logging"] = o.VerboseOidcLogging

	return json.Marshal(toSerialize)
}
