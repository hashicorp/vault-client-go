// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// QueryTokenSelfCapabilitiesRequest struct for QueryTokenSelfCapabilitiesRequest
type QueryTokenSelfCapabilitiesRequest struct {
	// Use 'paths' instead.
	// Deprecated
	Path []string `json:"path"`

	// Paths on which capabilities are being queried.
	Paths []string `json:"paths"`

	// Token for which capabilities are being queried.
	Token string `json:"token"`
}

// NewQueryTokenSelfCapabilitiesRequestWithDefaults instantiates a new QueryTokenSelfCapabilitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTokenSelfCapabilitiesRequestWithDefaults() *QueryTokenSelfCapabilitiesRequest {
	var this QueryTokenSelfCapabilitiesRequest

	return &this
}

func (o QueryTokenSelfCapabilitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["path"] = o.Path
	toSerialize["paths"] = o.Paths
	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
