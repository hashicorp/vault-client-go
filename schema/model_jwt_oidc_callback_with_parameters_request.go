// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// JwtOidcCallbackWithParametersRequest struct for JwtOidcCallbackWithParametersRequest
type JwtOidcCallbackWithParametersRequest struct {
	ClientNonce string `json:"client_nonce,omitempty"`

	Code string `json:"code,omitempty"`

	IdToken string `json:"id_token,omitempty"`

	State string `json:"state,omitempty"`
}
