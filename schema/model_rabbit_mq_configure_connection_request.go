// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RabbitMqConfigureConnectionRequest struct for RabbitMqConfigureConnectionRequest
type RabbitMqConfigureConnectionRequest struct {
	// RabbitMQ Management URI
	ConnectionUri string `json:"connection_uri,omitempty"`

	// Password of the provided RabbitMQ management user
	Password string `json:"password,omitempty"`

	// Name of the password policy to use to generate passwords for dynamic credentials.
	PasswordPolicy string `json:"password_policy,omitempty"`

	// Username of a RabbitMQ management administrator
	Username string `json:"username,omitempty"`

	// Template describing how dynamic usernames are generated.
	UsernameTemplate string `json:"username_template,omitempty"`

	// If set, connection_uri is verified by actually connecting to the RabbitMQ management API
	VerifyConnection bool `json:"verify_connection,omitempty"`
}
