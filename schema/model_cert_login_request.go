// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CertLoginRequest struct for CertLoginRequest
type CertLoginRequest struct {
	// The name of the certificate role to authenticate against.
	Name string `json:"name"`
}

// NewCertLoginRequestWithDefaults instantiates a new CertLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertLoginRequestWithDefaults() *CertLoginRequest {
	var this CertLoginRequest

	return &this
}

func (o CertLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["name"] = o.Name

	return json.Marshal(toSerialize)
}
