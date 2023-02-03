// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteConfigCORSRequest struct for WriteConfigCORSRequest
type WriteConfigCORSRequest struct {
	// A comma-separated string or array of strings indicating headers that are allowed on cross-origin requests.
	AllowedHeaders []string `json:"allowed_headers"`

	// A comma-separated string or array of strings indicating origins that may make cross-origin requests.
	AllowedOrigins []string `json:"allowed_origins"`

	// Enables or disables CORS headers on requests.
	Enable bool `json:"enable"`
}

// NewWriteConfigCORSRequestWithDefaults instantiates a new WriteConfigCORSRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteConfigCORSRequestWithDefaults() *WriteConfigCORSRequest {
	var this WriteConfigCORSRequest

	return &this
}

func (o WriteConfigCORSRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
