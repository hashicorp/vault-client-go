// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MfaCreateDuoMethodRequest struct for MfaCreateDuoMethodRequest
type MfaCreateDuoMethodRequest struct {
	// API host name for Duo.
	ApiHostname string `json:"api_hostname,omitempty"`

	// Integration key for Duo.
	IntegrationKey string `json:"integration_key,omitempty"`

	// The unique name identifier for this MFA method.
	MethodName string `json:"method_name,omitempty"`

	// Push information for Duo.
	PushInfo string `json:"push_info,omitempty"`

	// Secret key for Duo.
	SecretKey string `json:"secret_key,omitempty"`

	// If true, the user is reminded to use the passcode upon MFA validation. This option does not enforce using the passcode. Defaults to false.
	UsePasscode bool `json:"use_passcode,omitempty"`

	// A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \"{{alias.name}}@example.com\". Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias's name field will be used as-is.
	UsernameFormat string `json:"username_format,omitempty"`
}
