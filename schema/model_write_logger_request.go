// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteLoggerRequest struct for WriteLoggerRequest
type WriteLoggerRequest struct {
	// Log verbosity level. Supported values (in order of detail) are \"trace\", \"debug\", \"info\", \"warn\", and \"error\".
	Level string `json:"level"`
}

// NewWriteLoggerRequestWithDefaults instantiates a new WriteLoggerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteLoggerRequestWithDefaults() *WriteLoggerRequest {
	var this WriteLoggerRequest

	return &this
}

func (o WriteLoggerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
