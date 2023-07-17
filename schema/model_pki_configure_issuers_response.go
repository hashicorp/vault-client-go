// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiConfigureIssuersResponse struct for PkiConfigureIssuersResponse
type PkiConfigureIssuersResponse struct {
	// Reference (name or identifier) to the default issuer.
	Default string `json:"default,omitempty"`

	// Whether the default issuer should automatically follow the latest generated or imported issuer. Defaults to false.
	DefaultFollowsLatestIssuer bool `json:"default_follows_latest_issuer,omitempty"`
}
