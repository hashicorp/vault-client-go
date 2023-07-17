// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// LeaderStatusResponse struct for LeaderStatusResponse
type LeaderStatusResponse struct {
	ActiveTime time.Time `json:"active_time,omitempty"`

	HaEnabled bool `json:"ha_enabled,omitempty"`

	IsSelf bool `json:"is_self,omitempty"`

	LastWal int64 `json:"last_wal,omitempty"`

	LeaderAddress string `json:"leader_address,omitempty"`

	LeaderClusterAddress string `json:"leader_cluster_address,omitempty"`

	PerformanceStandby bool `json:"performance_standby,omitempty"`

	PerformanceStandbyLastRemoteWal int64 `json:"performance_standby_last_remote_wal,omitempty"`

	RaftAppliedIndex int64 `json:"raft_applied_index,omitempty"`

	RaftCommittedIndex int64 `json:"raft_committed_index,omitempty"`
}
