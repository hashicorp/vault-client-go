// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteSecretIDNumUsesRequest struct for AppRoleWriteSecretIDNumUsesRequest
type AppRoleWriteSecretIDNumUsesRequest struct {
	// Number of times a SecretID can access the role, after which the SecretID will expire.
	SecretIdNumUses int32 `json:"secret_id_num_uses"`
}

// NewAppRoleWriteSecretIDNumUsesRequestWithDefaults instantiates a new AppRoleWriteSecretIDNumUsesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDNumUsesRequestWithDefaults() *AppRoleWriteSecretIDNumUsesRequest {
	var this AppRoleWriteSecretIDNumUsesRequest

	return &this
}

func (o AppRoleWriteSecretIDNumUsesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id_num_uses"] = o.SecretIdNumUses

	return json.Marshal(toSerialize)
}
