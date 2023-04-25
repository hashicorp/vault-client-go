// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadSecretIdTtlResponse struct for AppRoleReadSecretIdTtlResponse
type AppRoleReadSecretIdTtlResponse struct {
	// Duration in seconds after which the issued secret ID should expire. Defaults to 0, meaning no expiration.
	SecretIdTtl int32 `json:"secret_id_ttl"`
}

// NewAppRoleReadSecretIdTtlResponseWithDefaults instantiates a new AppRoleReadSecretIdTtlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadSecretIdTtlResponseWithDefaults() *AppRoleReadSecretIdTtlResponse {
	var this AppRoleReadSecretIdTtlResponse

	return &this
}

func (o AppRoleReadSecretIdTtlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id_ttl"] = o.SecretIdTtl

	return json.Marshal(toSerialize)
}
