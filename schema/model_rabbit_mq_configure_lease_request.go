// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RabbitMqConfigureLeaseRequest struct for RabbitMqConfigureLeaseRequest
type RabbitMqConfigureLeaseRequest struct {
	// Duration after which the issued credentials should not be allowed to be renewed
	MaxTtl string `json:"max_ttl,omitempty"`

	// Duration before which the issued credentials needs renewal
	Ttl string `json:"ttl,omitempty"`
}
