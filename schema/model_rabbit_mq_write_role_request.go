// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RabbitMqWriteRoleRequest struct for RabbitMqWriteRoleRequest
type RabbitMqWriteRoleRequest struct {
	// Comma-separated list of tags for this role.
	Tags string `json:"tags,omitempty"`

	// A nested map of virtual hosts and exchanges to topic permissions.
	VhostTopics string `json:"vhost_topics,omitempty"`

	// A map of virtual hosts to permissions.
	Vhosts string `json:"vhosts,omitempty"`
}

// NewRabbitMqWriteRoleRequestWithDefaults instantiates a new RabbitMqWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRabbitMqWriteRoleRequestWithDefaults() *RabbitMqWriteRoleRequest {
	var this RabbitMqWriteRoleRequest

	return &this
}
