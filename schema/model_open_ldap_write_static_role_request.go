// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OpenLDAPWriteStaticRoleRequest struct for OpenLDAPWriteStaticRoleRequest
type OpenLDAPWriteStaticRoleRequest struct {
	// The distinguished name of the entry to manage.
	Dn string `json:"dn"`

	// Period for automatic credential rotation of the given entry.
	RotationPeriod int32 `json:"rotation_period"`

	// The username/logon name for the entry with which this role will be associated.
	Username string `json:"username"`
}

// NewOpenLDAPWriteStaticRoleRequestWithDefaults instantiates a new OpenLDAPWriteStaticRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenLDAPWriteStaticRoleRequestWithDefaults() *OpenLDAPWriteStaticRoleRequest {
	var this OpenLDAPWriteStaticRoleRequest

	return &this
}

func (o OpenLDAPWriteStaticRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["dn"] = o.Dn
	toSerialize["rotation_period"] = o.RotationPeriod
	toSerialize["username"] = o.Username

	return json.Marshal(toSerialize)
}
