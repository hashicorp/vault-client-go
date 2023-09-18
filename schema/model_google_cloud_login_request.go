// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudLoginRequest struct for GoogleCloudLoginRequest
type GoogleCloudLoginRequest struct {
	// A signed JWT. This is either a self-signed service account JWT ('iam' roles only) or a GCE identity metadata token ('iam', 'gce' roles).
	Jwt string `json:"jwt,omitempty"`

	// Name of the role against which the login is being attempted. Required.
	Role string `json:"role,omitempty"`
}
