// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KubernetesLoginRequest struct for KubernetesLoginRequest
type KubernetesLoginRequest struct {
	// A signed JWT for authenticating a service account. This field is required.
	Jwt string `json:"jwt,omitempty"`

	// Name of the role against which the login is being attempted. This field is required
	Role string `json:"role,omitempty"`
}
