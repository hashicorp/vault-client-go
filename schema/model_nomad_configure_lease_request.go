// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// NomadConfigureLeaseRequest struct for NomadConfigureLeaseRequest
type NomadConfigureLeaseRequest struct {
	// Duration after which the issued token should not be allowed to be renewed
	MaxTtl string `json:"max_ttl,omitempty"`

	// Duration before which the issued token needs renewal
	Ttl string `json:"ttl,omitempty"`
}

// NewNomadConfigureLeaseRequestWithDefaults instantiates a new NomadConfigureLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNomadConfigureLeaseRequestWithDefaults() *NomadConfigureLeaseRequest {
	var this NomadConfigureLeaseRequest

	return &this
}
