// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenWriteRoleRequest struct for TokenWriteRoleRequest
type TokenWriteRoleRequest struct {
	// String or JSON list of allowed entity aliases. If set, specifies the entity aliases which are allowed to be used during token generation. This field supports globbing.
	AllowedEntityAliases []string `json:"allowed_entity_aliases"`

	// If set, tokens can be created with any subset of the policies in this list, rather than the normal semantics of tokens being a subset of the calling token's policies. The parameter is a comma-delimited string of policy names.
	AllowedPolicies []string `json:"allowed_policies"`

	// If set, tokens can be created with any subset of glob matched policies in this list, rather than the normal semantics of tokens being a subset of the calling token's policies. The parameter is a comma-delimited string of policy name globs.
	AllowedPoliciesGlob []string `json:"allowed_policies_glob"`

	// Use 'token_bound_cidrs' instead.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs"`

	// If set, successful token creation via this role will require that no policies in the given list are requested. The parameter is a comma-delimited string of policy names.
	DisallowedPolicies []string `json:"disallowed_policies"`

	// If set, successful token creation via this role will require that no requested policies glob match any of policies in this list. The parameter is a comma-delimited string of policy name globs.
	DisallowedPoliciesGlob []string `json:"disallowed_policies_glob"`

	// Use 'token_explicit_max_ttl' instead.
	// Deprecated
	ExplicitMaxTtl int32 `json:"explicit_max_ttl"`

	// If true, tokens created via this role will be orphan tokens (have no parent)
	Orphan bool `json:"orphan"`

	// If set, tokens created via this role will contain the given suffix as a part of their path. This can be used to assist use of the 'revoke-prefix' endpoint later on. The given suffix must match the regular expression.\\w[\\w-.]+\\w
	PathSuffix string `json:"path_suffix"`

	// Use 'token_period' instead.
	// Deprecated
	Period int32 `json:"period"`

	// Tokens created via this role will be renewable or not according to this value. Defaults to \"true\".
	Renewable bool `json:"renewable"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`

	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl int32 `json:"token_explicit_max_ttl"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod int32 `json:"token_period"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type"`
}

// NewTokenWriteRoleRequestWithDefaults instantiates a new TokenWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWriteRoleRequestWithDefaults() *TokenWriteRoleRequest {
	var this TokenWriteRoleRequest

	this.Renewable = true
	this.TokenType = "default-service"

	return &this
}

func (o TokenWriteRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
