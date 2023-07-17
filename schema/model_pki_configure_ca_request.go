// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiConfigureCaRequest struct for PkiConfigureCaRequest
type PkiConfigureCaRequest struct {
	// PEM-format, concatenated unencrypted secret key and certificate.
	PemBundle string `json:"pem_bundle,omitempty"`
}
