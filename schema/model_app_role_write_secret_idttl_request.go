// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteSecretIDTTLRequest struct for AppRoleWriteSecretIDTTLRequest
type AppRoleWriteSecretIDTTLRequest struct {
	// Duration in seconds after which the issued SecretID should expire. Defaults to 0, meaning no expiration.
	SecretIdTtl int32 `json:"secret_id_ttl"`
}

// NewAppRoleWriteSecretIDTTLRequestWithDefaults instantiates a new AppRoleWriteSecretIDTTLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDTTLRequestWithDefaults() *AppRoleWriteSecretIDTTLRequest {
	var this AppRoleWriteSecretIDTTLRequest

	return &this
}

func (o AppRoleWriteSecretIDTTLRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id_ttl"] = o.SecretIdTtl

	return json.Marshal(toSerialize)
}
