// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// UserpassWriteUserRequest struct for UserpassWriteUserRequest
type UserpassWriteUserRequest struct {
	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.

	// Deprecated

	BoundCidrs []string `json:"bound_cidrs"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.

	// Deprecated

	MaxTtl int32 `json:"max_ttl"`

	// Password for this user.

	Password string `json:"password"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.

	// Deprecated

	Policies []string `json:"policies"`

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
}

// NewUserpassWriteUserRequestWithDefaults instantiates a new UserpassWriteUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserpassWriteUserRequestWithDefaults() *UserpassWriteUserRequest {
	var this UserpassWriteUserRequest

	this.TokenType = "default-service"

	return &this
}

func (o UserpassWriteUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bound_cidrs"] = o.BoundCidrs

	toSerialize["max_ttl"] = o.MaxTtl

	toSerialize["password"] = o.Password

	toSerialize["policies"] = o.Policies

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

	return json.Marshal(toSerialize)
}
