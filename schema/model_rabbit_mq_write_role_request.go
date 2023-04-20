// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RabbitMqWriteRoleRequest struct for RabbitMqWriteRoleRequest
type RabbitMqWriteRoleRequest struct {
	// Comma-separated list of tags for this role.
	Tags string `json:"tags"`

	// A nested map of virtual hosts and exchanges to topic permissions.
	VhostTopics string `json:"vhost_topics"`

	// A map of virtual hosts to permissions.
	Vhosts string `json:"vhosts"`
}

// NewRabbitMqWriteRoleRequestWithDefaults instantiates a new RabbitMqWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRabbitMqWriteRoleRequestWithDefaults() *RabbitMqWriteRoleRequest {
	var this RabbitMqWriteRoleRequest

	return &this
}

func (o RabbitMqWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["tags"] = o.Tags
	toSerialize["vhost_topics"] = o.VhostTopics
	toSerialize["vhosts"] = o.Vhosts

	return json.Marshal(toSerialize)
}
