// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudWriteConfigRequest struct for GoogleCloudWriteConfigRequest
type GoogleCloudWriteConfigRequest struct {
	// GCP IAM service account credentials JSON with permissions to create new service accounts and set IAM policies
	Credentials string `json:"credentials"`

	// Maximum time a service account key is valid for. If <= 0, will use system default.
	MaxTtl int32 `json:"max_ttl"`

	// Default lease for generated keys. If <= 0, will use system default.
	Ttl int32 `json:"ttl"`
}

// NewGoogleCloudWriteConfigRequestWithDefaults instantiates a new GoogleCloudWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteConfigRequestWithDefaults() *GoogleCloudWriteConfigRequest {
	var this GoogleCloudWriteConfigRequest

	return &this
}

func (o GoogleCloudWriteConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["credentials"] = o.Credentials
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
