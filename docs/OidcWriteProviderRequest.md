# OidcWriteProviderRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedClientIds** | Pointer to **[]string** | The client IDs that are permitted to use the provider | [optional] 
**Issuer** | Pointer to **string** | Specifies what will be used for the iss claim of ID tokens. | [optional] 
**ScopesSupported** | Pointer to **[]string** | The scopes supported for requesting on the provider | [optional] 



## Methods


### NewOidcWriteProviderRequest

`func NewOidcWriteProviderRequest() *OidcWriteProviderRequest`

NewOidcWriteProviderRequest instantiates a new OidcWriteProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcWriteProviderRequestWithDefaults

`func NewOidcWriteProviderRequestWithDefaults() *OidcWriteProviderRequest`

NewOidcWriteProviderRequestWithDefaults instantiates a new OidcWriteProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedClientIds

`func (o *OidcWriteProviderRequest) GetAllowedClientIds() []string`

GetAllowedClientIds returns the AllowedClientIds field if non-nil, zero value otherwise.

### GetAllowedClientIdsOk

`func (o *OidcWriteProviderRequest) GetAllowedClientIdsOk() (*[]string, bool)`

GetAllowedClientIdsOk returns a tuple with the AllowedClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientIds

`func (o *OidcWriteProviderRequest) SetAllowedClientIds(v []string)`

SetAllowedClientIds sets AllowedClientIds field to given value.


### HasAllowedClientIds

`func (o *OidcWriteProviderRequest) HasAllowedClientIds() bool`

HasAllowedClientIds returns a boolean if a field has been set.




### GetIssuer

`func (o *OidcWriteProviderRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OidcWriteProviderRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OidcWriteProviderRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### HasIssuer

`func (o *OidcWriteProviderRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.




### GetScopesSupported

`func (o *OidcWriteProviderRequest) GetScopesSupported() []string`

GetScopesSupported returns the ScopesSupported field if non-nil, zero value otherwise.

### GetScopesSupportedOk

`func (o *OidcWriteProviderRequest) GetScopesSupportedOk() (*[]string, bool)`

GetScopesSupportedOk returns a tuple with the ScopesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesSupported

`func (o *OidcWriteProviderRequest) SetScopesSupported(v []string)`

SetScopesSupported sets ScopesSupported field to given value.


### HasScopesSupported

`func (o *OidcWriteProviderRequest) HasScopesSupported() bool`

HasScopesSupported returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


