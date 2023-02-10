// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadSecretIDBoundCIDRsResponse struct for AppRoleReadSecretIDBoundCIDRsResponse
type AppRoleReadSecretIDBoundCIDRsResponse struct {
	// Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.
	SecretIdBoundCidrs []string `json:"secret_id_bound_cidrs"`
}

// NewAppRoleReadSecretIDBoundCIDRsResponseWithDefaults instantiates a new AppRoleReadSecretIDBoundCIDRsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadSecretIDBoundCIDRsResponseWithDefaults() *AppRoleReadSecretIDBoundCIDRsResponse {
	var this AppRoleReadSecretIDBoundCIDRsResponse

	return &this
}

func (o AppRoleReadSecretIDBoundCIDRsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id_bound_cidrs"] = o.SecretIdBoundCidrs

	return json.Marshal(toSerialize)
}
