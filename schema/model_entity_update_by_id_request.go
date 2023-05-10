// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EntityUpdateByIdRequest struct for EntityUpdateByIdRequest
type EntityUpdateByIdRequest struct {
	// If set true, tokens tied to this identity will not be able to be used (but will not be revoked).
	Disabled bool `json:"disabled"`

	// Metadata to be associated with the entity. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata"`

	// Name of the entity
	Name string `json:"name"`

	// Policies to be tied to the entity.
	Policies []string `json:"policies"`
}

// NewEntityUpdateByIdRequestWithDefaults instantiates a new EntityUpdateByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityUpdateByIdRequestWithDefaults() *EntityUpdateByIdRequest {
	var this EntityUpdateByIdRequest

	return &this
}

func (o EntityUpdateByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["disabled"] = o.Disabled
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["policies"] = o.Policies

	return json.Marshal(toSerialize)
}
