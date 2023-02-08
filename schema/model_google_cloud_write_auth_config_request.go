// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudWriteAuthConfigRequest struct for GoogleCloudWriteAuthConfigRequest
type GoogleCloudWriteAuthConfigRequest struct {
	// Google credentials JSON that Vault will use to verify users against GCP APIs. If not specified, will use application default credentials
	Credentials string `json:"credentials"`

	// Specifies overrides for various Google API Service Endpoints used in requests.
	CustomEndpoint map[string]interface{} `json:"custom_endpoint"`

	// Indicates what value to use when generating an alias for GCE authentications.
	GceAlias string `json:"gce_alias"`

	// The metadata to include on the aliases and audit logs generated by this plugin. When set to 'default', includes: instance_creation_timestamp, instance_id, instance_name, project_id, project_number, role, service_account_id, service_account_email, zone. Not editing this field means the 'default' fields are included. Explicitly setting this field to empty overrides the 'default' and means no metadata will be included. If not using 'default', explicit fields must be sent like: 'field1,field2'.
	GceMetadata []string `json:"gce_metadata"`

	// Deprecated. This field does nothing and be removed in a future release
	// Deprecated
	GoogleCertsEndpoint string `json:"google_certs_endpoint"`

	// Indicates what value to use when generating an alias for IAM authentications.
	IamAlias string `json:"iam_alias"`

	// The metadata to include on the aliases and audit logs generated by this plugin. When set to 'default', includes: project_id, role, service_account_id, service_account_email. Not editing this field means the 'default' fields are included. Explicitly setting this field to empty overrides the 'default' and means no metadata will be included. If not using 'default', explicit fields must be sent like: 'field1,field2'.
	IamMetadata []string `json:"iam_metadata"`
}

// NewGoogleCloudWriteAuthConfigRequestWithDefaults instantiates a new GoogleCloudWriteAuthConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteAuthConfigRequestWithDefaults() *GoogleCloudWriteAuthConfigRequest {
	var this GoogleCloudWriteAuthConfigRequest

	this.GceAlias = "role_id"
	this.IamAlias = "role_id"

	return &this
}

func (o GoogleCloudWriteAuthConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
