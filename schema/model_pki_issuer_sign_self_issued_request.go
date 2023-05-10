// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuerSignSelfIssuedRequest struct for PkiIssuerSignSelfIssuedRequest
type PkiIssuerSignSelfIssuedRequest struct {
	// PEM-format self-issued certificate to be signed.
	Certificate string `json:"certificate,omitempty"`

	// If true, require the public key algorithm of the signer to match that of the self issued certificate.
	RequireMatchingCertificateAlgorithms bool `json:"require_matching_certificate_algorithms,omitempty"`
}

// NewPkiIssuerSignSelfIssuedRequestWithDefaults instantiates a new PkiIssuerSignSelfIssuedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerSignSelfIssuedRequestWithDefaults() *PkiIssuerSignSelfIssuedRequest {
	var this PkiIssuerSignSelfIssuedRequest

	this.RequireMatchingCertificateAlgorithms = false

	return &this
}
