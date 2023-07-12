// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AliCloudConfigureRequest struct for AliCloudConfigureRequest
type AliCloudConfigureRequest struct {
	// Access key with appropriate permissions.
	AccessKey string `json:"access_key,omitempty"`

	// Secret key with appropriate permissions.
	SecretKey string `json:"secret_key,omitempty"`
}
