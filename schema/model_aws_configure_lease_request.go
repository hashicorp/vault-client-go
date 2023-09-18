// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsConfigureLeaseRequest struct for AwsConfigureLeaseRequest
type AwsConfigureLeaseRequest struct {
	// Default lease for roles.
	Lease string `json:"lease,omitempty"`

	// Maximum time a credential is valid for.
	LeaseMax string `json:"lease_max,omitempty"`
}
