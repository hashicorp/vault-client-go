// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiReadCertDeltaCrlResponse struct for PkiReadCertDeltaCrlResponse
type PkiReadCertDeltaCrlResponse struct {
	// Issuing CA Chain
	CaChain []string `json:"ca_chain"`

	// Certificate
	Certificate string `json:"certificate"`

	// ID of the issuer
	IssuerId string `json:"issuer_id"`

	// Revocation time
	RevocationTime string `json:"revocation_time"`

	// Revocation time RFC 3339 formatted
	RevocationTimeRfc3339 string `json:"revocation_time_rfc3339"`
}

// NewPkiReadCertDeltaCrlResponseWithDefaults instantiates a new PkiReadCertDeltaCrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReadCertDeltaCrlResponseWithDefaults() *PkiReadCertDeltaCrlResponse {
	var this PkiReadCertDeltaCrlResponse

	return &this
}

func (o PkiReadCertDeltaCrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ca_chain"] = o.CaChain
	toSerialize["certificate"] = o.Certificate
	toSerialize["issuer_id"] = o.IssuerId
	toSerialize["revocation_time"] = o.RevocationTime
	toSerialize["revocation_time_rfc3339"] = o.RevocationTimeRfc3339

	return json.Marshal(toSerialize)
}
