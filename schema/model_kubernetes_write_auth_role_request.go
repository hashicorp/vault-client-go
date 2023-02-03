// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KubernetesWriteAuthRoleRequest struct for KubernetesWriteAuthRoleRequest
type KubernetesWriteAuthRoleRequest struct {
	// Source to use when deriving the Alias name. valid choices: \"serviceaccount_uid\" : <token.uid> e.g. 474b11b5-0f20-4f9d-8ca5-65715ab325e0 (most secure choice) \"serviceaccount_name\" : <namespace>/<serviceaccount> e.g. vault/vault-agent default: \"serviceaccount_uid\"
	AliasNameSource string `json:"alias_name_source"`

	// Optional Audience claim to verify in the jwt.
	Audience string `json:"audience"`

	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs"`

	// List of service account names able to access this role. If set to \"*\" all names are allowed.
	BoundServiceAccountNames []string `json:"bound_service_account_names"`

	// List of namespaces allowed to access this role. If set to \"*\" all namespaces are allowed.
	BoundServiceAccountNamespaces []string `json:"bound_service_account_namespaces"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl int32 `json:"max_ttl"`

	// Use \"token_num_uses\" instead. If this and \"token_num_uses\" are both specified, only \"token_num_uses\" will be used.
	// Deprecated
	NumUses int32 `json:"num_uses"`

	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period int32 `json:"period"`

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

// NewKubernetesWriteAuthRoleRequestWithDefaults instantiates a new KubernetesWriteAuthRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesWriteAuthRoleRequestWithDefaults() *KubernetesWriteAuthRoleRequest {
	var this KubernetesWriteAuthRoleRequest

	this.AliasNameSource = "serviceaccount_uid"
	this.TokenType = "default-service"

	return &this
}

func (o KubernetesWriteAuthRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
