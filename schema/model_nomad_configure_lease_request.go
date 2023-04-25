// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// NomadConfigureLeaseRequest struct for NomadConfigureLeaseRequest
type NomadConfigureLeaseRequest struct {
	// Duration after which the issued token should not be allowed to be renewed
	MaxTtl int32 `json:"max_ttl"`

	// Duration before which the issued token needs renewal
	Ttl int32 `json:"ttl"`
}

// NewNomadConfigureLeaseRequestWithDefaults instantiates a new NomadConfigureLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNomadConfigureLeaseRequestWithDefaults() *NomadConfigureLeaseRequest {
	var this NomadConfigureLeaseRequest

	return &this
}

func (o NomadConfigureLeaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
