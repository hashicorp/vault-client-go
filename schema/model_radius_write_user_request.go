// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RadiusWriteUserRequest struct for RadiusWriteUserRequest
type RadiusWriteUserRequest struct {
	// Comma-separated list of policies associated to the user.
	Policies []string `json:"policies,omitempty"`
}
