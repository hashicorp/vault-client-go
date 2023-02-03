// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MongoDBAtlasWriteConfigRequest struct for MongoDBAtlasWriteConfigRequest
type MongoDBAtlasWriteConfigRequest struct {
	// MongoDB Atlas Programmatic Private Key

	PrivateKey string `json:"private_key"`

	// MongoDB Atlas Programmatic Public Key

	PublicKey string `json:"public_key"`
}

// NewMongoDBAtlasWriteConfigRequestWithDefaults instantiates a new MongoDBAtlasWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongoDBAtlasWriteConfigRequestWithDefaults() *MongoDBAtlasWriteConfigRequest {
	var this MongoDBAtlasWriteConfigRequest

	return &this
}

func (o MongoDBAtlasWriteConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["private_key"] = o.PrivateKey

	toSerialize["public_key"] = o.PublicKey

	return json.Marshal(toSerialize)
}
