# OIDCWriteAuthURLRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientNonce** | Pointer to **string** | Optional client-provided nonce that must match during callback, if present. | [optional] 
**RedirectUri** | Pointer to **string** | The OAuth redirect_uri to use in the authorization URL. | [optional] 
**Role** | Pointer to **string** | The role to issue an OIDC authorization URL against. | [optional] 



## Methods


### NewOIDCWriteAuthURLRequest

`func NewOIDCWriteAuthURLRequest() *OIDCWriteAuthURLRequest`

NewOIDCWriteAuthURLRequest instantiates a new OIDCWriteAuthURLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCWriteAuthURLRequestWithDefaults

`func NewOIDCWriteAuthURLRequestWithDefaults() *OIDCWriteAuthURLRequest`

NewOIDCWriteAuthURLRequestWithDefaults instantiates a new OIDCWriteAuthURLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientNonce

`func (o *OIDCWriteAuthURLRequest) GetClientNonce() string`

GetClientNonce returns the ClientNonce field if non-nil, zero value otherwise.

### GetClientNonceOk

`func (o *OIDCWriteAuthURLRequest) GetClientNonceOk() (*string, bool)`

GetClientNonceOk returns a tuple with the ClientNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNonce

`func (o *OIDCWriteAuthURLRequest) SetClientNonce(v string)`

SetClientNonce sets ClientNonce field to given value.


### HasClientNonce

`func (o *OIDCWriteAuthURLRequest) HasClientNonce() bool`

HasClientNonce returns a boolean if a field has been set.




### GetRedirectUri

`func (o *OIDCWriteAuthURLRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OIDCWriteAuthURLRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OIDCWriteAuthURLRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### HasRedirectUri

`func (o *OIDCWriteAuthURLRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.




### GetRole

`func (o *OIDCWriteAuthURLRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OIDCWriteAuthURLRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OIDCWriteAuthURLRequest) SetRole(v string)`

SetRole sets Role field to given value.


### HasRole

`func (o *OIDCWriteAuthURLRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


