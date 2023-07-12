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
	MaxTtl string `json:"max_ttl,omitempty"`

	// Default lease for generated keys. If <= 0, will use system default.
	Ttl string `json:"ttl,omitempty"`
}
