// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteCustomSecretIDRequest struct for AppRoleWriteCustomSecretIDRequest
type AppRoleWriteCustomSecretIDRequest struct {
	// Comma separated string or list of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If 'bound_cidr_list' is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role.
	CidrList []string `json:"cidr_list"`

	// Metadata to be tied to the SecretID. This should be a JSON formatted string containing metadata in key value pairs.
	Metadata string `json:"metadata"`

	// Number of times this SecretID can be used, after which the SecretID expires. Overrides secret_id_num_uses role option when supplied. May not be higher than role's secret_id_num_uses.
	NumUses int32 `json:"num_uses"`

	// SecretID to be attached to the role.
	SecretId string `json:"secret_id"`

	// Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`

	// Duration in seconds after which this SecretID expires. Overrides secret_id_ttl role option when supplied. May not be longer than role's secret_id_ttl.
	Ttl int32 `json:"ttl"`
}

// NewAppRoleWriteCustomSecretIDRequestWithDefaults instantiates a new AppRoleWriteCustomSecretIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteCustomSecretIDRequestWithDefaults() *AppRoleWriteCustomSecretIDRequest {
	var this AppRoleWriteCustomSecretIDRequest

	return &this
}

func (o AppRoleWriteCustomSecretIDRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cidr_list"] = o.CidrList
	toSerialize["metadata"] = o.Metadata
	toSerialize["num_uses"] = o.NumUses
	toSerialize["secret_id"] = o.SecretId
	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
