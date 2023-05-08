// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OidcWriteAssignmentRequest struct for OidcWriteAssignmentRequest
type OidcWriteAssignmentRequest struct {
	// Comma separated string or array of identity entity IDs
	EntityIds []string `json:"entity_ids,omitempty"`

	// Comma separated string or array of identity group IDs
	GroupIds []string `json:"group_ids,omitempty"`
}

// NewOidcWriteAssignmentRequestWithDefaults instantiates a new OidcWriteAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcWriteAssignmentRequestWithDefaults() *OidcWriteAssignmentRequest {
	var this OidcWriteAssignmentRequest

	return &this
}
