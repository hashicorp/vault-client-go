// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CertWriteCrlRequest struct for CertWriteCrlRequest
type CertWriteCrlRequest struct {
	// The public CRL that should be trusted to attest to certificates' validity statuses. May be DER or PEM encoded. Note: the expiration time is ignored; if the CRL is no longer valid, delete it using the same name as specified here.
	Crl string `json:"crl"`

	// The URL of a CRL distribution point. Only one of 'crl' or 'url' parameters should be specified.
	Url string `json:"url"`
}

// NewCertWriteCrlRequestWithDefaults instantiates a new CertWriteCrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertWriteCrlRequestWithDefaults() *CertWriteCrlRequest {
	var this CertWriteCrlRequest

	return &this
}

func (o CertWriteCrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["crl"] = o.Crl
	toSerialize["url"] = o.Url

	return json.Marshal(toSerialize)
}
