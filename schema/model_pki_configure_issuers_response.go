// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiConfigureIssuersResponse struct for PkiConfigureIssuersResponse
type PkiConfigureIssuersResponse struct {
	// Reference (name or identifier) to the default issuer.
	Default string `json:"default"`

	// Whether the default issuer should automatically follow the latest generated or imported issuer. Defaults to false.
	DefaultFollowsLatestIssuer bool `json:"default_follows_latest_issuer"`
}

// NewPkiConfigureIssuersResponseWithDefaults instantiates a new PkiConfigureIssuersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiConfigureIssuersResponseWithDefaults() *PkiConfigureIssuersResponse {
	var this PkiConfigureIssuersResponse

	return &this
}

func (o PkiConfigureIssuersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["default"] = o.Default
	toSerialize["default_follows_latest_issuer"] = o.DefaultFollowsLatestIssuer

	return json.Marshal(toSerialize)
}
