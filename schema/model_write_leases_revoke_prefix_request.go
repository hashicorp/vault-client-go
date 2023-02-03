// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteLeasesRevokePrefixRequest struct for WriteLeasesRevokePrefixRequest
type WriteLeasesRevokePrefixRequest struct {
	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`
}

// NewWriteLeasesRevokePrefixRequestWithDefaults instantiates a new WriteLeasesRevokePrefixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteLeasesRevokePrefixRequestWithDefaults() *WriteLeasesRevokePrefixRequest {
	var this WriteLeasesRevokePrefixRequest

	this.Sync = true

	return &this
}
