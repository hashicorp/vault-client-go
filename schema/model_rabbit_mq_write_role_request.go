// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RabbitMQWriteRoleRequest struct for RabbitMQWriteRoleRequest
type RabbitMQWriteRoleRequest struct {
	// Comma-separated list of tags for this role.
	Tags string `json:"tags"`

	// A nested map of virtual hosts and exchanges to topic permissions.
	VhostTopics string `json:"vhost_topics"`

	// A map of virtual hosts to permissions.
	Vhosts string `json:"vhosts"`
}

// NewRabbitMQWriteRoleRequestWithDefaults instantiates a new RabbitMQWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRabbitMQWriteRoleRequestWithDefaults() *RabbitMQWriteRoleRequest {
	var this RabbitMQWriteRoleRequest

	return &this
}

func (o RabbitMQWriteRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
