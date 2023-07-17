// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// ConsulWriteRoleRequest struct for ConsulWriteRoleRequest
type ConsulWriteRoleRequest struct {
	// Indicates which namespace that the token will be created within. Defaults to 'default'. Available in Consul 1.7 and above.
	ConsulNamespace string `json:"consul_namespace,omitempty"`

	// List of policies to attach to the token. Either \"consul_policies\" or \"consul_roles\" are required for Consul 1.5 and above, or just \"consul_policies\" if using Consul 1.4.
	ConsulPolicies []string `json:"consul_policies,omitempty"`

	// List of Consul roles to attach to the token. Either \"policies\" or \"consul_roles\" are required for Consul 1.5 and above.
	ConsulRoles []string `json:"consul_roles,omitempty"`

	// Use \"ttl\" instead.
	// Deprecated
	Lease string `json:"lease,omitempty"`

	// Indicates that the token should not be replicated globally and instead be local to the current datacenter. Available in Consul 1.4 and above.
	Local bool `json:"local,omitempty"`

	// Max TTL for the Consul token created from the role.
	MaxTtl string `json:"max_ttl,omitempty"`

	// List of Node Identities to attach to the token. Available in Consul 1.8.1 or above.
	NodeIdentities []string `json:"node_identities,omitempty"`

	// Indicates which admin partition that the token will be created within. Defaults to 'default'. Available in Consul 1.11 and above.
	Partition string `json:"partition,omitempty"`

	// Use \"consul_policies\" instead.
	// Deprecated
	Policies []string `json:"policies,omitempty"`

	// Policy document, base64 encoded. Required for 'client' tokens. Required for Consul pre-1.4.
	// Deprecated
	Policy string `json:"policy,omitempty"`

	// List of Service Identities to attach to the token, separated by semicolons. Available in Consul 1.5 or above.
	ServiceIdentities []string `json:"service_identities,omitempty"`

	// Which type of token to create: 'client' or 'management'. If a 'management' token, the \"policy\", \"policies\", and \"consul_roles\" parameters are not required. Defaults to 'client'.
	// Deprecated
	TokenType string `json:"token_type,omitempty"`

	// TTL for the Consul token created from the role.
	Ttl string `json:"ttl,omitempty"`
}
