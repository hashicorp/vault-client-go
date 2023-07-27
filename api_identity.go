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

// Identity is a simple wrapper around the client for Identity requests
type Identity struct {
	client *Client
}

// AliasCreate Create a new alias.
func (i *Identity) AliasCreate(ctx context.Context, request schema.AliasCreateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/alias"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AliasDeleteById
// id: ID of the alias
func (i *Identity) AliasDeleteById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliasListById List all the alias IDs.
func (i *Identity) AliasListById(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/alias/id/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliasReadById
// id: ID of the alias
func (i *Identity) AliasReadById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliasUpdateById
// id: ID of the alias
func (i *Identity) AliasUpdateById(ctx context.Context, id string, request schema.AliasUpdateByIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityBatchDelete Delete all of the entities provided
func (i *Identity) EntityBatchDelete(ctx context.Context, request schema.EntityBatchDeleteRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/batch-delete"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityCreate Create a new entity
func (i *Identity) EntityCreate(ctx context.Context, request schema.EntityCreateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityCreateAlias Create a new alias.
func (i *Identity) EntityCreateAlias(ctx context.Context, request schema.EntityCreateAliasRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity-alias"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityDeleteAliasById
// id: ID of the alias
func (i *Identity) EntityDeleteAliasById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity-alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityDeleteById
// id: ID of the entity. If set, updates the corresponding existing entity.
func (i *Identity) EntityDeleteById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityDeleteByName
// name: Name of the entity
func (i *Identity) EntityDeleteByName(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/name/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityListAliasesById List all the alias IDs.
func (i *Identity) EntityListAliasesById(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity-alias/id/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityListById List all the entity IDs
func (i *Identity) EntityListById(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/id/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityListByName List all the entity names
func (i *Identity) EntityListByName(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/name/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityLookUp Query entities based on various properties.
func (i *Identity) EntityLookUp(ctx context.Context, request schema.EntityLookUpRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/lookup/entity"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityMerge Merge two or more entities together
func (i *Identity) EntityMerge(ctx context.Context, request schema.EntityMergeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/merge"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityReadAliasById
// id: ID of the alias
func (i *Identity) EntityReadAliasById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity-alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityReadById
// id: ID of the entity. If set, updates the corresponding existing entity.
func (i *Identity) EntityReadById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityReadByName
// name: Name of the entity
func (i *Identity) EntityReadByName(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/name/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityUpdateAliasById
// id: ID of the alias
func (i *Identity) EntityUpdateAliasById(ctx context.Context, id string, request schema.EntityUpdateAliasByIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity-alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityUpdateById
// id: ID of the entity. If set, updates the corresponding existing entity.
func (i *Identity) EntityUpdateById(ctx context.Context, id string, request schema.EntityUpdateByIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// EntityUpdateByName
// name: Name of the entity
func (i *Identity) EntityUpdateByName(ctx context.Context, name string, request schema.EntityUpdateByNameRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/entity/name/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupCreate
func (i *Identity) GroupCreate(ctx context.Context, request schema.GroupCreateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupCreateAlias Creates a new group alias, or updates an existing one.
func (i *Identity) GroupCreateAlias(ctx context.Context, request schema.GroupCreateAliasRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group-alias"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupDeleteAliasById
// id: ID of the group alias.
func (i *Identity) GroupDeleteAliasById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group-alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupDeleteById
// id: ID of the group. If set, updates the corresponding existing group.
func (i *Identity) GroupDeleteById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupDeleteByName
// name: Name of the group.
func (i *Identity) GroupDeleteByName(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group/name/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupListAliasesById List all the group alias IDs.
func (i *Identity) GroupListAliasesById(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group-alias/id/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupListById List all the group IDs.
func (i *Identity) GroupListById(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group/id/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupListByName
func (i *Identity) GroupListByName(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group/name/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupLookUp Query groups based on various properties.
func (i *Identity) GroupLookUp(ctx context.Context, request schema.GroupLookUpRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/lookup/group"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupReadAliasById
// id: ID of the group alias.
func (i *Identity) GroupReadAliasById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group-alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupReadById
// id: ID of the group. If set, updates the corresponding existing group.
func (i *Identity) GroupReadById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupReadByName
// name: Name of the group.
func (i *Identity) GroupReadByName(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group/name/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupUpdateAliasById
// id: ID of the group alias.
func (i *Identity) GroupUpdateAliasById(ctx context.Context, id string, request schema.GroupUpdateAliasByIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group-alias/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupUpdateById
// id: ID of the group. If set, updates the corresponding existing group.
func (i *Identity) GroupUpdateById(ctx context.Context, id string, request schema.GroupUpdateByIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GroupUpdateByName
// name: Name of the group.
func (i *Identity) GroupUpdateByName(ctx context.Context, name string, request schema.GroupUpdateByNameRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/group/name/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaAdminDestroyTotpSecret Destroys a TOTP secret for the given MFA method ID on the given entity
func (i *Identity) MfaAdminDestroyTotpSecret(ctx context.Context, request schema.MfaAdminDestroyTotpSecretRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/totp/admin-destroy"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaAdminGenerateTotpSecret Update or create TOTP secret for the given method ID on the given entity.
func (i *Identity) MfaAdminGenerateTotpSecret(ctx context.Context, request schema.MfaAdminGenerateTotpSecretRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/totp/admin-generate"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaCreateDuoMethod Create the given MFA method
func (i *Identity) MfaCreateDuoMethod(ctx context.Context, request schema.MfaCreateDuoMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/duo"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaCreateOktaMethod Create the given MFA method
func (i *Identity) MfaCreateOktaMethod(ctx context.Context, request schema.MfaCreateOktaMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/okta"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaCreatePingIdMethod Create the given MFA method
func (i *Identity) MfaCreatePingIdMethod(ctx context.Context, request schema.MfaCreatePingIdMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/pingid"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaCreateTotpMethod Create the given MFA method
func (i *Identity) MfaCreateTotpMethod(ctx context.Context, request schema.MfaCreateTotpMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/totp"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaDeleteDuoMethod Delete the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaDeleteDuoMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/duo/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaDeleteLoginEnforcement Delete a login enforcement
// name: Name for this login enforcement configuration
func (i *Identity) MfaDeleteLoginEnforcement(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/login-enforcement/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaDeleteOktaMethod Delete the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaDeleteOktaMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/okta/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaDeletePingIdMethod Delete the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaDeletePingIdMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/pingid/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaDeleteTotpMethod Delete the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaDeleteTotpMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/totp/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaGenerateTotpSecret Update or create TOTP secret for the given method ID on the given entity.
func (i *Identity) MfaGenerateTotpSecret(ctx context.Context, request schema.MfaGenerateTotpSecretRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/totp/generate"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaListDuoMethods List MFA method configurations for the given MFA method
func (i *Identity) MfaListDuoMethods(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/duo/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaListLoginEnforcements List login enforcements
func (i *Identity) MfaListLoginEnforcements(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/login-enforcement/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaListMethods List MFA method configurations for all MFA methods
func (i *Identity) MfaListMethods(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaListOktaMethods List MFA method configurations for the given MFA method
func (i *Identity) MfaListOktaMethods(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/okta/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaListPingIdMethods List MFA method configurations for the given MFA method
func (i *Identity) MfaListPingIdMethods(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/pingid/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaListTotpMethods List MFA method configurations for the given MFA method
func (i *Identity) MfaListTotpMethods(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/totp/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaReadDuoMethod Read the current configuration for the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaReadDuoMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/duo/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaReadLoginEnforcement Read the current login enforcement
// name: Name for this login enforcement configuration
func (i *Identity) MfaReadLoginEnforcement(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/login-enforcement/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaReadMethod Read the current configuration for the given ID regardless of the MFA method type
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaReadMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaReadOktaMethod Read the current configuration for the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaReadOktaMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/okta/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaReadPingIdMethod Read the current configuration for the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaReadPingIdMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/pingid/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaReadTotpMethod Read the current configuration for the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaReadTotpMethod(ctx context.Context, methodId string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/totp/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaUpdateDuoMethod Update the configuration for the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaUpdateDuoMethod(ctx context.Context, methodId string, request schema.MfaUpdateDuoMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/duo/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaUpdateOktaMethod Update the configuration for the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaUpdateOktaMethod(ctx context.Context, methodId string, request schema.MfaUpdateOktaMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/okta/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaUpdatePingIdMethod Update the configuration for the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaUpdatePingIdMethod(ctx context.Context, methodId string, request schema.MfaUpdatePingIdMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/pingid/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaUpdateTotpMethod Update the configuration for the given MFA method
// methodId: The unique identifier for this MFA method.
func (i *Identity) MfaUpdateTotpMethod(ctx context.Context, methodId string, request schema.MfaUpdateTotpMethodRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/method/totp/{method_id}"
	requestPath = strings.Replace(requestPath, "{"+"method_id"+"}", url.PathEscape(methodId), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MfaWriteLoginEnforcement Create or update a login enforcement
// name: Name for this login enforcement configuration
func (i *Identity) MfaWriteLoginEnforcement(ctx context.Context, name string, request schema.MfaWriteLoginEnforcementRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/mfa/login-enforcement/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcConfigure
func (i *Identity) OidcConfigure(ctx context.Context, request schema.OidcConfigureRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/config"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcDeleteAssignment
// name: Name of the assignment
func (i *Identity) OidcDeleteAssignment(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/assignment/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcDeleteClient
// name: Name of the client.
func (i *Identity) OidcDeleteClient(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/client/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcDeleteKey CRUD operations for OIDC keys.
// name: Name of the key
func (i *Identity) OidcDeleteKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/key/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcDeleteProvider
// name: Name of the provider
func (i *Identity) OidcDeleteProvider(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcDeleteRole CRUD operations on OIDC Roles
// name: Name of the role
func (i *Identity) OidcDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcDeleteScope
// name: Name of the scope
func (i *Identity) OidcDeleteScope(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/scope/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcGenerateToken Generate an OIDC token
// name: Name of the role
func (i *Identity) OidcGenerateToken(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/token/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcIntrospect Verify the authenticity of an OIDC token
func (i *Identity) OidcIntrospect(ctx context.Context, request schema.OidcIntrospectRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/introspect"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcListAssignments
func (i *Identity) OidcListAssignments(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/assignment/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcListClients
func (i *Identity) OidcListClients(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/client/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcListKeys List OIDC keys
func (i *Identity) OidcListKeys(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/key/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcListProviders
// allowedClientId: Filters the list of OIDC providers to those that allow the given client ID in their set of allowed_client_ids.
func (i *Identity) OidcListProviders(ctx context.Context, allowedClientId string, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("allowed_client_id", parameterToString(allowedClientId))
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcListRoles List configured OIDC roles
func (i *Identity) OidcListRoles(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/role/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcListScopes
func (i *Identity) OidcListScopes(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/scope/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcProviderAuthorize
// name: Name of the provider
// clientId: The ID of the requesting client.
// codeChallenge: The code challenge derived from the code verifier.
// codeChallengeMethod: The method that was used to derive the code challenge. The following methods are supported: &#x27;S256&#x27;, &#x27;plain&#x27;. Defaults to &#x27;plain&#x27;.
// maxAge: The allowable elapsed time in seconds since the last time the end-user was actively authenticated.
// nonce: The value that will be returned in the ID token nonce claim after a token exchange.
// redirectUri: The redirection URI to which the response will be sent.
// responseType: The OIDC authentication flow to be used. The following response types are supported: &#x27;code&#x27;
// scope: A space-delimited, case-sensitive list of scopes to be requested. The &#x27;openid&#x27; scope is required.
// state: The value used to maintain state between the authentication request and client.
func (i *Identity) OidcProviderAuthorize(ctx context.Context, name string, clientId string, codeChallenge string, codeChallengeMethod string, maxAge int32, nonce string, redirectUri string, responseType string, scope string, state string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}/authorize"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("client_id", parameterToString(clientId))
	requestQueryParameters.Add("code_challenge", parameterToString(codeChallenge))
	requestQueryParameters.Add("code_challenge_method", parameterToString(codeChallengeMethod))
	requestQueryParameters.Add("max_age", parameterToString(maxAge))
	requestQueryParameters.Add("nonce", parameterToString(nonce))
	requestQueryParameters.Add("redirect_uri", parameterToString(redirectUri))
	requestQueryParameters.Add("response_type", parameterToString(responseType))
	requestQueryParameters.Add("scope", parameterToString(scope))
	requestQueryParameters.Add("state", parameterToString(state))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcProviderAuthorizeWithParameters
// name: Name of the provider
func (i *Identity) OidcProviderAuthorizeWithParameters(ctx context.Context, name string, request schema.OidcProviderAuthorizeWithParametersRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}/authorize"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcProviderToken
// name: Name of the provider
func (i *Identity) OidcProviderToken(ctx context.Context, name string, request schema.OidcProviderTokenRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcProviderUserInfo
// name: Name of the provider
func (i *Identity) OidcProviderUserInfo(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}/userinfo"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadAssignment
// name: Name of the assignment
func (i *Identity) OidcReadAssignment(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/assignment/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadClient
// name: Name of the client.
func (i *Identity) OidcReadClient(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/client/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadConfiguration
func (i *Identity) OidcReadConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/config"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadKey CRUD operations for OIDC keys.
// name: Name of the key
func (i *Identity) OidcReadKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/key/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadOpenIdConfiguration Query OIDC configurations
func (i *Identity) OidcReadOpenIdConfiguration(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/.well-known/openid-configuration"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadProvider
// name: Name of the provider
func (i *Identity) OidcReadProvider(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadProviderOpenIdConfiguration
// name: Name of the provider
func (i *Identity) OidcReadProviderOpenIdConfiguration(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}/.well-known/openid-configuration"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadProviderPublicKeys
// name: Name of the provider
func (i *Identity) OidcReadProviderPublicKeys(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}/.well-known/keys"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadPublicKeys Retrieve public keys
func (i *Identity) OidcReadPublicKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/.well-known/keys"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadRole CRUD operations on OIDC Roles
// name: Name of the role
func (i *Identity) OidcReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcReadScope
// name: Name of the scope
func (i *Identity) OidcReadScope(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/scope/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcRotateKey Rotate a named OIDC key.
// name: Name of the key
func (i *Identity) OidcRotateKey(ctx context.Context, name string, request schema.OidcRotateKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/key/{name}/rotate"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcWriteAssignment
// name: Name of the assignment
func (i *Identity) OidcWriteAssignment(ctx context.Context, name string, request schema.OidcWriteAssignmentRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/assignment/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcWriteClient
// name: Name of the client.
func (i *Identity) OidcWriteClient(ctx context.Context, name string, request schema.OidcWriteClientRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/client/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcWriteKey CRUD operations for OIDC keys.
// name: Name of the key
func (i *Identity) OidcWriteKey(ctx context.Context, name string, request schema.OidcWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/key/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcWriteProvider
// name: Name of the provider
func (i *Identity) OidcWriteProvider(ctx context.Context, name string, request schema.OidcWriteProviderRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/provider/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcWriteRole CRUD operations on OIDC Roles
// name: Name of the role
func (i *Identity) OidcWriteRole(ctx context.Context, name string, request schema.OidcWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OidcWriteScope
// name: Name of the scope
func (i *Identity) OidcWriteScope(ctx context.Context, name string, request schema.OidcWriteScopeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/oidc/scope/{name}"
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PersonaCreate Create a new alias.
func (i *Identity) PersonaCreate(ctx context.Context, request schema.PersonaCreateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/persona"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PersonaDeleteById
// id: ID of the persona
func (i *Identity) PersonaDeleteById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/persona/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PersonaListById List all the alias IDs.
func (i *Identity) PersonaListById(ctx context.Context, options ...RequestOption) (*Response[schema.StandardListResponse], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/persona/id/"

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[schema.StandardListResponse](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PersonaReadById
// id: ID of the persona
func (i *Identity) PersonaReadById(ctx context.Context, id string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/persona/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PersonaUpdateById
// id: ID of the persona
func (i *Identity) PersonaUpdateById(ctx context.Context, id string, request schema.PersonaUpdateByIdRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/identity/persona/id/{id}"
	requestPath = strings.Replace(requestPath, "{"+"id"+"}", url.PathEscape(id), -1)

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		i.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}
