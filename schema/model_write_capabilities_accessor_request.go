// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// WriteCapabilitiesAccessorRequest struct for WriteCapabilitiesAccessorRequest
type WriteCapabilitiesAccessorRequest struct {
	// Accessor of the token for which capabilities are being queried.
	Accessor string `json:"accessor"`

	// Use 'paths' instead.
	// Deprecated
	Path []string `json:"path"`

	// Paths on which capabilities are being queried.
	Paths []string `json:"paths"`
}

// NewWriteCapabilitiesAccessorRequestWithDefaults instantiates a new WriteCapabilitiesAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteCapabilitiesAccessorRequestWithDefaults() *WriteCapabilitiesAccessorRequest {
	var this WriteCapabilitiesAccessorRequest

	return &this
}
