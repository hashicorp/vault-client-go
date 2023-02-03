// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIRootSignSelfIssuedRequest struct for PKIRootSignSelfIssuedRequest
type PKIRootSignSelfIssuedRequest struct {
	// PEM-format self-issued certificate to be signed.
	Certificate string `json:"certificate"`

	// Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer.
	IssuerRef string `json:"issuer_ref"`

	// If true, require the public key algorithm of the signer to match that of the self issued certificate.
	RequireMatchingCertificateAlgorithms bool `json:"require_matching_certificate_algorithms"`
}

// NewPKIRootSignSelfIssuedRequestWithDefaults instantiates a new PKIRootSignSelfIssuedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIRootSignSelfIssuedRequestWithDefaults() *PKIRootSignSelfIssuedRequest {
	var this PKIRootSignSelfIssuedRequest

	this.IssuerRef = "default"
	this.RequireMatchingCertificateAlgorithms = false

	return &this
}

func (o PKIRootSignSelfIssuedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
