// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiRootSignSelfIssuedRequest struct for PkiRootSignSelfIssuedRequest
type PkiRootSignSelfIssuedRequest struct {
	// PEM-format self-issued certificate to be signed.
	Certificate string `json:"certificate,omitempty"`

	// Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer.
	IssuerRef string `json:"issuer_ref,omitempty"`

	// If true, require the public key algorithm of the signer to match that of the self issued certificate.
	RequireMatchingCertificateAlgorithms bool `json:"require_matching_certificate_algorithms,omitempty"`
}
