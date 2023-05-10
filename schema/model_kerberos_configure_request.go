// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KerberosConfigureRequest struct for KerberosConfigureRequest
type KerberosConfigureRequest struct {
	// If set to true, returns any groups found in LDAP as a group alias.
	AddGroupAliases bool `json:"add_group_aliases"`

	// Base64 encoded keytab
	Keytab string `json:"keytab"`

	// Remove instance/FQDN from keytab principal names.
	RemoveInstanceName bool `json:"remove_instance_name"`

	// Service Account
	ServiceAccount string `json:"service_account"`
}

// NewKerberosConfigureRequestWithDefaults instantiates a new KerberosConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosConfigureRequestWithDefaults() *KerberosConfigureRequest {
	var this KerberosConfigureRequest

	return &this
}

func (o KerberosConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["add_group_aliases"] = o.AddGroupAliases
	toSerialize["keytab"] = o.Keytab
	toSerialize["remove_instance_name"] = o.RemoveInstanceName
	toSerialize["service_account"] = o.ServiceAccount

	return json.Marshal(toSerialize)
}
