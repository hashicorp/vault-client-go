// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RemountRequest struct for RemountRequest
type RemountRequest struct {
	// The previous mount point.
	From string `json:"from"`

	// The new mount point.
	To string `json:"to"`
}

// NewRemountRequestWithDefaults instantiates a new RemountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemountRequestWithDefaults() *RemountRequest {
	var this RemountRequest

	return &this
}

func (o RemountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
