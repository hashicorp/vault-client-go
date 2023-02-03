// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RenewForRequest struct for RenewForRequest
type RenewForRequest struct {
	// The desired increment in seconds to the lease
	Increment int32 `json:"increment"`

	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`
}

// NewRenewForRequestWithDefaults instantiates a new RenewForRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewForRequestWithDefaults() *RenewForRequest {
	var this RenewForRequest

	return &this
}
