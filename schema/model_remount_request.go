// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RemountRequest struct for RemountRequest
type RemountRequest struct {
	// The previous mount point.
	From string `json:"from,omitempty"`

	// The new mount point.
	To string `json:"to,omitempty"`
}
