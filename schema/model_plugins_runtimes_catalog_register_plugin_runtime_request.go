// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PluginsRuntimesCatalogRegisterPluginRuntimeRequest struct for PluginsRuntimesCatalogRegisterPluginRuntimeRequest
type PluginsRuntimesCatalogRegisterPluginRuntimeRequest struct {
	// Optional parent cgroup for the container
	CgroupParent string `json:"cgroup_parent,omitempty"`

	// The limit of runtime CPU in nanos
	CpuNanos int64 `json:"cpu_nanos,omitempty"`

	// The limit of runtime memory in bytes
	MemoryBytes int64 `json:"memory_bytes,omitempty"`

	// The OCI-compatible runtime (default \"runsc\")
	OciRuntime string `json:"oci_runtime,omitempty"`
}
