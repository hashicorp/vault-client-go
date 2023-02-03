// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitGenerateRandomRequest struct for TransitGenerateRandomRequest
type TransitGenerateRandomRequest struct {
	// The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits).
	Bytes int32 `json:"bytes"`

	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"base64\".
	Format string `json:"format"`

	// Which system to source random data from, ether \"platform\", \"seal\", or \"all\".
	Source string `json:"source"`

	// The number of bytes to generate (POST URL parameter)
	Urlbytes string `json:"urlbytes"`
}

// NewTransitGenerateRandomRequestWithDefaults instantiates a new TransitGenerateRandomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitGenerateRandomRequestWithDefaults() *TransitGenerateRandomRequest {
	var this TransitGenerateRandomRequest

	this.Bytes = 32
	this.Format = "base64"
	this.Source = "platform"

	return &this
}
