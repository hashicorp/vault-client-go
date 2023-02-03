// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AliCloudWriteRoleRequest struct for AliCloudWriteRoleRequest
type AliCloudWriteRoleRequest struct {
	// JSON of policies to be dynamically applied to users of this role.
	InlinePolicies string `json:"inline_policies"`

	// The maximum allowed lifetime of tokens issued using this role.
	MaxTtl int32 `json:"max_ttl"`

	// The name and type of each remote policy to be applied. Example: \"name:AliyunRDSReadOnlyAccess,type:System\".
	RemotePolicies []string `json:"remote_policies"`

	// ARN of the role to be assumed. If provided, inline_policies and remote_policies should be blank. At creation time, this role must have configured trusted actors, and the access key and secret that will be used to assume the role (in /config) must qualify as a trusted actor.
	RoleArn string `json:"role_arn"`

	// Duration in seconds after which the issued token should expire. Defaults to 0, in which case the value will fallback to the system/mount defaults.
	Ttl int32 `json:"ttl"`
}

// NewAliCloudWriteRoleRequestWithDefaults instantiates a new AliCloudWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliCloudWriteRoleRequestWithDefaults() *AliCloudWriteRoleRequest {
	var this AliCloudWriteRoleRequest

	return &this
}
