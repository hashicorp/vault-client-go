// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIRevokeRequest struct for PKIRevokeRequest
type PKIRevokeRequest struct {
	// Certificate to revoke in PEM format; must be signed by an issuer in this mount.
	Certificate string `json:"certificate"`

	// Certificate serial number, in colon- or hyphen-separated octal
	SerialNumber string `json:"serial_number"`
}

// NewPKIRevokeRequestWithDefaults instantiates a new PKIRevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIRevokeRequestWithDefaults() *PKIRevokeRequest {
	var this PKIRevokeRequest

	return &this
}

func (o PKIRevokeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
