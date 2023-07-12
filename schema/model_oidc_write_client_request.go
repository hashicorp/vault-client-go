// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OidcWriteClientRequest struct for OidcWriteClientRequest
type OidcWriteClientRequest struct {
	// The time-to-live for access tokens obtained by the client.
	AccessTokenTtl string `json:"access_token_ttl,omitempty"`

	// Comma separated string or array of assignment resources.
	Assignments []string `json:"assignments,omitempty"`

	// The client type based on its ability to maintain confidentiality of credentials. The following client types are supported: 'confidential', 'public'. Defaults to 'confidential'.
	ClientType string `json:"client_type,omitempty"`

	// The time-to-live for ID tokens obtained by the client.
	IdTokenTtl string `json:"id_token_ttl,omitempty"`

	// A reference to a named key resource. Cannot be modified after creation. Defaults to the 'default' key.
	Key string `json:"key,omitempty"`

	// Comma separated string or array of redirect URIs used by the client. One of these values must exactly match the redirect_uri parameter value used in each authentication request.
	RedirectUris []string `json:"redirect_uris,omitempty"`
}
