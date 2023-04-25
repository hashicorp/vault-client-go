// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MountsReadTuningInformationResponse struct for MountsReadTuningInformationResponse
type MountsReadTuningInformationResponse struct {
	AllowedManagedKeys []string `json:"allowed_managed_keys"`

	// A list of headers to whitelist and allow a plugin to set on responses.
	AllowedResponseHeaders []string `json:"allowed_response_headers"`

	AuditNonHmacRequestKeys []string `json:"audit_non_hmac_request_keys"`

	AuditNonHmacResponseKeys []string `json:"audit_non_hmac_response_keys"`

	// The default lease TTL for this mount.
	DefaultLeaseTtl int32 `json:"default_lease_ttl"`

	// User-friendly description for this credential backend.
	Description string `json:"description"`

	ExternalEntropyAccess bool `json:"external_entropy_access"`

	ForceNoCache bool `json:"force_no_cache"`

	ListingVisibility string `json:"listing_visibility"`

	// The max lease TTL for this mount.
	MaxLeaseTtl int32 `json:"max_lease_ttl"`

	// The options to pass into the backend. Should be a json object with string keys and values.
	Options map[string]interface{} `json:"options"`

	PassthroughRequestHeaders []string `json:"passthrough_request_headers"`

	// The semantic version of the plugin to use.
	PluginVersion string `json:"plugin_version"`

	// The type of token to issue (service or batch).
	TokenType string `json:"token_type"`

	UserLockoutCounterResetDuration int64 `json:"user_lockout_counter_reset_duration"`

	UserLockoutDisable bool `json:"user_lockout_disable"`

	UserLockoutDuration int64 `json:"user_lockout_duration"`

	UserLockoutThreshold int64 `json:"user_lockout_threshold"`
}

// NewMountsReadTuningInformationResponseWithDefaults instantiates a new MountsReadTuningInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountsReadTuningInformationResponseWithDefaults() *MountsReadTuningInformationResponse {
	var this MountsReadTuningInformationResponse

	return &this
}

func (o MountsReadTuningInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allowed_managed_keys"] = o.AllowedManagedKeys
	toSerialize["allowed_response_headers"] = o.AllowedResponseHeaders
	toSerialize["audit_non_hmac_request_keys"] = o.AuditNonHmacRequestKeys
	toSerialize["audit_non_hmac_response_keys"] = o.AuditNonHmacResponseKeys
	toSerialize["default_lease_ttl"] = o.DefaultLeaseTtl
	toSerialize["description"] = o.Description
	toSerialize["external_entropy_access"] = o.ExternalEntropyAccess
	toSerialize["force_no_cache"] = o.ForceNoCache
	toSerialize["listing_visibility"] = o.ListingVisibility
	toSerialize["max_lease_ttl"] = o.MaxLeaseTtl
	toSerialize["options"] = o.Options
	toSerialize["passthrough_request_headers"] = o.PassthroughRequestHeaders
	toSerialize["plugin_version"] = o.PluginVersion
	toSerialize["token_type"] = o.TokenType
	toSerialize["user_lockout_counter_reset_duration"] = o.UserLockoutCounterResetDuration
	toSerialize["user_lockout_disable"] = o.UserLockoutDisable
	toSerialize["user_lockout_duration"] = o.UserLockoutDuration
	toSerialize["user_lockout_threshold"] = o.UserLockoutThreshold

	return json.Marshal(toSerialize)
}
