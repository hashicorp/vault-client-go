// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSConfigWriteLeaseRequest struct for AWSConfigWriteLeaseRequest
type AWSConfigWriteLeaseRequest struct {
	// Default lease for roles.
	Lease string `json:"lease"`

	// Maximum time a credential is valid for.
	LeaseMax string `json:"lease_max"`
}

// NewAWSConfigWriteLeaseRequestWithDefaults instantiates a new AWSConfigWriteLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSConfigWriteLeaseRequestWithDefaults() *AWSConfigWriteLeaseRequest {
	var this AWSConfigWriteLeaseRequest

	return &this
}

func (o AWSConfigWriteLeaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["lease"] = o.Lease
	toSerialize["lease_max"] = o.LeaseMax

	return json.Marshal(toSerialize)
}
