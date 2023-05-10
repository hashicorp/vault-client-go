// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadPeriodResponse struct for AppRoleReadPeriodResponse
type AppRoleReadPeriodResponse struct {
	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period int32 `json:"period"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod int32 `json:"token_period"`
}

// NewAppRoleReadPeriodResponseWithDefaults instantiates a new AppRoleReadPeriodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadPeriodResponseWithDefaults() *AppRoleReadPeriodResponse {
	var this AppRoleReadPeriodResponse

	return &this
}

func (o AppRoleReadPeriodResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["period"] = o.Period
	toSerialize["token_period"] = o.TokenPeriod

	return json.Marshal(toSerialize)
}
