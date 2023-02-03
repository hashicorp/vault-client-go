// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCIntrospectRequest struct for OIDCIntrospectRequest
type OIDCIntrospectRequest struct {
	// Optional client_id to verify
	ClientId string `json:"client_id"`

	// Token to verify
	Token string `json:"token"`
}

// NewOIDCIntrospectRequestWithDefaults instantiates a new OIDCIntrospectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCIntrospectRequestWithDefaults() *OIDCIntrospectRequest {
	var this OIDCIntrospectRequest

	return &this
}
