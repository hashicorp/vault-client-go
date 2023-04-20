// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuerSignSelfIssuedRequest struct for PkiIssuerSignSelfIssuedRequest
type PkiIssuerSignSelfIssuedRequest struct {
	// PEM-format self-issued certificate to be signed.
	Certificate string `json:"certificate"`

	// If true, require the public key algorithm of the signer to match that of the self issued certificate.
	RequireMatchingCertificateAlgorithms bool `json:"require_matching_certificate_algorithms"`
}

// NewPkiIssuerSignSelfIssuedRequestWithDefaults instantiates a new PkiIssuerSignSelfIssuedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerSignSelfIssuedRequestWithDefaults() *PkiIssuerSignSelfIssuedRequest {
	var this PkiIssuerSignSelfIssuedRequest

	this.RequireMatchingCertificateAlgorithms = false

	return &this
}

func (o PkiIssuerSignSelfIssuedRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["certificate"] = o.Certificate
	toSerialize["require_matching_certificate_algorithms"] = o.RequireMatchingCertificateAlgorithms

	return json.Marshal(toSerialize)
}
