// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteAssignmentRequest struct for OIDCWriteAssignmentRequest
type OIDCWriteAssignmentRequest struct {
	// Comma separated string or array of identity entity IDs
	EntityIds []string `json:"entity_ids"`

	// Comma separated string or array of identity group IDs
	GroupIds []string `json:"group_ids"`
}

// NewOIDCWriteAssignmentRequestWithDefaults instantiates a new OIDCWriteAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteAssignmentRequestWithDefaults() *OIDCWriteAssignmentRequest {
	var this OIDCWriteAssignmentRequest

	return &this
}
