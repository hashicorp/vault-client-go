// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OpenLDAPWriteRoleRequest struct for OpenLDAPWriteRoleRequest
type OpenLDAPWriteRoleRequest struct {
	// LDIF string used to create new entities within the LDAP system. This LDIF can be templated.
	CreationLdif string `json:"creation_ldif"`

	// Default TTL for dynamic credentials
	DefaultTtl int32 `json:"default_ttl"`

	// LDIF string used to delete entities created within the LDAP system. This LDIF can be templated.
	DeletionLdif string `json:"deletion_ldif"`

	// Max TTL a dynamic credential can be extended to
	MaxTtl int32 `json:"max_ttl"`

	// LDIF string used to rollback changes in the event of a failure to create credentials. This LDIF can be templated.
	RollbackLdif string `json:"rollback_ldif"`

	// The template used to create a username
	UsernameTemplate string `json:"username_template"`
}

// NewOpenLDAPWriteRoleRequestWithDefaults instantiates a new OpenLDAPWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenLDAPWriteRoleRequestWithDefaults() *OpenLDAPWriteRoleRequest {
	var this OpenLDAPWriteRoleRequest

	return &this
}

func (o OpenLDAPWriteRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
