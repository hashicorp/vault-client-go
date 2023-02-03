// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteBoundCIDRListRequest struct for AppRoleWriteBoundCIDRListRequest
type AppRoleWriteBoundCIDRListRequest struct {
	// Deprecated: Please use \"secret_id_bound_cidrs\" instead. Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.

	BoundCidrList []string `json:"bound_cidr_list"`
}

// NewAppRoleWriteBoundCIDRListRequestWithDefaults instantiates a new AppRoleWriteBoundCIDRListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteBoundCIDRListRequestWithDefaults() *AppRoleWriteBoundCIDRListRequest {
	var this AppRoleWriteBoundCIDRListRequest

	return &this
}

func (o AppRoleWriteBoundCIDRListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bound_cidr_list"] = o.BoundCidrList

	return json.Marshal(toSerialize)
}
