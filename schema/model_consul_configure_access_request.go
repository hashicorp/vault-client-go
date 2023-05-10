// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// ConsulConfigureAccessRequest struct for ConsulConfigureAccessRequest
type ConsulConfigureAccessRequest struct {
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

// NewConsulConfigureAccessRequestWithDefaults instantiates a new ConsulConfigureAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulConfigureAccessRequestWithDefaults() *ConsulConfigureAccessRequest {
	var this ConsulConfigureAccessRequest

	this.Scheme = "http"

	return &this
}

func (o ConsulConfigureAccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["address"] = o.Address
	toSerialize["ca_cert"] = o.CaCert
	toSerialize["client_cert"] = o.ClientCert
	toSerialize["client_key"] = o.ClientKey
	toSerialize["scheme"] = o.Scheme
	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
