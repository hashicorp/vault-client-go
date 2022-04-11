# SystemAuthTuneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedResponseHeaders** | Pointer to **[]string** | A list of headers to whitelist and allow a plugin to set on responses. | [optional] 
**AuditNonHmacRequestKeys** | Pointer to **[]string** | The list of keys in the request data object that will not be HMAC&#39;ed by audit devices. | [optional] 
**AuditNonHmacResponseKeys** | Pointer to **[]string** | The list of keys in the response data object that will not be HMAC&#39;ed by audit devices. | [optional] 
**DefaultLeaseTtl** | Pointer to **string** | The default lease TTL for this mount. | [optional] 
**Description** | Pointer to **string** | User-friendly description for this credential backend. | [optional] 
**ListingVisibility** | Pointer to **string** | Determines the visibility of the mount in the UI-specific listing endpoint. Accepted value are &#39;unauth&#39; and &#39;&#39;. | [optional] 
**MaxLeaseTtl** | Pointer to **string** | The max lease TTL for this mount. | [optional] 
**Options** | Pointer to **map[string]interface{}** | The options to pass into the backend. Should be a json object with string keys and values. | [optional] 
**PassthroughRequestHeaders** | Pointer to **[]string** | A list of headers to whitelist and pass from the request to the plugin. | [optional] 
**TokenType** | Pointer to **string** | The type of token to issue (service or batch). | [optional] 

## Methods

### NewSystemAuthTuneRequest

`func NewSystemAuthTuneRequest() *SystemAuthTuneRequest`

NewSystemAuthTuneRequest instantiates a new SystemAuthTuneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemAuthTuneRequestWithDefaults

`func NewSystemAuthTuneRequestWithDefaults() *SystemAuthTuneRequest`

NewSystemAuthTuneRequestWithDefaults instantiates a new SystemAuthTuneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedResponseHeaders

`func (o *SystemAuthTuneRequest) GetAllowedResponseHeaders() []string`

GetAllowedResponseHeaders returns the AllowedResponseHeaders field if non-nil, zero value otherwise.

### GetAllowedResponseHeadersOk

`func (o *SystemAuthTuneRequest) GetAllowedResponseHeadersOk() (*[]string, bool)`

GetAllowedResponseHeadersOk returns a tuple with the AllowedResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedResponseHeaders

`func (o *SystemAuthTuneRequest) SetAllowedResponseHeaders(v []string)`

SetAllowedResponseHeaders sets AllowedResponseHeaders field to given value.

### HasAllowedResponseHeaders

`func (o *SystemAuthTuneRequest) HasAllowedResponseHeaders() bool`

HasAllowedResponseHeaders returns a boolean if a field has been set.

### GetAuditNonHmacRequestKeys

`func (o *SystemAuthTuneRequest) GetAuditNonHmacRequestKeys() []string`

GetAuditNonHmacRequestKeys returns the AuditNonHmacRequestKeys field if non-nil, zero value otherwise.

### GetAuditNonHmacRequestKeysOk

`func (o *SystemAuthTuneRequest) GetAuditNonHmacRequestKeysOk() (*[]string, bool)`

GetAuditNonHmacRequestKeysOk returns a tuple with the AuditNonHmacRequestKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNonHmacRequestKeys

`func (o *SystemAuthTuneRequest) SetAuditNonHmacRequestKeys(v []string)`

SetAuditNonHmacRequestKeys sets AuditNonHmacRequestKeys field to given value.

### HasAuditNonHmacRequestKeys

`func (o *SystemAuthTuneRequest) HasAuditNonHmacRequestKeys() bool`

HasAuditNonHmacRequestKeys returns a boolean if a field has been set.

### GetAuditNonHmacResponseKeys

`func (o *SystemAuthTuneRequest) GetAuditNonHmacResponseKeys() []string`

GetAuditNonHmacResponseKeys returns the AuditNonHmacResponseKeys field if non-nil, zero value otherwise.

### GetAuditNonHmacResponseKeysOk

`func (o *SystemAuthTuneRequest) GetAuditNonHmacResponseKeysOk() (*[]string, bool)`

