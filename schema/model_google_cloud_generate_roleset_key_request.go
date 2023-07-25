// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudGenerateRolesetKeyRequest struct for GoogleCloudGenerateRolesetKeyRequest
type GoogleCloudGenerateRolesetKeyRequest struct {
	// Private key algorithm for service account key - defaults to KEY_ALG_RSA_2048\"
	KeyAlgorithm string `json:"key_algorithm,omitempty"`

	// Private key type for service account key - defaults to TYPE_GOOGLE_CREDENTIALS_FILE\"
	KeyType string `json:"key_type,omitempty"`

	// Lifetime of the service account key
	Ttl string `json:"ttl,omitempty"`
}
