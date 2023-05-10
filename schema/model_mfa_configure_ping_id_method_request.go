// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MfaConfigurePingIdMethodRequest struct for MfaConfigurePingIdMethodRequest
type MfaConfigurePingIdMethodRequest struct {
	// The unique name identifier for this MFA method.
	MethodName string `json:"method_name"`

	// The settings file provided by Ping, Base64-encoded. This must be a settings file suitable for third-party clients, not the PingID SDK or PingFederate.
	SettingsFileBase64 string `json:"settings_file_base64"`

	// A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \"{{alias.name}}@example.com\". Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias's name field will be used as-is.
	UsernameFormat string `json:"username_format"`
}

// NewMfaConfigurePingIdMethodRequestWithDefaults instantiates a new MfaConfigurePingIdMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaConfigurePingIdMethodRequestWithDefaults() *MfaConfigurePingIdMethodRequest {
	var this MfaConfigurePingIdMethodRequest

	return &this
}

func (o MfaConfigurePingIdMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["method_name"] = o.MethodName
	toSerialize["settings_file_base64"] = o.SettingsFileBase64
	toSerialize["username_format"] = o.UsernameFormat

	return json.Marshal(toSerialize)
}
