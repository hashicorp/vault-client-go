// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CertificatesLoginRequest struct for CertificatesLoginRequest
type CertificatesLoginRequest struct {
	// The name of the certificate role to authenticate against.
	Name string `json:"name"`
}

// NewCertificatesLoginRequestWithDefaults instantiates a new CertificatesLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatesLoginRequestWithDefaults() *CertificatesLoginRequest {
	var this CertificatesLoginRequest

	return &this
}

func (o CertificatesLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
