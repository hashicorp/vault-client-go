// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteInternalCountersConfigRequest struct for WriteInternalCountersConfigRequest
type WriteInternalCountersConfigRequest struct {
	// Number of months to report if no start date specified.
	DefaultReportMonths int32 `json:"default_report_months"`

	// Enable or disable collection of client count: enable, disable, or default.
	Enabled string `json:"enabled"`

	// Number of months of client data to retain. Setting to 0 will clear all existing data.
	RetentionMonths int32 `json:"retention_months"`
}

// NewWriteInternalCountersConfigRequestWithDefaults instantiates a new WriteInternalCountersConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteInternalCountersConfigRequestWithDefaults() *WriteInternalCountersConfigRequest {
	var this WriteInternalCountersConfigRequest

	this.DefaultReportMonths = 12
	this.Enabled = "default"
	this.RetentionMonths = 24

	return &this
}

func (o WriteInternalCountersConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["default_report_months"] = o.DefaultReportMonths
	toSerialize["enabled"] = o.Enabled
	toSerialize["retention_months"] = o.RetentionMonths

	return json.Marshal(toSerialize)
}
