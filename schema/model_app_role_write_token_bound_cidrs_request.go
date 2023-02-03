// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteTokenBoundCIDRsRequest struct for AppRoleWriteTokenBoundCIDRsRequest
type AppRoleWriteTokenBoundCIDRsRequest struct {
	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`
}

// NewAppRoleWriteTokenBoundCIDRsRequestWithDefaults instantiates a new AppRoleWriteTokenBoundCIDRsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteTokenBoundCIDRsRequestWithDefaults() *AppRoleWriteTokenBoundCIDRsRequest {
	var this AppRoleWriteTokenBoundCIDRsRequest

	return &this
}
