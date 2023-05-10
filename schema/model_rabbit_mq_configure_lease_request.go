// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RabbitMqConfigureLeaseRequest struct for RabbitMqConfigureLeaseRequest
type RabbitMqConfigureLeaseRequest struct {
	// Duration after which the issued credentials should not be allowed to be renewed
	MaxTtl int32 `json:"max_ttl"`

	// Duration before which the issued credentials needs renewal
	Ttl int32 `json:"ttl"`
}

// NewRabbitMqConfigureLeaseRequestWithDefaults instantiates a new RabbitMqConfigureLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRabbitMqConfigureLeaseRequestWithDefaults() *RabbitMqConfigureLeaseRequest {
	var this RabbitMqConfigureLeaseRequest

	this.MaxTtl = 0
	this.Ttl = 0

	return &this
}

func (o RabbitMqConfigureLeaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
