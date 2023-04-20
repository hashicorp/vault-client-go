// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RateLimitQuotasReadConfigurationResponse struct for RateLimitQuotasReadConfigurationResponse
type RateLimitQuotasReadConfigurationResponse struct {
	EnableRateLimitAuditLogging bool `json:"enable_rate_limit_audit_logging"`

	EnableRateLimitResponseHeaders bool `json:"enable_rate_limit_response_headers"`

	RateLimitExemptPaths []string `json:"rate_limit_exempt_paths"`
}

// NewRateLimitQuotasReadConfigurationResponseWithDefaults instantiates a new RateLimitQuotasReadConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitQuotasReadConfigurationResponseWithDefaults() *RateLimitQuotasReadConfigurationResponse {
	var this RateLimitQuotasReadConfigurationResponse

	return &this
}

func (o RateLimitQuotasReadConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["enable_rate_limit_audit_logging"] = o.EnableRateLimitAuditLogging
	toSerialize["enable_rate_limit_response_headers"] = o.EnableRateLimitResponseHeaders
	toSerialize["rate_limit_exempt_paths"] = o.RateLimitExemptPaths

	return json.Marshal(toSerialize)
}
