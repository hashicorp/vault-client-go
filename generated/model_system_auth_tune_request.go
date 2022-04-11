/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SystemAuthTuneRequest struct for SystemAuthTuneRequest
type SystemAuthTuneRequest struct {
	// A list of headers to whitelist and allow a plugin to set on responses.
	AllowedResponseHeaders []string `json:"allowed_response_headers,omitempty"`
	// The list of keys in the request data object that will not be HMAC'ed by audit devices.
	AuditNonHmacRequestKeys []string `json:"audit_non_hmac_request_keys,omitempty"`
	// The list of keys in the response data object that will not be HMAC'ed by audit devices.
	AuditNonHmacResponseKeys []string `json:"audit_non_hmac_response_keys,omitempty"`
	// The default lease TTL for this mount.
	DefaultLeaseTtl *string `json:"default_lease_ttl,omitempty"`
	// User-friendly description for this credential backend.
	Description *string `json:"description,omitempty"`
	// Determines the visibility of the mount in the UI-specific listing endpoint. Accepted value are 'unauth' and ''.
	ListingVisibility *string `json:"listing_visibility,omitempty"`
	// The max lease TTL for this mount.
	MaxLeaseTtl *string `json:"max_lease_ttl,omitempty"`
	// The options to pass into the backend. Should be a json object with string keys and values.
	Options map[string]interface{} `json:"options,omitempty"`
	// A list of headers to whitelist and pass from the request to the plugin.
	PassthroughRequestHeaders []string `json:"passthrough_request_headers,omitempty"`
	// The type of token to issue (service or batch).
	TokenType *string `json:"token_type,omitempty"`
}

// NewSystemAuthTuneRequest instantiates a new SystemAuthTuneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemAuthTuneRequest() *SystemAuthTuneRequest {
	this := SystemAuthTuneRequest{}
	return &this
}

// NewSystemAuthTuneRequestWithDefaults instantiates a new SystemAuthTuneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemAuthTuneRequestWithDefaults() *SystemAuthTuneRequest {
	this := SystemAuthTuneRequest{}
	return &this
}

// GetAllowedResponseHeaders returns the AllowedResponseHeaders field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetAllowedResponseHeaders() []string {
	if o == nil || o.AllowedResponseHeaders == nil {
		var ret []string
		return ret
	}
	return o.AllowedResponseHeaders
}

// GetAllowedResponseHeadersOk returns a tuple with the AllowedResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetAllowedResponseHeadersOk() ([]string, bool) {
	if o == nil || o.AllowedResponseHeaders == nil {
		return nil, false
	}
	return o.AllowedResponseHeaders, true
}

// HasAllowedResponseHeaders returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasAllowedResponseHeaders() bool {
	if o != nil && o.AllowedResponseHeaders != nil {
		return true
	}

	return false
}

// SetAllowedResponseHeaders gets a reference to the given []string and assigns it to the AllowedResponseHeaders field.
func (o *SystemAuthTuneRequest) SetAllowedResponseHeaders(v []string) {
	o.AllowedResponseHeaders = v
}

// GetAuditNonHmacRequestKeys returns the AuditNonHmacRequestKeys field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetAuditNonHmacRequestKeys() []string {
	if o == nil || o.AuditNonHmacRequestKeys == nil {
		var ret []string
		return ret
	}
	return o.AuditNonHmacRequestKeys
}

// GetAuditNonHmacRequestKeysOk returns a tuple with the AuditNonHmacRequestKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetAuditNonHmacRequestKeysOk() ([]string, bool) {
	if o == nil || o.AuditNonHmacRequestKeys == nil {
		return nil, false
	}
	return o.AuditNonHmacRequestKeys, true
}

// HasAuditNonHmacRequestKeys returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasAuditNonHmacRequestKeys() bool {
	if o != nil && o.AuditNonHmacRequestKeys != nil {
		return true
	}

	return false
}

// SetAuditNonHmacRequestKeys gets a reference to the given []string and assigns it to the AuditNonHmacRequestKeys field.
func (o *SystemAuthTuneRequest) SetAuditNonHmacRequestKeys(v []string) {
	o.AuditNonHmacRequestKeys = v
}

// GetAuditNonHmacResponseKeys returns the AuditNonHmacResponseKeys field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetAuditNonHmacResponseKeys() []string {
	if o == nil || o.AuditNonHmacResponseKeys == nil {
		var ret []string
		return ret
	}
	return o.AuditNonHmacResponseKeys
}

// GetAuditNonHmacResponseKeysOk returns a tuple with the AuditNonHmacResponseKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetAuditNonHmacResponseKeysOk() ([]string, bool) {
	if o == nil || o.AuditNonHmacResponseKeys == nil {
		return nil, false
	}
	return o.AuditNonHmacResponseKeys, true
}

// HasAuditNonHmacResponseKeys returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasAuditNonHmacResponseKeys() bool {
	if o != nil && o.AuditNonHmacResponseKeys != nil {
		return true
	}

	return false
}

// SetAuditNonHmacResponseKeys gets a reference to the given []string and assigns it to the AuditNonHmacResponseKeys field.
func (o *SystemAuthTuneRequest) SetAuditNonHmacResponseKeys(v []string) {
	o.AuditNonHmacResponseKeys = v
}

// GetDefaultLeaseTtl returns the DefaultLeaseTtl field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetDefaultLeaseTtl() string {
	if o == nil || o.DefaultLeaseTtl == nil {
		var ret string
		return ret
	}
	return *o.DefaultLeaseTtl
}

