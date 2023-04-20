// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuerSignRevocationListRequest struct for PkiIssuerSignRevocationListRequest
type PkiIssuerSignRevocationListRequest struct {
	// The sequence number to be written within the CRL Number extension.
	CrlNumber int32 `json:"crl_number"`

	// Using a zero or greater value specifies the base CRL revision number to encode within a Delta CRL indicator extension, otherwise the extension will not be added.
	DeltaCrlBaseNumber int32 `json:"delta_crl_base_number"`

	// A list of maps containing extensions with keys id (string), critical (bool), value (string)
	Extensions []map[string]interface{} `json:"extensions"`

	// The format of the combined CRL, can be \"pem\" or \"der\". If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format string `json:"format"`

	// The amount of time the generated CRL should be valid; defaults to 72 hours.
	NextUpdate string `json:"next_update"`

	// A list of maps containing the keys serial_number (string), revocation_time (string), and extensions (map with keys id (string), critical (bool), value (string))
	RevokedCerts []map[string]interface{} `json:"revoked_certs"`
}

// NewPkiIssuerSignRevocationListRequestWithDefaults instantiates a new PkiIssuerSignRevocationListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerSignRevocationListRequestWithDefaults() *PkiIssuerSignRevocationListRequest {
	var this PkiIssuerSignRevocationListRequest

	this.DeltaCrlBaseNumber = -1
	this.Format = "pem"
	this.NextUpdate = "72h"

	return &this
}

func (o PkiIssuerSignRevocationListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["crl_number"] = o.CrlNumber
	toSerialize["delta_crl_base_number"] = o.DeltaCrlBaseNumber
	toSerialize["extensions"] = o.Extensions
	toSerialize["format"] = o.Format
	toSerialize["next_update"] = o.NextUpdate
	toSerialize["revoked_certs"] = o.RevokedCerts

	return json.Marshal(toSerialize)
}
