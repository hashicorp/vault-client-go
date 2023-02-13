// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteQuotasRateLimitRequest struct for WriteQuotasRateLimitRequest
type WriteQuotasRateLimitRequest struct {
	// If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' has elapsed.
	BlockInterval int32 `json:"block_interval"`

	// The duration to enforce rate limiting for (default '1s').
	Interval int32 `json:"interval"`

	// Path of the mount or namespace to apply the quota. A blank path configures a global quota. For example namespace1/ adds a quota to a full namespace, namespace1/auth/userpass adds a quota to userpass in namespace1.
	Path string `json:"path"`

	// The maximum number of requests in a given interval to be allowed by the quota rule. The 'rate' must be positive.
	Rate float32 `json:"rate"`

	// Login role to apply this quota to. Note that when set, path must be configured to a valid auth method with a concept of roles.
	Role string `json:"role"`

	// Type of the quota rule.
	Type string `json:"type"`
}

// NewWriteQuotasRateLimitRequestWithDefaults instantiates a new WriteQuotasRateLimitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteQuotasRateLimitRequestWithDefaults() *WriteQuotasRateLimitRequest {
	var this WriteQuotasRateLimitRequest

	return &this
}

func (o WriteQuotasRateLimitRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["block_interval"] = o.BlockInterval
	toSerialize["interval"] = o.Interval
	toSerialize["path"] = o.Path
	toSerialize["rate"] = o.Rate
	toSerialize["role"] = o.Role
	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
