// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LoggersUpdateVerbosityLevelRequest struct for LoggersUpdateVerbosityLevelRequest
type LoggersUpdateVerbosityLevelRequest struct {
	// Log verbosity level. Supported values (in order of detail) are \"trace\", \"debug\", \"info\", \"warn\", and \"error\".
	Level string `json:"level,omitempty"`
}
