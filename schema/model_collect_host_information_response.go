// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// CollectHostInformationResponse struct for CollectHostInformationResponse
type CollectHostInformationResponse struct {
	Cpu []map[string]interface{} `json:"cpu,omitempty"`

	CpuTimes []map[string]interface{} `json:"cpu_times,omitempty"`

	Disk []map[string]interface{} `json:"disk,omitempty"`

	Host map[string]interface{} `json:"host,omitempty"`

	Memory map[string]interface{} `json:"memory,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`
}

// NewCollectHostInformationResponseWithDefaults instantiates a new CollectHostInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectHostInformationResponseWithDefaults() *CollectHostInformationResponse {
	var this CollectHostInformationResponse

	return &this
}
