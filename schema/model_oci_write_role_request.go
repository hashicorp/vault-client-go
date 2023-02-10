// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OCIWriteRoleRequest struct for OCIWriteRoleRequest
type OCIWriteRoleRequest struct {
	// A comma separated list of Group or Dynamic Group OCIDs that are allowed to take this role.
	OcidList []string `json:"ocid_list"`

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
}

// NewOCIWriteRoleRequestWithDefaults instantiates a new OCIWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCIWriteRoleRequestWithDefaults() *OCIWriteRoleRequest {
	var this OCIWriteRoleRequest

	this.TokenType = "default-service"

	return &this
}

func (o OCIWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ocid_list"] = o.OcidList
	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	toSerialize["token_explicit_max_ttl"] = o.TokenExplicitMaxTtl
	toSerialize["token_max_ttl"] = o.TokenMaxTtl
	toSerialize["token_no_default_policy"] = o.TokenNoDefaultPolicy
	toSerialize["token_num_uses"] = o.TokenNumUses
	toSerialize["token_period"] = o.TokenPeriod
	toSerialize["token_policies"] = o.TokenPolicies
	toSerialize["token_ttl"] = o.TokenTtl
	toSerialize["token_type"] = o.TokenType

	return json.Marshal(toSerialize)
}
