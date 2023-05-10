// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiRevokeWithKeyRequest struct for PkiRevokeWithKeyRequest
type PkiRevokeWithKeyRequest struct {
	// Certificate to revoke in PEM format; must be signed by an issuer in this mount.
	Certificate string `json:"certificate"`

	// Key to use to verify revocation permission; must be in PEM format.
	PrivateKey string `json:"private_key"`

	// Certificate serial number, in colon- or hyphen-separated octal
	SerialNumber string `json:"serial_number"`
}

// NewPkiRevokeWithKeyRequestWithDefaults instantiates a new PkiRevokeWithKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRevokeWithKeyRequestWithDefaults() *PkiRevokeWithKeyRequest {
	var this PkiRevokeWithKeyRequest

	return &this
}

func (o PkiRevokeWithKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["certificate"] = o.Certificate
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["serial_number"] = o.SerialNumber

	return json.Marshal(toSerialize)
}
