// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RevokeLeaseRequest struct for RevokeLeaseRequest
type RevokeLeaseRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`

	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`
}

// NewRevokeLeaseRequestWithDefaults instantiates a new RevokeLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeLeaseRequestWithDefaults() *RevokeLeaseRequest {
	var this RevokeLeaseRequest

	this.Sync = true

	return &this
}

func (o RevokeLeaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
