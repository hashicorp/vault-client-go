// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudConfigureRequest struct for GoogleCloudConfigureRequest
type GoogleCloudConfigureRequest struct {
	// GCP IAM service account credentials JSON with permissions to create new service accounts and set IAM policies
	Credentials string `json:"credentials,omitempty"`

	// Maximum time a service account key is valid for. If <= 0, will use system default.
	MaxTtl int32 `json:"max_ttl,omitempty"`

	// Default lease for generated keys. If <= 0, will use system default.
	Ttl int32 `json:"ttl,omitempty"`
}

// NewGoogleCloudConfigureRequestWithDefaults instantiates a new GoogleCloudConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudConfigureRequestWithDefaults() *GoogleCloudConfigureRequest {
	var this GoogleCloudConfigureRequest

	return &this
}
