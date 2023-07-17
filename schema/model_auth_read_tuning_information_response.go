// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AuthReadTuningInformationResponse struct for AuthReadTuningInformationResponse
type AuthReadTuningInformationResponse struct {
	AllowedManagedKeys []string `json:"allowed_managed_keys,omitempty"`

	AllowedResponseHeaders []string `json:"allowed_response_headers,omitempty"`

	AuditNonHmacRequestKeys []string `json:"audit_non_hmac_request_keys,omitempty"`

	AuditNonHmacResponseKeys []string `json:"audit_non_hmac_response_keys,omitempty"`

	DefaultLeaseTtl int32 `json:"default_lease_ttl,omitempty"`

	Description string `json:"description,omitempty"`

	ExternalEntropyAccess bool `json:"external_entropy_access,omitempty"`

	ForceNoCache bool `json:"force_no_cache,omitempty"`

	ListingVisibility string `json:"listing_visibility,omitempty"`

	MaxLeaseTtl int32 `json:"max_lease_ttl,omitempty"`

	Options map[string]interface{} `json:"options,omitempty"`

	PassthroughRequestHeaders []string `json:"passthrough_request_headers,omitempty"`

	PluginVersion string `json:"plugin_version,omitempty"`

	TokenType string `json:"token_type,omitempty"`

	UserLockoutCounterResetDuration int64 `json:"user_lockout_counter_reset_duration,omitempty"`

	UserLockoutDisable bool `json:"user_lockout_disable,omitempty"`

	UserLockoutDuration int64 `json:"user_lockout_duration,omitempty"`

	UserLockoutThreshold int64 `json:"user_lockout_threshold,omitempty"`
}
