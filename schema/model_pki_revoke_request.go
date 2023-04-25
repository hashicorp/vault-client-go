// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiRevokeRequest struct for PkiRevokeRequest
type PkiRevokeRequest struct {
	// Certificate to revoke in PEM format; must be signed by an issuer in this mount.
	Certificate string `json:"certificate"`

	// Certificate serial number, in colon- or hyphen-separated octal
	SerialNumber string `json:"serial_number"`
}

// NewPkiRevokeRequestWithDefaults instantiates a new PkiRevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRevokeRequestWithDefaults() *PkiRevokeRequest {
	var this PkiRevokeRequest

	return &this
}

func (o PkiRevokeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["certificate"] = o.Certificate
	toSerialize["serial_number"] = o.SerialNumber

	return json.Marshal(toSerialize)
}
