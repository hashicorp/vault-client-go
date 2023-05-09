// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MongoDbAtlasConfigureRequest struct for MongoDbAtlasConfigureRequest
type MongoDbAtlasConfigureRequest struct {
	// MongoDB Atlas Programmatic Private Key
	PrivateKey string `json:"private_key"`

	// MongoDB Atlas Programmatic Public Key
	PublicKey string `json:"public_key"`
}

// NewMongoDbAtlasConfigureRequestWithDefaults instantiates a new MongoDbAtlasConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongoDbAtlasConfigureRequestWithDefaults() *MongoDbAtlasConfigureRequest {
	var this MongoDbAtlasConfigureRequest

	return &this
}
