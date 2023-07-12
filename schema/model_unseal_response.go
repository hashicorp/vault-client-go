// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// UnsealResponse struct for UnsealResponse
type UnsealResponse struct {
	BuildDate string `json:"build_date,omitempty"`

	ClusterId string `json:"cluster_id,omitempty"`

	ClusterName string `json:"cluster_name,omitempty"`

	HcpLinkResourceID string `json:"hcp_link_resource_ID,omitempty"`

	HcpLinkStatus string `json:"hcp_link_status,omitempty"`

	Initialized bool `json:"initialized,omitempty"`

	Migration bool `json:"migration,omitempty"`

	N int32 `json:"n,omitempty"`

	Nonce string `json:"nonce,omitempty"`

	Progress int32 `json:"progress,omitempty"`

	RecoverySeal bool `json:"recovery_seal,omitempty"`

	Sealed bool `json:"sealed,omitempty"`

	StorageType string `json:"storage_type,omitempty"`

	T int32 `json:"t,omitempty"`

	Type string `json:"type,omitempty"`

	Version string `json:"version,omitempty"`
}
