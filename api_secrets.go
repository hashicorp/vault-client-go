// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/vault-client-go/schema"
)

// Secrets is a simple wrapper around the client for Secrets requests
type Secrets struct {
	client *Client
}

// AliCloudConfigure
// alicloudMountPath: Path that the backend was mounted at
func (s *Secrets) AliCloudConfigure(ctx context.Context, alicloudMountPath string, request schema.AliCloudConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudDeleteConfiguration
// alicloudMountPath: Path that the backend was mounted at
func (s *Secrets) AliCloudDeleteConfiguration(ctx context.Context, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudDeleteRole Read, write and reference policies and roles that API keys or STS credentials can be made for.
// name: The name of the role.
// alicloudMountPath: Path that the backend was mounted at
func (s *Secrets) AliCloudDeleteRole(ctx context.Context, name string, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudGenerateCredentials Generate an API key or STS credential using the given role's configuration.'
// name: The name of the role.
// alicloudMountPath: Path that the backend was mounted at
func (s *Secrets) AliCloudGenerateCredentials(ctx context.Context, name string, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudListRoles List the existing roles in this backend.
// alicloudMountPath: Path that the backend was mounted at
func (s *Secrets) AliCloudListRoles(ctx context.Context, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudReadConfiguration
// alicloudMountPath: Path that the backend was mounted at
func (s *Secrets) AliCloudReadConfiguration(ctx context.Context, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudReadRole Read, write and reference policies and roles that API keys or STS credentials can be made for.
// name: The name of the role.
// alicloudMountPath: Path that the backend was mounted at
func (s *Secrets) AliCloudReadRole(ctx context.Context, name string, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudWriteRole Read, write and reference policies and roles that API keys or STS credentials can be made for.
// name: The name of the role.
// alicloudMountPath: Path that the backend was mounted at
func (s *Secrets) AliCloudWriteRole(ctx context.Context, name string, alicloudMountPath string, request schema.AliCloudWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureLease
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsConfigureLease(ctx context.Context, awsMountPath string, request schema.AwsConfigureLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureRootIamCredentials
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsConfigureRootIamCredentials(ctx context.Context, awsMountPath string, request schema.AwsConfigureRootIamCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsDeleteRole(ctx context.Context, name string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsGenerateCredentials
// name: Name of the role
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsGenerateCredentials(ctx context.Context, name string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsGenerateCredentials2
// name: Name of the role
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsGenerateCredentials2(ctx context.Context, name string, awsMountPath string, request schema.AwsGenerateCredentials2Request, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsGenerateStsCredentials
// name: Name of the role
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsGenerateStsCredentials(ctx context.Context, name string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/sts/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsGenerateStsCredentials2
// name: Name of the role
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsGenerateStsCredentials2(ctx context.Context, name string, awsMountPath string, request schema.AwsGenerateStsCredentials2Request, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/sts/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsListRoles List the existing roles in this backend
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsListRoles(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadLeaseConfiguration
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsReadLeaseConfiguration(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsReadRole(ctx context.Context, name string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadRootIamCredentialsConfiguration
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsReadRootIamCredentialsConfiguration(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsRotateRootIamCredentials
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsRotateRootIamCredentials(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsWriteRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
// awsMountPath: Path that the backend was mounted at
func (s *Secrets) AwsWriteRole(ctx context.Context, name string, awsMountPath string, request schema.AwsWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureConfigure
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureConfigure(ctx context.Context, azureMountPath string, request schema.AzureConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureDeleteConfiguration
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureDeleteConfiguration(ctx context.Context, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureDeleteRole Manage the Vault roles used to generate Azure credentials.
// name: Name of the role.
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureDeleteRole(ctx context.Context, name string, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureListRoles List existing roles.
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureListRoles(ctx context.Context, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadConfiguration
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureReadConfiguration(ctx context.Context, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadRole Manage the Vault roles used to generate Azure credentials.
// name: Name of the role.
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureReadRole(ctx context.Context, name string, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureRequestServicePrincipalCredentials
// role: Name of the Vault role
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureRequestServicePrincipalCredentials(ctx context.Context, role string, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureRotateRoot
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureRotateRoot(ctx context.Context, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureWriteRole Manage the Vault roles used to generate Azure credentials.
// name: Name of the role.
// azureMountPath: Path that the backend was mounted at
func (s *Secrets) AzureWriteRole(ctx context.Context, name string, azureMountPath string, request schema.AzureWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulConfigureAccess
// consulMountPath: Path that the backend was mounted at
func (s *Secrets) ConsulConfigureAccess(ctx context.Context, consulMountPath string, request schema.ConsulConfigureAccessRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulDeleteRole
// name: Name of the role.
// consulMountPath: Path that the backend was mounted at
func (s *Secrets) ConsulDeleteRole(ctx context.Context, name string, consulMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulGenerateCredentials
// role: Name of the role.
// consulMountPath: Path that the backend was mounted at
func (s *Secrets) ConsulGenerateCredentials(ctx context.Context, role string, consulMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulListRoles
// consulMountPath: Path that the backend was mounted at
func (s *Secrets) ConsulListRoles(ctx context.Context, consulMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulReadAccessConfiguration
// consulMountPath: Path that the backend was mounted at
func (s *Secrets) ConsulReadAccessConfiguration(ctx context.Context, consulMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulReadRole
// name: Name of the role.
// consulMountPath: Path that the backend was mounted at
func (s *Secrets) ConsulReadRole(ctx context.Context, name string, consulMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulWriteRole
// name: Name of the role.
// consulMountPath: Path that the backend was mounted at
func (s *Secrets) ConsulWriteRole(ctx context.Context, name string, consulMountPath string, request schema.ConsulWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CubbyholeDelete Deletes the secret at the specified location.
// path: Specifies the path of the secret.
func (s *Secrets) CubbyholeDelete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/cubbyhole/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CubbyholeRead Retrieve the secret at the specified location.
// path: Specifies the path of the secret.
func (s *Secrets) CubbyholeRead(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/cubbyhole/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CubbyholeWrite Store a secret at the specified location.
// path: Specifies the path of the secret.
func (s *Secrets) CubbyholeWrite(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/cubbyhole/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseConfigureConnection
// name: Name of this database connection
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseConfigureConnection(ctx context.Context, name string, databaseMountPath string, request schema.DatabaseConfigureConnectionRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/config/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseDeleteConnectionConfiguration
// name: Name of this database connection
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseDeleteConnectionConfiguration(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/config/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseDeleteRole Manage the roles that can be created with this backend.
// name: Name of the role.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseDeleteRole(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseDeleteStaticRole Manage the static roles that can be created with this backend.
// name: Name of the role.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseDeleteStaticRole(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseGenerateCredentials Request database credentials for a certain role.
// name: Name of the role.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseGenerateCredentials(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseListConnections Configure connection details to a database plugin.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseListConnections(ctx context.Context, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseListRoles Manage the roles that can be created with this backend.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseListRoles(ctx context.Context, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseListStaticRoles Manage the static roles that can be created with this backend.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseListStaticRoles(ctx context.Context, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-roles"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseReadConnectionConfiguration
// name: Name of this database connection
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseReadConnectionConfiguration(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/config/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseReadRole Manage the roles that can be created with this backend.
// name: Name of the role.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseReadRole(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseReadStaticRole Manage the static roles that can be created with this backend.
// name: Name of the role.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseReadStaticRole(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseReadStaticRoleCredentials Request database credentials for a certain static role. These credentials are rotated periodically.
// name: Name of the static role.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseReadStaticRoleCredentials(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseResetConnection Resets a database plugin.
// name: Name of this database connection
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseResetConnection(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/reset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseRotateRootCredentials
// name: Name of this database connection
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseRotateRootCredentials(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/rotate-root/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseRotateStaticRoleCredentials
// name: Name of the static role
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseRotateStaticRoleCredentials(ctx context.Context, name string, databaseMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseWriteRole Manage the roles that can be created with this backend.
// name: Name of the role.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseWriteRole(ctx context.Context, name string, databaseMountPath string, request schema.DatabaseWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// DatabaseWriteStaticRole Manage the static roles that can be created with this backend.
// name: Name of the role.
// databaseMountPath: Path that the backend was mounted at
func (s *Secrets) DatabaseWriteStaticRole(ctx context.Context, name string, databaseMountPath string, request schema.DatabaseWriteStaticRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudConfigure
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudConfigure(ctx context.Context, gcpMountPath string, request schema.GoogleCloudConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudDeleteImpersonatedAccount
// name: Required. Name to refer to this impersonated account in Vault. Cannot be updated.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudDeleteImpersonatedAccount(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudDeleteRoleset
// name: Required. Name of the role.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudDeleteRoleset(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudDeleteStaticAccount
// name: Required. Name to refer to this static account in Vault. Cannot be updated.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudDeleteStaticAccount(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateImpersonatedAccountAccessToken
// name: Required. Name of the impersonated account.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateImpersonatedAccountAccessToken(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateImpersonatedAccountAccessToken2
// name: Required. Name of the impersonated account.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateImpersonatedAccountAccessToken2(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateRolesetAccessToken
// roleset: Required. Name of the role set.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateRolesetAccessToken(ctx context.Context, roleset string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateRolesetAccessToken2
// roleset: Required. Name of the role set.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateRolesetAccessToken2(ctx context.Context, roleset string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/token/{roleset}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateRolesetAccessTokenWithParameters
// roleset: Required. Name of the role set.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateRolesetAccessTokenWithParameters(ctx context.Context, roleset string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateRolesetAccessTokenWithParameters2
// roleset: Required. Name of the role set.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateRolesetAccessTokenWithParameters2(ctx context.Context, roleset string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/token/{roleset}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateRolesetKey
// roleset: Required. Name of the role set.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateRolesetKey(ctx context.Context, roleset string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateRolesetKey2
// roleset: Required. Name of the role set.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateRolesetKey2(ctx context.Context, roleset string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/key/{roleset}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateRolesetKeyWithParameters
// roleset: Required. Name of the role set.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateRolesetKeyWithParameters(ctx context.Context, roleset string, gcpMountPath string, request schema.GoogleCloudGenerateRolesetKeyWithParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateRolesetKeyWithParameters2
// roleset: Required. Name of the role set.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateRolesetKeyWithParameters2(ctx context.Context, roleset string, gcpMountPath string, request schema.GoogleCloudGenerateRolesetKeyWithParameters2Request, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/key/{roleset}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateStaticAccountAccessToken
// name: Required. Name of the static account.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateStaticAccountAccessToken(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateStaticAccountAccessTokenWithParameters
// name: Required. Name of the static account.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateStaticAccountAccessTokenWithParameters(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateStaticAccountKey
// name: Required. Name of the static account.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateStaticAccountKey(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudGenerateStaticAccountKeyWithParameters
// name: Required. Name of the static account.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudGenerateStaticAccountKeyWithParameters(ctx context.Context, name string, gcpMountPath string, request schema.GoogleCloudGenerateStaticAccountKeyWithParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsConfigure
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsConfigure(ctx context.Context, gcpkmsMountPath string, request schema.GoogleCloudKmsConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsConfigureKey
// key: Name of the key in Vault.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsConfigureKey(ctx context.Context, key string, gcpkmsMountPath string, request schema.GoogleCloudKmsConfigureKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/config/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsDecrypt Decrypt a ciphertext value using a named key
// key: Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsDecrypt(ctx context.Context, key string, gcpkmsMountPath string, request schema.GoogleCloudKmsDecryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/decrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsDeleteConfiguration
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsDeleteConfiguration(ctx context.Context, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsDeleteKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsDeleteKey(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsDeregisterKey
// key: Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsDeregisterKey(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/deregister/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsDeregisterKey2
// key: Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsDeregisterKey2(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/deregister/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsEncrypt Encrypt a plaintext value using a named key
// key: Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsEncrypt(ctx context.Context, key string, gcpkmsMountPath string, request schema.GoogleCloudKmsEncryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/encrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsListKeys List named keys
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsListKeys(ctx context.Context, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsReadConfiguration
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsReadConfiguration(ctx context.Context, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsReadKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsReadKey(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsReadKeyConfiguration
// key: Name of the key in Vault.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsReadKeyConfiguration(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/config/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsReencrypt Re-encrypt existing ciphertext data to a new version
// key: Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsReencrypt(ctx context.Context, key string, gcpkmsMountPath string, request schema.GoogleCloudKmsReencryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/reencrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsRegisterKey Register an existing crypto key in Google Cloud KMS
// key: Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsRegisterKey(ctx context.Context, key string, gcpkmsMountPath string, request schema.GoogleCloudKmsRegisterKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/register/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsRetrievePublicKey Retrieve the public key associated with the named key
// key: Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsRetrievePublicKey(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/pubkey/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsRotateKey Rotate a crypto key to a new primary version
// key: Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsRotateKey(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/rotate/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsSign Signs a message or digest using a named key
// key: Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsSign(ctx context.Context, key string, gcpkmsMountPath string, request schema.GoogleCloudKmsSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/sign/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsTrimKeyVersions
// key: Name of the key in Vault.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsTrimKeyVersions(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/trim/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsTrimKeyVersions2
// key: Name of the key in Vault.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsTrimKeyVersions2(ctx context.Context, key string, gcpkmsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/trim/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsVerify Verify a signature using a named key
// key: Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsVerify(ctx context.Context, key string, gcpkmsMountPath string, request schema.GoogleCloudKmsVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/verify/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKmsWriteKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
// gcpkmsMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudKmsWriteKey(ctx context.Context, key string, gcpkmsMountPath string, request schema.GoogleCloudKmsWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListImpersonatedAccounts
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudListImpersonatedAccounts(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListImpersonatedAccounts2
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudListImpersonatedAccounts2(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-accounts"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListRolesets
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudListRolesets(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListRolesets2
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudListRolesets2(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/rolesets"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListStaticAccounts
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudListStaticAccounts(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListStaticAccounts2
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudListStaticAccounts2(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-accounts"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadConfiguration
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudReadConfiguration(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadImpersonatedAccount
// name: Required. Name to refer to this impersonated account in Vault. Cannot be updated.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudReadImpersonatedAccount(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadRoleset
// name: Required. Name of the role.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudReadRoleset(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadStaticAccount
// name: Required. Name to refer to this static account in Vault. Cannot be updated.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudReadStaticAccount(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudRotateRoleset
// name: Name of the role.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudRotateRoleset(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}/rotate"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudRotateRolesetKey
// name: Name of the role.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudRotateRolesetKey(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}/rotate-key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudRotateRootCredentials
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudRotateRootCredentials(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudRotateStaticAccountKey
// name: Name of the account.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudRotateStaticAccountKey(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/rotate-key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteImpersonatedAccount
// name: Required. Name to refer to this impersonated account in Vault. Cannot be updated.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudWriteImpersonatedAccount(ctx context.Context, name string, gcpMountPath string, request schema.GoogleCloudWriteImpersonatedAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteRoleset
// name: Required. Name of the role.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudWriteRoleset(ctx context.Context, name string, gcpMountPath string, request schema.GoogleCloudWriteRolesetRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteStaticAccount
// name: Required. Name to refer to this static account in Vault. Cannot be updated.
// gcpMountPath: Path that the backend was mounted at
func (s *Secrets) GoogleCloudWriteStaticAccount(ctx context.Context, name string, gcpMountPath string, request schema.GoogleCloudWriteStaticAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesCheckConfiguration
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesCheckConfiguration(ctx context.Context, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/check"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesConfigure
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesConfigure(ctx context.Context, kubernetesMountPath string, request schema.KubernetesConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesDeleteConfiguration
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesDeleteConfiguration(ctx context.Context, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesDeleteRole
// name: Name of the role
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesDeleteRole(ctx context.Context, name string, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesGenerateCredentials
// name: Name of the Vault role
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesGenerateCredentials(ctx context.Context, name string, kubernetesMountPath string, request schema.KubernetesGenerateCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesListRoles
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesListRoles(ctx context.Context, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesReadConfiguration
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesReadConfiguration(ctx context.Context, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesReadRole
// name: Name of the role
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesReadRole(ctx context.Context, name string, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesWriteRole
// name: Name of the role
// kubernetesMountPath: Path that the backend was mounted at
func (s *Secrets) KubernetesWriteRole(ctx context.Context, name string, kubernetesMountPath string, request schema.KubernetesWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV1Delete
// path: Location of the secret.
// kvV1MountPath: Path that the backend was mounted at
func (s *Secrets) KvV1Delete(ctx context.Context, path string, kvV1MountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v1_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v1_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v1")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV1Read
// path: Location of the secret.
// kvV1MountPath: Path that the backend was mounted at
func (s *Secrets) KvV1Read(ctx context.Context, path string, kvV1MountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v1_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v1_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v1")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV1Write
// path: Location of the secret.
// kvV1MountPath: Path that the backend was mounted at
func (s *Secrets) KvV1Write(ctx context.Context, path string, kvV1MountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v1_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v1_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v1")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2Configure Configure backend level settings that are applied to every key in the key-value store.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2Configure(ctx context.Context, kvV2MountPath string, request schema.KvV2ConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2Delete
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2Delete(ctx context.Context, path string, kvV2MountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2DeleteMetadata
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2DeleteMetadata(ctx context.Context, path string, kvV2MountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2DeleteVersions
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2DeleteVersions(ctx context.Context, path string, kvV2MountPath string, request schema.KvV2DeleteVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/delete/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2DestroyVersions
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2DestroyVersions(ctx context.Context, path string, kvV2MountPath string, request schema.KvV2DestroyVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/destroy/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2Read
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2Read(ctx context.Context, path string, kvV2MountPath string, options ...RequestOption) (*Response[schema.KvV2ReadResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.KvV2ReadResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2ReadConfiguration Read the backend level settings.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2ReadConfiguration(ctx context.Context, kvV2MountPath string, options ...RequestOption) (*Response[schema.KvV2ReadConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.KvV2ReadConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2ReadMetadata
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2ReadMetadata(ctx context.Context, path string, kvV2MountPath string, options ...RequestOption) (*Response[schema.KvV2ReadMetadataResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.KvV2ReadMetadataResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2ReadSubkeys
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2ReadSubkeys(ctx context.Context, path string, kvV2MountPath string, options ...RequestOption) (*Response[schema.KvV2ReadSubkeysResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/subkeys/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.KvV2ReadSubkeysResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2UndeleteVersions
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2UndeleteVersions(ctx context.Context, path string, kvV2MountPath string, request schema.KvV2UndeleteVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/undelete/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2Write
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2Write(ctx context.Context, path string, kvV2MountPath string, request schema.KvV2WriteRequest, options ...RequestOption) (*Response[schema.KvV2WriteResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.KvV2WriteResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2WriteMetadata
// path: Location of the secret.
// kvV2MountPath: Path that the backend was mounted at
func (s *Secrets) KvV2WriteMetadata(ctx context.Context, path string, kvV2MountPath string, request schema.KvV2WriteMetadataRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv-v2_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv-v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapConfigure
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapConfigure(ctx context.Context, ldapMountPath string, request schema.LdapConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapDeleteConfiguration
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapDeleteConfiguration(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapDeleteDynamicRole
// name: Name of the role (lowercase)
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapDeleteDynamicRole(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapDeleteStaticRole
// name: Name of the role
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapDeleteStaticRole(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLibraryCheckIn Check service accounts in to the library.
// name: Name of the set.
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapLibraryCheckIn(ctx context.Context, name string, ldapMountPath string, request schema.LdapLibraryCheckInRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLibraryCheckOut Check a service account out from the library.
// name: Name of the set
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapLibraryCheckOut(ctx context.Context, name string, ldapMountPath string, request schema.LdapLibraryCheckOutRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/check-out"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLibraryCheckStatus Check the status of the service accounts in a library set.
// name: Name of the set.
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapLibraryCheckStatus(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/status"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLibraryConfigure Update a library set.
// name: Name of the set.
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapLibraryConfigure(ctx context.Context, name string, ldapMountPath string, request schema.LdapLibraryConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLibraryDelete Delete a library set.
// name: Name of the set.
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapLibraryDelete(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLibraryForceCheckIn Check service accounts in to the library.
// name: Name of the set.
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapLibraryForceCheckIn(ctx context.Context, name string, ldapMountPath string, request schema.LdapLibraryForceCheckInRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/manage/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLibraryList
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapLibraryList(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLibraryRead Read a library set.
// name: Name of the set.
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapLibraryRead(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapListDynamicRoles
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapListDynamicRoles(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapListStaticRoles
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapListStaticRoles(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapReadConfiguration
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapReadConfiguration(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapReadDynamicRole
// name: Name of the role (lowercase)
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapReadDynamicRole(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapReadStaticRole
// name: Name of the role
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapReadStaticRole(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapRequestDynamicRoleCredentials
// name: Name of the dynamic role.
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapRequestDynamicRoleCredentials(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapRequestStaticRoleCredentials
// name: Name of the static role.
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapRequestStaticRoleCredentials(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-cred/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapRotateRootCredentials
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapRotateRootCredentials(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapRotateStaticRole
// name: Name of the static role
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapRotateStaticRole(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapWriteDynamicRole
// name: Name of the role (lowercase)
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapWriteDynamicRole(ctx context.Context, name string, ldapMountPath string, request schema.LdapWriteDynamicRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapWriteStaticRole
// name: Name of the role
// ldapMountPath: Path that the backend was mounted at
func (s *Secrets) LdapWriteStaticRole(ctx context.Context, name string, ldapMountPath string, request schema.LdapWriteStaticRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDbAtlasConfigure
// mongodbatlasMountPath: Path that the backend was mounted at
func (s *Secrets) MongoDbAtlasConfigure(ctx context.Context, mongodbatlasMountPath string, request schema.MongoDbAtlasConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDbAtlasDeleteRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
// mongodbatlasMountPath: Path that the backend was mounted at
func (s *Secrets) MongoDbAtlasDeleteRole(ctx context.Context, name string, mongodbatlasMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDbAtlasGenerateCredentials
// name: Name of the role
// mongodbatlasMountPath: Path that the backend was mounted at
func (s *Secrets) MongoDbAtlasGenerateCredentials(ctx context.Context, name string, mongodbatlasMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDbAtlasGenerateCredentials2
// name: Name of the role
// mongodbatlasMountPath: Path that the backend was mounted at
func (s *Secrets) MongoDbAtlasGenerateCredentials2(ctx context.Context, name string, mongodbatlasMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDbAtlasListRoles List the existing roles in this backend
// mongodbatlasMountPath: Path that the backend was mounted at
func (s *Secrets) MongoDbAtlasListRoles(ctx context.Context, mongodbatlasMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDbAtlasReadConfiguration
// mongodbatlasMountPath: Path that the backend was mounted at
func (s *Secrets) MongoDbAtlasReadConfiguration(ctx context.Context, mongodbatlasMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDbAtlasReadRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
// mongodbatlasMountPath: Path that the backend was mounted at
func (s *Secrets) MongoDbAtlasReadRole(ctx context.Context, name string, mongodbatlasMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDbAtlasWriteRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
// mongodbatlasMountPath: Path that the backend was mounted at
func (s *Secrets) MongoDbAtlasWriteRole(ctx context.Context, name string, mongodbatlasMountPath string, request schema.MongoDbAtlasWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadConfigureAccess
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadConfigureAccess(ctx context.Context, nomadMountPath string, request schema.NomadConfigureAccessRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadConfigureLease
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadConfigureLease(ctx context.Context, nomadMountPath string, request schema.NomadConfigureLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadDeleteAccessConfiguration
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadDeleteAccessConfiguration(ctx context.Context, nomadMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadDeleteLeaseConfiguration
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadDeleteLeaseConfiguration(ctx context.Context, nomadMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadDeleteRole
// name: Name of the role
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadDeleteRole(ctx context.Context, name string, nomadMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadGenerateCredentials
// name: Name of the role
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadGenerateCredentials(ctx context.Context, name string, nomadMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadListRoles
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadListRoles(ctx context.Context, nomadMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadReadAccessConfiguration
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadReadAccessConfiguration(ctx context.Context, nomadMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadReadLeaseConfiguration
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadReadLeaseConfiguration(ctx context.Context, nomadMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadReadRole
// name: Name of the role
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadReadRole(ctx context.Context, name string, nomadMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadWriteRole
// name: Name of the role
// nomadMountPath: Path that the backend was mounted at
func (s *Secrets) NomadWriteRole(ctx context.Context, name string, nomadMountPath string, request schema.NomadWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiConfigureAutoTidy
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiConfigureAutoTidy(ctx context.Context, pkiMountPath string, request schema.PkiConfigureAutoTidyRequest, options ...RequestOption) (*Response[schema.PkiConfigureAutoTidyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/auto-tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiConfigureAutoTidyResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiConfigureCa
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiConfigureCa(ctx context.Context, pkiMountPath string, request schema.PkiConfigureCaRequest, options ...RequestOption) (*Response[schema.PkiConfigureCaResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiConfigureCaResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiConfigureCluster
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiConfigureCluster(ctx context.Context, pkiMountPath string, request schema.PkiConfigureClusterRequest, options ...RequestOption) (*Response[schema.PkiConfigureClusterResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/cluster"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiConfigureClusterResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiConfigureCrl
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiConfigureCrl(ctx context.Context, pkiMountPath string, request schema.PkiConfigureCrlRequest, options ...RequestOption) (*Response[schema.PkiConfigureCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiConfigureCrlResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiConfigureIssuers
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiConfigureIssuers(ctx context.Context, pkiMountPath string, request schema.PkiConfigureIssuersRequest, options ...RequestOption) (*Response[schema.PkiConfigureIssuersResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiConfigureIssuersResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiConfigureKeys
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiConfigureKeys(ctx context.Context, pkiMountPath string, request schema.PkiConfigureKeysRequest, options ...RequestOption) (*Response[schema.PkiConfigureKeysResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiConfigureKeysResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiConfigureUrls
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiConfigureUrls(ctx context.Context, pkiMountPath string, request schema.PkiConfigureUrlsRequest, options ...RequestOption) (*Response[schema.PkiConfigureUrlsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/urls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiConfigureUrlsResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiCrossSignIntermediate
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiCrossSignIntermediate(ctx context.Context, pkiMountPath string, request schema.PkiCrossSignIntermediateRequest, options ...RequestOption) (*Response[schema.PkiCrossSignIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/cross-sign"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiCrossSignIntermediateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiDeleteIssuer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiDeleteIssuer(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiDeleteKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiDeleteKey(ctx context.Context, keyRef string, pkiMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiDeleteRole
// name: Name of the role
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiDeleteRole(ctx context.Context, name string, pkiMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiDeleteRoot
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiDeleteRoot(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateExportedKey
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiGenerateExportedKey(ctx context.Context, pkiMountPath string, request schema.PkiGenerateExportedKeyRequest, options ...RequestOption) (*Response[schema.PkiGenerateExportedKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/generate/exported"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiGenerateExportedKeyResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateIntermediate
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiGenerateIntermediate(ctx context.Context, exported string, pkiMountPath string, request schema.PkiGenerateIntermediateRequest, options ...RequestOption) (*Response[schema.PkiGenerateIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/generate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiGenerateIntermediateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateInternalKey
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiGenerateInternalKey(ctx context.Context, pkiMountPath string, request schema.PkiGenerateInternalKeyRequest, options ...RequestOption) (*Response[schema.PkiGenerateInternalKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/generate/internal"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiGenerateInternalKeyResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateKmsKey
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiGenerateKmsKey(ctx context.Context, pkiMountPath string, request schema.PkiGenerateKmsKeyRequest, options ...RequestOption) (*Response[schema.PkiGenerateKmsKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/generate/kms"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiGenerateKmsKeyResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiGenerateRoot(ctx context.Context, exported string, pkiMountPath string, request schema.PkiGenerateRootRequest, options ...RequestOption) (*Response[schema.PkiGenerateRootResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/generate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiGenerateRootResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiImportKey
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiImportKey(ctx context.Context, pkiMountPath string, request schema.PkiImportKeyRequest, options ...RequestOption) (*Response[schema.PkiImportKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/import"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiImportKeyResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssueWithRole
// role: The desired role with configuration for this request
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssueWithRole(ctx context.Context, role string, pkiMountPath string, request schema.PkiIssueWithRoleRequest, options ...RequestOption) (*Response[schema.PkiIssueWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssueWithRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerIssueWithRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerIssueWithRole(ctx context.Context, issuerRef string, role string, pkiMountPath string, request schema.PkiIssuerIssueWithRoleRequest, options ...RequestOption) (*Response[schema.PkiIssuerIssueWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuerIssueWithRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerReadCrl
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerReadCrl(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiIssuerReadCrlResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerReadCrlDelta
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerReadCrlDelta(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlDeltaResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiIssuerReadCrlDeltaResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerReadCrlDeltaDer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerReadCrlDeltaDer(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlDeltaDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/delta/der"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiIssuerReadCrlDeltaDerResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerReadCrlDeltaPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerReadCrlDeltaPem(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlDeltaPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/delta/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiIssuerReadCrlDeltaPemResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerReadCrlDer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerReadCrlDer(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/der"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiIssuerReadCrlDerResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerReadCrlPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerReadCrlPem(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiIssuerReadCrlPemResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerResignCrls
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerResignCrls(ctx context.Context, issuerRef string, pkiMountPath string, request schema.PkiIssuerResignCrlsRequest, options ...RequestOption) (*Response[schema.PkiIssuerResignCrlsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/resign-crls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuerResignCrlsResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerSignIntermediate
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerSignIntermediate(ctx context.Context, issuerRef string, pkiMountPath string, request schema.PkiIssuerSignIntermediateRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-intermediate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuerSignIntermediateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerSignRevocationList
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerSignRevocationList(ctx context.Context, issuerRef string, pkiMountPath string, request schema.PkiIssuerSignRevocationListRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignRevocationListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-revocation-list"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuerSignRevocationListResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerSignSelfIssued
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerSignSelfIssued(ctx context.Context, issuerRef string, pkiMountPath string, request schema.PkiIssuerSignSelfIssuedRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignSelfIssuedResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-self-issued"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuerSignSelfIssuedResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerSignVerbatim
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerSignVerbatim(ctx context.Context, issuerRef string, pkiMountPath string, request schema.PkiIssuerSignVerbatimRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignVerbatimResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuerSignVerbatimResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerSignVerbatimWithRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerSignVerbatimWithRole(ctx context.Context, issuerRef string, role string, pkiMountPath string, request schema.PkiIssuerSignVerbatimWithRoleRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignVerbatimWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuerSignVerbatimWithRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuerSignWithRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuerSignWithRole(ctx context.Context, issuerRef string, role string, pkiMountPath string, request schema.PkiIssuerSignWithRoleRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuerSignWithRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuersGenerateIntermediate
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuersGenerateIntermediate(ctx context.Context, exported string, pkiMountPath string, request schema.PkiIssuersGenerateIntermediateRequest, options ...RequestOption) (*Response[schema.PkiIssuersGenerateIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/generate/intermediate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuersGenerateIntermediateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuersGenerateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuersGenerateRoot(ctx context.Context, exported string, pkiMountPath string, request schema.PkiIssuersGenerateRootRequest, options ...RequestOption) (*Response[schema.PkiIssuersGenerateRootResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/generate/root/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuersGenerateRootResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuersImportBundle
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuersImportBundle(ctx context.Context, pkiMountPath string, request schema.PkiIssuersImportBundleRequest, options ...RequestOption) (*Response[schema.PkiIssuersImportBundleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/import/bundle"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuersImportBundleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuersImportCert
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuersImportCert(ctx context.Context, pkiMountPath string, request schema.PkiIssuersImportCertRequest, options ...RequestOption) (*Response[schema.PkiIssuersImportCertResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/import/cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuersImportCertResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiIssuersRotateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiIssuersRotateRoot(ctx context.Context, exported string, pkiMountPath string, request schema.PkiIssuersRotateRootRequest, options ...RequestOption) (*Response[schema.PkiIssuersRotateRootResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/rotate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiIssuersRotateRootResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiListCerts
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiListCerts(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiListCertsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/certs"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.PkiListCertsResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiListIssuers
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiListIssuers(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiListIssuersResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.PkiListIssuersResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiListKeys
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiListKeys(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiListKeysResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.PkiListKeysResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiListRevokedCerts
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiListRevokedCerts(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiListRevokedCertsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/certs/revoked"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.PkiListRevokedCertsResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiListRoles
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiListRoles(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiListRolesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.PkiListRolesResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiQueryOcsp
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiQueryOcsp(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ocsp"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiQueryOcspWithGetReq
// req: base-64 encoded ocsp request
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiQueryOcspWithGetReq(ctx context.Context, req string, pkiMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ocsp/{req}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"req"+"}", url.PathEscape(req), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadAutoTidyConfiguration
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadAutoTidyConfiguration(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadAutoTidyConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/auto-tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadAutoTidyConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCaChainPem
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCaChainPem(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCaChainPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca_chain"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCaChainPemResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCaDer
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCaDer(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCaDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCaDerResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCaPem
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCaPem(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCaPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCaPemResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCert
// serial: Certificate serial number, in colon- or hyphen-separated octal
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCert(ctx context.Context, serial string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCertResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCertResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCertCaChain
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCertCaChain(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCertCaChainResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/ca_chain"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCertCaChainResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCertCrl
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCertCrl(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCertCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCertCrlResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCertDeltaCrl
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCertDeltaCrl(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCertDeltaCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/delta-crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCertDeltaCrlResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCertRawDer
// serial: Certificate serial number, in colon- or hyphen-separated octal
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCertRawDer(ctx context.Context, serial string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCertRawDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}/raw"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCertRawDerResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCertRawPem
// serial: Certificate serial number, in colon- or hyphen-separated octal
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCertRawPem(ctx context.Context, serial string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCertRawPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}/raw/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCertRawPemResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadClusterConfiguration
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadClusterConfiguration(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadClusterConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/cluster"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadClusterConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCrlConfiguration
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCrlConfiguration(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCrlConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCrlConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCrlDelta
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCrlDelta(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCrlDeltaResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCrlDeltaResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCrlDeltaPem
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCrlDeltaPem(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCrlDeltaPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/delta/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCrlDeltaPemResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCrlDer
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCrlDer(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCrlDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCrlDerResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadCrlPem
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadCrlPem(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadCrlPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadCrlPemResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadIssuer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadIssuer(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadIssuerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadIssuerResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadIssuerDer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadIssuerDer(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadIssuerDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/der"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadIssuerDerResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadIssuerJson
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadIssuerJson(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadIssuerJsonResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/json"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadIssuerJsonResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadIssuerPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadIssuerPem(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadIssuerPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadIssuerPemResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadIssuersConfiguration
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadIssuersConfiguration(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadIssuersConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadIssuersConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadKey(ctx context.Context, keyRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadKeyResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadKeysConfiguration
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadKeysConfiguration(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadKeysConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadKeysConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadRole
// name: Name of the role
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadRole(ctx context.Context, name string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadRoleResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadUrlsConfiguration
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReadUrlsConfiguration(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiReadUrlsConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/urls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiReadUrlsConfigurationResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReplaceRoot
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiReplaceRoot(ctx context.Context, pkiMountPath string, request schema.PkiReplaceRootRequest, options ...RequestOption) (*Response[schema.PkiReplaceRootResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/replace"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiReplaceRootResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiRevoke
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiRevoke(ctx context.Context, pkiMountPath string, request schema.PkiRevokeRequest, options ...RequestOption) (*Response[schema.PkiRevokeResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/revoke"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiRevokeResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiRevokeIssuer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiRevokeIssuer(ctx context.Context, issuerRef string, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiRevokeIssuerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/revoke"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiRevokeIssuerResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiRevokeWithKey
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiRevokeWithKey(ctx context.Context, pkiMountPath string, request schema.PkiRevokeWithKeyRequest, options ...RequestOption) (*Response[schema.PkiRevokeWithKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/revoke-with-key"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiRevokeWithKeyResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiRootSignIntermediate
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiRootSignIntermediate(ctx context.Context, pkiMountPath string, request schema.PkiRootSignIntermediateRequest, options ...RequestOption) (*Response[schema.PkiRootSignIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/sign-intermediate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiRootSignIntermediateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiRootSignSelfIssued
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiRootSignSelfIssued(ctx context.Context, pkiMountPath string, request schema.PkiRootSignSelfIssuedRequest, options ...RequestOption) (*Response[schema.PkiRootSignSelfIssuedResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/sign-self-issued"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiRootSignSelfIssuedResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiRotateCrl
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiRotateCrl(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiRotateCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/rotate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiRotateCrlResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiRotateDeltaCrl
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiRotateDeltaCrl(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiRotateDeltaCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/rotate-delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiRotateDeltaCrlResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiSetSignedIntermediate
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiSetSignedIntermediate(ctx context.Context, pkiMountPath string, request schema.PkiSetSignedIntermediateRequest, options ...RequestOption) (*Response[schema.PkiSetSignedIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/set-signed"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiSetSignedIntermediateResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiSignVerbatim
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiSignVerbatim(ctx context.Context, pkiMountPath string, request schema.PkiSignVerbatimRequest, options ...RequestOption) (*Response[schema.PkiSignVerbatimResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign-verbatim"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiSignVerbatimResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiSignVerbatimWithRole
// role: The desired role with configuration for this request
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiSignVerbatimWithRole(ctx context.Context, role string, pkiMountPath string, request schema.PkiSignVerbatimWithRoleRequest, options ...RequestOption) (*Response[schema.PkiSignVerbatimWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign-verbatim/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiSignVerbatimWithRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiSignWithRole
// role: The desired role with configuration for this request
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiSignWithRole(ctx context.Context, role string, pkiMountPath string, request schema.PkiSignWithRoleRequest, options ...RequestOption) (*Response[schema.PkiSignWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiSignWithRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiTidy
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiTidy(ctx context.Context, pkiMountPath string, request schema.PkiTidyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiTidyCancel
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiTidyCancel(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiTidyCancelResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy-cancel"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiTidyCancelResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiTidyStatus
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiTidyStatus(ctx context.Context, pkiMountPath string, options ...RequestOption) (*Response[schema.PkiTidyStatusResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy-status"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[schema.PkiTidyStatusResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiWriteIssuer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiWriteIssuer(ctx context.Context, issuerRef string, pkiMountPath string, request schema.PkiWriteIssuerRequest, options ...RequestOption) (*Response[schema.PkiWriteIssuerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiWriteIssuerResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiWriteKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiWriteKey(ctx context.Context, keyRef string, pkiMountPath string, request schema.PkiWriteKeyRequest, options ...RequestOption) (*Response[schema.PkiWriteKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiWriteKeyResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiWriteRole
// name: Name of the role
// pkiMountPath: Path that the backend was mounted at
func (s *Secrets) PkiWriteRole(ctx context.Context, name string, pkiMountPath string, request schema.PkiWriteRoleRequest, options ...RequestOption) (*Response[schema.PkiWriteRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[schema.PkiWriteRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMqConfigureConnection Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
// rabbitmqMountPath: Path that the backend was mounted at
func (s *Secrets) RabbitMqConfigureConnection(ctx context.Context, rabbitmqMountPath string, request schema.RabbitMqConfigureConnectionRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/connection"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMqConfigureLease
// rabbitmqMountPath: Path that the backend was mounted at
func (s *Secrets) RabbitMqConfigureLease(ctx context.Context, rabbitmqMountPath string, request schema.RabbitMqConfigureLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMqDeleteRole Manage the roles that can be created with this backend.
// name: Name of the role.
// rabbitmqMountPath: Path that the backend was mounted at
func (s *Secrets) RabbitMqDeleteRole(ctx context.Context, name string, rabbitmqMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMqListRoles Manage the roles that can be created with this backend.
// rabbitmqMountPath: Path that the backend was mounted at
func (s *Secrets) RabbitMqListRoles(ctx context.Context, rabbitmqMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMqReadLeaseConfiguration
// rabbitmqMountPath: Path that the backend was mounted at
func (s *Secrets) RabbitMqReadLeaseConfiguration(ctx context.Context, rabbitmqMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMqReadRole Manage the roles that can be created with this backend.
// name: Name of the role.
// rabbitmqMountPath: Path that the backend was mounted at
func (s *Secrets) RabbitMqReadRole(ctx context.Context, name string, rabbitmqMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMqRequestCredentials Request RabbitMQ credentials for a certain role.
// name: Name of the role.
// rabbitmqMountPath: Path that the backend was mounted at
func (s *Secrets) RabbitMqRequestCredentials(ctx context.Context, name string, rabbitmqMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMqWriteRole Manage the roles that can be created with this backend.
// name: Name of the role.
// rabbitmqMountPath: Path that the backend was mounted at
func (s *Secrets) RabbitMqWriteRole(ctx context.Context, name string, rabbitmqMountPath string, request schema.RabbitMqWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SshConfigureCa
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshConfigureCa(ctx context.Context, sshMountPath string, request schema.SshConfigureCaRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SshConfigureZeroAddress
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshConfigureZeroAddress(ctx context.Context, sshMountPath string, request schema.SshConfigureZeroAddressRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SshDeleteCaConfiguration
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshDeleteCaConfiguration(ctx context.Context, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshDeleteRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshDeleteRole(ctx context.Context, role string, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshDeleteZeroAddressConfiguration
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshDeleteZeroAddressConfiguration(ctx context.Context, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshGenerateCredentials Creates a credential for establishing SSH connection with the remote host.
// role: [Required] Name of the role
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshGenerateCredentials(ctx context.Context, role string, sshMountPath string, request schema.SshGenerateCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SshIssueCertificate
// role: The desired role with configuration for this request.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshIssueCertificate(ctx context.Context, role string, sshMountPath string, request schema.SshIssueCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SshListRoles Manage the 'roles' that can be created with this backend.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshListRoles(ctx context.Context, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshListRolesByIp List all the roles associated with the given IP address.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshListRolesByIp(ctx context.Context, sshMountPath string, request schema.SshListRolesByIpRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/lookup"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SshReadCaConfiguration
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshReadCaConfiguration(ctx context.Context, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshReadPublicKey Retrieve the public key.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshReadPublicKey(ctx context.Context, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/public_key"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshReadRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshReadRole(ctx context.Context, role string, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshReadZeroAddressConfiguration
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshReadZeroAddressConfiguration(ctx context.Context, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshSignCertificate Request signing an SSH key using a certain role with the provided details.
// role: The desired role with configuration for this request.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshSignCertificate(ctx context.Context, role string, sshMountPath string, request schema.SshSignCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SshTidyDynamicHostKeys This endpoint removes the stored host keys used for the removed Dynamic Key feature, if present.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshTidyDynamicHostKeys(ctx context.Context, sshMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/tidy/dynamic-keys"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SshVerifyOtp Validate the OTP provided by Vault SSH Agent.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshVerifyOtp(ctx context.Context, sshMountPath string, request schema.SshVerifyOtpRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/verify"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SshWriteRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
// sshMountPath: Path that the backend was mounted at
func (s *Secrets) SshWriteRole(ctx context.Context, role string, sshMountPath string, request schema.SshWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudConfigure
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudConfigure(ctx context.Context, terraformMountPath string, request schema.TerraformCloudConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudDeleteConfiguration
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudDeleteConfiguration(ctx context.Context, terraformMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudDeleteRole
// name: Name of the role
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudDeleteRole(ctx context.Context, name string, terraformMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudGenerateCredentials
// name: Name of the role
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudGenerateCredentials(ctx context.Context, name string, terraformMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudGenerateCredentials2
// name: Name of the role
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudGenerateCredentials2(ctx context.Context, name string, terraformMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudListRoles
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudListRoles(ctx context.Context, terraformMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudReadConfiguration
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudReadConfiguration(ctx context.Context, terraformMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudReadRole
// name: Name of the role
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudReadRole(ctx context.Context, name string, terraformMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudRotateRole
// name: Name of the team or organization role
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudRotateRole(ctx context.Context, name string, terraformMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformCloudWriteRole
// name: Name of the role
// terraformMountPath: Path that the backend was mounted at
func (s *Secrets) TerraformCloudWriteRole(ctx context.Context, name string, terraformMountPath string, request schema.TerraformCloudWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TotpCreateKey
// name: Name of the key.
// totpMountPath: Path that the backend was mounted at
func (s *Secrets) TotpCreateKey(ctx context.Context, name string, totpMountPath string, request schema.TotpCreateKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TotpDeleteKey
// name: Name of the key.
// totpMountPath: Path that the backend was mounted at
func (s *Secrets) TotpDeleteKey(ctx context.Context, name string, totpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TotpGenerateCode
// name: Name of the key.
// totpMountPath: Path that the backend was mounted at
func (s *Secrets) TotpGenerateCode(ctx context.Context, name string, totpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/code/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TotpListKeys Manage the keys that can be created with this backend.
// totpMountPath: Path that the backend was mounted at
func (s *Secrets) TotpListKeys(ctx context.Context, totpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TotpReadKey
// name: Name of the key.
// totpMountPath: Path that the backend was mounted at
func (s *Secrets) TotpReadKey(ctx context.Context, name string, totpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TotpValidateCode
// name: Name of the key.
// totpMountPath: Path that the backend was mounted at
func (s *Secrets) TotpValidateCode(ctx context.Context, name string, totpMountPath string, request schema.TotpValidateCodeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/code/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitBackUpKey Backup the named key
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitBackUpKey(ctx context.Context, name string, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/backup/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitConfigureCache Configures a new cache of the specified size
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitConfigureCache(ctx context.Context, transitMountPath string, request schema.TransitConfigureCacheRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/cache-config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitConfigureKey Configure a named encryption key
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitConfigureKey(ctx context.Context, name string, transitMountPath string, request schema.TransitConfigureKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitConfigureKeys
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitConfigureKeys(ctx context.Context, transitMountPath string, request schema.TransitConfigureKeysRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitCreateKey
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitCreateKey(ctx context.Context, name string, transitMountPath string, request schema.TransitCreateKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitDecrypt Decrypt a ciphertext value using a named key
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitDecrypt(ctx context.Context, name string, transitMountPath string, request schema.TransitDecryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/decrypt/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitDeleteKey
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitDeleteKey(ctx context.Context, name string, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitEncrypt Encrypt a plaintext value or a batch of plaintext blocks using a named key
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitEncrypt(ctx context.Context, name string, transitMountPath string, request schema.TransitEncryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/encrypt/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitExportKey Export named encryption or signing key
// name: Name of the key
// type_: Type of key to export (encryption-key, signing-key, hmac-key)
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitExportKey(ctx context.Context, name string, type_ string, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/export/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitExportKeyVersion Export named encryption or signing key
// name: Name of the key
// type_: Type of key to export (encryption-key, signing-key, hmac-key)
// version: Version of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitExportKeyVersion(ctx context.Context, name string, type_ string, version string, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/export/{type}/{name}/{version}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)
	requestPath = strings.Replace(requestPath, "{"+"version"+"}", url.PathEscape(version), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateDataKey Generate a data key
// name: The backend key used for encrypting the data key
// plaintext: \&quot;plaintext\&quot; will return the key in both plaintext and ciphertext; \&quot;wrapped\&quot; will return the ciphertext only.
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitGenerateDataKey(ctx context.Context, name string, plaintext string, transitMountPath string, request schema.TransitGenerateDataKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/datakey/{plaintext}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"plaintext"+"}", url.PathEscape(plaintext), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateHmac Generate an HMAC for input data using the named key
// name: The key to use for the HMAC function
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitGenerateHmac(ctx context.Context, name string, transitMountPath string, request schema.TransitGenerateHmacRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hmac/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateHmacWithAlgorithm Generate an HMAC for input data using the named key
// name: The key to use for the HMAC function
// urlalgorithm: Algorithm to use (POST URL parameter)
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitGenerateHmacWithAlgorithm(ctx context.Context, name string, urlalgorithm string, transitMountPath string, request schema.TransitGenerateHmacWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hmac/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateRandom Generate random bytes
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitGenerateRandom(ctx context.Context, transitMountPath string, request schema.TransitGenerateRandomRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateRandomWithBytes Generate random bytes
// urlbytes: The number of bytes to generate (POST URL parameter)
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitGenerateRandomWithBytes(ctx context.Context, urlbytes string, transitMountPath string, request schema.TransitGenerateRandomWithBytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateRandomWithSource Generate random bytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitGenerateRandomWithSource(ctx context.Context, source string, transitMountPath string, request schema.TransitGenerateRandomWithSourceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{source}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateRandomWithSourceAndBytes Generate random bytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
// urlbytes: The number of bytes to generate (POST URL parameter)
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitGenerateRandomWithSourceAndBytes(ctx context.Context, source string, urlbytes string, transitMountPath string, request schema.TransitGenerateRandomWithSourceAndBytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{source}/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitHash Generate a hash sum for input data
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitHash(ctx context.Context, transitMountPath string, request schema.TransitHashRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hash"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitHashWithAlgorithm Generate a hash sum for input data
// urlalgorithm: Algorithm to use (POST URL parameter)
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitHashWithAlgorithm(ctx context.Context, urlalgorithm string, transitMountPath string, request schema.TransitHashWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hash/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitImportKey Imports an externally-generated key into a new transit key
// name: The name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitImportKey(ctx context.Context, name string, transitMountPath string, request schema.TransitImportKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/import"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitImportKeyVersion Imports an externally-generated key into an existing imported key
// name: The name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitImportKeyVersion(ctx context.Context, name string, transitMountPath string, request schema.TransitImportKeyVersionRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/import_version"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitListKeys Managed named encryption keys
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitListKeys(ctx context.Context, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitReadCacheConfiguration Returns the size of the active cache
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitReadCacheConfiguration(ctx context.Context, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/cache-config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitReadKey
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitReadKey(ctx context.Context, name string, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitReadKeysConfiguration
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitReadKeysConfiguration(ctx context.Context, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitReadWrappingKey Returns the public key to use for wrapping imported keys
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitReadWrappingKey(ctx context.Context, transitMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/wrapping_key"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitRestoreAndRenameKey Restore the named key
// name: If set, this will be the name of the restored key.
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitRestoreAndRenameKey(ctx context.Context, name string, transitMountPath string, request schema.TransitRestoreAndRenameKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/restore/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitRestoreKey Restore the named key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitRestoreKey(ctx context.Context, transitMountPath string, request schema.TransitRestoreKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/restore"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitRewrap Rewrap ciphertext
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitRewrap(ctx context.Context, name string, transitMountPath string, request schema.TransitRewrapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/rewrap/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitRotateKey Rotate named encryption key
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitRotateKey(ctx context.Context, name string, transitMountPath string, request schema.TransitRotateKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/rotate"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitSign Generate a signature for input data using the named key
// name: The key to use
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitSign(ctx context.Context, name string, transitMountPath string, request schema.TransitSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/sign/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitSignWithAlgorithm Generate a signature for input data using the named key
// name: The key to use
// urlalgorithm: Hash algorithm to use (POST URL parameter)
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitSignWithAlgorithm(ctx context.Context, name string, urlalgorithm string, transitMountPath string, request schema.TransitSignWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/sign/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitTrimKey Trim key versions of a named key
// name: Name of the key
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitTrimKey(ctx context.Context, name string, transitMountPath string, request schema.TransitTrimKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/trim"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitVerify Verify a signature or HMAC for input data created using the named key
// name: The key to use
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitVerify(ctx context.Context, name string, transitMountPath string, request schema.TransitVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/verify/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitVerifyWithAlgorithm Verify a signature or HMAC for input data created using the named key
// name: The key to use
// urlalgorithm: Hash algorithm to use (POST URL parameter)
// transitMountPath: Path that the backend was mounted at
func (s *Secrets) TransitVerifyWithAlgorithm(ctx context.Context, name string, urlalgorithm string, transitMountPath string, request schema.TransitVerifyWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/verify/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}
