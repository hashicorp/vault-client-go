// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIWriteCertsRequest struct for PKIWriteCertsRequest
type PKIWriteCertsRequest struct {
	// PEM-format, concatenated unencrypted secret-key (optional) and certificates.
	PemBundle string `json:"pem_bundle"`
}

// NewPKIWriteCertsRequestWithDefaults instantiates a new PKIWriteCertsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteCertsRequestWithDefaults() *PKIWriteCertsRequest {
	var this PKIWriteCertsRequest

	return &this
}

func (o PKIWriteCertsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
