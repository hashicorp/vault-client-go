// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LoggersUpdateVerbosityLevelForRequest struct for LoggersUpdateVerbosityLevelForRequest
type LoggersUpdateVerbosityLevelForRequest struct {
	// Log verbosity level. Supported values (in order of detail) are \"trace\", \"debug\", \"info\", \"warn\", and \"error\".
	Level string `json:"level,omitempty"`
}

// NewLoggersUpdateVerbosityLevelForRequestWithDefaults instantiates a new LoggersUpdateVerbosityLevelForRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggersUpdateVerbosityLevelForRequestWithDefaults() *LoggersUpdateVerbosityLevelForRequest {
	var this LoggersUpdateVerbosityLevelForRequest

	return &this
}
