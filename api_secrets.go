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
func (s *Secrets) AliCloudConfigure(ctx context.Context, request schema.AliCloudConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AliCloudDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AliCloudDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AliCloudGenerateCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AliCloudListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) AliCloudReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AliCloudReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AliCloudWriteRole(ctx context.Context, name string, request schema.AliCloudWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AwsConfigureLease(ctx context.Context, request schema.AwsConfigureLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AwsConfigureRootIamCredentials(ctx context.Context, request schema.AwsConfigureRootIamCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
// name: Name of the role
func (s *Secrets) AwsDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// AwsDeleteStaticRolesName
// name: The name of this role.
func (s *Secrets) AwsDeleteStaticRolesName(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
// roleArn: ARN of role to assume when credential_type is assumed_role
// roleSessionName: Session name to use when assuming role. Max chars: 64
// ttl: Lifetime of the returned credentials in seconds
func (s *Secrets) AwsGenerateCredentials(ctx context.Context, name string, roleArn string, roleSessionName string, ttl string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("role_arn", parameterToString(roleArn))
	requestQueryParameters.Add("role_session_name", parameterToString(roleSessionName))
	requestQueryParameters.Add("ttl", parameterToString(ttl))

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

// AwsGenerateCredentialsWithParameters
// name: Name of the role
func (s *Secrets) AwsGenerateCredentialsWithParameters(ctx context.Context, name string, request schema.AwsGenerateCredentialsWithParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
// roleArn: ARN of role to assume when credential_type is assumed_role
// roleSessionName: Session name to use when assuming role. Max chars: 64
// ttl: Lifetime of the returned credentials in seconds
func (s *Secrets) AwsGenerateStsCredentials(ctx context.Context, name string, roleArn string, roleSessionName string, ttl string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/sts/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("role_arn", parameterToString(roleArn))
	requestQueryParameters.Add("role_session_name", parameterToString(roleSessionName))
	requestQueryParameters.Add("ttl", parameterToString(ttl))

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

// AwsGenerateStsCredentialsWithParameters
// name: Name of the role
func (s *Secrets) AwsGenerateStsCredentialsWithParameters(ctx context.Context, name string, request schema.AwsGenerateStsCredentialsWithParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/sts/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AwsListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) AwsReadLeaseConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
// name: Name of the role
func (s *Secrets) AwsReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AwsReadRootIamCredentialsConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// AwsReadStaticCredsName
// name: The name of this role.
func (s *Secrets) AwsReadStaticCredsName(ctx context.Context, name string, options ...RequestOption) (*Response[schema.AwsReadStaticCredsNameResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/static-creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AwsReadStaticCredsNameResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadStaticRolesName
// name: The name of this role.
func (s *Secrets) AwsReadStaticRolesName(ctx context.Context, name string, options ...RequestOption) (*Response[schema.AwsReadStaticRolesNameResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AwsReadStaticRolesNameResponse](
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
func (s *Secrets) AwsRotateRootIamCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
// name: Name of the role
func (s *Secrets) AwsWriteRole(ctx context.Context, name string, request schema.AwsWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// AwsWriteStaticRolesName
// name: The name of this role.
func (s *Secrets) AwsWriteStaticRolesName(ctx context.Context, name string, request schema.AwsWriteStaticRolesNameRequest, options ...RequestOption) (*Response[schema.AwsWriteStaticRolesNameResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.AwsWriteStaticRolesNameResponse](
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
func (s *Secrets) AzureConfigure(ctx context.Context, request schema.AzureConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AzureDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AzureDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AzureListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) AzureReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AzureReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AzureRequestServicePrincipalCredentials(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AzureRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) AzureWriteRole(ctx context.Context, name string, request schema.AzureWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) ConsulConfigureAccess(ctx context.Context, request schema.ConsulConfigureAccessRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) ConsulDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) ConsulGenerateCredentials(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) ConsulListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) ConsulReadAccessConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) ConsulReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) ConsulWriteRole(ctx context.Context, name string, request schema.ConsulWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// CubbyholeList List secret entries at the specified location.
// Folders are suffixed with /. The input must be a folder; list on a file will not return a value. The values themselves are not accessible via this command.
// path: Specifies the path of the secret.
func (s *Secrets) CubbyholeList(ctx context.Context, path string, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/cubbyhole/{path}/"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		s.client,
		http.MethodGet,
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

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) CubbyholeWrite(ctx context.Context, path string, request map[string]interface{}, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/cubbyhole/{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// DatabaseConfigureConnection
// name: Name of this database connection
func (s *Secrets) DatabaseConfigureConnection(ctx context.Context, name string, request schema.DatabaseConfigureConnectionRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/config/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseDeleteConnectionConfiguration(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/config/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseDeleteStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseGenerateCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseListConnections(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/config/"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) DatabaseListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) DatabaseListStaticRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-roles/"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) DatabaseReadConnectionConfiguration(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/config/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseReadStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseReadStaticRoleCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseResetConnection(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/reset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseRotateRootCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/rotate-root/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseRotateStaticRoleCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseWriteRole(ctx context.Context, name string, request schema.DatabaseWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) DatabaseWriteStaticRole(ctx context.Context, name string, request schema.DatabaseWriteStaticRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{database_mount_path}/static-roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"database_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("database")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudConfigure(ctx context.Context, request schema.GoogleCloudConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudDeleteImpersonatedAccount(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudDeleteRoleset(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudDeleteStaticAccount(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudGenerateImpersonatedAccountAccessToken(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// GoogleCloudGenerateRolesetAccessToken
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudGenerateRolesetAccessToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudGenerateRolesetKey(ctx context.Context, roleset string, request schema.GoogleCloudGenerateRolesetKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudGenerateStaticAccountAccessToken(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudGenerateStaticAccountKey(ctx context.Context, name string, request schema.GoogleCloudGenerateStaticAccountKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsConfigure(ctx context.Context, request schema.GoogleCloudKmsConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsConfigureKey(ctx context.Context, key string, request schema.GoogleCloudKmsConfigureKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/config/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsDecrypt(ctx context.Context, key string, request schema.GoogleCloudKmsDecryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/decrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsDeleteKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsDeregisterKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/deregister/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// GoogleCloudKmsEncrypt Encrypt a plaintext value using a named key
// key: Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (s *Secrets) GoogleCloudKmsEncrypt(ctx context.Context, key string, request schema.GoogleCloudKmsEncryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/encrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsListKeys(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) GoogleCloudKmsReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsReadKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsReadKeyConfiguration(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/config/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsReencrypt(ctx context.Context, key string, request schema.GoogleCloudKmsReencryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/reencrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsRegisterKey(ctx context.Context, key string, request schema.GoogleCloudKmsRegisterKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/register/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsRetrievePublicKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/pubkey/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsRotateKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/rotate/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsSign(ctx context.Context, key string, request schema.GoogleCloudKmsSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/sign/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsTrimKeyVersions(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/trim/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// GoogleCloudKmsVerify Verify a signature using a named key
// key: Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (s *Secrets) GoogleCloudKmsVerify(ctx context.Context, key string, request schema.GoogleCloudKmsVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/verify/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudKmsWriteKey(ctx context.Context, key string, request schema.GoogleCloudKmsWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudListImpersonatedAccounts(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) GoogleCloudListRolesets(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) GoogleCloudListStaticAccounts(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) GoogleCloudReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudReadImpersonatedAccount(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudReadRoleset(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudReadStaticAccount(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudRotateRoleset(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}/rotate"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudRotateRolesetKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}/rotate-key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudRotateRootCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudRotateStaticAccountKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/rotate-key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudWriteImpersonatedAccount(ctx context.Context, name string, request schema.GoogleCloudWriteImpersonatedAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/impersonated-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudWriteRoleset(ctx context.Context, name string, request schema.GoogleCloudWriteRolesetRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) GoogleCloudWriteStaticAccount(ctx context.Context, name string, request schema.GoogleCloudWriteStaticAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KubernetesCheckConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/check"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KubernetesConfigure(ctx context.Context, request schema.KubernetesConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KubernetesDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KubernetesDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KubernetesGenerateCredentials(ctx context.Context, name string, request schema.KubernetesGenerateCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KubernetesListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) KubernetesReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KubernetesReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KubernetesWriteRole(ctx context.Context, name string, request schema.KubernetesWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV1Delete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v1_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v1_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v1")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// KvV1List
// path: Location of the secret.
func (s *Secrets) KvV1List(ctx context.Context, path string, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v1_mount_path}/{path}/"
	requestPath = strings.Replace(requestPath, "{"+"kv_v1_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v1")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV1Read
// path: Location of the secret.
func (s *Secrets) KvV1Read(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v1_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v1_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v1")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV1Write(ctx context.Context, path string, request map[string]interface{}, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v1_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v1_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v1")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// KvV2Configure Configure backend level settings that are applied to every key in the key-value store.
func (s *Secrets) KvV2Configure(ctx context.Context, request schema.KvV2ConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2Delete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// KvV2DeleteMetadataAndAllVersions
// path: Location of the secret.
func (s *Secrets) KvV2DeleteMetadataAndAllVersions(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2DeleteVersions(ctx context.Context, path string, request schema.KvV2DeleteVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/delete/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2DestroyVersions(ctx context.Context, path string, request schema.KvV2DestroyVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/destroy/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// KvV2List
// path: Location of the secret.
func (s *Secrets) KvV2List(ctx context.Context, path string, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/metadata/{path}/"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KvV2Read
// path: Location of the secret.
func (s *Secrets) KvV2Read(ctx context.Context, path string, options ...RequestOption) (*Response[schema.KvV2ReadResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2ReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.KvV2ReadConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2ReadMetadata(ctx context.Context, path string, options ...RequestOption) (*Response[schema.KvV2ReadMetadataResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2ReadSubkeys(ctx context.Context, path string, options ...RequestOption) (*Response[schema.KvV2ReadSubkeysResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/subkeys/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2UndeleteVersions(ctx context.Context, path string, request schema.KvV2UndeleteVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/undelete/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2Write(ctx context.Context, path string, request schema.KvV2WriteRequest, options ...RequestOption) (*Response[schema.KvV2WriteResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) KvV2WriteMetadata(ctx context.Context, path string, request schema.KvV2WriteMetadataRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_v2_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_v2_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv-v2")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapConfigure(ctx context.Context, request schema.LdapConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapDeleteDynamicRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapDeleteStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapLibraryCheckIn(ctx context.Context, name string, request schema.LdapLibraryCheckInRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapLibraryCheckOut(ctx context.Context, name string, request schema.LdapLibraryCheckOutRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/check-out"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapLibraryCheckStatus(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/status"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapLibraryConfigure(ctx context.Context, name string, request schema.LdapLibraryConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapLibraryDelete(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapLibraryForceCheckIn(ctx context.Context, name string, request schema.LdapLibraryForceCheckInRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/manage/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapLibraryList(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) LdapLibraryRead(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapListDynamicRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) LdapListStaticRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) LdapReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapReadDynamicRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapReadStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapRequestDynamicRoleCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapRequestStaticRoleCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-cred/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapRotateRootCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapRotateStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapWriteDynamicRole(ctx context.Context, name string, request schema.LdapWriteDynamicRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) LdapWriteStaticRole(ctx context.Context, name string, request schema.LdapWriteStaticRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) MongoDbAtlasConfigure(ctx context.Context, request schema.MongoDbAtlasConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) MongoDbAtlasDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) MongoDbAtlasGenerateCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// MongoDbAtlasListRoles List the existing roles in this backend
func (s *Secrets) MongoDbAtlasListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) MongoDbAtlasReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) MongoDbAtlasReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) MongoDbAtlasWriteRole(ctx context.Context, name string, request schema.MongoDbAtlasWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadConfigureAccess(ctx context.Context, request schema.NomadConfigureAccessRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadConfigureLease(ctx context.Context, request schema.NomadConfigureLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadDeleteAccessConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadDeleteLeaseConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadGenerateCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) NomadReadAccessConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadReadLeaseConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) NomadWriteRole(ctx context.Context, name string, request schema.NomadWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiConfigureAcme
func (s *Secrets) PkiConfigureAcme(ctx context.Context, request schema.PkiConfigureAcmeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/acme"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiConfigureAutoTidy(ctx context.Context, request schema.PkiConfigureAutoTidyRequest, options ...RequestOption) (*Response[schema.PkiConfigureAutoTidyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/auto-tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiConfigureCa(ctx context.Context, request schema.PkiConfigureCaRequest, options ...RequestOption) (*Response[schema.PkiConfigureCaResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiConfigureCluster(ctx context.Context, request schema.PkiConfigureClusterRequest, options ...RequestOption) (*Response[schema.PkiConfigureClusterResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/cluster"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiConfigureCrl(ctx context.Context, request schema.PkiConfigureCrlRequest, options ...RequestOption) (*Response[schema.PkiConfigureCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiConfigureIssuers(ctx context.Context, request schema.PkiConfigureIssuersRequest, options ...RequestOption) (*Response[schema.PkiConfigureIssuersResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiConfigureKeys(ctx context.Context, request schema.PkiConfigureKeysRequest, options ...RequestOption) (*Response[schema.PkiConfigureKeysResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiConfigureUrls(ctx context.Context, request schema.PkiConfigureUrlsRequest, options ...RequestOption) (*Response[schema.PkiConfigureUrlsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/urls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiCrossSignIntermediate(ctx context.Context, request schema.PkiCrossSignIntermediateRequest, options ...RequestOption) (*Response[schema.PkiCrossSignIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/cross-sign"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiDeleteEabKey
// keyId: EAB key identifier
func (s *Secrets) PkiDeleteEabKey(ctx context.Context, keyId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/eab/{key_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_id"+"}", url.PathEscape(keyId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiDeleteIssuer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PkiDeleteIssuer(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiDeleteKey(ctx context.Context, keyRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiDeleteRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiGenerateEabKey
func (s *Secrets) PkiGenerateEabKey(ctx context.Context, options ...RequestOption) (*Response[schema.PkiGenerateEabKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/new-eab"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PkiGenerateEabKeyResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateEabKeyForIssuer
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiGenerateEabKeyForIssuer(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiGenerateEabKeyForIssuerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/new-eab"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PkiGenerateEabKeyForIssuerResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateEabKeyForIssuerAndRole
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiGenerateEabKeyForIssuerAndRole(ctx context.Context, issuerRef string, role string, options ...RequestOption) (*Response[schema.PkiGenerateEabKeyForIssuerAndRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/new-eab"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PkiGenerateEabKeyForIssuerAndRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateEabKeyForRole
// role: The desired role for the acme request
func (s *Secrets) PkiGenerateEabKeyForRole(ctx context.Context, role string, options ...RequestOption) (*Response[schema.PkiGenerateEabKeyForRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/new-eab"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.PkiGenerateEabKeyForRoleResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiGenerateExportedKey
func (s *Secrets) PkiGenerateExportedKey(ctx context.Context, request schema.PkiGenerateExportedKeyRequest, options ...RequestOption) (*Response[schema.PkiGenerateExportedKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/generate/exported"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiGenerateIntermediate(ctx context.Context, exported string, request schema.PkiGenerateIntermediateRequest, options ...RequestOption) (*Response[schema.PkiGenerateIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/generate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiGenerateInternalKey(ctx context.Context, request schema.PkiGenerateInternalKeyRequest, options ...RequestOption) (*Response[schema.PkiGenerateInternalKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/generate/internal"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiGenerateKmsKey(ctx context.Context, request schema.PkiGenerateKmsKeyRequest, options ...RequestOption) (*Response[schema.PkiGenerateKmsKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/generate/kms"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiGenerateRoot(ctx context.Context, exported string, request schema.PkiGenerateRootRequest, options ...RequestOption) (*Response[schema.PkiGenerateRootResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/generate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiImportKey(ctx context.Context, request schema.PkiImportKeyRequest, options ...RequestOption) (*Response[schema.PkiImportKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/import"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssueWithRole(ctx context.Context, role string, request schema.PkiIssueWithRoleRequest, options ...RequestOption) (*Response[schema.PkiIssueWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerIssueWithRole(ctx context.Context, issuerRef string, role string, request schema.PkiIssuerIssueWithRoleRequest, options ...RequestOption) (*Response[schema.PkiIssuerIssueWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerReadCrl(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerReadCrlDelta(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlDeltaResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerReadCrlDeltaDer(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlDeltaDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/delta/der"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerReadCrlDeltaPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlDeltaPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/delta/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerReadCrlDer(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/der"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerReadCrlPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiIssuerReadCrlPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/crl/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerResignCrls(ctx context.Context, issuerRef string, request schema.PkiIssuerResignCrlsRequest, options ...RequestOption) (*Response[schema.PkiIssuerResignCrlsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/resign-crls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerSignIntermediate(ctx context.Context, issuerRef string, request schema.PkiIssuerSignIntermediateRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-intermediate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerSignRevocationList(ctx context.Context, issuerRef string, request schema.PkiIssuerSignRevocationListRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignRevocationListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-revocation-list"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerSignSelfIssued(ctx context.Context, issuerRef string, request schema.PkiIssuerSignSelfIssuedRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignSelfIssuedResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-self-issued"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerSignVerbatim(ctx context.Context, issuerRef string, request schema.PkiIssuerSignVerbatimRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignVerbatimResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerSignVerbatimWithRole(ctx context.Context, issuerRef string, role string, request schema.PkiIssuerSignVerbatimWithRoleRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignVerbatimWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuerSignWithRole(ctx context.Context, issuerRef string, role string, request schema.PkiIssuerSignWithRoleRequest, options ...RequestOption) (*Response[schema.PkiIssuerSignWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuersGenerateIntermediate(ctx context.Context, exported string, request schema.PkiIssuersGenerateIntermediateRequest, options ...RequestOption) (*Response[schema.PkiIssuersGenerateIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/generate/intermediate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuersGenerateRoot(ctx context.Context, exported string, request schema.PkiIssuersGenerateRootRequest, options ...RequestOption) (*Response[schema.PkiIssuersGenerateRootResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/generate/root/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuersImportBundle(ctx context.Context, request schema.PkiIssuersImportBundleRequest, options ...RequestOption) (*Response[schema.PkiIssuersImportBundleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/import/bundle"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiIssuersImportCert(ctx context.Context, request schema.PkiIssuersImportCertRequest, options ...RequestOption) (*Response[schema.PkiIssuersImportCertResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/import/cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiListCerts
func (s *Secrets) PkiListCerts(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/certs/"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		s.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiListEabKeys
func (s *Secrets) PkiListEabKeys(ctx context.Context, options ...RequestOption) (*Response[schema.PkiListEabKeysResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/eab/"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.PkiListEabKeysResponse](
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
func (s *Secrets) PkiListIssuers(ctx context.Context, options ...RequestOption) (*Response[schema.PkiListIssuersResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

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
func (s *Secrets) PkiListKeys(ctx context.Context, options ...RequestOption) (*Response[schema.PkiListKeysResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

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
func (s *Secrets) PkiListRevokedCerts(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/certs/revoked/"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) PkiListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) PkiQueryOcsp(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ocsp"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiQueryOcspWithGetReq(ctx context.Context, req string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ocsp/{req}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"req"+"}", url.PathEscape(req), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadAcmeConfiguration
func (s *Secrets) PkiReadAcmeConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/acme"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadAcmeDirectory
func (s *Secrets) PkiReadAcmeDirectory(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/directory"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadAcmeNewNonce
func (s *Secrets) PkiReadAcmeNewNonce(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/new-nonce"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadAutoTidyConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadAutoTidyConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/auto-tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCaChainPem(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCaChainPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca_chain"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCaDer(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCaDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCaPem(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCaPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCert(ctx context.Context, serial string, options ...RequestOption) (*Response[schema.PkiReadCertResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCertCaChain(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCertCaChainResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/ca_chain"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCertCrl(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCertCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCertDeltaCrl(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCertDeltaCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/delta-crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCertRawDer(ctx context.Context, serial string, options ...RequestOption) (*Response[schema.PkiReadCertRawDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}/raw"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCertRawPem(ctx context.Context, serial string, options ...RequestOption) (*Response[schema.PkiReadCertRawPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}/raw/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadClusterConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadClusterConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/cluster"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCrlConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCrlConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCrlDelta(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCrlDeltaResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCrlDeltaPem(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCrlDeltaPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/delta/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCrlDer(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCrlDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadCrlPem(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadCrlPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadIssuer(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiReadIssuerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadIssuerDer(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiReadIssuerDerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/der"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadIssuerIssuerRefAcmeDirectory
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiReadIssuerIssuerRefAcmeDirectory(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/directory"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadIssuerIssuerRefAcmeNewNonce
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiReadIssuerIssuerRefAcmeNewNonce(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/new-nonce"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadIssuerIssuerRefRolesRoleAcmeDirectory
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiReadIssuerIssuerRefRolesRoleAcmeDirectory(ctx context.Context, issuerRef string, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/directory"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadIssuerIssuerRefRolesRoleAcmeNewNonce
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiReadIssuerIssuerRefRolesRoleAcmeNewNonce(ctx context.Context, issuerRef string, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/new-nonce"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadIssuerJson
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PkiReadIssuerJson(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiReadIssuerJsonResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/json"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadIssuerPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiReadIssuerPemResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadIssuersConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadIssuersConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadKey(ctx context.Context, keyRef string, options ...RequestOption) (*Response[schema.PkiReadKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadKeysConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadKeysConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[schema.PkiReadRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadRolesRoleAcmeDirectory
// role: The desired role for the acme request
func (s *Secrets) PkiReadRolesRoleAcmeDirectory(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/directory"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadRolesRoleAcmeNewNonce
// role: The desired role for the acme request
func (s *Secrets) PkiReadRolesRoleAcmeNewNonce(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/new-nonce"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiReadUrlsConfiguration
func (s *Secrets) PkiReadUrlsConfiguration(ctx context.Context, options ...RequestOption) (*Response[schema.PkiReadUrlsConfigurationResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/urls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiReplaceRoot(ctx context.Context, request schema.PkiReplaceRootRequest, options ...RequestOption) (*Response[schema.PkiReplaceRootResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/replace"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiRevoke(ctx context.Context, request schema.PkiRevokeRequest, options ...RequestOption) (*Response[schema.PkiRevokeResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/revoke"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiRevokeIssuer(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[schema.PkiRevokeIssuerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/revoke"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiRevokeWithKey(ctx context.Context, request schema.PkiRevokeWithKeyRequest, options ...RequestOption) (*Response[schema.PkiRevokeWithKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/revoke-with-key"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiRootSignIntermediate(ctx context.Context, request schema.PkiRootSignIntermediateRequest, options ...RequestOption) (*Response[schema.PkiRootSignIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/sign-intermediate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiRootSignSelfIssued(ctx context.Context, request schema.PkiRootSignSelfIssuedRequest, options ...RequestOption) (*Response[schema.PkiRootSignSelfIssuedResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/sign-self-issued"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiRotateCrl(ctx context.Context, options ...RequestOption) (*Response[schema.PkiRotateCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/rotate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiRotateDeltaCrl(ctx context.Context, options ...RequestOption) (*Response[schema.PkiRotateDeltaCrlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/rotate-delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiRotateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (s *Secrets) PkiRotateRoot(ctx context.Context, exported string, request schema.PkiRotateRootRequest, options ...RequestOption) (*Response[schema.PkiRotateRootResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/rotate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.PkiRotateRootResponse](
		ctx,
		s.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiSetSignedIntermediate
func (s *Secrets) PkiSetSignedIntermediate(ctx context.Context, request schema.PkiSetSignedIntermediateRequest, options ...RequestOption) (*Response[schema.PkiSetSignedIntermediateResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/set-signed"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiSignVerbatim(ctx context.Context, request schema.PkiSignVerbatimRequest, options ...RequestOption) (*Response[schema.PkiSignVerbatimResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign-verbatim"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiSignVerbatimWithRole(ctx context.Context, role string, request schema.PkiSignVerbatimWithRoleRequest, options ...RequestOption) (*Response[schema.PkiSignVerbatimWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign-verbatim/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiSignWithRole(ctx context.Context, role string, request schema.PkiSignWithRoleRequest, options ...RequestOption) (*Response[schema.PkiSignWithRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiTidy(ctx context.Context, request schema.PkiTidyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiTidyCancel(ctx context.Context, options ...RequestOption) (*Response[schema.PkiTidyCancelResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy-cancel"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiTidyStatus(ctx context.Context, options ...RequestOption) (*Response[schema.PkiTidyStatusResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy-status"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeAccountKid
// kid: The key identifier provided by the CA
func (s *Secrets) PkiWriteAcmeAccountKid(ctx context.Context, kid string, request schema.PkiWriteAcmeAccountKidRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/account/{kid}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"kid"+"}", url.PathEscape(kid), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeAuthorizationAuthId
// authId: ACME authorization identifier value
func (s *Secrets) PkiWriteAcmeAuthorizationAuthId(ctx context.Context, authId string, request schema.PkiWriteAcmeAuthorizationAuthIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/authorization/{auth_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"auth_id"+"}", url.PathEscape(authId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeChallengeAuthIdChallengeType
// authId: ACME authorization identifier value
// challengeType: ACME challenge type
func (s *Secrets) PkiWriteAcmeChallengeAuthIdChallengeType(ctx context.Context, authId string, challengeType string, request schema.PkiWriteAcmeChallengeAuthIdChallengeTypeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/challenge/{auth_id}/{challenge_type}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"auth_id"+"}", url.PathEscape(authId), -1)
	requestPath = strings.Replace(requestPath, "{"+"challenge_type"+"}", url.PathEscape(challengeType), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeNewAccount
func (s *Secrets) PkiWriteAcmeNewAccount(ctx context.Context, request schema.PkiWriteAcmeNewAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/new-account"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeNewOrder
func (s *Secrets) PkiWriteAcmeNewOrder(ctx context.Context, request schema.PkiWriteAcmeNewOrderRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/new-order"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeOrderOrderId
// orderId: The ACME order identifier to fetch
func (s *Secrets) PkiWriteAcmeOrderOrderId(ctx context.Context, orderId string, request schema.PkiWriteAcmeOrderOrderIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/order/{order_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeOrderOrderIdCert
// orderId: The ACME order identifier to fetch
func (s *Secrets) PkiWriteAcmeOrderOrderIdCert(ctx context.Context, orderId string, request schema.PkiWriteAcmeOrderOrderIdCertRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/order/{order_id}/cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeOrderOrderIdFinalize
// orderId: The ACME order identifier to fetch
func (s *Secrets) PkiWriteAcmeOrderOrderIdFinalize(ctx context.Context, orderId string, request schema.PkiWriteAcmeOrderOrderIdFinalizeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/order/{order_id}/finalize"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeOrders
func (s *Secrets) PkiWriteAcmeOrders(ctx context.Context, request schema.PkiWriteAcmeOrdersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/orders"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteAcmeRevokeCert
func (s *Secrets) PkiWriteAcmeRevokeCert(ctx context.Context, request schema.PkiWriteAcmeRevokeCertRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/acme/revoke-cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuer
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PkiWriteIssuer(ctx context.Context, issuerRef string, request schema.PkiWriteIssuerRequest, options ...RequestOption) (*Response[schema.PkiWriteIssuerResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeAccountKid
// issuerRef: Reference to an existing issuer name or issuer id
// kid: The key identifier provided by the CA
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeAccountKid(ctx context.Context, issuerRef string, kid string, request schema.PkiWriteIssuerIssuerRefAcmeAccountKidRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/account/{kid}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"kid"+"}", url.PathEscape(kid), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeAuthorizationAuthId
// authId: ACME authorization identifier value
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeAuthorizationAuthId(ctx context.Context, authId string, issuerRef string, request schema.PkiWriteIssuerIssuerRefAcmeAuthorizationAuthIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/authorization/{auth_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"auth_id"+"}", url.PathEscape(authId), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeChallengeAuthIdChallengeType
// authId: ACME authorization identifier value
// challengeType: ACME challenge type
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeChallengeAuthIdChallengeType(ctx context.Context, authId string, challengeType string, issuerRef string, request schema.PkiWriteIssuerIssuerRefAcmeChallengeAuthIdChallengeTypeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/challenge/{auth_id}/{challenge_type}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"auth_id"+"}", url.PathEscape(authId), -1)
	requestPath = strings.Replace(requestPath, "{"+"challenge_type"+"}", url.PathEscape(challengeType), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeNewAccount
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeNewAccount(ctx context.Context, issuerRef string, request schema.PkiWriteIssuerIssuerRefAcmeNewAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/new-account"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeNewOrder
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeNewOrder(ctx context.Context, issuerRef string, request schema.PkiWriteIssuerIssuerRefAcmeNewOrderRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/new-order"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeOrderOrderId
// issuerRef: Reference to an existing issuer name or issuer id
// orderId: The ACME order identifier to fetch
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeOrderOrderId(ctx context.Context, issuerRef string, orderId string, request schema.PkiWriteIssuerIssuerRefAcmeOrderOrderIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/order/{order_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeOrderOrderIdCert
// issuerRef: Reference to an existing issuer name or issuer id
// orderId: The ACME order identifier to fetch
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeOrderOrderIdCert(ctx context.Context, issuerRef string, orderId string, request schema.PkiWriteIssuerIssuerRefAcmeOrderOrderIdCertRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/order/{order_id}/cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeOrderOrderIdFinalize
// issuerRef: Reference to an existing issuer name or issuer id
// orderId: The ACME order identifier to fetch
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeOrderOrderIdFinalize(ctx context.Context, issuerRef string, orderId string, request schema.PkiWriteIssuerIssuerRefAcmeOrderOrderIdFinalizeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/order/{order_id}/finalize"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeOrders
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeOrders(ctx context.Context, issuerRef string, request schema.PkiWriteIssuerIssuerRefAcmeOrdersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/orders"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefAcmeRevokeCert
// issuerRef: Reference to an existing issuer name or issuer id
func (s *Secrets) PkiWriteIssuerIssuerRefAcmeRevokeCert(ctx context.Context, issuerRef string, request schema.PkiWriteIssuerIssuerRefAcmeRevokeCertRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/acme/revoke-cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeAccountKid
// issuerRef: Reference to an existing issuer name or issuer id
// kid: The key identifier provided by the CA
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeAccountKid(ctx context.Context, issuerRef string, kid string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeAccountKidRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/account/{kid}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"kid"+"}", url.PathEscape(kid), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeAuthorizationAuthId
// authId: ACME authorization identifier value
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeAuthorizationAuthId(ctx context.Context, authId string, issuerRef string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeAuthorizationAuthIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/authorization/{auth_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"auth_id"+"}", url.PathEscape(authId), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeChallengeAuthIdChallengeType
// authId: ACME authorization identifier value
// challengeType: ACME challenge type
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeChallengeAuthIdChallengeType(ctx context.Context, authId string, challengeType string, issuerRef string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeChallengeAuthIdChallengeTypeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/challenge/{auth_id}/{challenge_type}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"auth_id"+"}", url.PathEscape(authId), -1)
	requestPath = strings.Replace(requestPath, "{"+"challenge_type"+"}", url.PathEscape(challengeType), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccount
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccount(ctx context.Context, issuerRef string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/new-account"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeNewOrder
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeNewOrder(ctx context.Context, issuerRef string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeNewOrderRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/new-order"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderId
// issuerRef: Reference to an existing issuer name or issuer id
// orderId: The ACME order identifier to fetch
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderId(ctx context.Context, issuerRef string, orderId string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/order/{order_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderIdCert
// issuerRef: Reference to an existing issuer name or issuer id
// orderId: The ACME order identifier to fetch
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderIdCert(ctx context.Context, issuerRef string, orderId string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderIdCertRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/order/{order_id}/cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderIdFinalize
// issuerRef: Reference to an existing issuer name or issuer id
// orderId: The ACME order identifier to fetch
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderIdFinalize(ctx context.Context, issuerRef string, orderId string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeOrderOrderIdFinalizeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/order/{order_id}/finalize"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeOrders
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeOrders(ctx context.Context, issuerRef string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeOrdersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/orders"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteIssuerIssuerRefRolesRoleAcmeRevokeCert
// issuerRef: Reference to an existing issuer name or issuer id
// role: The desired role for the acme request
func (s *Secrets) PkiWriteIssuerIssuerRefRolesRoleAcmeRevokeCert(ctx context.Context, issuerRef string, role string, request schema.PkiWriteIssuerIssuerRefRolesRoleAcmeRevokeCertRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/roles/{role}/acme/revoke-cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
func (s *Secrets) PkiWriteKey(ctx context.Context, keyRef string, request schema.PkiWriteKeyRequest, options ...RequestOption) (*Response[schema.PkiWriteKeyResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) PkiWriteRole(ctx context.Context, name string, request schema.PkiWriteRoleRequest, options ...RequestOption) (*Response[schema.PkiWriteRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeAccountKid
// kid: The key identifier provided by the CA
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeAccountKid(ctx context.Context, kid string, role string, request schema.PkiWriteRolesRoleAcmeAccountKidRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/account/{kid}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"kid"+"}", url.PathEscape(kid), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeAuthorizationAuthId
// authId: ACME authorization identifier value
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeAuthorizationAuthId(ctx context.Context, authId string, role string, request schema.PkiWriteRolesRoleAcmeAuthorizationAuthIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/authorization/{auth_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"auth_id"+"}", url.PathEscape(authId), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeChallengeAuthIdChallengeType
// authId: ACME authorization identifier value
// challengeType: ACME challenge type
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeChallengeAuthIdChallengeType(ctx context.Context, authId string, challengeType string, role string, request schema.PkiWriteRolesRoleAcmeChallengeAuthIdChallengeTypeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/challenge/{auth_id}/{challenge_type}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"auth_id"+"}", url.PathEscape(authId), -1)
	requestPath = strings.Replace(requestPath, "{"+"challenge_type"+"}", url.PathEscape(challengeType), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeNewAccount
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeNewAccount(ctx context.Context, role string, request schema.PkiWriteRolesRoleAcmeNewAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/new-account"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeNewOrder
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeNewOrder(ctx context.Context, role string, request schema.PkiWriteRolesRoleAcmeNewOrderRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/new-order"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeOrderOrderId
// orderId: The ACME order identifier to fetch
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeOrderOrderId(ctx context.Context, orderId string, role string, request schema.PkiWriteRolesRoleAcmeOrderOrderIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/order/{order_id}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeOrderOrderIdCert
// orderId: The ACME order identifier to fetch
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeOrderOrderIdCert(ctx context.Context, orderId string, role string, request schema.PkiWriteRolesRoleAcmeOrderOrderIdCertRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/order/{order_id}/cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeOrderOrderIdFinalize
// orderId: The ACME order identifier to fetch
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeOrderOrderIdFinalize(ctx context.Context, orderId string, role string, request schema.PkiWriteRolesRoleAcmeOrderOrderIdFinalizeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/order/{order_id}/finalize"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"order_id"+"}", url.PathEscape(orderId), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeOrders
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeOrders(ctx context.Context, role string, request schema.PkiWriteRolesRoleAcmeOrdersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/orders"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// PkiWriteRolesRoleAcmeRevokeCert
// role: The desired role for the acme request
func (s *Secrets) PkiWriteRolesRoleAcmeRevokeCert(ctx context.Context, role string, request schema.PkiWriteRolesRoleAcmeRevokeCertRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{role}/acme/revoke-cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// RabbitMqConfigureConnection Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
func (s *Secrets) RabbitMqConfigureConnection(ctx context.Context, request schema.RabbitMqConfigureConnectionRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/connection"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) RabbitMqConfigureLease(ctx context.Context, request schema.RabbitMqConfigureLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) RabbitMqDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) RabbitMqListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) RabbitMqReadLeaseConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) RabbitMqReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) RabbitMqRequestCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) RabbitMqWriteRole(ctx context.Context, name string, request schema.RabbitMqWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshConfigureCa(ctx context.Context, request schema.SshConfigureCaRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshConfigureZeroAddress(ctx context.Context, request schema.SshConfigureZeroAddressRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshDeleteCaConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshDeleteRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshDeleteZeroAddressConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshGenerateCredentials(ctx context.Context, role string, request schema.SshGenerateCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshIssueCertificate(ctx context.Context, role string, request schema.SshIssueCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) SshListRolesByIp(ctx context.Context, request schema.SshListRolesByIpRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/lookup"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshReadCaConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshReadPublicKey(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/public_key"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshReadRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshReadZeroAddressConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshSignCertificate(ctx context.Context, role string, request schema.SshSignCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshTidyDynamicHostKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/tidy/dynamic-keys"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshVerifyOtp(ctx context.Context, request schema.SshVerifyOtpRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/verify"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) SshWriteRole(ctx context.Context, role string, request schema.SshWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TerraformCloudConfigure(ctx context.Context, request schema.TerraformCloudConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TerraformCloudDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TerraformCloudDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TerraformCloudGenerateCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// TerraformCloudListRoles
func (s *Secrets) TerraformCloudListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) TerraformCloudReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TerraformCloudReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TerraformCloudRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TerraformCloudWriteRole(ctx context.Context, name string, request schema.TerraformCloudWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TotpCreateKey(ctx context.Context, name string, request schema.TotpCreateKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TotpDeleteKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TotpGenerateCode(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/code/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TotpListKeys(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) TotpReadKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TotpValidateCode(ctx context.Context, name string, request schema.TotpValidateCodeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/code/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitBackUpKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/backup/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// TransitByokKey Securely export named encryption or signing key
// destination: Destination key to export to; usually the public wrapping key of another Transit instance.
// source: Source key to export; could be any present key within Transit.
func (s *Secrets) TransitByokKey(ctx context.Context, destination string, source string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/byok-export/{destination}/{source}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"destination"+"}", url.PathEscape(destination), -1)
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// TransitByokKeyVersion Securely export named encryption or signing key
// destination: Destination key to export to; usually the public wrapping key of another Transit instance.
// source: Source key to export; could be any present key within Transit.
// version: Optional version of the key to export, else all key versions are exported.
func (s *Secrets) TransitByokKeyVersion(ctx context.Context, destination string, source string, version string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/byok-export/{destination}/{source}/{version}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"destination"+"}", url.PathEscape(destination), -1)
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)
	requestPath = strings.Replace(requestPath, "{"+"version"+"}", url.PathEscape(version), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitConfigureCache(ctx context.Context, request schema.TransitConfigureCacheRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/cache-config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitConfigureKey(ctx context.Context, name string, request schema.TransitConfigureKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitConfigureKeys(ctx context.Context, request schema.TransitConfigureKeysRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitCreateKey(ctx context.Context, name string, request schema.TransitCreateKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitDecrypt(ctx context.Context, name string, request schema.TransitDecryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/decrypt/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitDeleteKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitEncrypt(ctx context.Context, name string, request schema.TransitEncryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/encrypt/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
// type_: Type of key to export (encryption-key, signing-key, hmac-key, public-key)
func (s *Secrets) TransitExportKey(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/export/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
// type_: Type of key to export (encryption-key, signing-key, hmac-key, public-key)
// version: Version of the key
func (s *Secrets) TransitExportKeyVersion(ctx context.Context, name string, type_ string, version string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/export/{type}/{name}/{version}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)
	requestPath = strings.Replace(requestPath, "{"+"version"+"}", url.PathEscape(version), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// TransitGenerateCsrForKey
// name: Name of the key
func (s *Secrets) TransitGenerateCsrForKey(ctx context.Context, name string, request schema.TransitGenerateCsrForKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/csr"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// TransitGenerateDataKey Generate a data key
// name: The backend key used for encrypting the data key
// plaintext: \&quot;plaintext\&quot; will return the key in both plaintext and ciphertext; \&quot;wrapped\&quot; will return the ciphertext only.
func (s *Secrets) TransitGenerateDataKey(ctx context.Context, name string, plaintext string, request schema.TransitGenerateDataKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/datakey/{plaintext}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"plaintext"+"}", url.PathEscape(plaintext), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitGenerateHmac(ctx context.Context, name string, request schema.TransitGenerateHmacRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hmac/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitGenerateHmacWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitGenerateHmacWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hmac/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitGenerateRandom(ctx context.Context, request schema.TransitGenerateRandomRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitGenerateRandomWithBytes(ctx context.Context, urlbytes string, request schema.TransitGenerateRandomWithBytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitGenerateRandomWithSource(ctx context.Context, source string, request schema.TransitGenerateRandomWithSourceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{source}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitGenerateRandomWithSourceAndBytes(ctx context.Context, source string, urlbytes string, request schema.TransitGenerateRandomWithSourceAndBytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{source}/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitHash(ctx context.Context, request schema.TransitHashRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hash"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitHashWithAlgorithm(ctx context.Context, urlalgorithm string, request schema.TransitHashWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hash/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitImportKey(ctx context.Context, name string, request schema.TransitImportKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/import"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitImportKeyVersion(ctx context.Context, name string, request schema.TransitImportKeyVersionRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/import_version"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitListKeys(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
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
func (s *Secrets) TransitReadCacheConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/cache-config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitReadKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitReadKeysConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitReadWrappingKey(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/wrapping_key"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitRestoreAndRenameKey(ctx context.Context, name string, request schema.TransitRestoreAndRenameKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/restore/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitRestoreKey(ctx context.Context, request schema.TransitRestoreKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/restore"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitRewrap(ctx context.Context, name string, request schema.TransitRewrapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/rewrap/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitRotateKey(ctx context.Context, name string, request schema.TransitRotateKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/rotate"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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

// TransitSetCertificateForKey
// name: Name of the key
func (s *Secrets) TransitSetCertificateForKey(ctx context.Context, name string, request schema.TransitSetCertificateForKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/set-certificate"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitSign(ctx context.Context, name string, request schema.TransitSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/sign/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitSignWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitSignWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/sign/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitTrimKey(ctx context.Context, name string, request schema.TransitTrimKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/trim"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitVerify(ctx context.Context, name string, request schema.TransitVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/verify/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
func (s *Secrets) TransitVerifyWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitVerifyWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/verify/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

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
