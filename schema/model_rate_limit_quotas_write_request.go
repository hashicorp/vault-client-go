// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RateLimitQuotasWriteRequest struct for RateLimitQuotasWriteRequest
type RateLimitQuotasWriteRequest struct {
	// If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' has elapsed.
	BlockInterval string `json:"block_interval,omitempty"`

	// Whether all child namespaces can inherit this namespace quota.
	Inheritable bool `json:"inheritable,omitempty"`

	// The duration to enforce rate limiting for (default '1s').
	Interval string `json:"interval,omitempty"`

	// Path of the mount or namespace to apply the quota. A blank path configures a global quota. For example namespace1/ adds a quota to a full namespace, namespace1/auth/userpass adds a quota to userpass in namespace1.
	Path string `json:"path,omitempty"`

	// The maximum number of requests in a given interval to be allowed by the quota rule. The 'rate' must be positive.
	Rate float32 `json:"rate,omitempty"`

	// Login role to apply this quota to. Note that when set, path must be configured to a valid auth method with a concept of roles.
	Role string `json:"role,omitempty"`

	// Type of the quota rule.
	Type string `json:"type,omitempty"`
}
