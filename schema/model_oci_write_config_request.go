// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OCIWriteConfigRequest struct for OCIWriteConfigRequest
type OCIWriteConfigRequest struct {
	// The tenancy id of the account.
	HomeTenancyId string `json:"home_tenancy_id"`
}

// NewOCIWriteConfigRequestWithDefaults instantiates a new OCIWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCIWriteConfigRequestWithDefaults() *OCIWriteConfigRequest {
	var this OCIWriteConfigRequest

	return &this
}

func (o OCIWriteConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["home_tenancy_id"] = o.HomeTenancyId

	return json.Marshal(toSerialize)
}
