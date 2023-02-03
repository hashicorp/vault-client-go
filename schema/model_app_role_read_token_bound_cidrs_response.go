// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadTokenBoundCIDRsResponse struct for AppRoleReadTokenBoundCIDRsResponse
type AppRoleReadTokenBoundCIDRsResponse struct {
	// Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`
}

// NewAppRoleReadTokenBoundCIDRsResponseWithDefaults instantiates a new AppRoleReadTokenBoundCIDRsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadTokenBoundCIDRsResponseWithDefaults() *AppRoleReadTokenBoundCIDRsResponse {
	var this AppRoleReadTokenBoundCIDRsResponse

	return &this
}
