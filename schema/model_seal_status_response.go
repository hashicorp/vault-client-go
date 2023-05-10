// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SealStatusResponse struct for SealStatusResponse
type SealStatusResponse struct {
	BuildDate string `json:"build_date"`

	ClusterId string `json:"cluster_id"`

	ClusterName string `json:"cluster_name"`

	HcpLinkResourceID string `json:"hcp_link_resource_ID"`

	HcpLinkStatus string `json:"hcp_link_status"`

	Initialized bool `json:"initialized"`

	Migration bool `json:"migration"`

	N int32 `json:"n"`

	Nonce string `json:"nonce"`

	Progress int32 `json:"progress"`

	RecoverySeal bool `json:"recovery_seal"`

	Sealed bool `json:"sealed"`

	StorageType string `json:"storage_type"`

	T int32 `json:"t"`

	Type string `json:"type"`

	Version string `json:"version"`
}

// NewSealStatusResponseWithDefaults instantiates a new SealStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSealStatusResponseWithDefaults() *SealStatusResponse {
	var this SealStatusResponse

	return &this
}

func (o SealStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["build_date"] = o.BuildDate
	toSerialize["cluster_id"] = o.ClusterId
	toSerialize["cluster_name"] = o.ClusterName
	toSerialize["hcp_link_resource_ID"] = o.HcpLinkResourceID
	toSerialize["hcp_link_status"] = o.HcpLinkStatus
	toSerialize["initialized"] = o.Initialized
	toSerialize["migration"] = o.Migration
	toSerialize["n"] = o.N
	toSerialize["nonce"] = o.Nonce
	toSerialize["progress"] = o.Progress
	toSerialize["recovery_seal"] = o.RecoverySeal
	toSerialize["sealed"] = o.Sealed
	toSerialize["storage_type"] = o.StorageType
	toSerialize["t"] = o.T
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version

	return json.Marshal(toSerialize)
}
