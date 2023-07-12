// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MfaUpdatePingIdMethodRequest struct for MfaUpdatePingIdMethodRequest
type MfaUpdatePingIdMethodRequest struct {
	// The unique name identifier for this MFA method.
	MethodName string `json:"method_name,omitempty"`

	// The settings file provided by Ping, Base64-encoded. This must be a settings file suitable for third-party clients, not the PingID SDK or PingFederate.
	SettingsFileBase64 string `json:"settings_file_base64,omitempty"`

	// A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \"{{alias.name}}@example.com\". Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias's name field will be used as-is.
	UsernameFormat string `json:"username_format,omitempty"`
}
