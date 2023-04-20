// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RabbitMqConfigureConnectionRequest struct for RabbitMqConfigureConnectionRequest
type RabbitMqConfigureConnectionRequest struct {
	// RabbitMQ Management URI
	ConnectionUri string `json:"connection_uri"`

	// Password of the provided RabbitMQ management user
	Password string `json:"password"`

	// Name of the password policy to use to generate passwords for dynamic credentials.
	PasswordPolicy string `json:"password_policy"`

	// Username of a RabbitMQ management administrator
	Username string `json:"username"`

	// Template describing how dynamic usernames are generated.
	UsernameTemplate string `json:"username_template"`

	// If set, connection_uri is verified by actually connecting to the RabbitMQ management API
	VerifyConnection bool `json:"verify_connection"`
}

// NewRabbitMqConfigureConnectionRequestWithDefaults instantiates a new RabbitMqConfigureConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRabbitMqConfigureConnectionRequestWithDefaults() *RabbitMqConfigureConnectionRequest {
	var this RabbitMqConfigureConnectionRequest

	this.VerifyConnection = true

	return &this
}

func (o RabbitMqConfigureConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["connection_uri"] = o.ConnectionUri
	toSerialize["password"] = o.Password
	toSerialize["password_policy"] = o.PasswordPolicy
	toSerialize["username"] = o.Username
	toSerialize["username_template"] = o.UsernameTemplate
	toSerialize["verify_connection"] = o.VerifyConnection

	return json.Marshal(toSerialize)
}
