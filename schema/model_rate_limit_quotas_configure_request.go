// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RateLimitQuotasConfigureRequest struct for RateLimitQuotasConfigureRequest
type RateLimitQuotasConfigureRequest struct {
	// If set, starts audit logging of requests that get rejected due to rate limit quota rule violations.
	EnableRateLimitAuditLogging bool `json:"enable_rate_limit_audit_logging"`

	// If set, additional rate limit quota HTTP headers will be added to responses.
	EnableRateLimitResponseHeaders bool `json:"enable_rate_limit_response_headers"`

	// Specifies the list of exempt paths from all rate limit quotas. If empty no paths will be exempt.
	RateLimitExemptPaths []string `json:"rate_limit_exempt_paths"`
}

// NewRateLimitQuotasConfigureRequestWithDefaults instantiates a new RateLimitQuotasConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitQuotasConfigureRequestWithDefaults() *RateLimitQuotasConfigureRequest {
	var this RateLimitQuotasConfigureRequest

	return &this
}

func (o RateLimitQuotasConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["enable_rate_limit_audit_logging"] = o.EnableRateLimitAuditLogging
	toSerialize["enable_rate_limit_response_headers"] = o.EnableRateLimitResponseHeaders
	toSerialize["rate_limit_exempt_paths"] = o.RateLimitExemptPaths

	return json.Marshal(toSerialize)
}
