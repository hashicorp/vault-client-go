// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIWriteCAConfigRequest struct for PKIWriteCAConfigRequest
type PKIWriteCAConfigRequest struct {
	// PEM-format, concatenated unencrypted secret key and certificate.
	PemBundle string `json:"pem_bundle"`
}

// NewPKIWriteCAConfigRequestWithDefaults instantiates a new PKIWriteCAConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteCAConfigRequestWithDefaults() *PKIWriteCAConfigRequest {
	var this PKIWriteCAConfigRequest

	return &this
}

func (o PKIWriteCAConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["pem_bundle"] = o.PemBundle

	return json.Marshal(toSerialize)
}
