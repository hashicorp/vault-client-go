// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MFAMethodWriteDuoRequest struct for MFAMethodWriteDuoRequest
type MFAMethodWriteDuoRequest struct {
	// API host name for Duo.
	ApiHostname string `json:"api_hostname"`

	// Integration key for Duo.
	IntegrationKey string `json:"integration_key"`

	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`

	// Push information for Duo.
	PushInfo string `json:"push_info"`

	// Secret key for Duo.
	SecretKey string `json:"secret_key"`

	// If true, the user is reminded to use the passcode upon MFA validation. This option does not enforce using the passcode. Defaults to false.
	UsePasscode bool `json:"use_passcode"`

	// A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \"{{alias.name}}@example.com\". Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias's name field will be used as-is.
	UsernameFormat string `json:"username_format"`
}

// NewMFAMethodWriteDuoRequestWithDefaults instantiates a new MFAMethodWriteDuoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAMethodWriteDuoRequestWithDefaults() *MFAMethodWriteDuoRequest {
	var this MFAMethodWriteDuoRequest

	return &this
}

func (o MFAMethodWriteDuoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
