// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadRoleResponse struct for AppRoleReadRoleResponse
type AppRoleReadRoleResponse struct {
	// Impose secret ID to be presented when logging in using this role.
	BindSecretId bool `json:"bind_secret_id"`

	// If true, the secret identifiers generated using this role will be cluster local. This can only be set during role creation and once set, it can't be reset later
	LocalSecretIds bool `json:"local_secret_ids"`

	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period int32 `json:"period"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`

	// Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.
	SecretIdBoundCidrs []string `json:"secret_id_bound_cidrs"`

	// Number of times a secret ID can access the role, after which the secret ID will expire.
	SecretIdNumUses int32 `json:"secret_id_num_uses"`

	// Duration in seconds after which the issued secret ID expires.
	SecretIdTtl int32 `json:"secret_id_ttl"`

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

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value.
	TokenPeriod int32 `json:"token_period"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`

	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type"`
}

// NewAppRoleReadRoleResponseWithDefaults instantiates a new AppRoleReadRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadRoleResponseWithDefaults() *AppRoleReadRoleResponse {
	var this AppRoleReadRoleResponse

	this.TokenType = "default-service"

	return &this
}

func (o AppRoleReadRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