// GetDefaultLeaseTtlOk returns a tuple with the DefaultLeaseTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetDefaultLeaseTtlOk() (*string, bool) {
	if o == nil || o.DefaultLeaseTtl == nil {
		return nil, false
	}
	return o.DefaultLeaseTtl, true
}

// HasDefaultLeaseTtl returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasDefaultLeaseTtl() bool {
	if o != nil && o.DefaultLeaseTtl != nil {
		return true
	}

	return false
}

// SetDefaultLeaseTtl gets a reference to the given string and assigns it to the DefaultLeaseTtl field.
func (o *SystemAuthTuneRequest) SetDefaultLeaseTtl(v string) {
	o.DefaultLeaseTtl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SystemAuthTuneRequest) SetDescription(v string) {
	o.Description = &v
}

// GetListingVisibility returns the ListingVisibility field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetListingVisibility() string {
	if o == nil || o.ListingVisibility == nil {
		var ret string
		return ret
	}
	return *o.ListingVisibility
}

// GetListingVisibilityOk returns a tuple with the ListingVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetListingVisibilityOk() (*string, bool) {
	if o == nil || o.ListingVisibility == nil {
		return nil, false
	}
	return o.ListingVisibility, true
}

// HasListingVisibility returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasListingVisibility() bool {
	if o != nil && o.ListingVisibility != nil {
		return true
	}

	return false
}

// SetListingVisibility gets a reference to the given string and assigns it to the ListingVisibility field.
func (o *SystemAuthTuneRequest) SetListingVisibility(v string) {
	o.ListingVisibility = &v
}

// GetMaxLeaseTtl returns the MaxLeaseTtl field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetMaxLeaseTtl() string {
	if o == nil || o.MaxLeaseTtl == nil {
		var ret string
		return ret
	}
	return *o.MaxLeaseTtl
}

// GetMaxLeaseTtlOk returns a tuple with the MaxLeaseTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetMaxLeaseTtlOk() (*string, bool) {
	if o == nil || o.MaxLeaseTtl == nil {
		return nil, false
	}
	return o.MaxLeaseTtl, true
}

// HasMaxLeaseTtl returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasMaxLeaseTtl() bool {
	if o != nil && o.MaxLeaseTtl != nil {
		return true
	}

	return false
}

// SetMaxLeaseTtl gets a reference to the given string and assigns it to the MaxLeaseTtl field.
func (o *SystemAuthTuneRequest) SetMaxLeaseTtl(v string) {
	o.MaxLeaseTtl = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *SystemAuthTuneRequest) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetPassthroughRequestHeaders returns the PassthroughRequestHeaders field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetPassthroughRequestHeaders() []string {
	if o == nil || o.PassthroughRequestHeaders == nil {
		var ret []string
		return ret
	}
	return o.PassthroughRequestHeaders
}

// GetPassthroughRequestHeadersOk returns a tuple with the PassthroughRequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetPassthroughRequestHeadersOk() ([]string, bool) {
	if o == nil || o.PassthroughRequestHeaders == nil {
		return nil, false
	}
	return o.PassthroughRequestHeaders, true
}

// HasPassthroughRequestHeaders returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasPassthroughRequestHeaders() bool {
	if o != nil && o.PassthroughRequestHeaders != nil {
		return true
	}

	return false
}

// SetPassthroughRequestHeaders gets a reference to the given []string and assigns it to the PassthroughRequestHeaders field.
func (o *SystemAuthTuneRequest) SetPassthroughRequestHeaders(v []string) {
	o.PassthroughRequestHeaders = v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *SystemAuthTuneRequest) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAuthTuneRequest) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *SystemAuthTuneRequest) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *SystemAuthTuneRequest) SetTokenType(v string) {
	o.TokenType = &v
}

func (o SystemAuthTuneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedResponseHeaders != nil {
		toSerialize["allowed_response_headers"] = o.AllowedResponseHeaders
	}
	if o.AuditNonHmacRequestKeys != nil {
		toSerialize["audit_non_hmac_request_keys"] = o.AuditNonHmacRequestKeys
	}
	if o.AuditNonHmacResponseKeys != nil {
		toSerialize["audit_non_hmac_response_keys"] = o.AuditNonHmacResponseKeys
	}
	if o.DefaultLeaseTtl != nil {
		toSerialize["default_lease_ttl"] = o.DefaultLeaseTtl
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ListingVisibility != nil {
		toSerialize["listing_visibility"] = o.ListingVisibility
	}
	if o.MaxLeaseTtl != nil {
		toSerialize["max_lease_ttl"] = o.MaxLeaseTtl
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.PassthroughRequestHeaders != nil {
		toSerialize["passthrough_request_headers"] = o.PassthroughRequestHeaders
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	return json.Marshal(toSerialize)
}

type NullableSystemAuthTuneRequest struct {
	value *SystemAuthTuneRequest
	isSet bool
}

func (v NullableSystemAuthTuneRequest) Get() *SystemAuthTuneRequest {
	return v.value
}

func (v *NullableSystemAuthTuneRequest) Set(val *SystemAuthTuneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemAuthTuneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemAuthTuneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemAuthTuneRequest(val *SystemAuthTuneRequest) *NullableSystemAuthTuneRequest {
	return &NullableSystemAuthTuneRequest{value: val, isSet: true}
}

func (v NullableSystemAuthTuneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemAuthTuneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


