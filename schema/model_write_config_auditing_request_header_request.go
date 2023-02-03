// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteConfigAuditingRequestHeaderRequest struct for WriteConfigAuditingRequestHeaderRequest
type WriteConfigAuditingRequestHeaderRequest struct {
	Hmac bool `json:"hmac"`
}

// NewWriteConfigAuditingRequestHeaderRequestWithDefaults instantiates a new WriteConfigAuditingRequestHeaderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteConfigAuditingRequestHeaderRequestWithDefaults() *WriteConfigAuditingRequestHeaderRequest {
	var this WriteConfigAuditingRequestHeaderRequest

	return &this
}

func (o WriteConfigAuditingRequestHeaderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["hmac"] = o.Hmac

	return json.Marshal(toSerialize)
}