GetAuditNonHmacResponseKeysOk returns a tuple with the AuditNonHmacResponseKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNonHmacResponseKeys

`func (o *SystemAuthTuneRequest) SetAuditNonHmacResponseKeys(v []string)`

SetAuditNonHmacResponseKeys sets AuditNonHmacResponseKeys field to given value.

### HasAuditNonHmacResponseKeys

`func (o *SystemAuthTuneRequest) HasAuditNonHmacResponseKeys() bool`

HasAuditNonHmacResponseKeys returns a boolean if a field has been set.

### GetDefaultLeaseTtl

`func (o *SystemAuthTuneRequest) GetDefaultLeaseTtl() string`

GetDefaultLeaseTtl returns the DefaultLeaseTtl field if non-nil, zero value otherwise.

### GetDefaultLeaseTtlOk

`func (o *SystemAuthTuneRequest) GetDefaultLeaseTtlOk() (*string, bool)`

GetDefaultLeaseTtlOk returns a tuple with the DefaultLeaseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLeaseTtl

`func (o *SystemAuthTuneRequest) SetDefaultLeaseTtl(v string)`

SetDefaultLeaseTtl sets DefaultLeaseTtl field to given value.

### HasDefaultLeaseTtl

`func (o *SystemAuthTuneRequest) HasDefaultLeaseTtl() bool`

HasDefaultLeaseTtl returns a boolean if a field has been set.

### GetDescription

`func (o *SystemAuthTuneRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemAuthTuneRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemAuthTuneRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemAuthTuneRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetListingVisibility

`func (o *SystemAuthTuneRequest) GetListingVisibility() string`

GetListingVisibility returns the ListingVisibility field if non-nil, zero value otherwise.

### GetListingVisibilityOk

`func (o *SystemAuthTuneRequest) GetListingVisibilityOk() (*string, bool)`

GetListingVisibilityOk returns a tuple with the ListingVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingVisibility

`func (o *SystemAuthTuneRequest) SetListingVisibility(v string)`

SetListingVisibility sets ListingVisibility field to given value.

### HasListingVisibility

`func (o *SystemAuthTuneRequest) HasListingVisibility() bool`

HasListingVisibility returns a boolean if a field has been set.

### GetMaxLeaseTtl

`func (o *SystemAuthTuneRequest) GetMaxLeaseTtl() string`

GetMaxLeaseTtl returns the MaxLeaseTtl field if non-nil, zero value otherwise.

### GetMaxLeaseTtlOk

`func (o *SystemAuthTuneRequest) GetMaxLeaseTtlOk() (*string, bool)`

GetMaxLeaseTtlOk returns a tuple with the MaxLeaseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLeaseTtl

`func (o *SystemAuthTuneRequest) SetMaxLeaseTtl(v string)`

SetMaxLeaseTtl sets MaxLeaseTtl field to given value.

### HasMaxLeaseTtl

`func (o *SystemAuthTuneRequest) HasMaxLeaseTtl() bool`

HasMaxLeaseTtl returns a boolean if a field has been set.

### GetOptions

`func (o *SystemAuthTuneRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SystemAuthTuneRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SystemAuthTuneRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SystemAuthTuneRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPassthroughRequestHeaders

`func (o *SystemAuthTuneRequest) GetPassthroughRequestHeaders() []string`

GetPassthroughRequestHeaders returns the PassthroughRequestHeaders field if non-nil, zero value otherwise.

### GetPassthroughRequestHeadersOk

`func (o *SystemAuthTuneRequest) GetPassthroughRequestHeadersOk() (*[]string, bool)`

GetPassthroughRequestHeadersOk returns a tuple with the PassthroughRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughRequestHeaders

`func (o *SystemAuthTuneRequest) SetPassthroughRequestHeaders(v []string)`

SetPassthroughRequestHeaders sets PassthroughRequestHeaders field to given value.

### HasPassthroughRequestHeaders

`func (o *SystemAuthTuneRequest) HasPassthroughRequestHeaders() bool`

HasPassthroughRequestHeaders returns a boolean if a field has been set.

### GetTokenType

`func (o *SystemAuthTuneRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *SystemAuthTuneRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *SystemAuthTuneRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *SystemAuthTuneRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


