// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AliCloudLoginRequest struct for AliCloudLoginRequest
type AliCloudLoginRequest struct {
	// The request headers. This must include the headers over which AliCloud has included a signature.
	IdentityRequestHeaders string `json:"identity_request_headers,omitempty"`

	// Base64-encoded full URL against which to make the AliCloud request.
	IdentityRequestUrl string `json:"identity_request_url,omitempty"`

	// Name of the role against which the login is being attempted. If 'role' is not specified, then the login endpoint looks for a role name in the ARN returned by the GetCallerIdentity request. If a matching role is not found, login fails.
	Role string `json:"role"`
}
