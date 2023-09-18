// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteCustomSecretIdRequest struct for AppRoleWriteCustomSecretIdRequest
type AppRoleWriteCustomSecretIdRequest struct {
	// Comma separated string or list of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If 'bound_cidr_list' is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role.
	CidrList []string `json:"cidr_list,omitempty"`

	// Metadata to be tied to the SecretID. This should be a JSON formatted string containing metadata in key value pairs.
	Metadata string `json:"metadata,omitempty"`

	// Number of times this SecretID can be used, after which the SecretID expires. Overrides secret_id_num_uses role option when supplied. May not be higher than role's secret_id_num_uses.
	NumUses int32 `json:"num_uses,omitempty"`

	// SecretID to be attached to the role.
	SecretId string `json:"secret_id,omitempty"`

	// Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any.
	TokenBoundCidrs []string `json:"token_bound_cidrs,omitempty"`

	// Duration in seconds after which this SecretID expires. Overrides secret_id_ttl role option when supplied. May not be longer than role's secret_id_ttl.
	Ttl string `json:"ttl,omitempty"`
}
