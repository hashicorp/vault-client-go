// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenWriteCreateWithRoleRequest struct for TokenWriteCreateWithRoleRequest
type TokenWriteCreateWithRoleRequest struct {
	// Name to associate with this token

	DisplayName string `json:"display_name"`

	// Name of the entity alias to associate with this token

	EntityAlias string `json:"entity_alias"`

	// Explicit Max TTL of this token

	ExplicitMaxTtl string `json:"explicit_max_ttl"`

	// Value for the token

	Id string `json:"id"`

	// Arbitrary key=value metadata to associate with the token

	Metadata map[string]interface{} `json:"metadata"`

	// Do not include default policy for this token

	NoDefaultPolicy bool `json:"no_default_policy"`

	// Create the token with no parent

	NoParent bool `json:"no_parent"`

	// Max number of uses for this token

	NumUses int32 `json:"num_uses"`

	// Renew period

	Period string `json:"period"`

	// List of policies for the token

	Policies []string `json:"policies"`

	// Allow token to be renewed past its initial TTL up to system/mount maximum TTL

	Renewable bool `json:"renewable"`

	// Time to live for this token

	Ttl string `json:"ttl"`

	// Token type

	Type string `json:"type"`
}

// NewTokenWriteCreateWithRoleRequestWithDefaults instantiates a new TokenWriteCreateWithRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWriteCreateWithRoleRequestWithDefaults() *TokenWriteCreateWithRoleRequest {
	var this TokenWriteCreateWithRoleRequest

	return &this
}

func (o TokenWriteCreateWithRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["display_name"] = o.DisplayName

	toSerialize["entity_alias"] = o.EntityAlias

	toSerialize["explicit_max_ttl"] = o.ExplicitMaxTtl

	toSerialize["id"] = o.Id

	toSerialize["metadata"] = o.Metadata

	toSerialize["no_default_policy"] = o.NoDefaultPolicy

	toSerialize["no_parent"] = o.NoParent

	toSerialize["num_uses"] = o.NumUses

	toSerialize["period"] = o.Period

	toSerialize["policies"] = o.Policies

	toSerialize["renewable"] = o.Renewable

	toSerialize["ttl"] = o.Ttl

	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
