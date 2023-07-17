// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LdapLibraryConfigureRequest struct for LdapLibraryConfigureRequest
type LdapLibraryConfigureRequest struct {
	// Disable the default behavior of requiring that check-ins are performed by the entity that checked them out.
	DisableCheckInEnforcement bool `json:"disable_check_in_enforcement,omitempty"`

	// In seconds, the max amount of time a check-out's renewals should last. Defaults to 24 hours.
	MaxTtl string `json:"max_ttl,omitempty"`

	// The username/logon name for the service accounts with which this set will be associated.
	ServiceAccountNames []string `json:"service_account_names,omitempty"`

	// In seconds, the amount of time a check-out should last. Defaults to 24 hours.
	Ttl string `json:"ttl,omitempty"`
}
