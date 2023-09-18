// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// DatabaseWriteStaticRoleRequest struct for DatabaseWriteStaticRoleRequest
type DatabaseWriteStaticRoleRequest struct {
	// The configuration for the given credential_type.
	CredentialConfig map[string]interface{} `json:"credential_config,omitempty"`

	// The type of credential to manage. Options include: 'password', 'rsa_private_key'. Defaults to 'password'.
	CredentialType string `json:"credential_type,omitempty"`

	// Name of the database this role acts on.
	DbName string `json:"db_name,omitempty"`

	// Period for automatic credential rotation of the given username. Not valid unless used with \"username\". Mutually exclusive with \"rotation_schedule.\"
	RotationPeriod string `json:"rotation_period,omitempty"`

	// Schedule for automatic credential rotation of the given username. Mutually exclusive with \"rotation_period.\"
	RotationSchedule string `json:"rotation_schedule,omitempty"`

	// Specifies the database statements to be executed to rotate the accounts credentials. Not every plugin type will support this functionality. See the plugin's API page for more information on support and formatting for this parameter.
	RotationStatements []string `json:"rotation_statements,omitempty"`

	// The window of time in which rotations are allowed to occur starting from a given \"rotation_schedule\". Requires \"rotation_schedule\" to be specified
	RotationWindow string `json:"rotation_window,omitempty"`

	// Name of the static user account for Vault to manage. Requires \"rotation_period\" to be specified
	Username string `json:"username,omitempty"`
}
