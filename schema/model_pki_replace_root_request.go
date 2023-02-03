// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIReplaceRootRequest struct for PKIReplaceRootRequest
type PKIReplaceRootRequest struct {
	// Reference (name or identifier) to the default issuer.
	Default string `json:"default"`
}

// NewPKIReplaceRootRequestWithDefaults instantiates a new PKIReplaceRootRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIReplaceRootRequestWithDefaults() *PKIReplaceRootRequest {
	var this PKIReplaceRootRequest

	this.Default = "next"

	return &this
}

func (o PKIReplaceRootRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
