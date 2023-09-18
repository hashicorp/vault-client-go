// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MfaCreateOktaMethodRequest struct for MfaCreateOktaMethodRequest
type MfaCreateOktaMethodRequest struct {
	// Okta API key.
	ApiToken string `json:"api_token,omitempty"`

	// The base domain to use for the Okta API. When not specified in the configuration, \"okta.com\" is used.
	BaseUrl string `json:"base_url,omitempty"`

	// The unique name identifier for this MFA method.
	MethodName string `json:"method_name,omitempty"`

	// Name of the organization to be used in the Okta API.
	OrgName string `json:"org_name,omitempty"`

	// If true, the username will only match the primary email for the account. Defaults to false.
	PrimaryEmail bool `json:"primary_email,omitempty"`

	// (DEPRECATED) Use base_url instead.
	Production bool `json:"production,omitempty"`

	// A template string for mapping Identity names to MFA method names. Values to substitute should be placed in {{}}. For example, \"{{entity.name}}@example.com\". If blank, the Entity's name field will be used as-is.
	UsernameFormat string `json:"username_format,omitempty"`
}
