// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// LeaderStatusResponse struct for LeaderStatusResponse
type LeaderStatusResponse struct {
	ActiveTime time.Time `json:"active_time"`

	HaEnabled bool `json:"ha_enabled"`

	IsSelf bool `json:"is_self"`

	LastWal int64 `json:"last_wal"`

	LeaderAddress string `json:"leader_address"`

	LeaderClusterAddress string `json:"leader_cluster_address"`

	PerformanceStandby bool `json:"performance_standby"`

	PerformanceStandbyLastRemoteWal int64 `json:"performance_standby_last_remote_wal"`

	RaftAppliedIndex int64 `json:"raft_applied_index"`

	RaftCommittedIndex int64 `json:"raft_committed_index"`
}

// NewLeaderStatusResponseWithDefaults instantiates a new LeaderStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeaderStatusResponseWithDefaults() *LeaderStatusResponse {
	var this LeaderStatusResponse

	return &this
}

func (o LeaderStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["active_time"] = o.ActiveTime
	toSerialize["ha_enabled"] = o.HaEnabled
	toSerialize["is_self"] = o.IsSelf
	toSerialize["last_wal"] = o.LastWal
	toSerialize["leader_address"] = o.LeaderAddress
	toSerialize["leader_cluster_address"] = o.LeaderClusterAddress
	toSerialize["performance_standby"] = o.PerformanceStandby
	toSerialize["performance_standby_last_remote_wal"] = o.PerformanceStandbyLastRemoteWal
	toSerialize["raft_applied_index"] = o.RaftAppliedIndex
	toSerialize["raft_committed_index"] = o.RaftCommittedIndex

	return json.Marshal(toSerialize)
}
