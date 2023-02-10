// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleListSecretIDResponse struct for AppRoleListSecretIDResponse
type AppRoleListSecretIDResponse struct {
	Keys []string `json:"keys"`
}

// NewAppRoleListSecretIDResponseWithDefaults instantiates a new AppRoleListSecretIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleListSecretIDResponseWithDefaults() *AppRoleListSecretIDResponse {
	var this AppRoleListSecretIDResponse

	return &this
}

func (o AppRoleListSecretIDResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
