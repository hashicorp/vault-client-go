// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// CollectHostInformationResponse struct for CollectHostInformationResponse
type CollectHostInformationResponse struct {
	Cpu []map[string]interface{} `json:"cpu"`

	CpuTimes []map[string]interface{} `json:"cpu_times"`

	Disk []map[string]interface{} `json:"disk"`

	Host map[string]interface{} `json:"host"`

	Memory map[string]interface{} `json:"memory"`

	Timestamp time.Time `json:"timestamp"`
}

// NewCollectHostInformationResponseWithDefaults instantiates a new CollectHostInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectHostInformationResponseWithDefaults() *CollectHostInformationResponse {
	var this CollectHostInformationResponse

	return &this
}

func (o CollectHostInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cpu"] = o.Cpu
	toSerialize["cpu_times"] = o.CpuTimes
	toSerialize["disk"] = o.Disk
	toSerialize["host"] = o.Host
	toSerialize["memory"] = o.Memory
	toSerialize["timestamp"] = o.Timestamp

	return json.Marshal(toSerialize)
}
