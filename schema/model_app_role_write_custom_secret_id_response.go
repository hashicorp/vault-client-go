// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteCustomSecretIdResponse struct for AppRoleWriteCustomSecretIdResponse
type AppRoleWriteCustomSecretIdResponse struct {
	// Secret ID attached to the role.
	SecretId string `json:"secret_id"`

	// Accessor of the secret ID
	SecretIdAccessor string `json:"secret_id_accessor"`

	// Number of times a secret ID can access the role, after which the secret ID will expire.
	SecretIdNumUses int32 `json:"secret_id_num_uses"`

	// Duration in seconds after which the issued secret ID expires.
	SecretIdTtl int32 `json:"secret_id_ttl"`
}

// NewAppRoleWriteCustomSecretIdResponseWithDefaults instantiates a new AppRoleWriteCustomSecretIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteCustomSecretIdResponseWithDefaults() *AppRoleWriteCustomSecretIdResponse {
	var this AppRoleWriteCustomSecretIdResponse

	return &this
}

func (o AppRoleWriteCustomSecretIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id"] = o.SecretId
	toSerialize["secret_id_accessor"] = o.SecretIdAccessor
	toSerialize["secret_id_num_uses"] = o.SecretIdNumUses
	toSerialize["secret_id_ttl"] = o.SecretIdTtl

	return json.Marshal(toSerialize)
}
