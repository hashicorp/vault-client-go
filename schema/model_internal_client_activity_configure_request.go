// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalClientActivityConfigureRequest struct for InternalClientActivityConfigureRequest
type InternalClientActivityConfigureRequest struct {
	// Number of months to report if no start date specified.
	DefaultReportMonths int32 `json:"default_report_months,omitempty"`

	// Enable or disable collection of client count: enable, disable, or default.
	Enabled string `json:"enabled,omitempty"`

	// Number of months of client data to retain. Setting to 0 will clear all existing data.
	RetentionMonths int32 `json:"retention_months,omitempty"`
}
