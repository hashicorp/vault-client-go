// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RabbitMQWriteLeaseConfigRequest struct for RabbitMQWriteLeaseConfigRequest
type RabbitMQWriteLeaseConfigRequest struct {
	// Duration after which the issued credentials should not be allowed to be renewed
	MaxTtl int32 `json:"max_ttl"`

	// Duration before which the issued credentials needs renewal
	Ttl int32 `json:"ttl"`
}

// NewRabbitMQWriteLeaseConfigRequestWithDefaults instantiates a new RabbitMQWriteLeaseConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRabbitMQWriteLeaseConfigRequestWithDefaults() *RabbitMQWriteLeaseConfigRequest {
	var this RabbitMQWriteLeaseConfigRequest

	this.MaxTtl = 0
	this.Ttl = 0

	return &this
}

func (o RabbitMQWriteLeaseConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
