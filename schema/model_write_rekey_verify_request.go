// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteRekeyVerifyRequest struct for WriteRekeyVerifyRequest
type WriteRekeyVerifyRequest struct {
	// Specifies a single unseal share key from the new set of shares.
	Key string `json:"key"`

	// Specifies the nonce of the rekey verification operation.
	Nonce string `json:"nonce"`
}

// NewWriteRekeyVerifyRequestWithDefaults instantiates a new WriteRekeyVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRekeyVerifyRequestWithDefaults() *WriteRekeyVerifyRequest {
	var this WriteRekeyVerifyRequest

	return &this
}

func (o WriteRekeyVerifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
