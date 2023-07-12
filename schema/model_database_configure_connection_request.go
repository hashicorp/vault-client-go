// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// DatabaseConfigureConnectionRequest struct for DatabaseConfigureConnectionRequest
type DatabaseConfigureConnectionRequest struct {
	// Comma separated string or array of the role names allowed to get creds from this database connection. If empty no roles are allowed. If \"*\" all roles are allowed.
	AllowedRoles []string `json:"allowed_roles,omitempty"`

	// Password policy to use when generating passwords.
	PasswordPolicy string `json:"password_policy,omitempty"`

	// The name of a builtin or previously registered plugin known to vault. This endpoint will create an instance of that plugin type.
	PluginName string `json:"plugin_name,omitempty"`

	// The version of the plugin to use.
	PluginVersion string `json:"plugin_version,omitempty"`

	// Specifies the database statements to be executed to rotate the root user's credentials. See the plugin's API page for more information on support and formatting for this parameter.
	RootRotationStatements []string `json:"root_rotation_statements,omitempty"`

	// If true, the connection details are verified by actually connecting to the database. Defaults to true.
	VerifyConnection bool `json:"verify_connection,omitempty"`
}
