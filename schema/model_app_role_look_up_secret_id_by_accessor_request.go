// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleLookUpSecretIdByAccessorRequest struct for AppRoleLookUpSecretIdByAccessorRequest
type AppRoleLookUpSecretIdByAccessorRequest struct {
	// Accessor of the SecretID
	SecretIdAccessor string `json:"secret_id_accessor"`
}

// NewAppRoleLookUpSecretIdByAccessorRequestWithDefaults instantiates a new AppRoleLookUpSecretIdByAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleLookUpSecretIdByAccessorRequestWithDefaults() *AppRoleLookUpSecretIdByAccessorRequest {
	var this AppRoleLookUpSecretIdByAccessorRequest

	return &this
}

func (o AppRoleLookUpSecretIdByAccessorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id_accessor"] = o.SecretIdAccessor

	return json.Marshal(toSerialize)
}
