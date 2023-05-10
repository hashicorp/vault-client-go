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

// NewAliCloudConfigureRequestWithDefaults instantiates a new AliCloudConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliCloudConfigureRequestWithDefaults() *AliCloudConfigureRequest {
	var this AliCloudConfigureRequest

	return &this
}
