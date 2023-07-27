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

// Auth is a simple wrapper around the client for Auth requests
type Auth struct {
	client *Client
}

// AliCloudDeleteAuthRole Create a role and associate policies to it.
// role: The name of the role as it should appear in Vault.
func (a *Auth) AliCloudDeleteAuthRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudListAuthRoles Lists all the roles that are registered with Vault.
func (a *Auth) AliCloudListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudLogin Authenticates an RAM entity with Vault.
func (a *Auth) AliCloudLogin(ctx context.Context, request schema.AliCloudLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudReadAuthRole Create a role and associate policies to it.
// role: The name of the role as it should appear in Vault.
func (a *Auth) AliCloudReadAuthRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudWriteAuthRole Create a role and associate policies to it.
// role: The name of the role as it should appear in Vault.
func (a *Auth) AliCloudWriteAuthRole(ctx context.Context, role string, request schema.AliCloudWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteBindSecretId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteBindSecretId(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteBoundCidrList
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteBoundCidrList(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeletePeriod
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeletePeriod(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeletePolicies
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeletePolicies(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteRole
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteRole(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteSecretIdBoundCidrs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteSecretIdBoundCidrs(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteSecretIdNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteSecretIdNumUses(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteSecretIdTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteSecretIdTtl(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteTokenBoundCidrs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteTokenBoundCidrs(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteTokenMaxTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteTokenMaxTtl(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteTokenNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteTokenNumUses(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDeleteTokenTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDeleteTokenTtl(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDestroySecretId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDestroySecretId(ctx context.Context, roleName string, request schema.AppRoleDestroySecretIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleDestroySecretIdByAccessor
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleDestroySecretIdByAccessor(ctx context.Context, roleName string, request schema.AppRoleDestroySecretIdByAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleListRoles
func (a *Auth) AppRoleListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleListSecretIds
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleListSecretIds(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleLogin
func (a *Auth) AppRoleLogin(ctx context.Context, request schema.AppRoleLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleLookUpSecretId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleLookUpSecretId(ctx context.Context, roleName string, request schema.AppRoleLookUpSecretIdRequest, options ...RequestOption) (*Response[schema.AppRoleLookUpSecretIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/lookup"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.AppRoleLookUpSecretIdResponse](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleLookUpSecretIdByAccessor
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleLookUpSecretIdByAccessor(ctx context.Context, roleName string, request schema.AppRoleLookUpSecretIdByAccessorRequest, options ...RequestOption) (*Response[schema.AppRoleLookUpSecretIdByAccessorResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/lookup"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.AppRoleLookUpSecretIdByAccessorResponse](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadBindSecretId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadBindSecretId(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadBindSecretIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadBindSecretIdResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadBoundCidrList
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadBoundCidrList(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadBoundCidrListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadBoundCidrListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadLocalSecretIds
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadLocalSecretIds(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadLocalSecretIdsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/local-secret-ids"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadLocalSecretIdsResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadPeriod
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadPeriod(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadPeriodResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadPeriodResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadPolicies
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadPolicies(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadPoliciesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadPoliciesResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadRole
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadRole(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadRoleResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadRoleId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadRoleId(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadRoleIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/role-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadRoleIdResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadSecretIdBoundCidrs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadSecretIdBoundCidrs(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIdBoundCidrsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadSecretIdBoundCidrsResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadSecretIdNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadSecretIdNumUses(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIdNumUsesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadSecretIdNumUsesResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadSecretIdTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadSecretIdTtl(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIdTtlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadSecretIdTtlResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadTokenBoundCidrs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadTokenBoundCidrs(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadTokenBoundCidrsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadTokenBoundCidrsResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadTokenMaxTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadTokenMaxTtl(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadTokenMaxTtlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadTokenMaxTtlResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadTokenNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadTokenNumUses(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadTokenNumUsesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadTokenNumUsesResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleReadTokenTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleReadTokenTtl(ctx context.Context, roleName string, options ...RequestOption) (*Response[schema.AppRoleReadTokenTtlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[schema.AppRoleReadTokenTtlResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleTidySecretId
func (a *Auth) AppRoleTidySecretId(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/tidy/secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteBindSecretId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteBindSecretId(ctx context.Context, roleName string, request schema.AppRoleWriteBindSecretIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteBoundCidrList
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteBoundCidrList(ctx context.Context, roleName string, request schema.AppRoleWriteBoundCidrListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteCustomSecretId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteCustomSecretId(ctx context.Context, roleName string, request schema.AppRoleWriteCustomSecretIdRequest, options ...RequestOption) (*Response[schema.AppRoleWriteCustomSecretIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/custom-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.AppRoleWriteCustomSecretIdResponse](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWritePeriod
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWritePeriod(ctx context.Context, roleName string, request schema.AppRoleWritePeriodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWritePolicies
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWritePolicies(ctx context.Context, roleName string, request schema.AppRoleWritePoliciesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteRole
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteRole(ctx context.Context, roleName string, request schema.AppRoleWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteRoleId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteRoleId(ctx context.Context, roleName string, request schema.AppRoleWriteRoleIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/role-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretId
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretId(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIdRequest, options ...RequestOption) (*Response[schema.AppRoleWriteSecretIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[schema.AppRoleWriteSecretIdResponse](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIdBoundCidrs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIdBoundCidrs(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIdBoundCidrsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIdNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIdNumUses(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIdNumUsesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteSecretIdTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteSecretIdTtl(ctx context.Context, roleName string, request schema.AppRoleWriteSecretIdTtlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteTokenBoundCidrs
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteTokenBoundCidrs(ctx context.Context, roleName string, request schema.AppRoleWriteTokenBoundCidrsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteTokenMaxTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteTokenMaxTtl(ctx context.Context, roleName string, request schema.AppRoleWriteTokenMaxTtlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteTokenNumUses
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteTokenNumUses(ctx context.Context, roleName string, request schema.AppRoleWriteTokenNumUsesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AppRoleWriteTokenTtl
// roleName: Name of the role. Must be less than 4096 bytes.
func (a *Auth) AppRoleWriteTokenTtl(ctx context.Context, roleName string, request schema.AppRoleWriteTokenTtlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureCertificate
// certName: Name of the certificate.
func (a *Auth) AwsConfigureCertificate(ctx context.Context, certName string, request schema.AwsConfigureCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureClient
func (a *Auth) AwsConfigureClient(ctx context.Context, request schema.AwsConfigureClientRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureIdentityAccessListTidyOperation
func (a *Auth) AwsConfigureIdentityAccessListTidyOperation(ctx context.Context, request schema.AwsConfigureIdentityAccessListTidyOperationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureIdentityIntegration
func (a *Auth) AwsConfigureIdentityIntegration(ctx context.Context, request schema.AwsConfigureIdentityIntegrationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/identity"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureIdentityWhitelistTidyOperation
func (a *Auth) AwsConfigureIdentityWhitelistTidyOperation(ctx context.Context, request schema.AwsConfigureIdentityWhitelistTidyOperationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureRoleTagBlacklistTidyOperation
func (a *Auth) AwsConfigureRoleTagBlacklistTidyOperation(ctx context.Context, request schema.AwsConfigureRoleTagBlacklistTidyOperationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsConfigureRoleTagDenyListTidyOperation
func (a *Auth) AwsConfigureRoleTagDenyListTidyOperation(ctx context.Context, request schema.AwsConfigureRoleTagDenyListTidyOperationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteAuthRole
// role: Name of the role.
func (a *Auth) AwsDeleteAuthRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteCertificateConfiguration
// certName: Name of the certificate.
func (a *Auth) AwsDeleteCertificateConfiguration(ctx context.Context, certName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteClientConfiguration
func (a *Auth) AwsDeleteClientConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteIdentityAccessList
// instanceId: EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
func (a *Auth) AwsDeleteIdentityAccessList(ctx context.Context, instanceId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteIdentityAccessListTidySettings
func (a *Auth) AwsDeleteIdentityAccessListTidySettings(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteIdentityWhitelist
// instanceId: EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
func (a *Auth) AwsDeleteIdentityWhitelist(ctx context.Context, instanceId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteIdentityWhitelistTidySettings
func (a *Auth) AwsDeleteIdentityWhitelistTidySettings(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteRoleTagBlacklist
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AwsDeleteRoleTagBlacklist(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteRoleTagBlacklistTidySettings
func (a *Auth) AwsDeleteRoleTagBlacklistTidySettings(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteRoleTagDenyList
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AwsDeleteRoleTagDenyList(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteRoleTagDenyListTidySettings
func (a *Auth) AwsDeleteRoleTagDenyListTidySettings(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsDeleteStsRole
// accountId: AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
func (a *Auth) AwsDeleteStsRole(ctx context.Context, accountId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsListAuthRoles
func (a *Auth) AwsListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsListCertificateConfigurations
func (a *Auth) AwsListCertificateConfigurations(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificates/"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsListIdentityAccessList
func (a *Auth) AwsListIdentityAccessList(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist/"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsListIdentityWhitelist
func (a *Auth) AwsListIdentityWhitelist(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist/"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsListRoleTagBlacklists
func (a *Auth) AwsListRoleTagBlacklists(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsListRoleTagDenyLists
func (a *Auth) AwsListRoleTagDenyLists(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsListStsRoleRelationships
func (a *Auth) AwsListStsRoleRelationships(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsLogin
func (a *Auth) AwsLogin(ctx context.Context, request schema.AwsLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadAuthRole
// role: Name of the role.
func (a *Auth) AwsReadAuthRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadCertificateConfiguration
// certName: Name of the certificate.
func (a *Auth) AwsReadCertificateConfiguration(ctx context.Context, certName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadClientConfiguration
func (a *Auth) AwsReadClientConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadIdentityAccessList
// instanceId: EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
func (a *Auth) AwsReadIdentityAccessList(ctx context.Context, instanceId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadIdentityAccessListTidySettings
func (a *Auth) AwsReadIdentityAccessListTidySettings(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadIdentityIntegrationConfiguration
func (a *Auth) AwsReadIdentityIntegrationConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/identity"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadIdentityWhitelist
// instanceId: EC2 instance ID. A successful login operation from an EC2 instance gets cached in this accesslist, keyed off of instance ID.
func (a *Auth) AwsReadIdentityWhitelist(ctx context.Context, instanceId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadIdentityWhitelistTidySettings
func (a *Auth) AwsReadIdentityWhitelistTidySettings(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadRoleTagBlacklist
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AwsReadRoleTagBlacklist(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadRoleTagBlacklistTidySettings
func (a *Auth) AwsReadRoleTagBlacklistTidySettings(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadRoleTagDenyList
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AwsReadRoleTagDenyList(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadRoleTagDenyListTidySettings
func (a *Auth) AwsReadRoleTagDenyListTidySettings(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsReadStsRole
// accountId: AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
func (a *Auth) AwsReadStsRole(ctx context.Context, accountId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsRotateRootCredentials
func (a *Auth) AwsRotateRootCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsTidyIdentityAccessList
func (a *Auth) AwsTidyIdentityAccessList(ctx context.Context, request schema.AwsTidyIdentityAccessListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsTidyIdentityWhitelist
func (a *Auth) AwsTidyIdentityWhitelist(ctx context.Context, request schema.AwsTidyIdentityWhitelistRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsTidyRoleTagBlacklist
func (a *Auth) AwsTidyRoleTagBlacklist(ctx context.Context, request schema.AwsTidyRoleTagBlacklistRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsTidyRoleTagDenyList
func (a *Auth) AwsTidyRoleTagDenyList(ctx context.Context, request schema.AwsTidyRoleTagDenyListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsWriteAuthRole
// role: Name of the role.
func (a *Auth) AwsWriteAuthRole(ctx context.Context, role string, request schema.AwsWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsWriteRoleTag
// role: Name of the role.
func (a *Auth) AwsWriteRoleTag(ctx context.Context, role string, request schema.AwsWriteRoleTagRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}/tag"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsWriteRoleTagBlacklist
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AwsWriteRoleTagBlacklist(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsWriteRoleTagDenyList
// roleTag: Role tag to be deny listed. The tag can be supplied as-is. In order to avoid any encoding problems, it can be base64 encoded.
func (a *Auth) AwsWriteRoleTagDenyList(ctx context.Context, roleTag string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AwsWriteStsRole
// accountId: AWS account ID to be associated with STS role. If set, Vault will use assumed credentials to verify any login attempts from EC2 instances in this account.
func (a *Auth) AwsWriteStsRole(ctx context.Context, accountId string, request schema.AwsWriteStsRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureConfigureAuth
func (a *Auth) AzureConfigureAuth(ctx context.Context, request schema.AzureConfigureAuthRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureDeleteAuthConfiguration
func (a *Auth) AzureDeleteAuthConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureDeleteAuthRole
// name: Name of the role.
func (a *Auth) AzureDeleteAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureListAuthRoles
func (a *Auth) AzureListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureLogin
func (a *Auth) AzureLogin(ctx context.Context, request schema.AzureLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadAuthConfiguration
func (a *Auth) AzureReadAuthConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadAuthRole
// name: Name of the role.
func (a *Auth) AzureReadAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureRotateRootCredentials
func (a *Auth) AzureRotateRootCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureWriteAuthRole
// name: Name of the role.
func (a *Auth) AzureWriteAuthRole(ctx context.Context, name string, request schema.AzureWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CentrifyConfigure
func (a *Auth) CentrifyConfigure(ctx context.Context, request schema.CentrifyConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CentrifyLogin Log in with a username and password.
func (a *Auth) CentrifyLogin(ctx context.Context, request schema.CentrifyLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CentrifyReadConfiguration
func (a *Auth) CentrifyReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertConfigure
func (a *Auth) CertConfigure(ctx context.Context, request schema.CertConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CertDeleteCertificate Manage trusted certificates used for authentication.
// name: The name of the certificate
func (a *Auth) CertDeleteCertificate(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertDeleteCrl Manage Certificate Revocation Lists checked during authentication.
// name: The name of the certificate
func (a *Auth) CertDeleteCrl(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertListCertificates Manage trusted certificates used for authentication.
func (a *Auth) CertListCertificates(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertListCrls
func (a *Auth) CertListCrls(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertLogin
func (a *Auth) CertLogin(ctx context.Context, request schema.CertLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CertReadCertificate Manage trusted certificates used for authentication.
// name: The name of the certificate
func (a *Auth) CertReadCertificate(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertReadConfiguration
func (a *Auth) CertReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertReadCrl Manage Certificate Revocation Lists checked during authentication.
// name: The name of the certificate
func (a *Auth) CertReadCrl(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CertWriteCertificate Manage trusted certificates used for authentication.
// name: The name of the certificate
func (a *Auth) CertWriteCertificate(ctx context.Context, name string, request schema.CertWriteCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CertWriteCrl Manage Certificate Revocation Lists checked during authentication.
// name: The name of the certificate
func (a *Auth) CertWriteCrl(ctx context.Context, name string, request schema.CertWriteCrlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryConfigure
func (a *Auth) CloudFoundryConfigure(ctx context.Context, request schema.CloudFoundryConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryDeleteConfiguration
func (a *Auth) CloudFoundryDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryDeleteRole
// role: The name of the role.
func (a *Auth) CloudFoundryDeleteRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryListRoles
func (a *Auth) CloudFoundryListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryLogin
func (a *Auth) CloudFoundryLogin(ctx context.Context, request schema.CloudFoundryLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryReadConfiguration
func (a *Auth) CloudFoundryReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryReadRole
// role: The name of the role.
func (a *Auth) CloudFoundryReadRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CloudFoundryWriteRole
// role: The name of the role.
func (a *Auth) CloudFoundryWriteRole(ctx context.Context, role string, request schema.CloudFoundryWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubConfigure
func (a *Auth) GithubConfigure(ctx context.Context, request schema.GithubConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubDeleteTeamMapping Read/write/delete a single teams mapping
// key: Key for the teams mapping
func (a *Auth) GithubDeleteTeamMapping(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubDeleteUserMapping Read/write/delete a single users mapping
// key: Key for the users mapping
func (a *Auth) GithubDeleteUserMapping(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubListTeams Read mappings for teams
func (a *Auth) GithubListTeams(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubListUsers Read mappings for users
func (a *Auth) GithubListUsers(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubLogin
func (a *Auth) GithubLogin(ctx context.Context, request schema.GithubLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubReadConfiguration
func (a *Auth) GithubReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubReadTeamMapping Read/write/delete a single teams mapping
// key: Key for the teams mapping
func (a *Auth) GithubReadTeamMapping(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubReadUserMapping Read/write/delete a single users mapping
// key: Key for the users mapping
func (a *Auth) GithubReadUserMapping(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubWriteTeamMapping Read/write/delete a single teams mapping
// key: Key for the teams mapping
func (a *Auth) GithubWriteTeamMapping(ctx context.Context, key string, request schema.GithubWriteTeamMappingRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GithubWriteUserMapping Read/write/delete a single users mapping
// key: Key for the users mapping
func (a *Auth) GithubWriteUserMapping(ctx context.Context, key string, request schema.GithubWriteUserMappingRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudConfigureAuth
func (a *Auth) GoogleCloudConfigureAuth(ctx context.Context, request schema.GoogleCloudConfigureAuthRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudDeleteRole Create a GCP role with associated policies and required attributes.
// name: Name of the role.
func (a *Auth) GoogleCloudDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudEditLabelsForRole Add or remove labels for an existing 'gce' role
// name: Name of the role.
func (a *Auth) GoogleCloudEditLabelsForRole(ctx context.Context, name string, request schema.GoogleCloudEditLabelsForRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}/labels"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudEditServiceAccountsForRole Add or remove service accounts for an existing `iam` role
// name: Name of the role.
func (a *Auth) GoogleCloudEditServiceAccountsForRole(ctx context.Context, name string, request schema.GoogleCloudEditServiceAccountsForRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}/service-accounts"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListRoles Lists all the roles that are registered with Vault.
func (a *Auth) GoogleCloudListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudLogin
func (a *Auth) GoogleCloudLogin(ctx context.Context, request schema.GoogleCloudLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadAuthConfiguration
func (a *Auth) GoogleCloudReadAuthConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadRole Create a GCP role with associated policies and required attributes.
// name: Name of the role.
func (a *Auth) GoogleCloudReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteRole Create a GCP role with associated policies and required attributes.
// name: Name of the role.
func (a *Auth) GoogleCloudWriteRole(ctx context.Context, name string, request schema.GoogleCloudWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtConfigure Configure the JWT authentication backend.
// The JWT authentication backend validates JWTs (or OIDC) using the configured credentials. If using OIDC Discovery, the URL must be provided, along with (optionally) the CA cert to use for the connection. If performing JWT validation locally, a set of public keys must be provided.
func (a *Auth) JwtConfigure(ctx context.Context, request schema.JwtConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtDeleteRole Delete an existing role.
// name: Name of the role.
func (a *Auth) JwtDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtListRoles Lists all the roles registered with the backend.
// The list will contain the names of the roles.
func (a *Auth) JwtListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtLogin Authenticates to Vault using a JWT (or OIDC) token.
func (a *Auth) JwtLogin(ctx context.Context, request schema.JwtLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtOidcCallback Callback endpoint to complete an OIDC login.
func (a *Auth) JwtOidcCallback(ctx context.Context, clientNonce string, code string, state string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/callback"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("client_nonce", parameterToString(clientNonce))
	requestQueryParameters.Add("code", parameterToString(code))
	requestQueryParameters.Add("state", parameterToString(state))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtOidcCallbackFormPost Callback endpoint to handle form_posts.
func (a *Auth) JwtOidcCallbackFormPost(ctx context.Context, request schema.JwtOidcCallbackFormPostRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/callback"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtOidcRequestAuthorizationUrl Request an authorization URL to start an OIDC login flow.
func (a *Auth) JwtOidcRequestAuthorizationUrl(ctx context.Context, request schema.JwtOidcRequestAuthorizationUrlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/auth_url"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtReadConfiguration Read the current JWT authentication backend configuration.
func (a *Auth) JwtReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtReadRole Read an existing role.
// name: Name of the role.
func (a *Auth) JwtReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// JwtWriteRole Register an role with the backend.
// A role is required to authenticate with this backend. The role binds   JWT token information with token policies and settings.   The bindings, token polices and token settings can all be configured   using this endpoint
// name: Name of the role.
func (a *Auth) JwtWriteRole(ctx context.Context, name string, request schema.JwtWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosConfigure
func (a *Auth) KerberosConfigure(ctx context.Context, request schema.KerberosConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosConfigureLdap
func (a *Auth) KerberosConfigureLdap(ctx context.Context, request schema.KerberosConfigureLdapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config/ldap"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosDeleteGroup
// name: Name of the LDAP group.
func (a *Auth) KerberosDeleteGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosListGroups
func (a *Auth) KerberosListGroups(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosLogin
func (a *Auth) KerberosLogin(ctx context.Context, request schema.KerberosLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosReadConfiguration
func (a *Auth) KerberosReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosReadGroup
// name: Name of the LDAP group.
func (a *Auth) KerberosReadGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosReadLdapConfiguration
func (a *Auth) KerberosReadLdapConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config/ldap"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KerberosWriteGroup
// name: Name of the LDAP group.
func (a *Auth) KerberosWriteGroup(ctx context.Context, name string, request schema.KerberosWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesConfigureAuth
func (a *Auth) KubernetesConfigureAuth(ctx context.Context, request schema.KubernetesConfigureAuthRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesDeleteAuthRole Register an role with the backend.
// name: Name of the role.
func (a *Auth) KubernetesDeleteAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesListAuthRoles Lists all the roles registered with the backend.
func (a *Auth) KubernetesListAuthRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesLogin Authenticates Kubernetes service accounts with Vault.
func (a *Auth) KubernetesLogin(ctx context.Context, request schema.KubernetesLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesReadAuthConfiguration
func (a *Auth) KubernetesReadAuthConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesReadAuthRole Register an role with the backend.
// name: Name of the role.
func (a *Auth) KubernetesReadAuthRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesWriteAuthRole Register an role with the backend.
// name: Name of the role.
func (a *Auth) KubernetesWriteAuthRole(ctx context.Context, name string, request schema.KubernetesWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapConfigureAuth
func (a *Auth) LdapConfigureAuth(ctx context.Context, request schema.LdapConfigureAuthRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapDeleteGroup Manage additional groups for users allowed to authenticate.
// name: Name of the LDAP group.
func (a *Auth) LdapDeleteGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapDeleteUser Manage users allowed to authenticate.
// name: Name of the LDAP user.
func (a *Auth) LdapDeleteUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapListGroups Manage additional groups for users allowed to authenticate.
func (a *Auth) LdapListGroups(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapListUsers Manage users allowed to authenticate.
func (a *Auth) LdapListUsers(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapLogin Log in with a username and password.
// username: DN (distinguished name) to be used for login.
func (a *Auth) LdapLogin(ctx context.Context, username string, request schema.LdapLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapReadAuthConfiguration
func (a *Auth) LdapReadAuthConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapReadGroup Manage additional groups for users allowed to authenticate.
// name: Name of the LDAP group.
func (a *Auth) LdapReadGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapReadUser Manage users allowed to authenticate.
// name: Name of the LDAP user.
func (a *Auth) LdapReadUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapWriteGroup Manage additional groups for users allowed to authenticate.
// name: Name of the LDAP group.
func (a *Auth) LdapWriteGroup(ctx context.Context, name string, request schema.LdapWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LdapWriteUser Manage users allowed to authenticate.
// name: Name of the LDAP user.
func (a *Auth) LdapWriteUser(ctx context.Context, name string, request schema.LdapWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OciConfigure
func (a *Auth) OciConfigure(ctx context.Context, request schema.OciConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OciDeleteConfiguration
func (a *Auth) OciDeleteConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OciDeleteRole Create a role and associate policies to it.
// role: Name of the role.
func (a *Auth) OciDeleteRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OciListRoles Lists all the roles that are registered with Vault.
func (a *Auth) OciListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OciLogin Authenticates to Vault using OCI credentials
// role: Name of the role.
func (a *Auth) OciLogin(ctx context.Context, role string, request schema.OciLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/login/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OciReadConfiguration
func (a *Auth) OciReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OciReadRole Create a role and associate policies to it.
// role: Name of the role.
func (a *Auth) OciReadRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OciWriteRole Create a role and associate policies to it.
// role: Name of the role.
func (a *Auth) OciWriteRole(ctx context.Context, role string, request schema.OciWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaConfigure
func (a *Auth) OktaConfigure(ctx context.Context, request schema.OktaConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaDeleteGroup Manage users allowed to authenticate.
// name: Name of the Okta group.
func (a *Auth) OktaDeleteGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaDeleteUser Manage additional groups for users allowed to authenticate.
// name: Name of the user.
func (a *Auth) OktaDeleteUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaListGroups Manage users allowed to authenticate.
func (a *Auth) OktaListGroups(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaListUsers Manage additional groups for users allowed to authenticate.
func (a *Auth) OktaListUsers(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaLogin Log in with a username and password.
// username: Username to be used for login.
func (a *Auth) OktaLogin(ctx context.Context, username string, request schema.OktaLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaReadConfiguration
func (a *Auth) OktaReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaReadGroup Manage users allowed to authenticate.
// name: Name of the Okta group.
func (a *Auth) OktaReadGroup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaReadUser Manage additional groups for users allowed to authenticate.
// name: Name of the user.
func (a *Auth) OktaReadUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaVerify
// nonce: Nonce provided during a login request to retrieve the number verification challenge for the matching request.
func (a *Auth) OktaVerify(ctx context.Context, nonce string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/verify/{nonce}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"nonce"+"}", url.PathEscape(nonce), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaWriteGroup Manage users allowed to authenticate.
// name: Name of the Okta group.
func (a *Auth) OktaWriteGroup(ctx context.Context, name string, request schema.OktaWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OktaWriteUser Manage additional groups for users allowed to authenticate.
// name: Name of the user.
func (a *Auth) OktaWriteUser(ctx context.Context, name string, request schema.OktaWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusConfigure
func (a *Auth) RadiusConfigure(ctx context.Context, request schema.RadiusConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusDeleteUser Manage users allowed to authenticate.
// name: Name of the RADIUS user.
func (a *Auth) RadiusDeleteUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusListUsers Manage users allowed to authenticate.
func (a *Auth) RadiusListUsers(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusLogin Log in with a username and password.
func (a *Auth) RadiusLogin(ctx context.Context, request schema.RadiusLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusLoginWithUsername Log in with a username and password.
// urlusername: Username to be used for login. (URL parameter)
func (a *Auth) RadiusLoginWithUsername(ctx context.Context, urlusername string, request schema.RadiusLoginWithUsernameRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/login/{urlusername}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlusername"+"}", url.PathEscape(urlusername), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusReadConfiguration
func (a *Auth) RadiusReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusReadUser Manage users allowed to authenticate.
// name: Name of the RADIUS user.
func (a *Auth) RadiusReadUser(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RadiusWriteUser Manage users allowed to authenticate.
// name: Name of the RADIUS user.
func (a *Auth) RadiusWriteUser(ctx context.Context, name string, request schema.RadiusWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenCreate The token create path is used to create new tokens.
func (a *Auth) TokenCreate(ctx context.Context, request schema.TokenCreateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/create"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenCreateAgainstRole This token create path is used to create new tokens adhering to the given role.
// roleName: Name of the role
func (a *Auth) TokenCreateAgainstRole(ctx context.Context, roleName string, request schema.TokenCreateAgainstRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/create/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenCreateOrphan The token create path is used to create new orphan tokens.
func (a *Auth) TokenCreateOrphan(ctx context.Context, request schema.TokenCreateOrphanRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/create-orphan"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenDeleteRole
// roleName: Name of the role
func (a *Auth) TokenDeleteRole(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/roles/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenListAccessors List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires 'sudo' capability in addition to 'list'.
func (a *Auth) TokenListAccessors(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/accessors/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenListRoles This endpoint lists configured roles.
func (a *Auth) TokenListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/roles/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenLookUp
func (a *Auth) TokenLookUp(ctx context.Context, request schema.TokenLookUpRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/lookup"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenLookUpAccessor This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
func (a *Auth) TokenLookUpAccessor(ctx context.Context, request schema.TokenLookUpAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/lookup-accessor"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenLookUpSelf
func (a *Auth) TokenLookUpSelf(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/lookup-self"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenReadRole
// roleName: Name of the role
func (a *Auth) TokenReadRole(ctx context.Context, roleName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/roles/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRenew This endpoint will renew the given token and prevent expiration.
func (a *Auth) TokenRenew(ctx context.Context, request schema.TokenRenewRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/renew"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRenewAccessor This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
func (a *Auth) TokenRenewAccessor(ctx context.Context, request schema.TokenRenewAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/renew-accessor"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRenewSelf This endpoint will renew the token used to call it and prevent expiration.
func (a *Auth) TokenRenewSelf(ctx context.Context, request schema.TokenRenewSelfRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/renew-self"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRevoke This endpoint will delete the given token and all of its child tokens.
func (a *Auth) TokenRevoke(ctx context.Context, request schema.TokenRevokeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/revoke"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRevokeAccessor This endpoint will delete the token associated with the accessor and all of its child tokens.
func (a *Auth) TokenRevokeAccessor(ctx context.Context, request schema.TokenRevokeAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/revoke-accessor"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRevokeOrphan This endpoint will delete the token and orphan its child tokens.
func (a *Auth) TokenRevokeOrphan(ctx context.Context, request schema.TokenRevokeOrphanRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/revoke-orphan"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenRevokeSelf This endpoint will delete the token used to call it and all of its child tokens.
func (a *Auth) TokenRevokeSelf(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/revoke-self"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenTidy This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
func (a *Auth) TokenTidy(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/tidy"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TokenWriteRole
// roleName: Name of the role
func (a *Auth) TokenWriteRole(ctx context.Context, roleName string, request schema.TokenWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/roles/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassDeleteUser Manage users allowed to authenticate.
// username: Username for this user.
func (a *Auth) UserpassDeleteUser(ctx context.Context, username string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassListUsers Manage users allowed to authenticate.
func (a *Auth) UserpassListUsers(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassLogin Log in with a username and password.
// username: Username of the user.
func (a *Auth) UserpassLogin(ctx context.Context, username string, request schema.UserpassLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassReadUser Manage users allowed to authenticate.
// username: Username for this user.
func (a *Auth) UserpassReadUser(ctx context.Context, username string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassResetPassword Reset user's password.
// username: Username for this user.
func (a *Auth) UserpassResetPassword(ctx context.Context, username string, request schema.UserpassResetPasswordRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}/password"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassUpdatePolicies Update the policies associated with the username.
// username: Username for this user.
func (a *Auth) UserpassUpdatePolicies(ctx context.Context, username string, request schema.UserpassUpdatePoliciesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}/policies"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// UserpassWriteUser Manage users allowed to authenticate.
// username: Username for this user.
func (a *Auth) UserpassWriteUser(ctx context.Context, username string, request schema.UserpassWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}
