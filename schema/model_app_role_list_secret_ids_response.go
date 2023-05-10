// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleListSecretIdsResponse struct for AppRoleListSecretIdsResponse
type AppRoleListSecretIdsResponse struct {
	Keys []string `json:"keys"`
}

// NewAppRoleListSecretIdsResponseWithDefaults instantiates a new AppRoleListSecretIdsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleListSecretIdsResponseWithDefaults() *AppRoleListSecretIdsResponse {
	var this AppRoleListSecretIdsResponse

	return &this
}

func (o AppRoleListSecretIdsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
