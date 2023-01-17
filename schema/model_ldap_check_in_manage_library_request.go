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

// LDAPCheckInManageLibraryRequest struct for LDAPCheckInManageLibraryRequest
type LDAPCheckInManageLibraryRequest struct {
	// The username/logon name for the service accounts to check in.
	ServiceAccountNames []string `json:"service_account_names"`
}

// NewLDAPCheckInManageLibraryRequestWithDefaults instantiates a new LDAPCheckInManageLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPCheckInManageLibraryRequestWithDefaults() *LDAPCheckInManageLibraryRequest {
	var this LDAPCheckInManageLibraryRequest

	return &this
}

func (o LDAPCheckInManageLibraryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["service_account_names"] = o.ServiceAccountNames

	return json.Marshal(toSerialize)
}
