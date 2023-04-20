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
// alicloudMountPath: Path that the backend was mounted at
func (a *Auth) AliCloudDeleteAuthRole(ctx context.Context, role string, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// alicloudMountPath: Path that the backend was mounted at
func (a *Auth) AliCloudListAuthRoles(ctx context.Context, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AliCloudListAuthRoles2 Lists all the roles that are registered with Vault.
// alicloudMountPath: Path that the backend was mounted at
func (a *Auth) AliCloudListAuthRoles2(ctx context.Context, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AliCloudLogin Authenticates an RAM entity with Vault.
// alicloudMountPath: Path that the backend was mounted at
func (a *Auth) AliCloudLogin(ctx context.Context, alicloudMountPath string, request schema.AliCloudLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)

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
// alicloudMountPath: Path that the backend was mounted at
func (a *Auth) AliCloudReadAuthRole(ctx context.Context, role string, alicloudMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// alicloudMountPath: Path that the backend was mounted at
func (a *Auth) AliCloudWriteAuthRole(ctx context.Context, role string, alicloudMountPath string, request schema.AliCloudWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{alicloud_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteBindSecretId(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteBoundCidrList(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeletePeriod(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeletePolicies(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteRole(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteSecretIdBoundCidrs(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteSecretIdNumUses(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteSecretIdTtl(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteTokenBoundCidrs(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteTokenMaxTtl(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteTokenNumUses(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDeleteTokenTtl(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDestroySecretId(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleDestroySecretIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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

// AppRoleDestroySecretId2
// roleName: Name of the role. Must be less than 4096 bytes.
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDestroySecretId2(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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

// AppRoleDestroySecretIdByAccessor
// roleName: Name of the role. Must be less than 4096 bytes.
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDestroySecretIdByAccessor(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleDestroySecretIdByAccessorRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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

// AppRoleDestroySecretIdByAccessor2
// roleName: Name of the role. Must be less than 4096 bytes.
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleDestroySecretIdByAccessor2(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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

// AppRoleListRoles
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleListRoles(ctx context.Context, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleListRolesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.AppRoleListRolesResponse](
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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleListSecretIds(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleListSecretIdsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[schema.AppRoleListSecretIdsResponse](
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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleLogin(ctx context.Context, approleMountPath string, request schema.AppRoleLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleLookUpSecretId(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleLookUpSecretIdRequest, options ...RequestOption) (*Response[schema.AppRoleLookUpSecretIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id/lookup"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleLookUpSecretIdByAccessor(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleLookUpSecretIdByAccessorRequest, options ...RequestOption) (*Response[schema.AppRoleLookUpSecretIdByAccessorResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/lookup"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadBindSecretId(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadBindSecretIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadBoundCidrList(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadBoundCidrListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadLocalSecretIds(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadLocalSecretIdsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/local-secret-ids"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadPeriod(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadPeriodResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadPolicies(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadPoliciesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadRole(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadRoleResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadRoleId(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadRoleIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/role-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadSecretIdBoundCidrs(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIdBoundCidrsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadSecretIdNumUses(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIdNumUsesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadSecretIdTtl(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadSecretIdTtlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadTokenBoundCidrs(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadTokenBoundCidrsResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadTokenMaxTtl(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadTokenMaxTtlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadTokenNumUses(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadTokenNumUsesResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleReadTokenTtl(ctx context.Context, roleName string, approleMountPath string, options ...RequestOption) (*Response[schema.AppRoleReadTokenTtlResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleTidySecretId(ctx context.Context, approleMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/tidy/secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteBindSecretId(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteBindSecretIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bind-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteBoundCidrList(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteBoundCidrListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/bound-cidr-list"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteCustomSecretId(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteCustomSecretIdRequest, options ...RequestOption) (*Response[schema.AppRoleWriteCustomSecretIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/custom-secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWritePeriod(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWritePeriodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/period"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWritePolicies(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWritePoliciesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/policies"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteRole(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteRoleId(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteRoleIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/role-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteSecretId(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteSecretIdRequest, options ...RequestOption) (*Response[schema.AppRoleWriteSecretIdResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteSecretIdBoundCidrs(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteSecretIdBoundCidrsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteSecretIdNumUses(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteSecretIdNumUsesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteSecretIdTtl(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteSecretIdTtlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/secret-id-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteTokenBoundCidrs(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteTokenBoundCidrsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteTokenMaxTtl(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteTokenMaxTtlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-max-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteTokenNumUses(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteTokenNumUsesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-num-uses"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// approleMountPath: Path that the backend was mounted at
func (a *Auth) AppRoleWriteTokenTtl(ctx context.Context, roleName string, approleMountPath string, request schema.AppRoleWriteTokenTtlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{approle_mount_path}/role/{role_name}/token-ttl"
	requestPath = strings.Replace(requestPath, "{"+"approle_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("approle")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsConfigureCertificate(ctx context.Context, certName string, awsMountPath string, request schema.AwsConfigureCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsConfigureClient(ctx context.Context, awsMountPath string, request schema.AwsConfigureClientRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsConfigureIdentityAccessListTidyOperation(ctx context.Context, awsMountPath string, request schema.AwsConfigureIdentityAccessListTidyOperationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsConfigureIdentityIntegration(ctx context.Context, awsMountPath string, request schema.AwsConfigureIdentityIntegrationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/identity"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsConfigureIdentityWhitelistTidyOperation(ctx context.Context, awsMountPath string, request schema.AwsConfigureIdentityWhitelistTidyOperationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsConfigureRoleTagBlacklistTidyOperation(ctx context.Context, awsMountPath string, request schema.AwsConfigureRoleTagBlacklistTidyOperationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsConfigureRoleTagDenyListTidyOperation(ctx context.Context, awsMountPath string, request schema.AwsConfigureRoleTagDenyListTidyOperationRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteAuthRole(ctx context.Context, role string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteCertificateConfiguration(ctx context.Context, certName string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteClientConfiguration(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteIdentityAccessList(ctx context.Context, instanceId string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteIdentityAccessListTidySettings(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteIdentityWhitelist(ctx context.Context, instanceId string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteIdentityWhitelistTidySettings(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteRoleTagBlacklist(ctx context.Context, roleTag string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteRoleTagBlacklistTidySettings(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteRoleTagDenyList(ctx context.Context, roleTag string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteRoleTagDenyListTidySettings(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsDeleteStsRole(ctx context.Context, accountId string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsListAuthRoles(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AwsListCertificateConfigurations
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsListCertificateConfigurations(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificates"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AwsListIdentityAccessList
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsListIdentityAccessList(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AwsListIdentityWhitelist
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsListIdentityWhitelist(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AwsListRoleTagBlacklists
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsListRoleTagBlacklists(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AwsListRoleTagDenyLists
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsListRoleTagDenyLists(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AwsListRoles2
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsListRoles2(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AwsListStsRoleRelationships
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsListStsRoleRelationships(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AwsLogin
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsLogin(ctx context.Context, awsMountPath string, request schema.AwsLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadAuthRole(ctx context.Context, role string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadCertificateConfiguration(ctx context.Context, certName string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/certificate/{cert_name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"cert_name"+"}", url.PathEscape(certName), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadClientConfiguration(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/client"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadIdentityAccessList(ctx context.Context, instanceId string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-accesslist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadIdentityAccessListTidySettings(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadIdentityIntegrationConfiguration(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/identity"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadIdentityWhitelist(ctx context.Context, instanceId string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/identity-whitelist/{instance_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"instance_id"+"}", url.PathEscape(instanceId), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadIdentityWhitelistTidySettings(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadRoleTagBlacklist(ctx context.Context, roleTag string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadRoleTagBlacklistTidySettings(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadRoleTagDenyList(ctx context.Context, roleTag string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadRoleTagDenyListTidySettings(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsReadStsRole(ctx context.Context, accountId string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsRotateRootCredentials(ctx context.Context, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsTidyIdentityAccessList(ctx context.Context, awsMountPath string, request schema.AwsTidyIdentityAccessListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/identity-accesslist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsTidyIdentityWhitelist(ctx context.Context, awsMountPath string, request schema.AwsTidyIdentityWhitelistRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/identity-whitelist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsTidyRoleTagBlacklist(ctx context.Context, awsMountPath string, request schema.AwsTidyRoleTagBlacklistRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/roletag-blacklist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsTidyRoleTagDenyList(ctx context.Context, awsMountPath string, request schema.AwsTidyRoleTagDenyListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/tidy/roletag-denylist"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsWriteAuthRole(ctx context.Context, role string, awsMountPath string, request schema.AwsWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsWriteRoleTag(ctx context.Context, role string, awsMountPath string, request schema.AwsWriteRoleTagRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/role/{role}/tag"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsWriteRoleTagBlacklist(ctx context.Context, roleTag string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-blacklist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsWriteRoleTagDenyList(ctx context.Context, roleTag string, awsMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/roletag-denylist/{role_tag}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role_tag"+"}", url.PathEscape(roleTag), -1)

	requestQueryParameters := make(url.Values)

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
// awsMountPath: Path that the backend was mounted at
func (a *Auth) AwsWriteStsRole(ctx context.Context, accountId string, awsMountPath string, request schema.AwsWriteStsRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{aws_mount_path}/config/sts/{account_id}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"account_id"+"}", url.PathEscape(accountId), -1)

	requestQueryParameters := make(url.Values)

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
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureConfigureAuth(ctx context.Context, azureMountPath string, request schema.AzureConfigureAuthRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

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
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureDeleteAuthConfiguration(ctx context.Context, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

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
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureDeleteAuthRole(ctx context.Context, name string, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureListAuthRoles(ctx context.Context, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// AzureLogin
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureLogin(ctx context.Context, azureMountPath string, request schema.AzureLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

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
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureReadAuthConfiguration(ctx context.Context, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

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
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureReadAuthRole(ctx context.Context, name string, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureRotateRootCredentials(ctx context.Context, azureMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

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
// azureMountPath: Path that the backend was mounted at
func (a *Auth) AzureWriteAuthRole(ctx context.Context, name string, azureMountPath string, request schema.AzureWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{azure_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// centrifyMountPath: Path that the backend was mounted at
func (a *Auth) CentrifyConfigure(ctx context.Context, centrifyMountPath string, request schema.CentrifyConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := make(url.Values)

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
// centrifyMountPath: Path that the backend was mounted at
func (a *Auth) CentrifyLogin(ctx context.Context, centrifyMountPath string, request schema.CentrifyLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := make(url.Values)

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
// centrifyMountPath: Path that the backend was mounted at
func (a *Auth) CentrifyReadConfiguration(ctx context.Context, centrifyMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{centrify_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"centrify_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("centrify")), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertConfigure(ctx context.Context, certMountPath string, request schema.CertConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertDeleteCertificate(ctx context.Context, name string, certMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertDeleteCrl(ctx context.Context, name string, certMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertListCertificates(ctx context.Context, certMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// CertListCrls
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertListCrls(ctx context.Context, certMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// CertLogin
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertLogin(ctx context.Context, certMountPath string, request schema.CertLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertReadCertificate(ctx context.Context, name string, certMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertReadConfiguration(ctx context.Context, certMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertReadCrl(ctx context.Context, name string, certMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertWriteCertificate(ctx context.Context, name string, certMountPath string, request schema.CertWriteCertificateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/certs/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// certMountPath: Path that the backend was mounted at
func (a *Auth) CertWriteCrl(ctx context.Context, name string, certMountPath string, request schema.CertWriteCrlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cert_mount_path}/crls/{name}"
	requestPath = strings.Replace(requestPath, "{"+"cert_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cert")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// cfMountPath: Path that the backend was mounted at
func (a *Auth) CloudFoundryConfigure(ctx context.Context, cfMountPath string, request schema.CloudFoundryConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)

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
// cfMountPath: Path that the backend was mounted at
func (a *Auth) CloudFoundryDeleteConfiguration(ctx context.Context, cfMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)

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
// cfMountPath: Path that the backend was mounted at
func (a *Auth) CloudFoundryDeleteRole(ctx context.Context, role string, cfMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// cfMountPath: Path that the backend was mounted at
func (a *Auth) CloudFoundryListRoles(ctx context.Context, cfMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// CloudFoundryLogin
// cfMountPath: Path that the backend was mounted at
func (a *Auth) CloudFoundryLogin(ctx context.Context, cfMountPath string, request schema.CloudFoundryLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)

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
// cfMountPath: Path that the backend was mounted at
func (a *Auth) CloudFoundryReadConfiguration(ctx context.Context, cfMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)

	requestQueryParameters := make(url.Values)

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
// cfMountPath: Path that the backend was mounted at
func (a *Auth) CloudFoundryReadRole(ctx context.Context, role string, cfMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// cfMountPath: Path that the backend was mounted at
func (a *Auth) CloudFoundryWriteRole(ctx context.Context, role string, cfMountPath string, request schema.CloudFoundryWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{cf_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"cf_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cf")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubConfigure(ctx context.Context, githubMountPath string, request schema.GithubConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)

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
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubDeleteTeamMapping(ctx context.Context, key string, githubMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

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
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubDeleteUserMapping(ctx context.Context, key string, githubMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

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

// GithubLogin
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubLogin(ctx context.Context, githubMountPath string, request schema.GithubLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)

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
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubReadConfiguration(ctx context.Context, githubMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)

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
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubReadTeamMapping(ctx context.Context, key string, githubMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

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

// GithubReadTeams Read mappings for teams
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubReadTeams(ctx context.Context, githubMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubReadUserMapping(ctx context.Context, key string, githubMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

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

// GithubReadUsers Read mappings for users
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubReadUsers(ctx context.Context, githubMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubWriteTeamMapping(ctx context.Context, key string, githubMountPath string, request schema.GithubWriteTeamMappingRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/teams/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

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
// githubMountPath: Path that the backend was mounted at
func (a *Auth) GithubWriteUserMapping(ctx context.Context, key string, githubMountPath string, request schema.GithubWriteUserMappingRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{github_mount_path}/map/users/{key}"
	requestPath = strings.Replace(requestPath, "{"+"github_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("github")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

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
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudConfigureAuth(ctx context.Context, gcpMountPath string, request schema.GoogleCloudConfigureAuthRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

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
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudDeleteRole(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudEditLabelsForRole(ctx context.Context, name string, gcpMountPath string, request schema.GoogleCloudEditLabelsForRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}/labels"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudEditServiceAccountsForRole(ctx context.Context, name string, gcpMountPath string, request schema.GoogleCloudEditServiceAccountsForRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}/service-accounts"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudListRoles(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// GoogleCloudListRoles2 Lists all the roles that are registered with Vault.
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudListRoles2(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// GoogleCloudLogin
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudLogin(ctx context.Context, gcpMountPath string, request schema.GoogleCloudLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

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
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudReadAuthConfiguration(ctx context.Context, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

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
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudReadRole(ctx context.Context, name string, gcpMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// gcpMountPath: Path that the backend was mounted at
func (a *Auth) GoogleCloudWriteRole(ctx context.Context, name string, gcpMountPath string, request schema.GoogleCloudWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{gcp_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtConfigure(ctx context.Context, jwtMountPath string, request schema.JwtConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

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
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtDeleteRole(ctx context.Context, name string, jwtMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtListRoles(ctx context.Context, jwtMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// JwtLogin Authenticates to Vault using a JWT (or OIDC) token.
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtLogin(ctx context.Context, jwtMountPath string, request schema.JwtLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

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
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtOidcCallback(ctx context.Context, jwtMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/callback"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

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

// JwtOidcCallbackWithParameters Callback endpoint to handle form_posts.
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtOidcCallbackWithParameters(ctx context.Context, jwtMountPath string, request schema.JwtOidcCallbackWithParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/callback"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

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
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtOidcRequestAuthorizationUrl(ctx context.Context, jwtMountPath string, request schema.JwtOidcRequestAuthorizationUrlRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/oidc/auth_url"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

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
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtReadConfiguration(ctx context.Context, jwtMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)

	requestQueryParameters := make(url.Values)

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
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtReadRole(ctx context.Context, name string, jwtMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// jwtMountPath: Path that the backend was mounted at
func (a *Auth) JwtWriteRole(ctx context.Context, name string, jwtMountPath string, request schema.JwtWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{jwt_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"jwt_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("jwt")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosConfigure(ctx context.Context, kerberosMountPath string, request schema.KerberosConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

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
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosConfigureLdap(ctx context.Context, kerberosMountPath string, request schema.KerberosConfigureLdapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config/ldap"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

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
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosDeleteGroup(ctx context.Context, name string, kerberosMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosListGroups(ctx context.Context, kerberosMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// KerberosLogin
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosLogin(ctx context.Context, kerberosMountPath string, request schema.KerberosLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

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

// KerberosLogin2
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosLogin2(ctx context.Context, kerberosMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

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

// KerberosReadConfiguration
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosReadConfiguration(ctx context.Context, kerberosMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

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
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosReadGroup(ctx context.Context, name string, kerberosMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosReadLdapConfiguration(ctx context.Context, kerberosMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/config/ldap"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)

	requestQueryParameters := make(url.Values)

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
// kerberosMountPath: Path that the backend was mounted at
func (a *Auth) KerberosWriteGroup(ctx context.Context, name string, kerberosMountPath string, request schema.KerberosWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kerberos_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kerberos_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kerberos")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// kubernetesMountPath: Path that the backend was mounted at
func (a *Auth) KubernetesConfigureAuth(ctx context.Context, kubernetesMountPath string, request schema.KubernetesConfigureAuthRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

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
// kubernetesMountPath: Path that the backend was mounted at
func (a *Auth) KubernetesDeleteAuthRole(ctx context.Context, name string, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// kubernetesMountPath: Path that the backend was mounted at
func (a *Auth) KubernetesListAuthRoles(ctx context.Context, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// KubernetesLogin Authenticates Kubernetes service accounts with Vault.
// kubernetesMountPath: Path that the backend was mounted at
func (a *Auth) KubernetesLogin(ctx context.Context, kubernetesMountPath string, request schema.KubernetesLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

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
// kubernetesMountPath: Path that the backend was mounted at
func (a *Auth) KubernetesReadAuthConfiguration(ctx context.Context, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

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
// kubernetesMountPath: Path that the backend was mounted at
func (a *Auth) KubernetesReadAuthRole(ctx context.Context, name string, kubernetesMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// kubernetesMountPath: Path that the backend was mounted at
func (a *Auth) KubernetesWriteAuthRole(ctx context.Context, name string, kubernetesMountPath string, request schema.KubernetesWriteAuthRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{kubernetes_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapConfigureAuth(ctx context.Context, ldapMountPath string, request schema.LdapConfigureAuthRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapDeleteGroup(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapDeleteUser(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapListGroups(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// LdapListUsers Manage users allowed to authenticate.
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapListUsers(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// LdapLogin Log in with a username and password.
// username: DN (distinguished name) to be used for login.
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapLogin(ctx context.Context, username string, ldapMountPath string, request schema.LdapLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapReadAuthConfiguration(ctx context.Context, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapReadGroup(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapReadUser(ctx context.Context, name string, ldapMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapWriteGroup(ctx context.Context, name string, ldapMountPath string, request schema.LdapWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// ldapMountPath: Path that the backend was mounted at
func (a *Auth) LdapWriteUser(ctx context.Context, name string, ldapMountPath string, request schema.LdapWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{ldap_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// ociMountPath: Path that the backend was mounted at
func (a *Auth) OciConfigure(ctx context.Context, ociMountPath string, request schema.OciConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := make(url.Values)

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
// ociMountPath: Path that the backend was mounted at
func (a *Auth) OciDeleteConfiguration(ctx context.Context, ociMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := make(url.Values)

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
// ociMountPath: Path that the backend was mounted at
func (a *Auth) OciDeleteRole(ctx context.Context, role string, ociMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// ociMountPath: Path that the backend was mounted at
func (a *Auth) OciListRoles(ctx context.Context, ociMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// OciLogin Authenticates to Vault using OCI credentials
// role: Name of the role.
// ociMountPath: Path that the backend was mounted at
func (a *Auth) OciLogin(ctx context.Context, role string, ociMountPath string, request schema.OciLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/login/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// ociMountPath: Path that the backend was mounted at
func (a *Auth) OciReadConfiguration(ctx context.Context, ociMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)

	requestQueryParameters := make(url.Values)

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
// ociMountPath: Path that the backend was mounted at
func (a *Auth) OciReadRole(ctx context.Context, role string, ociMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// ociMountPath: Path that the backend was mounted at
func (a *Auth) OciWriteRole(ctx context.Context, role string, ociMountPath string, request schema.OciWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{oci_mount_path}/role/{role}"
	requestPath = strings.Replace(requestPath, "{"+"oci_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("oci")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaConfigure(ctx context.Context, oktaMountPath string, request schema.OktaConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaDeleteGroup(ctx context.Context, name string, oktaMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaDeleteUser(ctx context.Context, name string, oktaMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaListGroups(ctx context.Context, oktaMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// OktaListUsers Manage additional groups for users allowed to authenticate.
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaListUsers(ctx context.Context, oktaMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// OktaLogin Log in with a username and password.
// username: Username to be used for login.
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaLogin(ctx context.Context, username string, oktaMountPath string, request schema.OktaLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaReadConfiguration(ctx context.Context, oktaMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaReadGroup(ctx context.Context, name string, oktaMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaReadUser(ctx context.Context, name string, oktaMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaVerify(ctx context.Context, nonce string, oktaMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/verify/{nonce}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"nonce"+"}", url.PathEscape(nonce), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaWriteGroup(ctx context.Context, name string, oktaMountPath string, request schema.OktaWriteGroupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/groups/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// oktaMountPath: Path that the backend was mounted at
func (a *Auth) OktaWriteUser(ctx context.Context, name string, oktaMountPath string, request schema.OktaWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{okta_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"okta_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("okta")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// radiusMountPath: Path that the backend was mounted at
func (a *Auth) RadiusConfigure(ctx context.Context, radiusMountPath string, request schema.RadiusConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := make(url.Values)

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
// radiusMountPath: Path that the backend was mounted at
func (a *Auth) RadiusDeleteUser(ctx context.Context, name string, radiusMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// radiusMountPath: Path that the backend was mounted at
func (a *Auth) RadiusListUsers(ctx context.Context, radiusMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// RadiusLogin Log in with a username and password.
// radiusMountPath: Path that the backend was mounted at
func (a *Auth) RadiusLogin(ctx context.Context, radiusMountPath string, request schema.RadiusLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/login"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := make(url.Values)

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
// radiusMountPath: Path that the backend was mounted at
func (a *Auth) RadiusLoginWithUsername(ctx context.Context, urlusername string, radiusMountPath string, request schema.RadiusLoginWithUsernameRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/login/{urlusername}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlusername"+"}", url.PathEscape(urlusername), -1)

	requestQueryParameters := make(url.Values)

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
// radiusMountPath: Path that the backend was mounted at
func (a *Auth) RadiusReadConfiguration(ctx context.Context, radiusMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)

	requestQueryParameters := make(url.Values)

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
// radiusMountPath: Path that the backend was mounted at
func (a *Auth) RadiusReadUser(ctx context.Context, name string, radiusMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// radiusMountPath: Path that the backend was mounted at
func (a *Auth) RadiusWriteUser(ctx context.Context, name string, radiusMountPath string, request schema.RadiusWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{radius_mount_path}/users/{name}"
	requestPath = strings.Replace(requestPath, "{"+"radius_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("radius")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

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
// format: Return json formatted output
func (a *Auth) TokenCreate(ctx context.Context, request schema.TokenCreateRequest, format string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/create"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("format", url.QueryEscape(format))

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
// format: Return json formatted output
func (a *Auth) TokenCreateAgainstRole(ctx context.Context, roleName string, request schema.TokenCreateAgainstRoleRequest, format string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/create/{role_name}"
	requestPath = strings.Replace(requestPath, "{"+"role_name"+"}", url.PathEscape(roleName), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("format", url.QueryEscape(format))

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
// format: Return json formatted output
func (a *Auth) TokenCreateOrphan(ctx context.Context, request schema.TokenCreateOrphanRequest, format string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/create-orphan"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("format", url.QueryEscape(format))

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

	requestQueryParameters := make(url.Values)

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
func (a *Auth) TokenListAccessors(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/accessors/"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// TokenListRoles This endpoint lists configured roles.
func (a *Auth) TokenListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/roles"

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// TokenLookUp
func (a *Auth) TokenLookUp(ctx context.Context, request schema.TokenLookUpRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/lookup"

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

// TokenLookUpSelf2
func (a *Auth) TokenLookUpSelf2(ctx context.Context, request schema.TokenLookUpSelf2Request, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/lookup-self"

	requestQueryParameters := make(url.Values)

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

// TokenLookUpSelf3
func (a *Auth) TokenLookUpSelf3(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/token/lookup"

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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

	requestQueryParameters := make(url.Values)

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
// userpassMountPath: Path that the backend was mounted at
func (a *Auth) UserpassDeleteUser(ctx context.Context, username string, userpassMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

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
// userpassMountPath: Path that the backend was mounted at
func (a *Auth) UserpassListUsers(ctx context.Context, userpassMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

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

// UserpassLogin Log in with a username and password.
// username: Username of the user.
// userpassMountPath: Path that the backend was mounted at
func (a *Auth) UserpassLogin(ctx context.Context, username string, userpassMountPath string, request schema.UserpassLoginRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/login/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

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
// userpassMountPath: Path that the backend was mounted at
func (a *Auth) UserpassReadUser(ctx context.Context, username string, userpassMountPath string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

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
// userpassMountPath: Path that the backend was mounted at
func (a *Auth) UserpassResetPassword(ctx context.Context, username string, userpassMountPath string, request schema.UserpassResetPasswordRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}/password"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

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
// userpassMountPath: Path that the backend was mounted at
func (a *Auth) UserpassUpdatePolicies(ctx context.Context, username string, userpassMountPath string, request schema.UserpassUpdatePoliciesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}/policies"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

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
// userpassMountPath: Path that the backend was mounted at
func (a *Auth) UserpassWriteUser(ctx context.Context, username string, userpassMountPath string, request schema.UserpassWriteUserRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/auth/{userpass_mount_path}/users/{username}"
	requestPath = strings.Replace(requestPath, "{"+"userpass_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("userpass")), -1)
	requestPath = strings.Replace(requestPath, "{"+"username"+"}", url.PathEscape(username), -1)

	requestQueryParameters := make(url.Values)

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
