// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIBundleWriteRequest struct for PKIBundleWriteRequest
type PKIBundleWriteRequest struct {
	// PEM-format, concatenated unencrypted secret-key (optional) and certificates.
	PemBundle string `json:"pem_bundle"`
}

// NewPKIBundleWriteRequestWithDefaults instantiates a new PKIBundleWriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIBundleWriteRequestWithDefaults() *PKIBundleWriteRequest {
	var this PKIBundleWriteRequest

	return &this
}

func (o PKIBundleWriteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
