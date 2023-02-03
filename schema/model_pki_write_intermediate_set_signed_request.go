// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PKIWriteIntermediateSetSignedRequest struct for PKIWriteIntermediateSetSignedRequest
type PKIWriteIntermediateSetSignedRequest struct {
	// PEM-format certificate. This must be a CA certificate with a public key matching the previously-generated key from the generation endpoint. Additional parent CAs may be optionally appended to the bundle.
	Certificate string `json:"certificate"`
}

// NewPKIWriteIntermediateSetSignedRequestWithDefaults instantiates a new PKIWriteIntermediateSetSignedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteIntermediateSetSignedRequestWithDefaults() *PKIWriteIntermediateSetSignedRequest {
	var this PKIWriteIntermediateSetSignedRequest

	return &this
}
