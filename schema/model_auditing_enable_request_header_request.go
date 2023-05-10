// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AuditingEnableRequestHeaderRequest struct for AuditingEnableRequestHeaderRequest
type AuditingEnableRequestHeaderRequest struct {
	Hmac bool `json:"hmac,omitempty"`
}

// NewAuditingEnableRequestHeaderRequestWithDefaults instantiates a new AuditingEnableRequestHeaderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditingEnableRequestHeaderRequestWithDefaults() *AuditingEnableRequestHeaderRequest {
	var this AuditingEnableRequestHeaderRequest

	return &this
}
