// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuersRotateRootResponse struct for PkiIssuersRotateRootResponse
type PkiIssuersRotateRootResponse struct {
	// The generated self-signed CA certificate.
	Certificate string `json:"certificate"`

	// The expiration of the given.
	Expiration string `json:"expiration"`

	// The ID of the issuer
	IssuerId string `json:"issuer_id"`

	// The name of the issuer.
	IssuerName string `json:"issuer_name"`

	// The issuing certificate authority.
	IssuingCa string `json:"issuing_ca"`

	// The ID of the key.
	KeyId string `json:"key_id"`

	// The key name if given.
	KeyName string `json:"key_name"`

	// The private key if exported was specified.
	PrivateKey string `json:"private_key"`

	// The requested Subject's named serial number.
	SerialNumber string `json:"serial_number"`
}

// NewPkiIssuersRotateRootResponseWithDefaults instantiates a new PkiIssuersRotateRootResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuersRotateRootResponseWithDefaults() *PkiIssuersRotateRootResponse {
	var this PkiIssuersRotateRootResponse

	return &this
}

func (o PkiIssuersRotateRootResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["certificate"] = o.Certificate
	toSerialize["expiration"] = o.Expiration
	toSerialize["issuer_id"] = o.IssuerId
	toSerialize["issuer_name"] = o.IssuerName
	toSerialize["issuing_ca"] = o.IssuingCa
	toSerialize["key_id"] = o.KeyId
	toSerialize["key_name"] = o.KeyName
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["serial_number"] = o.SerialNumber

	return json.Marshal(toSerialize)
}
