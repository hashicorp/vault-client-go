// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// DatabaseWriteRoleRequest struct for DatabaseWriteRoleRequest
type DatabaseWriteRoleRequest struct {
	// Specifies the database statements executed to create and configure a user. See the plugin's API page for more information on support and formatting for this parameter.
	CreationStatements []string `json:"creation_statements,omitempty"`

	// The configuration for the given credential_type.
	CredentialConfig map[string]interface{} `json:"credential_config,omitempty"`

	// The type of credential to manage. Options include: 'password', 'rsa_private_key'. Defaults to 'password'.
	CredentialType string `json:"credential_type,omitempty"`

	// Name of the database this role acts on.
	DbName string `json:"db_name,omitempty"`

	// Default ttl for role.
	DefaultTtl string `json:"default_ttl,omitempty"`

	// Maximum time a credential is valid for
	MaxTtl string `json:"max_ttl,omitempty"`

	// Specifies the database statements to be executed to renew a user. Not every plugin type will support this functionality. See the plugin's API page for more information on support and formatting for this parameter.
	RenewStatements []string `json:"renew_statements,omitempty"`

	// Specifies the database statements to be executed to revoke a user. See the plugin's API page for more information on support and formatting for this parameter.
	RevocationStatements []string `json:"revocation_statements,omitempty"`

	// Specifies the database statements to be executed rollback a create operation in the event of an error. Not every plugin type will support this functionality. See the plugin's API page for more information on support and formatting for this parameter.
	RollbackStatements []string `json:"rollback_statements,omitempty"`
}
