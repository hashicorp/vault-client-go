// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// AppRoleLookUpSecretIdResponse struct for AppRoleLookUpSecretIdResponse
type AppRoleLookUpSecretIdResponse struct {
	// List of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If 'bound_cidr_list' is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role.
	CidrList []string `json:"cidr_list,omitempty"`

	CreationTime time.Time `json:"creation_time,omitempty"`

	ExpirationTime time.Time `json:"expiration_time,omitempty"`

	LastUpdatedTime time.Time `json:"last_updated_time,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Accessor of the secret ID
	SecretIdAccessor string `json:"secret_id_accessor,omitempty"`

	// Number of times a secret ID can access the role, after which the secret ID will expire.
	SecretIdNumUses int32 `json:"secret_id_num_uses,omitempty"`

	// Duration in seconds after which the issued secret ID expires.
	SecretIdTtl int64 `json:"secret_id_ttl,omitempty"`

	// List of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any.
	TokenBoundCidrs []string `json:"token_bound_cidrs,omitempty"`
}
