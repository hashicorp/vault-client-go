// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KubernetesWriteAuthRoleRequest struct for KubernetesWriteAuthRoleRequest
type KubernetesWriteAuthRoleRequest struct {
	// Source to use when deriving the Alias name. valid choices: \"serviceaccount_uid\" : <token.uid> e.g. 474b11b5-0f20-4f9d-8ca5-65715ab325e0 (most secure choice) \"serviceaccount_name\" : <namespace>/<serviceaccount> e.g. vault/vault-agent default: \"serviceaccount_uid\"
	AliasNameSource string `json:"alias_name_source,omitempty"`

	// Optional Audience claim to verify in the jwt.
	Audience string `json:"audience,omitempty"`

	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs,omitempty"`

	// List of service account names able to access this role. If set to \"*\" all names are allowed.
	BoundServiceAccountNames []string `json:"bound_service_account_names,omitempty"`

	// List of namespaces allowed to access this role. If set to \"*\" all namespaces are allowed.
	BoundServiceAccountNamespaces []string `json:"bound_service_account_namespaces,omitempty"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl string `json:"max_ttl,omitempty"`

	// Use \"token_num_uses\" instead. If this and \"token_num_uses\" are both specified, only \"token_num_uses\" will be used.
	// Deprecated
	NumUses int32 `json:"num_uses,omitempty"`

	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period string `json:"period,omitempty"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies,omitempty"`

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
}
