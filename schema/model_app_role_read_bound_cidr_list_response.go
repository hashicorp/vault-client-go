// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadBoundCidrListResponse struct for AppRoleReadBoundCidrListResponse
type AppRoleReadBoundCidrListResponse struct {
	// Deprecated: Please use \"secret_id_bound_cidrs\" instead. Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.
	// Deprecated
	BoundCidrList []string `json:"bound_cidr_list"`
}

// NewAppRoleReadBoundCidrListResponseWithDefaults instantiates a new AppRoleReadBoundCidrListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadBoundCidrListResponseWithDefaults() *AppRoleReadBoundCidrListResponse {
	var this AppRoleReadBoundCidrListResponse

	return &this
}

func (o AppRoleReadBoundCidrListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bound_cidr_list"] = o.BoundCidrList

	return json.Marshal(toSerialize)
}
