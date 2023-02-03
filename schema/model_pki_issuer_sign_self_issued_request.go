// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIIssuerSignSelfIssuedRequest struct for PKIIssuerSignSelfIssuedRequest
type PKIIssuerSignSelfIssuedRequest struct {
	// PEM-format self-issued certificate to be signed.
	Certificate string `json:"certificate"`

	// If true, require the public key algorithm of the signer to match that of the self issued certificate.
	RequireMatchingCertificateAlgorithms bool `json:"require_matching_certificate_algorithms"`
}

// NewPKIIssuerSignSelfIssuedRequestWithDefaults instantiates a new PKIIssuerSignSelfIssuedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIIssuerSignSelfIssuedRequestWithDefaults() *PKIIssuerSignSelfIssuedRequest {
	var this PKIIssuerSignSelfIssuedRequest

	this.RequireMatchingCertificateAlgorithms = false

	return &this
}
