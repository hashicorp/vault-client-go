// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleDestroySecretIdRequest struct for AppRoleDestroySecretIdRequest
type AppRoleDestroySecretIdRequest struct {
	// SecretID attached to the role.
	SecretId string `json:"secret_id"`
}

// NewAppRoleDestroySecretIdRequestWithDefaults instantiates a new AppRoleDestroySecretIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleDestroySecretIdRequestWithDefaults() *AppRoleDestroySecretIdRequest {
	var this AppRoleDestroySecretIdRequest

	return &this
}

func (o AppRoleDestroySecretIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id"] = o.SecretId

	return json.Marshal(toSerialize)
}
