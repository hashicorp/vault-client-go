// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// ReadWrappingPropertiesResponse struct for ReadWrappingPropertiesResponse
type ReadWrappingPropertiesResponse struct {
	CreationPath string `json:"creation_path,omitempty"`

	CreationTime time.Time `json:"creation_time,omitempty"`

	CreationTtl string `json:"creation_ttl,omitempty"`
}
