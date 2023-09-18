// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitConfigureKeysRequest struct for TransitConfigureKeysRequest
type TransitConfigureKeysRequest struct {
	// Whether to allow automatic upserting (creation) of keys on the encrypt endpoint.
	DisableUpsert bool `json:"disable_upsert,omitempty"`
}
