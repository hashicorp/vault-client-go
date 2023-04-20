// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MfaConfigureOktaMethodRequest struct for MfaConfigureOktaMethodRequest
type MfaConfigureOktaMethodRequest struct {
	// Okta API key.
	ApiToken string `json:"api_token"`

	// The base domain to use for the Okta API. When not specified in the configuration, \"okta.com\" is used.
	BaseUrl string `json:"base_url"`

	// The unique name identifier for this MFA method.
	MethodName string `json:"method_name"`

	// Name of the organization to be used in the Okta API.
	OrgName string `json:"org_name"`

	// If true, the username will only match the primary email for the account. Defaults to false.
	PrimaryEmail bool `json:"primary_email"`

	// (DEPRECATED) Use base_url instead.
	Production bool `json:"production"`

	// A template string for mapping Identity names to MFA method names. Values to substitute should be placed in {{}}. For example, \"{{entity.name}}@example.com\". If blank, the Entity's name field will be used as-is.
	UsernameFormat string `json:"username_format"`
}

// NewMfaConfigureOktaMethodRequestWithDefaults instantiates a new MfaConfigureOktaMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaConfigureOktaMethodRequestWithDefaults() *MfaConfigureOktaMethodRequest {
	var this MfaConfigureOktaMethodRequest

	return &this
}

func (o MfaConfigureOktaMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["api_token"] = o.ApiToken
	toSerialize["base_url"] = o.BaseUrl
	toSerialize["method_name"] = o.MethodName
	toSerialize["org_name"] = o.OrgName
	toSerialize["primary_email"] = o.PrimaryEmail
	toSerialize["production"] = o.Production
	toSerialize["username_format"] = o.UsernameFormat

	return json.Marshal(toSerialize)
}
