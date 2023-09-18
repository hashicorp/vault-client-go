// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitSetCertificateForKeyRequest struct for TransitSetCertificateForKeyRequest
type TransitSetCertificateForKeyRequest struct {
	// PEM encoded certificate chain. It should be composed by one or more concatenated PEM blocks and ordered starting from the end-entity certificate.
	CertificateChain string `json:"certificate_chain"`

	// Optional version of key, 'latest' if not set
	Version int32 `json:"version,omitempty"`
}
