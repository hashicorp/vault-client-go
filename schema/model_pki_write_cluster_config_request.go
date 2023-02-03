// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PKIWriteClusterConfigRequest struct for PKIWriteClusterConfigRequest
type PKIWriteClusterConfigRequest struct {
	// Canonical URI to this mount on this performance replication cluster's external address. This is for resolving AIA URLs and providing the {{cluster_path}} template parameter but might be used for other purposes in the future. This should only point back to this particular PR replica and should not ever point to another PR cluster. It may point to any node in the PR replica, including standby nodes, and need not always point to the active node. For example: https://pr1.vault.example.com:8200/v1/pki
	Path string `json:"path"`
}

// NewPKIWriteClusterConfigRequestWithDefaults instantiates a new PKIWriteClusterConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteClusterConfigRequestWithDefaults() *PKIWriteClusterConfigRequest {
	var this PKIWriteClusterConfigRequest

	return &this
}
