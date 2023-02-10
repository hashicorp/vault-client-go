// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AliCloudWriteConfigRequest struct for AliCloudWriteConfigRequest
type AliCloudWriteConfigRequest struct {
	// Access key with appropriate permissions.
	AccessKey string `json:"access_key"`

	// Secret key with appropriate permissions.
	SecretKey string `json:"secret_key"`
}

// NewAliCloudWriteConfigRequestWithDefaults instantiates a new AliCloudWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliCloudWriteConfigRequestWithDefaults() *AliCloudWriteConfigRequest {
	var this AliCloudWriteConfigRequest

	return &this
}

func (o AliCloudWriteConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["access_key"] = o.AccessKey
	toSerialize["secret_key"] = o.SecretKey

	return json.Marshal(toSerialize)
}
