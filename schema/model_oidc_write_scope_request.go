// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteScopeRequest struct for OIDCWriteScopeRequest
type OIDCWriteScopeRequest struct {
	// The description of the scope
	Description string `json:"description"`

	// The template string to use for the scope. This may be in string-ified JSON or base64 format.
	Template string `json:"template"`
}

// NewOIDCWriteScopeRequestWithDefaults instantiates a new OIDCWriteScopeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteScopeRequestWithDefaults() *OIDCWriteScopeRequest {
	var this OIDCWriteScopeRequest

	return &this
}

func (o OIDCWriteScopeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
