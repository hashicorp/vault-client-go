// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiReadIssuerJsonResponse struct for PkiReadIssuerJsonResponse
type PkiReadIssuerJsonResponse struct {
	// CA Chain
	CaChain []string `json:"ca_chain"`

	// Certificate
	Certificate string `json:"certificate"`

	// Issuer Id
	IssuerId string `json:"issuer_id"`

	// Issuer Name
	IssuerName string `json:"issuer_name"`
}

// NewPkiReadIssuerJsonResponseWithDefaults instantiates a new PkiReadIssuerJsonResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReadIssuerJsonResponseWithDefaults() *PkiReadIssuerJsonResponse {
	var this PkiReadIssuerJsonResponse

	return &this
}

func (o PkiReadIssuerJsonResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ca_chain"] = o.CaChain
	toSerialize["certificate"] = o.Certificate
	toSerialize["issuer_id"] = o.IssuerId
	toSerialize["issuer_name"] = o.IssuerName

	return json.Marshal(toSerialize)
}
