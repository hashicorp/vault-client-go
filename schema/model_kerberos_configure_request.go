// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KerberosConfigureRequest struct for KerberosConfigureRequest
type KerberosConfigureRequest struct {
	// If set to true, returns any groups found in LDAP as a group alias.
	AddGroupAliases bool `json:"add_group_aliases,omitempty"`

	// Base64 encoded keytab
	Keytab string `json:"keytab,omitempty"`

	// Remove instance/FQDN from keytab principal names.
	RemoveInstanceName bool `json:"remove_instance_name,omitempty"`

	// Service Account
	ServiceAccount string `json:"service_account,omitempty"`
}
