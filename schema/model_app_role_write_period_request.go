// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWritePeriodRequest struct for AppRoleWritePeriodRequest
type AppRoleWritePeriodRequest struct {
	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period int32 `json:"period,omitempty"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod int32 `json:"token_period,omitempty"`
}

// NewAppRoleWritePeriodRequestWithDefaults instantiates a new AppRoleWritePeriodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWritePeriodRequestWithDefaults() *AppRoleWritePeriodRequest {
	var this AppRoleWritePeriodRequest

	return &this
}
