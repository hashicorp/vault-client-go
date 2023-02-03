// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// NomadWriteAccessConfigRequest struct for NomadWriteAccessConfigRequest
type NomadWriteAccessConfigRequest struct {
	// Nomad server address
	Address string `json:"address"`

	// CA certificate to use when verifying Nomad server certificate, must be x509 PEM encoded.
	CaCert string `json:"ca_cert"`

	// Client certificate used for Nomad's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert string `json:"client_cert"`

	// Client key used for Nomad's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey string `json:"client_key"`

	// Max length for name of generated Nomad tokens
	MaxTokenNameLength int32 `json:"max_token_name_length"`

	// Token for API calls
	Token string `json:"token"`
}

// NewNomadWriteAccessConfigRequestWithDefaults instantiates a new NomadWriteAccessConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNomadWriteAccessConfigRequestWithDefaults() *NomadWriteAccessConfigRequest {
	var this NomadWriteAccessConfigRequest

	return &this
}

func (o NomadWriteAccessConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
