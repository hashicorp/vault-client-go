// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIIssuerResignCRLsRequest struct for PKIIssuerResignCRLsRequest
type PKIIssuerResignCRLsRequest struct {
	// The sequence number to be written within the CRL Number extension.
	CrlNumber int32 `json:"crl_number"`

	// A list of PEM encoded CRLs to combine, originally signed by the requested issuer.
	Crls []string `json:"crls"`

	// Using a zero or greater value specifies the base CRL revision number to encode within a Delta CRL indicator extension, otherwise the extension will not be added.
	DeltaCrlBaseNumber int32 `json:"delta_crl_base_number"`

	// The format of the combined CRL, can be \"pem\" or \"der\". If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format string `json:"format"`

	// The amount of time the generated CRL should be valid; defaults to 72 hours.
	NextUpdate string `json:"next_update"`
}

// NewPKIIssuerResignCRLsRequestWithDefaults instantiates a new PKIIssuerResignCRLsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIIssuerResignCRLsRequestWithDefaults() *PKIIssuerResignCRLsRequest {
	var this PKIIssuerResignCRLsRequest

	this.DeltaCrlBaseNumber = -1
	this.Format = "pem"
	this.NextUpdate = "72h"

	return &this
}

func (o PKIIssuerResignCRLsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
