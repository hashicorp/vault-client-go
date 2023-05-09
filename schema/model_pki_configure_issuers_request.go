// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiConfigureIssuersRequest struct for PkiConfigureIssuersRequest
type PkiConfigureIssuersRequest struct {
	// Reference (name or identifier) to the default issuer.
	Default string `json:"default,omitempty"`

	// Whether the default issuer should automatically follow the latest generated or imported issuer. Defaults to false.
	DefaultFollowsLatestIssuer bool `json:"default_follows_latest_issuer,omitempty"`
}

// NewPkiConfigureIssuersRequestWithDefaults instantiates a new PkiConfigureIssuersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiConfigureIssuersRequestWithDefaults() *PkiConfigureIssuersRequest {
	var this PkiConfigureIssuersRequest

	this.DefaultFollowsLatestIssuer = false

	return &this
}
