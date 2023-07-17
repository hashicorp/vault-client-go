// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RateLimitQuotasReadConfigurationResponse struct for RateLimitQuotasReadConfigurationResponse
type RateLimitQuotasReadConfigurationResponse struct {
	EnableRateLimitAuditLogging bool `json:"enable_rate_limit_audit_logging,omitempty"`

	EnableRateLimitResponseHeaders bool `json:"enable_rate_limit_response_headers,omitempty"`

	RateLimitExemptPaths []string `json:"rate_limit_exempt_paths,omitempty"`
}
