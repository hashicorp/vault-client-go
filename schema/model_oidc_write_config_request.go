// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteConfigRequest struct for OIDCWriteConfigRequest
type OIDCWriteConfigRequest struct {
	// Issuer URL to be used in the iss claim of the token. If not set, Vault's app_addr will be used.
	Issuer string `json:"issuer"`
}

// NewOIDCWriteConfigRequestWithDefaults instantiates a new OIDCWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteConfigRequestWithDefaults() *OIDCWriteConfigRequest {
	var this OIDCWriteConfigRequest

	return &this
}

func (o OIDCWriteConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
