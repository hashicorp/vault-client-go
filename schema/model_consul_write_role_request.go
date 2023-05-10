// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// ConsulWriteRoleRequest struct for ConsulWriteRoleRequest
type ConsulWriteRoleRequest struct {
	// Indicates which namespace that the token will be created within. Defaults to 'default'. Available in Consul 1.7 and above.
	ConsulNamespace string `json:"consul_namespace"`

	// List of policies to attach to the token. Either \"consul_policies\" or \"consul_roles\" are required for Consul 1.5 and above, or just \"consul_policies\" if using Consul 1.4.
	ConsulPolicies []string `json:"consul_policies"`

	// List of Consul roles to attach to the token. Either \"policies\" or \"consul_roles\" are required for Consul 1.5 and above.
	ConsulRoles []string `json:"consul_roles"`

	// Use \"ttl\" instead.
	// Deprecated
	Lease int32 `json:"lease"`

	// Indicates that the token should not be replicated globally and instead be local to the current datacenter. Available in Consul 1.4 and above.
	Local bool `json:"local"`

	// Max TTL for the Consul token created from the role.
	MaxTtl int32 `json:"max_ttl"`

	// List of Node Identities to attach to the token. Available in Consul 1.8.1 or above.
	NodeIdentities []string `json:"node_identities"`

	// Indicates which admin partition that the token will be created within. Defaults to 'default'. Available in Consul 1.11 and above.
	Partition string `json:"partition"`

	// Use \"consul_policies\" instead.
	// Deprecated
	Policies []string `json:"policies"`

	// Policy document, base64 encoded. Required for 'client' tokens. Required for Consul pre-1.4.
	// Deprecated
	Policy string `json:"policy"`

	// List of Service Identities to attach to the token, separated by semicolons. Available in Consul 1.5 or above.
	ServiceIdentities []string `json:"service_identities"`

	// Which type of token to create: 'client' or 'management'. If a 'management' token, the \"policy\", \"policies\", and \"consul_roles\" parameters are not required. Defaults to 'client'.
	// Deprecated
	TokenType string `json:"token_type"`

	// TTL for the Consul token created from the role.
	Ttl int32 `json:"ttl"`
}

// NewConsulWriteRoleRequestWithDefaults instantiates a new ConsulWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulWriteRoleRequestWithDefaults() *ConsulWriteRoleRequest {
	var this ConsulWriteRoleRequest

	this.TokenType = "client"

	return &this
}

func (o ConsulWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["consul_namespace"] = o.ConsulNamespace
	toSerialize["consul_policies"] = o.ConsulPolicies
	toSerialize["consul_roles"] = o.ConsulRoles
	toSerialize["lease"] = o.Lease
	toSerialize["local"] = o.Local
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["node_identities"] = o.NodeIdentities
	toSerialize["partition"] = o.Partition
	toSerialize["policies"] = o.Policies
	toSerialize["policy"] = o.Policy
	toSerialize["service_identities"] = o.ServiceIdentities
	toSerialize["token_type"] = o.TokenType
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
