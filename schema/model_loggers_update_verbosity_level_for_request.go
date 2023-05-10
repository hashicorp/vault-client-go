// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LoggersUpdateVerbosityLevelForRequest struct for LoggersUpdateVerbosityLevelForRequest
type LoggersUpdateVerbosityLevelForRequest struct {
	// Log verbosity level. Supported values (in order of detail) are \"trace\", \"debug\", \"info\", \"warn\", and \"error\".
	Level string `json:"level"`
}

// NewLoggersUpdateVerbosityLevelForRequestWithDefaults instantiates a new LoggersUpdateVerbosityLevelForRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggersUpdateVerbosityLevelForRequestWithDefaults() *LoggersUpdateVerbosityLevelForRequest {
	var this LoggersUpdateVerbosityLevelForRequest

	return &this
}

func (o LoggersUpdateVerbosityLevelForRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["level"] = o.Level

	return json.Marshal(toSerialize)
}
