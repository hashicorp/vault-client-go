// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleLookUpSecretIdRequest struct for AppRoleLookUpSecretIdRequest
type AppRoleLookUpSecretIdRequest struct {
	// SecretID attached to the role.
	SecretId string `json:"secret_id"`
}

// NewAppRoleLookUpSecretIdRequestWithDefaults instantiates a new AppRoleLookUpSecretIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleLookUpSecretIdRequestWithDefaults() *AppRoleLookUpSecretIdRequest {
	var this AppRoleLookUpSecretIdRequest

	return &this
}

func (o AppRoleLookUpSecretIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id"] = o.SecretId

	return json.Marshal(toSerialize)
}
