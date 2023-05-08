// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AuditingListRequestHeadersResponse struct for AuditingListRequestHeadersResponse
type AuditingListRequestHeadersResponse struct {
	Headers map[string]interface{} `json:"headers,omitempty"`
}

// NewAuditingListRequestHeadersResponseWithDefaults instantiates a new AuditingListRequestHeadersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditingListRequestHeadersResponseWithDefaults() *AuditingListRequestHeadersResponse {
	var this AuditingListRequestHeadersResponse

	return &this
}
