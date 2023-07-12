// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LdapWriteStaticRoleRequest struct for LdapWriteStaticRoleRequest
type LdapWriteStaticRoleRequest struct {
	// The distinguished name of the entry to manage.
	Dn string `json:"dn,omitempty"`

	// Period for automatic credential rotation of the given entry.
	RotationPeriod string `json:"rotation_period,omitempty"`

	// The username/logon name for the entry with which this role will be associated.
	Username string `json:"username,omitempty"`
}
