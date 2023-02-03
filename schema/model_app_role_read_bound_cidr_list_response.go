// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadBoundCIDRListResponse struct for AppRoleReadBoundCIDRListResponse
type AppRoleReadBoundCIDRListResponse struct {
	// Deprecated: Please use \"secret_id_bound_cidrs\" instead. Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.
	// Deprecated
	BoundCidrList []string `json:"bound_cidr_list"`
}

// NewAppRoleReadBoundCIDRListResponseWithDefaults instantiates a new AppRoleReadBoundCIDRListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadBoundCIDRListResponseWithDefaults() *AppRoleReadBoundCIDRListResponse {
	var this AppRoleReadBoundCIDRListResponse

	return &this
}
