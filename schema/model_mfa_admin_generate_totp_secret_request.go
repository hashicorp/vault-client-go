// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MfaAdminGenerateTotpSecretRequest struct for MfaAdminGenerateTotpSecretRequest
type MfaAdminGenerateTotpSecretRequest struct {
	// Entity ID on which the generated secret needs to get stored.
	EntityId string `json:"entity_id"`

	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`
}
