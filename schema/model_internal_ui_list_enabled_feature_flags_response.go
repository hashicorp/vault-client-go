// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalUiListEnabledFeatureFlagsResponse struct for InternalUiListEnabledFeatureFlagsResponse
type InternalUiListEnabledFeatureFlagsResponse struct {
	FeatureFlags []string `json:"feature_flags,omitempty"`
}

// NewInternalUiListEnabledFeatureFlagsResponseWithDefaults instantiates a new InternalUiListEnabledFeatureFlagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalUiListEnabledFeatureFlagsResponseWithDefaults() *InternalUiListEnabledFeatureFlagsResponse {
	var this InternalUiListEnabledFeatureFlagsResponse

	return &this
}
