// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CertificatesWriteCRLRequest struct for CertificatesWriteCRLRequest
type CertificatesWriteCRLRequest struct {
	// The public CRL that should be trusted to attest to certificates' validity statuses. May be DER or PEM encoded. Note: the expiration time is ignored; if the CRL is no longer valid, delete it using the same name as specified here.
	Crl string `json:"crl"`

	// The URL of a CRL distribution point. Only one of 'crl' or 'url' parameters should be specified.
	Url string `json:"url"`
}

// NewCertificatesWriteCRLRequestWithDefaults instantiates a new CertificatesWriteCRLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatesWriteCRLRequestWithDefaults() *CertificatesWriteCRLRequest {
	var this CertificatesWriteCRLRequest

	return &this
}
