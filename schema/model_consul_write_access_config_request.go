// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// ConsulWriteAccessConfigRequest struct for ConsulWriteAccessConfigRequest
type ConsulWriteAccessConfigRequest struct {
	// Consul server address
	Address string `json:"address"`

	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert string `json:"ca_cert"`

	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	ClientCert string `json:"client_cert"`

	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	ClientKey string `json:"client_key"`

	// URI scheme for the Consul address
	Scheme string `json:"scheme"`

	// Token for API calls
	Token string `json:"token"`
}

// NewConsulWriteAccessConfigRequestWithDefaults instantiates a new ConsulWriteAccessConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulWriteAccessConfigRequestWithDefaults() *ConsulWriteAccessConfigRequest {
	var this ConsulWriteAccessConfigRequest

	this.Scheme = "http"

	return &this
}

func (o ConsulWriteAccessConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
