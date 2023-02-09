// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OktaWriteGroupRequest struct for OktaWriteGroupRequest
type OktaWriteGroupRequest struct {
	// Comma-separated list of policies associated to the group.
	Policies []string `json:"policies"`
}

// NewOktaWriteGroupRequestWithDefaults instantiates a new OktaWriteGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaWriteGroupRequestWithDefaults() *OktaWriteGroupRequest {
	var this OktaWriteGroupRequest

	return &this
}

func (o OktaWriteGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
