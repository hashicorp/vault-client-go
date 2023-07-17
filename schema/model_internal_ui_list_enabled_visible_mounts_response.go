// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalUiListEnabledVisibleMountsResponse struct for InternalUiListEnabledVisibleMountsResponse
type InternalUiListEnabledVisibleMountsResponse struct {
	// auth mounts
	Auth map[string]interface{} `json:"auth,omitempty"`

	// secret mounts
	Secret map[string]interface{} `json:"secret,omitempty"`
}
