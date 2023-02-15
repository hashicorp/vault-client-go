// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// DatabaseWriteConfigRequest struct for DatabaseWriteConfigRequest
type DatabaseWriteConfigRequest struct {
	// Comma separated string or array of the role names allowed to get creds from this database connection. If empty no roles are allowed. If \"*\" all roles are allowed.
	AllowedRoles []string `json:"allowed_roles"`

	// Password policy to use when generating passwords.
	PasswordPolicy string `json:"password_policy"`

	// The name of a builtin or previously registered plugin known to vault. This endpoint will create an instance of that plugin type.
	PluginName string `json:"plugin_name"`

	// The version of the plugin to use.
	PluginVersion string `json:"plugin_version"`

	// Specifies the database statements to be executed to rotate the root user's credentials. See the plugin's API page for more information on support and formatting for this parameter.
	RootRotationStatements []string `json:"root_rotation_statements"`

	// If true, the connection details are verified by actually connecting to the database. Defaults to true.
	VerifyConnection bool `json:"verify_connection"`
}

// NewDatabaseWriteConfigRequestWithDefaults instantiates a new DatabaseWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseWriteConfigRequestWithDefaults() *DatabaseWriteConfigRequest {
	var this DatabaseWriteConfigRequest

	this.VerifyConnection = true

	return &this
}

func (o DatabaseWriteConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allowed_roles"] = o.AllowedRoles
	toSerialize["password_policy"] = o.PasswordPolicy
	toSerialize["plugin_name"] = o.PluginName
	toSerialize["plugin_version"] = o.PluginVersion
	toSerialize["root_rotation_statements"] = o.RootRotationStatements
	toSerialize["verify_connection"] = o.VerifyConnection

	return json.Marshal(toSerialize)
}
