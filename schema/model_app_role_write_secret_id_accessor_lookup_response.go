// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// AppRoleWriteSecretIDAccessorLookupResponse struct for AppRoleWriteSecretIDAccessorLookupResponse
type AppRoleWriteSecretIDAccessorLookupResponse struct {
	// List of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If 'bound_cidr_list' is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role.

	CidrList []string `json:"cidr_list"`

	CreationTime time.Time `json:"creation_time"`

	ExpirationTime time.Time `json:"expiration_time"`

	LastUpdatedTime time.Time `json:"last_updated_time"`

	Metadata map[string]interface{} `json:"metadata"`

	// Accessor of the secret ID

	SecretIdAccessor string `json:"secret_id_accessor"`

	// Number of times a secret ID can access the role, after which the secret ID will expire.

	SecretIdNumUses int32 `json:"secret_id_num_uses"`

	// Duration in seconds after which the issued secret ID expires.

	SecretIdTtl int32 `json:"secret_id_ttl"`

	// List of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any.

	TokenBoundCidrs []string `json:"token_bound_cidrs"`
}

// NewAppRoleWriteSecretIDAccessorLookupResponseWithDefaults instantiates a new AppRoleWriteSecretIDAccessorLookupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDAccessorLookupResponseWithDefaults() *AppRoleWriteSecretIDAccessorLookupResponse {
	var this AppRoleWriteSecretIDAccessorLookupResponse

	return &this
}

func (o AppRoleWriteSecretIDAccessorLookupResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cidr_list"] = o.CidrList

	toSerialize["creation_time"] = o.CreationTime

	toSerialize["expiration_time"] = o.ExpirationTime

	toSerialize["last_updated_time"] = o.LastUpdatedTime

	toSerialize["metadata"] = o.Metadata

	toSerialize["secret_id_accessor"] = o.SecretIdAccessor

	toSerialize["secret_id_num_uses"] = o.SecretIdNumUses

	toSerialize["secret_id_ttl"] = o.SecretIdTtl

	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs

	return json.Marshal(toSerialize)
}
