// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// JwtWriteRoleRequest struct for JwtWriteRoleRequest
type JwtWriteRoleRequest struct {
	// Comma-separated list of allowed values for redirect_uri
	AllowedRedirectUris []string `json:"allowed_redirect_uris,omitempty"`

	// Comma-separated list of 'aud' claims that are valid for login; any match is sufficient
	BoundAudiences []string `json:"bound_audiences,omitempty"`

	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs,omitempty"`

	// Map of claims/values which must match for login
	BoundClaims map[string]interface{} `json:"bound_claims,omitempty"`

	// How to interpret values in the map of claims/values (which must match for login): allowed values are 'string' or 'glob'
	BoundClaimsType string `json:"bound_claims_type,omitempty"`

	// The 'sub' claim that is valid for login. Optional.
	BoundSubject string `json:"bound_subject,omitempty"`

	// Mappings of claims (key) that will be copied to a metadata field (value)
	ClaimMappings map[string]interface{} `json:"claim_mappings,omitempty"`

	// Duration in seconds of leeway when validating all claims to account for clock skew. Defaults to 60 (1 minute) if set to 0 and can be disabled if set to -1.
	ClockSkewLeeway string `json:"clock_skew_leeway,omitempty"`

	// Duration in seconds of leeway when validating expiration of a token to account for clock skew. Defaults to 150 (2.5 minutes) if set to 0 and can be disabled if set to -1.
	ExpirationLeeway string `json:"expiration_leeway,omitempty"`

	// The claim to use for the Identity group alias names
	GroupsClaim string `json:"groups_claim,omitempty"`

	// Specifies the allowable elapsed time in seconds since the last time the user was actively authenticated.
	MaxAge string `json:"max_age,omitempty"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl string `json:"max_ttl,omitempty"`

	// Duration in seconds of leeway when validating not before values of a token to account for clock skew. Defaults to 150 (2.5 minutes) if set to 0 and can be disabled if set to -1.
	NotBeforeLeeway string `json:"not_before_leeway,omitempty"`

	// Use \"token_num_uses\" instead. If this and \"token_num_uses\" are both specified, only \"token_num_uses\" will be used.
	// Deprecated
	NumUses int32 `json:"num_uses,omitempty"`

	// Comma-separated list of OIDC scopes
	OidcScopes []string `json:"oidc_scopes,omitempty"`

	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period string `json:"period,omitempty"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies,omitempty"`

	// Type of the role, either 'jwt' or 'oidc'.
	RoleType string `json:"role_type,omitempty"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs,omitempty"`

	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl string `json:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	TokenMaxTtl string `json:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses,omitempty"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod string `json:"token_period,omitempty"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies,omitempty"`

	// The initial ttl of the token to generate
	TokenTtl string `json:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type,omitempty"`

	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl string `json:"ttl,omitempty"`

	// The claim to use for the Identity entity alias name
	UserClaim string `json:"user_claim,omitempty"`

	// If true, the user_claim value will use JSON pointer syntax for referencing claims.
	UserClaimJsonPointer bool `json:"user_claim_json_pointer,omitempty"`

	// Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive information may be present in OIDC responses.
	VerboseOidcLogging bool `json:"verbose_oidc_logging,omitempty"`
}
