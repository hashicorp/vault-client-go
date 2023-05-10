// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiConfigureClusterRequest struct for PkiConfigureClusterRequest
type PkiConfigureClusterRequest struct {
	// Optional URI to this mount's AIA distribution point; may refer to an external non-Vault responder. This is for resolving AIA URLs and providing the {{cluster_aia_path}} template parameter and will not be used for other purposes. As such, unlike path above, this could safely be an insecure transit mechanism (like HTTP without TLS). For example: http://cdn.example.com/pr1/pki
	AiaPath string `json:"aia_path"`

	// Canonical URI to this mount on this performance replication cluster's external address. This is for resolving AIA URLs and providing the {{cluster_path}} template parameter but might be used for other purposes in the future. This should only point back to this particular PR replica and should not ever point to another PR cluster. It may point to any node in the PR replica, including standby nodes, and need not always point to the active node. For example: https://pr1.vault.example.com:8200/v1/pki
	Path string `json:"path"`
}

// NewPkiConfigureClusterRequestWithDefaults instantiates a new PkiConfigureClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiConfigureClusterRequestWithDefaults() *PkiConfigureClusterRequest {
	var this PkiConfigureClusterRequest

	return &this
}

func (o PkiConfigureClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["aia_path"] = o.AiaPath
	toSerialize["path"] = o.Path

	return json.Marshal(toSerialize)
}
