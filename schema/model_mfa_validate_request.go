// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MfaValidateRequest struct for MfaValidateRequest
type MfaValidateRequest struct {
	// A map from MFA method ID to a slice of passcodes or an empty slice if the method does not use passcodes
	MfaPayload map[string]interface{} `json:"mfa_payload"`

	// ID for this MFA request
	MfaRequestId string `json:"mfa_request_id"`
}
