// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudKmsConfigureRequest struct for GoogleCloudKmsConfigureRequest
type GoogleCloudKmsConfigureRequest struct {
	// The credentials to use for authenticating to Google Cloud. Leave this blank to use the Default Application Credentials or instance metadata authentication.
	Credentials string `json:"credentials"`

	// The list of full-URL scopes to request when authenticating. By default, this requests https://www.googleapis.com/auth/cloudkms.
	Scopes []string `json:"scopes"`
}

// NewGoogleCloudKmsConfigureRequestWithDefaults instantiates a new GoogleCloudKmsConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKmsConfigureRequestWithDefaults() *GoogleCloudKmsConfigureRequest {
	var this GoogleCloudKmsConfigureRequest

	return &this
}

func (o GoogleCloudKmsConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["credentials"] = o.Credentials
	toSerialize["scopes"] = o.Scopes

	return json.Marshal(toSerialize)
}
