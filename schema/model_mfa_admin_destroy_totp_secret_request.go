// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MfaAdminDestroyTotpSecretRequest struct for MfaAdminDestroyTotpSecretRequest
type MfaAdminDestroyTotpSecretRequest struct {
	// Identifier of the entity from which the MFA method secret needs to be removed.
	EntityId string `json:"entity_id"`

	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`
}
