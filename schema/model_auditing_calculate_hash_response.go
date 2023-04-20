// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AuditingCalculateHashResponse struct for AuditingCalculateHashResponse
type AuditingCalculateHashResponse struct {
	Hash string `json:"hash"`
}

// NewAuditingCalculateHashResponseWithDefaults instantiates a new AuditingCalculateHashResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditingCalculateHashResponseWithDefaults() *AuditingCalculateHashResponse {
	var this AuditingCalculateHashResponse

	return &this
}

func (o AuditingCalculateHashResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["hash"] = o.Hash

	return json.Marshal(toSerialize)
}
