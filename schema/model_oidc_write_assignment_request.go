// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteAssignmentRequest struct for OIDCWriteAssignmentRequest
type OIDCWriteAssignmentRequest struct {
	// Comma separated string or array of identity entity IDs
	EntityIds []string `json:"entity_ids"`

	// Comma separated string or array of identity group IDs
	GroupIds []string `json:"group_ids"`
}

// NewOIDCWriteAssignmentRequestWithDefaults instantiates a new OIDCWriteAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteAssignmentRequestWithDefaults() *OIDCWriteAssignmentRequest {
	var this OIDCWriteAssignmentRequest

	return &this
}

func (o OIDCWriteAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["entity_ids"] = o.EntityIds
	toSerialize["group_ids"] = o.GroupIds

	return json.Marshal(toSerialize)
}
