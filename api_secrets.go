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

// AWSConfigReadLease Configure the default lease information for generated credentials.
func (s *Secrets) AWSConfigReadLease(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSConfigReadRootIAMCredentials Configure the root credentials that are used to manage IAM.
func (s *Secrets) AWSConfigReadRootIAMCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSConfigRotateRootIAMCredentials
func (s *Secrets) AWSConfigRotateRootIAMCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSConfigWriteLease Configure the default lease information for generated credentials.
func (s *Secrets) AWSConfigWriteLease(ctx context.Context, request schema.AWSConfigWriteLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSConfigWriteRootIAMCredentials Configure the root credentials that are used to manage IAM.
func (s *Secrets) AWSConfigWriteRootIAMCredentials(ctx context.Context, request schema.AWSConfigWriteRootIAMCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSDeleteRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
func (s *Secrets) AWSDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSListRoles List the existing roles in this backend
func (s *Secrets) AWSListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSReadCredentials Generate AWS credentials from a specific Vault role.
func (s *Secrets) AWSReadCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/creds"
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

// AWSReadRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
func (s *Secrets) AWSReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSReadSecurityTokenService Generate AWS credentials from a specific Vault role.
// name: Name of the role
func (s *Secrets) AWSReadSecurityTokenService(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSWriteCredentials Generate AWS credentials from a specific Vault role.
func (s *Secrets) AWSWriteCredentials(ctx context.Context, request schema.AWSWriteCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/creds"
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

// AWSWriteRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
func (s *Secrets) AWSWriteRole(ctx context.Context, name string, request schema.AWSWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AWSWriteSecurityTokenService Generate AWS credentials from a specific Vault role.
// name: Name of the role
func (s *Secrets) AWSWriteSecurityTokenService(ctx context.Context, name string, request schema.AWSWriteSecurityTokenServiceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// ActiveDirectoryCheckInLibrary Check service accounts in to the library.
// name: Name of the set.
func (s *Secrets) ActiveDirectoryCheckInLibrary(ctx context.Context, name string, request schema.ActiveDirectoryCheckInLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryCheckInManageLibrary Check service accounts in to the library.
// name: Name of the set.
func (s *Secrets) ActiveDirectoryCheckInManageLibrary(ctx context.Context, name string, request schema.ActiveDirectoryCheckInManageLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/manage/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryCheckOutLibrary Check a service account out from the library.
// name: Name of the set
func (s *Secrets) ActiveDirectoryCheckOutLibrary(ctx context.Context, name string, request schema.ActiveDirectoryCheckOutLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}/check-out"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryDeleteConfig Configure the AD server to connect to, along with password options.
func (s *Secrets) ActiveDirectoryDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

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

// ActiveDirectoryDeleteLibrary Delete a library set.
// name: Name of the set.
func (s *Secrets) ActiveDirectoryDeleteLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryDeleteRole Manage roles to build links between Vault and Active Directory service accounts.
// name: Name of the role
func (s *Secrets) ActiveDirectoryDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryListLibraries
func (s *Secrets) ActiveDirectoryListLibraries(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

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

// ActiveDirectoryListRoles List the name of each role currently stored.
func (s *Secrets) ActiveDirectoryListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

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

// ActiveDirectoryReadConfig Configure the AD server to connect to, along with password options.
func (s *Secrets) ActiveDirectoryReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

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

// ActiveDirectoryReadCredentials
// name: Name of the role
func (s *Secrets) ActiveDirectoryReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryReadLibrary Read a library set.
// name: Name of the set.
func (s *Secrets) ActiveDirectoryReadLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryReadLibraryStatus Check the status of the service accounts in a library set.
// name: Name of the set.
func (s *Secrets) ActiveDirectoryReadLibraryStatus(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}/status"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryReadRole Manage roles to build links between Vault and Active Directory service accounts.
// name: Name of the role
func (s *Secrets) ActiveDirectoryReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryRotateRole
// name: Name of the static role
func (s *Secrets) ActiveDirectoryRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryRotateRoot
func (s *Secrets) ActiveDirectoryRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

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

// ActiveDirectoryWriteConfig Configure the AD server to connect to, along with password options.
func (s *Secrets) ActiveDirectoryWriteConfig(ctx context.Context, request schema.ActiveDirectoryWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

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

// ActiveDirectoryWriteLibrary Update a library set.
// name: Name of the set.
func (s *Secrets) ActiveDirectoryWriteLibrary(ctx context.Context, name string, request schema.ActiveDirectoryWriteLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// ActiveDirectoryWriteRole Manage roles to build links between Vault and Active Directory service accounts.
// name: Name of the role
func (s *Secrets) ActiveDirectoryWriteRole(ctx context.Context, name string, request schema.ActiveDirectoryWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
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

// AliCloudDeleteConfig Configure the access key and secret to use for RAM and STS calls.
func (s *Secrets) AliCloudDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) AliCloudDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AliCloudListRoles List the existing roles in this backend.
func (s *Secrets) AliCloudListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AliCloudReadConfig Configure the access key and secret to use for RAM and STS calls.
func (s *Secrets) AliCloudReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AliCloudReadCredentials Generate an API key or STS credential using the given role's configuration.'
// name: The name of the role.
func (s *Secrets) AliCloudReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AliCloudWriteConfig Configure the access key and secret to use for RAM and STS calls.
func (s *Secrets) AliCloudWriteConfig(ctx context.Context, request schema.AliCloudWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AzureDeleteConfig
func (s *Secrets) AzureDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) AzureDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) AzureListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AzureReadConfig
func (s *Secrets) AzureReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AzureReadCredentials
// role: Name of the Vault role
func (s *Secrets) AzureReadCredentials(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) AzureRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// AzureWriteConfig
func (s *Secrets) AzureWriteConfig(ctx context.Context, request schema.AzureWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) ConsulDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// ConsulListRoles
func (s *Secrets) ConsulListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// ConsulReadAccessConfig
func (s *Secrets) ConsulReadAccessConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// ConsulReadCredentials
// role: Name of the role.
func (s *Secrets) ConsulReadCredentials(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// ConsulWriteAccessConfig
func (s *Secrets) ConsulWriteAccessConfig(ctx context.Context, request schema.ConsulWriteAccessConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

	requestPath := "/v1/{cubbyhole_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"cubbyhole_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cubbyhole")), -1)
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

	requestPath := "/v1/{cubbyhole_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"cubbyhole_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cubbyhole")), -1)
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

	requestPath := "/v1/{cubbyhole_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"cubbyhole_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cubbyhole")), -1)
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
func (s *Secrets) GoogleCloudDeleteStaticAccount(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSDecrypt Decrypt a ciphertext value using a named key
// key: Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (s *Secrets) GoogleCloudKMSDecrypt(ctx context.Context, key string, request schema.GoogleCloudKMSDecryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSDeleteConfig Configure the GCP KMS secrets engine
func (s *Secrets) GoogleCloudKMSDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSDeleteKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
func (s *Secrets) GoogleCloudKMSDeleteKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSDeregisterKey Deregister an existing key in Vault
// key: Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
func (s *Secrets) GoogleCloudKMSDeregisterKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSEncrypt Encrypt a plaintext value using a named key
// key: Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (s *Secrets) GoogleCloudKMSEncrypt(ctx context.Context, key string, request schema.GoogleCloudKMSEncryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSListKeys List named keys
func (s *Secrets) GoogleCloudKMSListKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSReadConfig Configure the GCP KMS secrets engine
func (s *Secrets) GoogleCloudKMSReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSReadKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
func (s *Secrets) GoogleCloudKMSReadKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSReadKeyConfig Configure the key in Vault
// key: Name of the key in Vault.
func (s *Secrets) GoogleCloudKMSReadKeyConfig(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSReadPubkey Retrieve the public key associated with the named key
// key: Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS.
func (s *Secrets) GoogleCloudKMSReadPubkey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSReencrypt Re-encrypt existing ciphertext data to a new version
// key: Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS.
func (s *Secrets) GoogleCloudKMSReencrypt(ctx context.Context, key string, request schema.GoogleCloudKMSReencryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSRegisterKey Register an existing crypto key in Google Cloud KMS
// key: Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data.
func (s *Secrets) GoogleCloudKMSRegisterKey(ctx context.Context, key string, request schema.GoogleCloudKMSRegisterKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSRotateKey Rotate a crypto key to a new primary version
// key: Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key.
func (s *Secrets) GoogleCloudKMSRotateKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSSign Signs a message or digest using a named key
// key: Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (s *Secrets) GoogleCloudKMSSign(ctx context.Context, key string, request schema.GoogleCloudKMSSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSTrimKey Delete old crypto key versions from Google Cloud KMS
// key: Name of the key in Vault.
func (s *Secrets) GoogleCloudKMSTrimKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSVerify Verify a signature using a named key
// key: Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (s *Secrets) GoogleCloudKMSVerify(ctx context.Context, key string, request schema.GoogleCloudKMSVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSWriteConfig Configure the GCP KMS secrets engine
func (s *Secrets) GoogleCloudKMSWriteConfig(ctx context.Context, request schema.GoogleCloudKMSWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSWriteKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
func (s *Secrets) GoogleCloudKMSWriteKey(ctx context.Context, key string, request schema.GoogleCloudKMSWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudKMSWriteKeyConfig Configure the key in Vault
// key: Name of the key in Vault.
func (s *Secrets) GoogleCloudKMSWriteKeyConfig(ctx context.Context, key string, request schema.GoogleCloudKMSWriteKeyConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudListRolesets
func (s *Secrets) GoogleCloudListRolesets(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) GoogleCloudListStaticAccounts(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudReadConfig
func (s *Secrets) GoogleCloudReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudReadKey
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudReadKey(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudReadRolesetKey
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudReadRolesetKey(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudReadRolesetToken
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudReadRolesetToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudReadStaticAccountKey
// name: Required. Name of the static account.
func (s *Secrets) GoogleCloudReadStaticAccountKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudReadStaticAccountToken
// name: Required. Name of the static account.
func (s *Secrets) GoogleCloudReadStaticAccountToken(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudReadToken
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudReadToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) GoogleCloudRotateRolesetKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudRotateRoot
func (s *Secrets) GoogleCloudRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) GoogleCloudRotateStaticAccountKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudWriteConfig
func (s *Secrets) GoogleCloudWriteConfig(ctx context.Context, request schema.GoogleCloudWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudWriteKey
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudWriteKey(ctx context.Context, roleset string, request schema.GoogleCloudWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudWriteRolesetKey
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudWriteRolesetKey(ctx context.Context, roleset string, request schema.GoogleCloudWriteRolesetKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudWriteRolesetToken
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudWriteRolesetToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudWriteStaticAccountKey
// name: Required. Name of the static account.
func (s *Secrets) GoogleCloudWriteStaticAccountKey(ctx context.Context, name string, request schema.GoogleCloudWriteStaticAccountKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudWriteStaticAccountToken
// name: Required. Name of the static account.
func (s *Secrets) GoogleCloudWriteStaticAccountToken(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// GoogleCloudWriteToken
// roleset: Required. Name of the role set.
func (s *Secrets) GoogleCloudWriteToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// KVv1Delete Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
// path: Location of the secret.
func (s *Secrets) KVv1Delete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv")), -1)
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

// KVv1Read Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
// path: Location of the secret.
func (s *Secrets) KVv1Read(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv")), -1)
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

// KVv1Write Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
// path: Location of the secret.
func (s *Secrets) KVv1Write(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv")), -1)
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

// KVv2Delete Write, Patch, Read, and Delete data in the Key-Value Store.
// path: Location of the secret.
func (s *Secrets) KVv2Delete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
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

// KVv2DeleteMetadata Configures settings for the KV store
// path: Location of the secret.
func (s *Secrets) KVv2DeleteMetadata(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
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

// KVv2DeleteVersions Marks one or more versions as deleted in the KV store.
// path: Location of the secret.
func (s *Secrets) KVv2DeleteVersions(ctx context.Context, path string, request schema.KVv2DeleteVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/delete/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
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

// KVv2DestroyVersions Permanently removes one or more versions in the KV store
// path: Location of the secret.
func (s *Secrets) KVv2DestroyVersions(ctx context.Context, path string, request schema.KVv2DestroyVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/destroy/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
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

// KVv2Read Write, Patch, Read, and Delete data in the Key-Value Store.
// path: Location of the secret.
func (s *Secrets) KVv2Read(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

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

// KVv2ReadConfig Read the backend level settings.
func (s *Secrets) KVv2ReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)

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

// KVv2ReadMetadata Configures settings for the KV store
// path: Location of the secret.
func (s *Secrets) KVv2ReadMetadata(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
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

// KVv2ReadSubkeys Read the structure of a secret entry from the Key-Value store with the values removed.
// path: Location of the secret.
func (s *Secrets) KVv2ReadSubkeys(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/subkeys/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

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

// KVv2UndeleteVersions Undeletes one or more versions from the KV store.
// path: Location of the secret.
func (s *Secrets) KVv2UndeleteVersions(ctx context.Context, path string, request schema.KVv2UndeleteVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/undelete/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
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

// KVv2Write Write, Patch, Read, and Delete data in the Key-Value Store.
// path: Location of the secret.
func (s *Secrets) KVv2Write(ctx context.Context, path string, request schema.KVv2WriteRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
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

// KVv2WriteConfig Configure backend level settings that are applied to every key in the key-value store.
func (s *Secrets) KVv2WriteConfig(ctx context.Context, request schema.KVv2WriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)

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

// KVv2WriteMetadata Configures settings for the KV store
// path: Location of the secret.
func (s *Secrets) KVv2WriteMetadata(ctx context.Context, path string, request schema.KVv2WriteMetadataRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
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

// KubernetesDeleteConfig
func (s *Secrets) KubernetesDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) KubernetesDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// KubernetesListRoles
func (s *Secrets) KubernetesListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// KubernetesReadConfig
func (s *Secrets) KubernetesReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) KubernetesReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// KubernetesWriteConfig
func (s *Secrets) KubernetesWriteConfig(ctx context.Context, request schema.KubernetesWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// KubernetesWriteCredentials
// name: Name of the Vault role
func (s *Secrets) KubernetesWriteCredentials(ctx context.Context, name string, request schema.KubernetesWriteCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPCheckInLibrary Check service accounts in to the library.
// name: Name of the set.
func (s *Secrets) LDAPCheckInLibrary(ctx context.Context, name string, request schema.LDAPCheckInLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPCheckInManageLibrary Check service accounts in to the library.
// name: Name of the set.
func (s *Secrets) LDAPCheckInManageLibrary(ctx context.Context, name string, request schema.LDAPCheckInManageLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPCheckOutLibrary Check a service account out from the library.
// name: Name of the set
func (s *Secrets) LDAPCheckOutLibrary(ctx context.Context, name string, request schema.LDAPCheckOutLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPDeleteConfig
func (s *Secrets) LDAPDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPDeleteLibrary Delete a library set.
// name: Name of the set.
func (s *Secrets) LDAPDeleteLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPDeleteRole
// name: Name of the role (lowercase)
func (s *Secrets) LDAPDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPDeleteStaticRole
// name: Name of the role
func (s *Secrets) LDAPDeleteStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPListLibraries
func (s *Secrets) LDAPListLibraries(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPListRoles
func (s *Secrets) LDAPListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPListStaticRoles
func (s *Secrets) LDAPListStaticRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPReadConfig
func (s *Secrets) LDAPReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPReadCredentials
// name: Name of the dynamic role.
func (s *Secrets) LDAPReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPReadLibrary Read a library set.
// name: Name of the set.
func (s *Secrets) LDAPReadLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPReadLibraryStatus Check the status of the service accounts in a library set.
// name: Name of the set.
func (s *Secrets) LDAPReadLibraryStatus(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPReadRole
// name: Name of the role (lowercase)
func (s *Secrets) LDAPReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPReadStaticCredentials
// name: Name of the static role.
func (s *Secrets) LDAPReadStaticCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPReadStaticRole
// name: Name of the role
func (s *Secrets) LDAPReadStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPRotateRole
// name: Name of the static role
func (s *Secrets) LDAPRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPRotateRoot
func (s *Secrets) LDAPRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPWriteConfig
func (s *Secrets) LDAPWriteConfig(ctx context.Context, request schema.LDAPWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPWriteLibrary Update a library set.
// name: Name of the set.
func (s *Secrets) LDAPWriteLibrary(ctx context.Context, name string, request schema.LDAPWriteLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPWriteRole
// name: Name of the role (lowercase)
func (s *Secrets) LDAPWriteRole(ctx context.Context, name string, request schema.LDAPWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// LDAPWriteStaticRole
// name: Name of the role
func (s *Secrets) LDAPWriteStaticRole(ctx context.Context, name string, request schema.LDAPWriteStaticRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// MongoDBAtlasDeleteRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
func (s *Secrets) MongoDBAtlasDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// MongoDBAtlasListRoles List the existing roles in this backend
func (s *Secrets) MongoDBAtlasListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// MongoDBAtlasReadConfig Configure the  credentials that are used to manage Database Users.
func (s *Secrets) MongoDBAtlasReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// MongoDBAtlasReadCredentials Generate MongoDB Atlas Programmatic API from a specific Vault role.
// name: Name of the role
func (s *Secrets) MongoDBAtlasReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// MongoDBAtlasReadRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
func (s *Secrets) MongoDBAtlasReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// MongoDBAtlasWriteConfig Configure the  credentials that are used to manage Database Users.
func (s *Secrets) MongoDBAtlasWriteConfig(ctx context.Context, request schema.MongoDBAtlasWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// MongoDBAtlasWriteCredentials Generate MongoDB Atlas Programmatic API from a specific Vault role.
// name: Name of the role
func (s *Secrets) MongoDBAtlasWriteCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// MongoDBAtlasWriteRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
func (s *Secrets) MongoDBAtlasWriteRole(ctx context.Context, name string, request schema.MongoDBAtlasWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// NomadDeleteAccessConfig
func (s *Secrets) NomadDeleteAccessConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// NomadDeleteLeaseConfig Configure the lease parameters for generated tokens
func (s *Secrets) NomadDeleteLeaseConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) NomadDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// NomadListRoles
func (s *Secrets) NomadListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// NomadReadAccessConfig
func (s *Secrets) NomadReadAccessConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// NomadReadCredentials
// name: Name of the role
func (s *Secrets) NomadReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// NomadReadLeaseConfig Configure the lease parameters for generated tokens
func (s *Secrets) NomadReadLeaseConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) NomadReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// NomadWriteAccessConfig
func (s *Secrets) NomadWriteAccessConfig(ctx context.Context, request schema.NomadWriteAccessConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// NomadWriteLeaseConfig Configure the lease parameters for generated tokens
func (s *Secrets) NomadWriteLeaseConfig(ctx context.Context, request schema.NomadWriteLeaseConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// OpenLDAPCheckInLibrary Check service accounts in to the library.
// name: Name of the set.
func (s *Secrets) OpenLDAPCheckInLibrary(ctx context.Context, name string, request schema.OpenLDAPCheckInLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPCheckInManageLibrary Check service accounts in to the library.
// name: Name of the set.
func (s *Secrets) OpenLDAPCheckInManageLibrary(ctx context.Context, name string, request schema.OpenLDAPCheckInManageLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/manage/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPCheckOutLibrary Check a service account out from the library.
// name: Name of the set
func (s *Secrets) OpenLDAPCheckOutLibrary(ctx context.Context, name string, request schema.OpenLDAPCheckOutLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}/check-out"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPDeleteConfig
func (s *Secrets) OpenLDAPDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

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

// OpenLDAPDeleteLibrary Delete a library set.
// name: Name of the set.
func (s *Secrets) OpenLDAPDeleteLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPDeleteRole
// name: Name of the role (lowercase)
func (s *Secrets) OpenLDAPDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPDeleteStaticRole
// name: Name of the role
func (s *Secrets) OpenLDAPDeleteStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPListLibraries
func (s *Secrets) OpenLDAPListLibraries(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

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

// OpenLDAPListRoles
func (s *Secrets) OpenLDAPListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

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

// OpenLDAPListStaticRoles
func (s *Secrets) OpenLDAPListStaticRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-role"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

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

// OpenLDAPReadConfig
func (s *Secrets) OpenLDAPReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

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

// OpenLDAPReadCredentials
// name: Name of the dynamic role.
func (s *Secrets) OpenLDAPReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPReadLibrary Read a library set.
// name: Name of the set.
func (s *Secrets) OpenLDAPReadLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPReadLibraryStatus Check the status of the service accounts in a library set.
// name: Name of the set.
func (s *Secrets) OpenLDAPReadLibraryStatus(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}/status"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPReadRole
// name: Name of the role (lowercase)
func (s *Secrets) OpenLDAPReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPReadStaticCredentials
// name: Name of the static role.
func (s *Secrets) OpenLDAPReadStaticCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-cred/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPReadStaticRole
// name: Name of the role
func (s *Secrets) OpenLDAPReadStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPRotateRole
// name: Name of the static role
func (s *Secrets) OpenLDAPRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPRotateRoot
func (s *Secrets) OpenLDAPRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

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

// OpenLDAPWriteConfig
func (s *Secrets) OpenLDAPWriteConfig(ctx context.Context, request schema.OpenLDAPWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

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

// OpenLDAPWriteLibrary Update a library set.
// name: Name of the set.
func (s *Secrets) OpenLDAPWriteLibrary(ctx context.Context, name string, request schema.OpenLDAPWriteLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPWriteRole
// name: Name of the role (lowercase)
func (s *Secrets) OpenLDAPWriteRole(ctx context.Context, name string, request schema.OpenLDAPWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// OpenLDAPWriteStaticRole
// name: Name of the role
func (s *Secrets) OpenLDAPWriteStaticRole(ctx context.Context, name string, request schema.OpenLDAPWriteStaticRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
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

// PKIBundleWrite
func (s *Secrets) PKIBundleWrite(ctx context.Context, request schema.PKIBundleWriteRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/bundle"
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

// PKIDeleteKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
func (s *Secrets) PKIDeleteKey(ctx context.Context, keyRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// PKIDeleteRole
// name: Name of the role
func (s *Secrets) PKIDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// PKIDeleteRoot
func (s *Secrets) PKIDeleteRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// PKIGenerateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (s *Secrets) PKIGenerateRoot(ctx context.Context, exported string, request schema.PKIGenerateRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/generate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

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

// PKIImportKeys
func (s *Secrets) PKIImportKeys(ctx context.Context, request schema.PKIImportKeysRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/import"
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

// PKIIssuerIssueRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
func (s *Secrets) PKIIssuerIssueRole(ctx context.Context, issuerRef string, role string, request schema.PKIIssuerIssueRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
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

// PKIIssuerResignCRLs
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PKIIssuerResignCRLs(ctx context.Context, issuerRef string, request schema.PKIIssuerResignCRLsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/resign-crls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PKIIssuerRevoke
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PKIIssuerRevoke(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/revoke"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PKIIssuerSignIntermediate
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PKIIssuerSignIntermediate(ctx context.Context, issuerRef string, request schema.PKIIssuerSignIntermediateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-intermediate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PKIIssuerSignRevocationList
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PKIIssuerSignRevocationList(ctx context.Context, issuerRef string, request schema.PKIIssuerSignRevocationListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-revocation-list"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PKIIssuerSignRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
func (s *Secrets) PKIIssuerSignRole(ctx context.Context, issuerRef string, role string, request schema.PKIIssuerSignRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
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

// PKIIssuerSignSelfIssued
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PKIIssuerSignSelfIssued(ctx context.Context, issuerRef string, request schema.PKIIssuerSignSelfIssuedRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-self-issued"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PKIIssuerSignVerbatim
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PKIIssuerSignVerbatim(ctx context.Context, issuerRef string, request schema.PKIIssuerSignVerbatimRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PKIIssuerSignVerbatimRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
func (s *Secrets) PKIIssuerSignVerbatimRole(ctx context.Context, issuerRef string, role string, request schema.PKIIssuerSignVerbatimRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
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

// PKIIssuersGenerateIntermediate
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (s *Secrets) PKIIssuersGenerateIntermediate(ctx context.Context, exported string, request schema.PKIIssuersGenerateIntermediateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/generate/intermediate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

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

// PKIIssuersGenerateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (s *Secrets) PKIIssuersGenerateRoot(ctx context.Context, exported string, request schema.PKIIssuersGenerateRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/generate/root/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

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

// PKIIssuersList
func (s *Secrets) PKIIssuersList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIListCerts
func (s *Secrets) PKIListCerts(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/certs"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIListCertsRevoked
func (s *Secrets) PKIListCertsRevoked(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/certs/revoked"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIListKeys
func (s *Secrets) PKIListKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIListRoles
func (s *Secrets) PKIListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadAutoTidyConfig
func (s *Secrets) PKIReadAutoTidyConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/auto-tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCA
func (s *Secrets) PKIReadCA(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCAChain
func (s *Secrets) PKIReadCAChain(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca_chain"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCAPem
func (s *Secrets) PKIReadCAPem(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCRL
func (s *Secrets) PKIReadCRL(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCRLConfig
func (s *Secrets) PKIReadCRLConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCRLRotate
func (s *Secrets) PKIReadCRLRotate(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/rotate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCRLRotateDelta
func (s *Secrets) PKIReadCRLRotateDelta(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/rotate-delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCert
// serial: Certificate serial number, in colon- or hyphen-separated octal
func (s *Secrets) PKIReadCert(ctx context.Context, serial string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

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

// PKIReadCertCAChain
func (s *Secrets) PKIReadCertCAChain(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/ca_chain"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadCertRaw
// serial: Certificate serial number, in colon- or hyphen-separated octal
func (s *Secrets) PKIReadCertRaw(ctx context.Context, serial string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}/raw"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

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

// PKIReadCertRawPem
// serial: Certificate serial number, in colon- or hyphen-separated octal
func (s *Secrets) PKIReadCertRawPem(ctx context.Context, serial string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}/raw/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

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

// PKIReadClusterConfig
func (s *Secrets) PKIReadClusterConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/cluster"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadDeltaCRL
func (s *Secrets) PKIReadDeltaCRL(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/delta-crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadIssuersConfig
func (s *Secrets) PKIReadIssuersConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
func (s *Secrets) PKIReadKey(ctx context.Context, keyRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadKeysConfig
func (s *Secrets) PKIReadKeysConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReadOCSPReq
// req: base-64 encoded ocsp request
func (s *Secrets) PKIReadOCSPReq(ctx context.Context, req string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// PKIReadRole
// name: Name of the role
func (s *Secrets) PKIReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadURLConfig
func (s *Secrets) PKIReadURLConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/urls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIReplaceRoot
func (s *Secrets) PKIReplaceRoot(ctx context.Context, request schema.PKIReplaceRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/replace"
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

// PKIRevoke
func (s *Secrets) PKIRevoke(ctx context.Context, request schema.PKIRevokeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/revoke"
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

// PKIRevokeWithKey
func (s *Secrets) PKIRevokeWithKey(ctx context.Context, request schema.PKIRevokeWithKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/revoke-with-key"
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

// PKIRootSignIntermediate
func (s *Secrets) PKIRootSignIntermediate(ctx context.Context, request schema.PKIRootSignIntermediateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/sign-intermediate"
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

// PKIRootSignSelfIssued
func (s *Secrets) PKIRootSignSelfIssued(ctx context.Context, request schema.PKIRootSignSelfIssuedRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/sign-self-issued"
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

// PKIRotateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (s *Secrets) PKIRotateRoot(ctx context.Context, exported string, request schema.PKIRotateRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/rotate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

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

// PKISignRole
// role: The desired role with configuration for this request
func (s *Secrets) PKISignRole(ctx context.Context, role string, request schema.PKISignRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
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

// PKISignVerbatim
func (s *Secrets) PKISignVerbatim(ctx context.Context, request schema.PKISignVerbatimRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign-verbatim"
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

// PKISignVerbatimRole
// role: The desired role with configuration for this request
func (s *Secrets) PKISignVerbatimRole(ctx context.Context, role string, request schema.PKISignVerbatimRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign-verbatim/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
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

// PKITidy
func (s *Secrets) PKITidy(ctx context.Context, request schema.PKITidyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// PKITidyCancel
func (s *Secrets) PKITidyCancel(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy-cancel"
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

// PKITidyStatus
func (s *Secrets) PKITidyStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy-status"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PKIWriteAutoTidyConfig
func (s *Secrets) PKIWriteAutoTidyConfig(ctx context.Context, request schema.PKIWriteAutoTidyConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/auto-tidy"
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

// PKIWriteCAConfig
func (s *Secrets) PKIWriteCAConfig(ctx context.Context, request schema.PKIWriteCAConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/ca"
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

// PKIWriteCRLConfig
func (s *Secrets) PKIWriteCRLConfig(ctx context.Context, request schema.PKIWriteCRLConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/crl"
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

// PKIWriteCerts
func (s *Secrets) PKIWriteCerts(ctx context.Context, request schema.PKIWriteCertsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert"
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

// PKIWriteClusterConfig
func (s *Secrets) PKIWriteClusterConfig(ctx context.Context, request schema.PKIWriteClusterConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/cluster"
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

// PKIWriteIntermediateCrossSign
func (s *Secrets) PKIWriteIntermediateCrossSign(ctx context.Context, request schema.PKIWriteIntermediateCrossSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/cross-sign"
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

// PKIWriteIntermediateGenerate
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (s *Secrets) PKIWriteIntermediateGenerate(ctx context.Context, exported string, request schema.PKIWriteIntermediateGenerateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/generate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

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

// PKIWriteIntermediateSetSigned
func (s *Secrets) PKIWriteIntermediateSetSigned(ctx context.Context, request schema.PKIWriteIntermediateSetSignedRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/set-signed"
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

// PKIWriteInternalExported
func (s *Secrets) PKIWriteInternalExported(ctx context.Context, request schema.PKIWriteInternalExportedRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/internal|exported"
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

// PKIWriteIssueRole
// role: The desired role with configuration for this request
func (s *Secrets) PKIWriteIssueRole(ctx context.Context, role string, request schema.PKIWriteIssueRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
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

// PKIWriteIssuersConfig
func (s *Secrets) PKIWriteIssuersConfig(ctx context.Context, request schema.PKIWriteIssuersConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/issuers"
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

// PKIWriteKMS
func (s *Secrets) PKIWriteKMS(ctx context.Context, request schema.PKIWriteKMSRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/kms"
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

// PKIWriteKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
func (s *Secrets) PKIWriteKey(ctx context.Context, keyRef string, request schema.PKIWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

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

// PKIWriteKeysConfig
func (s *Secrets) PKIWriteKeysConfig(ctx context.Context, request schema.PKIWriteKeysConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/keys"
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

// PKIWriteOCSP
func (s *Secrets) PKIWriteOCSP(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// PKIWriteRole
// name: Name of the role
func (s *Secrets) PKIWriteRole(ctx context.Context, name string, request schema.PKIWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
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

// PKIWriteURLConfig
func (s *Secrets) PKIWriteURLConfig(ctx context.Context, request schema.PKIWriteURLConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/urls"
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

// PkiDeleteIssuerRefDerPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PkiDeleteIssuerRefDerPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/{issuer_ref}/der|/pem"
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

// PkiDeleteJson
func (s *Secrets) PkiDeleteJson(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//json"
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

// PkiReadDelta
func (s *Secrets) PkiReadDelta(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PkiReadDeltaPem
func (s *Secrets) PkiReadDeltaPem(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//delta/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PkiReadDer
func (s *Secrets) PkiReadDer(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//der"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PkiReadIssuerRefCrlPemDerDeltaPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PkiReadIssuerRefCrlPemDerDeltaPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/{issuer_ref}/crl/pem|/der|/delta/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PkiReadIssuerRefDerPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PkiReadIssuerRefDerPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/{issuer_ref}/der|/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PkiReadJson
func (s *Secrets) PkiReadJson(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//json"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PkiReadPem
func (s *Secrets) PkiReadPem(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

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

// PkiWriteIssuerRefDerPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (s *Secrets) PkiWriteIssuerRefDerPem(ctx context.Context, issuerRef string, request schema.PkiWriteIssuerRefDerPemRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/{issuer_ref}/der|/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

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

// PkiWriteJson
func (s *Secrets) PkiWriteJson(ctx context.Context, request schema.PkiWriteJsonRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//json"
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

// RabbitMQDeleteRole Manage the roles that can be created with this backend.
// name: Name of the role.
func (s *Secrets) RabbitMQDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// RabbitMQListRoles Manage the roles that can be created with this backend.
func (s *Secrets) RabbitMQListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// RabbitMQReadCredentials Request RabbitMQ credentials for a certain role.
// name: Name of the role.
func (s *Secrets) RabbitMQReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// RabbitMQReadLeaseConfig Configure the lease parameters for generated credentials
func (s *Secrets) RabbitMQReadLeaseConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// RabbitMQReadRole Manage the roles that can be created with this backend.
// name: Name of the role.
func (s *Secrets) RabbitMQReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// RabbitMQWriteConnectionConfig Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
func (s *Secrets) RabbitMQWriteConnectionConfig(ctx context.Context, request schema.RabbitMQWriteConnectionConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// RabbitMQWriteLeaseConfig Configure the lease parameters for generated credentials
func (s *Secrets) RabbitMQWriteLeaseConfig(ctx context.Context, request schema.RabbitMQWriteLeaseConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// RabbitMQWriteRole Manage the roles that can be created with this backend.
// name: Name of the role.
func (s *Secrets) RabbitMQWriteRole(ctx context.Context, name string, request schema.RabbitMQWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHDeleteCAConfig Set the SSH private key used for signing certificates.
func (s *Secrets) SSHDeleteCAConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHDeleteKeys Register a shared private key with Vault.
// keyName: [Required] Name of the key
func (s *Secrets) SSHDeleteKeys(ctx context.Context, keyName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/keys/{key_name}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_name"+"}", url.PathEscape(keyName), -1)

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

// SSHDeleteRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
func (s *Secrets) SSHDeleteRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHDeleteZeroAddressConfig Assign zero address as default CIDR block for select roles.
func (s *Secrets) SSHDeleteZeroAddressConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHListRoles Manage the 'roles' that can be created with this backend.
func (s *Secrets) SSHListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHLookup List all the roles associated with the given IP address.
func (s *Secrets) SSHLookup(ctx context.Context, request schema.SSHLookupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHReadCAConfig Set the SSH private key used for signing certificates.
func (s *Secrets) SSHReadCAConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHReadPublicKey Retrieve the public key.
func (s *Secrets) SSHReadPublicKey(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHReadRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
func (s *Secrets) SSHReadRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHReadZeroAddressConfig Assign zero address as default CIDR block for select roles.
func (s *Secrets) SSHReadZeroAddressConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHSign Request signing an SSH key using a certain role with the provided details.
// role: The desired role with configuration for this request.
func (s *Secrets) SSHSign(ctx context.Context, role string, request schema.SSHSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHVerify Validate the OTP provided by Vault SSH Agent.
func (s *Secrets) SSHVerify(ctx context.Context, request schema.SSHVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHWriteCAConfig Set the SSH private key used for signing certificates.
func (s *Secrets) SSHWriteCAConfig(ctx context.Context, request schema.SSHWriteCAConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHWriteCredentials Creates a credential for establishing SSH connection with the remote host.
// role: [Required] Name of the role
func (s *Secrets) SSHWriteCredentials(ctx context.Context, role string, request schema.SSHWriteCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHWriteIssue
// role: The desired role with configuration for this request.
func (s *Secrets) SSHWriteIssue(ctx context.Context, role string, request schema.SSHWriteIssueRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHWriteKeys Register a shared private key with Vault.
// keyName: [Required] Name of the key
func (s *Secrets) SSHWriteKeys(ctx context.Context, keyName string, request schema.SSHWriteKeysRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/keys/{key_name}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_name"+"}", url.PathEscape(keyName), -1)

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

// SSHWriteRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
func (s *Secrets) SSHWriteRole(ctx context.Context, role string, request schema.SSHWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// SSHWriteZeroAddressConfig Assign zero address as default CIDR block for select roles.
func (s *Secrets) SSHWriteZeroAddressConfig(ctx context.Context, request schema.SSHWriteZeroAddressConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TOTPDeleteKey Manage the keys that can be created with this backend.
// name: Name of the key.
func (s *Secrets) TOTPDeleteKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TOTPListKeys Manage the keys that can be created with this backend.
func (s *Secrets) TOTPListKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TOTPReadCode Request time-based one-time use password or validate a password for a certain key .
// name: Name of the key.
func (s *Secrets) TOTPReadCode(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TOTPReadKey Manage the keys that can be created with this backend.
// name: Name of the key.
func (s *Secrets) TOTPReadKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TOTPWriteCode Request time-based one-time use password or validate a password for a certain key .
// name: Name of the key.
func (s *Secrets) TOTPWriteCode(ctx context.Context, name string, request schema.TOTPWriteCodeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TOTPWriteKey Manage the keys that can be created with this backend.
// name: Name of the key.
func (s *Secrets) TOTPWriteKey(ctx context.Context, name string, request schema.TOTPWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformDeleteConfig
func (s *Secrets) TerraformDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformDeleteRole
// name: Name of the role
func (s *Secrets) TerraformDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformListRoles
func (s *Secrets) TerraformListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformReadConfig
func (s *Secrets) TerraformReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformReadCredentials Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
// name: Name of the role
func (s *Secrets) TerraformReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformReadRole
// name: Name of the role
func (s *Secrets) TerraformReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformRotateRole
// name: Name of the team or organization role
func (s *Secrets) TerraformRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformWriteConfig
func (s *Secrets) TerraformWriteConfig(ctx context.Context, request schema.TerraformWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformWriteCredentials Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
// name: Name of the role
func (s *Secrets) TerraformWriteCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TerraformWriteRole
// name: Name of the role
func (s *Secrets) TerraformWriteRole(ctx context.Context, name string, request schema.TerraformWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitBackup Backup the named key
// name: Name of the key
func (s *Secrets) TransitBackup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitDeleteKey Managed named encryption keys
// name: Name of the key
func (s *Secrets) TransitDeleteKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitEncrypt(ctx context.Context, name string, request schema.TransitEncryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitExport Export named encryption or signing key
// name: Name of the key
// type_: Type of key to export (encryption-key, signing-key, hmac-key)
func (s *Secrets) TransitExport(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitExportVersion Export named encryption or signing key
// name: Name of the key
// type_: Type of key to export (encryption-key, signing-key, hmac-key)
// version: Version of the key
func (s *Secrets) TransitExportVersion(ctx context.Context, name string, type_ string, version string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitGenerateDataKey(ctx context.Context, name string, plaintext string, request schema.TransitGenerateDataKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitGenerateHMAC Generate an HMAC for input data using the named key
// name: The key to use for the HMAC function
func (s *Secrets) TransitGenerateHMAC(ctx context.Context, name string, request schema.TransitGenerateHMACRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitGenerateHMACWithAlgorithm Generate an HMAC for input data using the named key
// name: The key to use for the HMAC function
// urlalgorithm: Algorithm to use (POST URL parameter)
func (s *Secrets) TransitGenerateHMACWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitGenerateHMACWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitGenerateRandom(ctx context.Context, request schema.TransitGenerateRandomRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitGenerateRandomSource Generate random bytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
func (s *Secrets) TransitGenerateRandomSource(ctx context.Context, source string, request schema.TransitGenerateRandomSourceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitGenerateRandomSourceBytes Generate random bytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
// urlbytes: The number of bytes to generate (POST URL parameter)
func (s *Secrets) TransitGenerateRandomSourceBytes(ctx context.Context, source string, urlbytes string, request schema.TransitGenerateRandomSourceBytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitHash(ctx context.Context, request schema.TransitHashRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitHashWithAlgorithm(ctx context.Context, urlalgorithm string, request schema.TransitHashWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitImportKey(ctx context.Context, name string, request schema.TransitImportKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitImportKeyVersion(ctx context.Context, name string, request schema.TransitImportKeyVersionRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitListKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitReadCacheConfig Returns the size of the active cache
func (s *Secrets) TransitReadCacheConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitReadConfigKeys Configuration common across all keys
func (s *Secrets) TransitReadConfigKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitReadKey Managed named encryption keys
// name: Name of the key
func (s *Secrets) TransitReadKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitReadWrappingKey Returns the public key to use for wrapping imported keys
func (s *Secrets) TransitReadWrappingKey(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitRestore Restore the named key
func (s *Secrets) TransitRestore(ctx context.Context, request schema.TransitRestoreRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitRestoreKey Restore the named key
// name: If set, this will be the name of the restored key.
func (s *Secrets) TransitRestoreKey(ctx context.Context, name string, request schema.TransitRestoreKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitRotateKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/rotate"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
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
func (s *Secrets) TransitSignWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitSignWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitTrimKey(ctx context.Context, name string, request schema.TransitTrimKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitVerify(ctx context.Context, name string, request schema.TransitVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
func (s *Secrets) TransitVerifyWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitVerifyWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitWriteCacheConfig Configures a new cache of the specified size
func (s *Secrets) TransitWriteCacheConfig(ctx context.Context, request schema.TransitWriteCacheConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitWriteConfigKeys Configuration common across all keys
func (s *Secrets) TransitWriteConfigKeys(ctx context.Context, request schema.TransitWriteConfigKeysRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitWriteKey Managed named encryption keys
// name: Name of the key
func (s *Secrets) TransitWriteKey(ctx context.Context, name string, request schema.TransitWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitWriteKeyConfig Configure a named encryption key
// name: Name of the key
func (s *Secrets) TransitWriteKeyConfig(ctx context.Context, name string, request schema.TransitWriteKeyConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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

// TransitWriteRandomUrlbytes Generate random bytes
// urlbytes: The number of bytes to generate (POST URL parameter)
func (s *Secrets) TransitWriteRandomUrlbytes(ctx context.Context, urlbytes string, request schema.TransitWriteRandomUrlbytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
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
