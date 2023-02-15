// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// DatabaseWriteRoleRequest struct for DatabaseWriteRoleRequest
type DatabaseWriteRoleRequest struct {
	// Specifies the database statements executed to create and configure a user. See the plugin's API page for more information on support and formatting for this parameter.
	CreationStatements []string `json:"creation_statements"`

	// The configuration for the given credential_type.
	CredentialConfig map[string]interface{} `json:"credential_config"`

	// The type of credential to manage. Options include: 'password', 'rsa_private_key'. Defaults to 'password'.
	CredentialType string `json:"credential_type"`

	// Name of the database this role acts on.
	DbName string `json:"db_name"`

	// Default ttl for role.
	DefaultTtl int32 `json:"default_ttl"`

	// Maximum time a credential is valid for
	MaxTtl int32 `json:"max_ttl"`

	// Specifies the database statements to be executed to renew a user. Not every plugin type will support this functionality. See the plugin's API page for more information on support and formatting for this parameter.
	RenewStatements []string `json:"renew_statements"`

	// Specifies the database statements to be executed to revoke a user. See the plugin's API page for more information on support and formatting for this parameter.
	RevocationStatements []string `json:"revocation_statements"`

	// Specifies the database statements to be executed rollback a create operation in the event of an error. Not every plugin type will support this functionality. See the plugin's API page for more information on support and formatting for this parameter.
	RollbackStatements []string `json:"rollback_statements"`
}

// NewDatabaseWriteRoleRequestWithDefaults instantiates a new DatabaseWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseWriteRoleRequestWithDefaults() *DatabaseWriteRoleRequest {
	var this DatabaseWriteRoleRequest

	this.CredentialType = "password"

	return &this
}

func (o DatabaseWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["creation_statements"] = o.CreationStatements
	toSerialize["credential_config"] = o.CredentialConfig
	toSerialize["credential_type"] = o.CredentialType
	toSerialize["db_name"] = o.DbName
	toSerialize["default_ttl"] = o.DefaultTtl
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["renew_statements"] = o.RenewStatements
	toSerialize["revocation_statements"] = o.RevocationStatements
	toSerialize["rollback_statements"] = o.RollbackStatements

	return json.Marshal(toSerialize)
}
