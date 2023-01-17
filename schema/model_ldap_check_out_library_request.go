/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LDAPCheckOutLibraryRequest struct for LDAPCheckOutLibraryRequest
type LDAPCheckOutLibraryRequest struct {
	// The length of time before the check-out will expire, in seconds.
	Ttl int32 `json:"ttl"`
}

// NewLDAPCheckOutLibraryRequestWithDefaults instantiates a new LDAPCheckOutLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPCheckOutLibraryRequestWithDefaults() *LDAPCheckOutLibraryRequest {
	var this LDAPCheckOutLibraryRequest

	return &this
}

func (o LDAPCheckOutLibraryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
