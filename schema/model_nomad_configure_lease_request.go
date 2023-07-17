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
