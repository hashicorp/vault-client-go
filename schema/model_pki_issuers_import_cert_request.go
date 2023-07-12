// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuersImportCertRequest struct for PkiIssuersImportCertRequest
type PkiIssuersImportCertRequest struct {
	// PEM-format, concatenated unencrypted secret-key (optional) and certificates.
	PemBundle string `json:"pem_bundle,omitempty"`
}
