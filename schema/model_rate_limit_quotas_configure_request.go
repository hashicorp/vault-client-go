// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RateLimitQuotasConfigureRequest struct for RateLimitQuotasConfigureRequest
type RateLimitQuotasConfigureRequest struct {
	// If set, starts audit logging of requests that get rejected due to rate limit quota rule violations.
	EnableRateLimitAuditLogging bool `json:"enable_rate_limit_audit_logging,omitempty"`

	// If set, additional rate limit quota HTTP headers will be added to responses.
	EnableRateLimitResponseHeaders bool `json:"enable_rate_limit_response_headers,omitempty"`

	// Specifies the list of exempt paths from all rate limit quotas. If empty no paths will be exempt.
	RateLimitExemptPaths []string `json:"rate_limit_exempt_paths,omitempty"`
}
