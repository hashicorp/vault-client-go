// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteLoggersRequest struct for WriteLoggersRequest
type WriteLoggersRequest struct {
	// Log verbosity level. Supported values (in order of detail) are \"trace\", \"debug\", \"info\", \"warn\", and \"error\".
	Level string `json:"level"`
}

// NewWriteLoggersRequestWithDefaults instantiates a new WriteLoggersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteLoggersRequestWithDefaults() *WriteLoggersRequest {
	var this WriteLoggersRequest

	return &this
}
