// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AuditingCalculateHashResponse struct for AuditingCalculateHashResponse
type AuditingCalculateHashResponse struct {
	Hash string `json:"hash,omitempty"`
}

// NewAuditingCalculateHashResponseWithDefaults instantiates a new AuditingCalculateHashResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditingCalculateHashResponseWithDefaults() *AuditingCalculateHashResponse {
	var this AuditingCalculateHashResponse

	return &this
}
