// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteCallbackRequest struct for OIDCWriteCallbackRequest
type OIDCWriteCallbackRequest struct {
	ClientNonce string `json:"client_nonce"`

	Code string `json:"code"`

	IdToken string `json:"id_token"`

	State string `json:"state"`
}

// NewOIDCWriteCallbackRequestWithDefaults instantiates a new OIDCWriteCallbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteCallbackRequestWithDefaults() *OIDCWriteCallbackRequest {
	var this OIDCWriteCallbackRequest

	return &this
}

func (o OIDCWriteCallbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["client_nonce"] = o.ClientNonce
	toSerialize["code"] = o.Code
	toSerialize["id_token"] = o.IdToken
	toSerialize["state"] = o.State

	return json.Marshal(toSerialize)
}
