// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MountsTuneConfigurationParametersRequest struct for MountsTuneConfigurationParametersRequest
type MountsTuneConfigurationParametersRequest struct {
	AllowedManagedKeys []string `json:"allowed_managed_keys,omitempty"`

	// A list of headers to whitelist and allow a plugin to set on responses.
	AllowedResponseHeaders []string `json:"allowed_response_headers,omitempty"`

	// The list of keys in the request data object that will not be HMAC'ed by audit devices.
	AuditNonHmacRequestKeys []string `json:"audit_non_hmac_request_keys,omitempty"`

	// The list of keys in the response data object that will not be HMAC'ed by audit devices.
	AuditNonHmacResponseKeys []string `json:"audit_non_hmac_response_keys,omitempty"`

	// The default lease TTL for this mount.
	DefaultLeaseTtl string `json:"default_lease_ttl,omitempty"`

	// User-friendly description for this credential backend.
	Description string `json:"description,omitempty"`

	// Determines the visibility of the mount in the UI-specific listing endpoint. Accepted value are 'unauth' and 'hidden', with the empty default ('') behaving like 'hidden'.
	ListingVisibility string `json:"listing_visibility,omitempty"`

	// The max lease TTL for this mount.
	MaxLeaseTtl string `json:"max_lease_ttl,omitempty"`

	// The options to pass into the backend. Should be a json object with string keys and values.
	Options map[string]interface{} `json:"options,omitempty"`

	// A list of headers to whitelist and pass from the request to the plugin.
	PassthroughRequestHeaders []string `json:"passthrough_request_headers,omitempty"`

	// The semantic version of the plugin to use, or image tag if oci_image is provided.
	PluginVersion string `json:"plugin_version,omitempty"`

	// The type of token to issue (service or batch).
	TokenType string `json:"token_type,omitempty"`

	// The user lockout configuration to pass into the backend. Should be a json object with string keys and values.
	UserLockoutConfig map[string]interface{} `json:"user_lockout_config,omitempty"`
}
