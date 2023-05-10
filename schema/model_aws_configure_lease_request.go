// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AwsConfigureLeaseRequest struct for AwsConfigureLeaseRequest
type AwsConfigureLeaseRequest struct {
	// Default lease for roles.
	Lease string `json:"lease"`

	// Maximum time a credential is valid for.
	LeaseMax string `json:"lease_max"`
}

// NewAwsConfigureLeaseRequestWithDefaults instantiates a new AwsConfigureLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsConfigureLeaseRequestWithDefaults() *AwsConfigureLeaseRequest {
	var this AwsConfigureLeaseRequest

	return &this
}

func (o AwsConfigureLeaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["lease"] = o.Lease
	toSerialize["lease_max"] = o.LeaseMax

	return json.Marshal(toSerialize)
}
