// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AuditingCalculateHashRequest struct for AuditingCalculateHashRequest
type AuditingCalculateHashRequest struct {
	Input string `json:"input,omitempty"`
}

// NewAuditingCalculateHashRequestWithDefaults instantiates a new AuditingCalculateHashRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditingCalculateHashRequestWithDefaults() *AuditingCalculateHashRequest {
	var this AuditingCalculateHashRequest

	return &this
}
