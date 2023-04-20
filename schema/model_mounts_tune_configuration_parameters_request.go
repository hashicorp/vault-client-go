// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MountsTuneConfigurationParametersRequest struct for MountsTuneConfigurationParametersRequest
type MountsTuneConfigurationParametersRequest struct {
	AllowedManagedKeys []string `json:"allowed_managed_keys"`

	// A list of headers to whitelist and allow a plugin to set on responses.
	AllowedResponseHeaders []string `json:"allowed_response_headers"`

	// The list of keys in the request data object that will not be HMAC'ed by audit devices.
	AuditNonHmacRequestKeys []string `json:"audit_non_hmac_request_keys"`

	// The list of keys in the response data object that will not be HMAC'ed by audit devices.
	AuditNonHmacResponseKeys []string `json:"audit_non_hmac_response_keys"`

	// The default lease TTL for this mount.
	DefaultLeaseTtl string `json:"default_lease_ttl"`

	// User-friendly description for this credential backend.
	Description string `json:"description"`

	// Determines the visibility of the mount in the UI-specific listing endpoint. Accepted value are 'unauth' and 'hidden', with the empty default ('') behaving like 'hidden'.
	ListingVisibility string `json:"listing_visibility"`

	// The max lease TTL for this mount.
	MaxLeaseTtl string `json:"max_lease_ttl"`

	// The options to pass into the backend. Should be a json object with string keys and values.
	Options map[string]interface{} `json:"options"`

	// A list of headers to whitelist and pass from the request to the plugin.
	PassthroughRequestHeaders []string `json:"passthrough_request_headers"`

	// The semantic version of the plugin to use.
	PluginVersion string `json:"plugin_version"`

	// The type of token to issue (service or batch).
	TokenType string `json:"token_type"`

	// The user lockout configuration to pass into the backend. Should be a json object with string keys and values.
	UserLockoutConfig map[string]interface{} `json:"user_lockout_config"`
}

// NewMountsTuneConfigurationParametersRequestWithDefaults instantiates a new MountsTuneConfigurationParametersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountsTuneConfigurationParametersRequestWithDefaults() *MountsTuneConfigurationParametersRequest {
	var this MountsTuneConfigurationParametersRequest

	return &this
}

func (o MountsTuneConfigurationParametersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allowed_managed_keys"] = o.AllowedManagedKeys
	toSerialize["allowed_response_headers"] = o.AllowedResponseHeaders
	toSerialize["audit_non_hmac_request_keys"] = o.AuditNonHmacRequestKeys
	toSerialize["audit_non_hmac_response_keys"] = o.AuditNonHmacResponseKeys
	toSerialize["default_lease_ttl"] = o.DefaultLeaseTtl
	toSerialize["description"] = o.Description
	toSerialize["listing_visibility"] = o.ListingVisibility
	toSerialize["max_lease_ttl"] = o.MaxLeaseTtl
	toSerialize["options"] = o.Options
	toSerialize["passthrough_request_headers"] = o.PassthroughRequestHeaders
	toSerialize["plugin_version"] = o.PluginVersion
	toSerialize["token_type"] = o.TokenType
	toSerialize["user_lockout_config"] = o.UserLockoutConfig

	return json.Marshal(toSerialize)
}
