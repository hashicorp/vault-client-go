// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MountsReadTuningInformationResponse struct for MountsReadTuningInformationResponse
type MountsReadTuningInformationResponse struct {
	AllowedManagedKeys []string `json:"allowed_managed_keys,omitempty"`

	// A list of headers to whitelist and allow a plugin to set on responses.
	AllowedResponseHeaders []string `json:"allowed_response_headers,omitempty"`

	AuditNonHmacRequestKeys []string `json:"audit_non_hmac_request_keys,omitempty"`

	AuditNonHmacResponseKeys []string `json:"audit_non_hmac_response_keys,omitempty"`

	// The default lease TTL for this mount.
	DefaultLeaseTtl int32 `json:"default_lease_ttl,omitempty"`

	// User-friendly description for this credential backend.
	Description string `json:"description,omitempty"`

	ExternalEntropyAccess bool `json:"external_entropy_access,omitempty"`

	ForceNoCache bool `json:"force_no_cache,omitempty"`

	ListingVisibility string `json:"listing_visibility,omitempty"`

	// The max lease TTL for this mount.
	MaxLeaseTtl int32 `json:"max_lease_ttl,omitempty"`

	// The options to pass into the backend. Should be a json object with string keys and values.
	Options map[string]interface{} `json:"options,omitempty"`

	PassthroughRequestHeaders []string `json:"passthrough_request_headers,omitempty"`

	// The semantic version of the plugin to use, or image tag if oci_image is provided.
	PluginVersion string `json:"plugin_version,omitempty"`

	// The type of token to issue (service or batch).
	TokenType string `json:"token_type,omitempty"`

	UserLockoutCounterResetDuration int64 `json:"user_lockout_counter_reset_duration,omitempty"`

	UserLockoutDisable bool `json:"user_lockout_disable,omitempty"`

	UserLockoutDuration int64 `json:"user_lockout_duration,omitempty"`

	UserLockoutThreshold int64 `json:"user_lockout_threshold,omitempty"`
}
