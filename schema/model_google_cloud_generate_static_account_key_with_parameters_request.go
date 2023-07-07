// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudGenerateStaticAccountKeyWithParametersRequest struct for GoogleCloudGenerateStaticAccountKeyWithParametersRequest
type GoogleCloudGenerateStaticAccountKeyWithParametersRequest struct {
	// Private key algorithm for service account key. Defaults to KEY_ALG_RSA_2048.\"
	KeyAlgorithm string `json:"key_algorithm,omitempty"`

	// Private key type for service account key. Defaults to TYPE_GOOGLE_CREDENTIALS_FILE.\"
	KeyType string `json:"key_type,omitempty"`

	// Lifetime of the service account key
	Ttl string `json:"ttl,omitempty"`
}

// NewGoogleCloudGenerateStaticAccountKeyWithParametersRequestWithDefaults instantiates a new GoogleCloudGenerateStaticAccountKeyWithParametersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudGenerateStaticAccountKeyWithParametersRequestWithDefaults() *GoogleCloudGenerateStaticAccountKeyWithParametersRequest {
	var this GoogleCloudGenerateStaticAccountKeyWithParametersRequest

	this.KeyAlgorithm = "KEY_ALG_RSA_2048"
	this.KeyType = "TYPE_GOOGLE_CREDENTIALS_FILE"

	return &this
}
