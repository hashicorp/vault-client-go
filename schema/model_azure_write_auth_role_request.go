// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AzureWriteAuthRoleRequest struct for AzureWriteAuthRoleRequest
type AzureWriteAuthRoleRequest struct {
	// Comma-separated list of group ids that login is restricted to.
	BoundGroupIds []string `json:"bound_group_ids"`

	// Comma-separated list of locations that login is restricted to.
	BoundLocations []string `json:"bound_locations"`

	// Comma-separated list of resource groups that login is restricted to.
	BoundResourceGroups []string `json:"bound_resource_groups"`

	// Comma-separated list of scale sets that login is restricted to.
	BoundScaleSets []string `json:"bound_scale_sets"`

	// Comma-separated list of service principal ids that login is restricted to.
	BoundServicePrincipalIds []string `json:"bound_service_principal_ids"`

	// Comma-separated list of subscription ids that login is restricted to.
	BoundSubscriptionIds []string `json:"bound_subscription_ids"`

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

// NewAzureWriteAuthRoleRequestWithDefaults instantiates a new AzureWriteAuthRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureWriteAuthRoleRequestWithDefaults() *AzureWriteAuthRoleRequest {
	var this AzureWriteAuthRoleRequest

	this.TokenType = "default-service"

	return &this
}

func (o AzureWriteAuthRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
