// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiImportKeyRequest struct for PkiImportKeyRequest
type PkiImportKeyRequest struct {
	// Optional name to be used for this key
	KeyName string `json:"key_name,omitempty"`

	// PEM-format, unencrypted secret key
	PemBundle string `json:"pem_bundle,omitempty"`
}
