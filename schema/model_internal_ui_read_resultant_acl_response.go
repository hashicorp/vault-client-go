// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// InternalUiReadResultantAclResponse struct for InternalUiReadResultantAclResponse
type InternalUiReadResultantAclResponse struct {
	ExactPaths map[string]interface{} `json:"exact_paths"`

	GlobPaths map[string]interface{} `json:"glob_paths"`

	Root bool `json:"root"`
}

// NewInternalUiReadResultantAclResponseWithDefaults instantiates a new InternalUiReadResultantAclResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalUiReadResultantAclResponseWithDefaults() *InternalUiReadResultantAclResponse {
	var this InternalUiReadResultantAclResponse

	return &this
}

func (o InternalUiReadResultantAclResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["exact_paths"] = o.ExactPaths
	toSerialize["glob_paths"] = o.GlobPaths
	toSerialize["root"] = o.Root

	return json.Marshal(toSerialize)
}
