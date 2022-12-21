# OIDCWriteProviderAuthorizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The ID of the requesting client. | 
**CodeChallenge** | Pointer to **string** | The code challenge derived from the code verifier. | [optional] 
**CodeChallengeMethod** | Pointer to **string** | The method that was used to derive the code challenge. The following methods are supported: &#39;S256&#39;, &#39;plain&#39;. Defaults to &#39;plain&#39;. | [optional] [default to "plain"]
**MaxAge** | Pointer to **int32** | The allowable elapsed time in seconds since the last time the end-user was actively authenticated. | [optional] 
**Nonce** | Pointer to **string** | The value that will be returned in the ID token nonce claim after a token exchange. | [optional] 
**RedirectUri** | **string** | The redirection URI to which the response will be sent. | 
**ResponseType** | **string** | The OIDC authentication flow to be used. The following response types are supported: &#39;code&#39; | 
**Scope** | **string** | A space-delimited, case-sensitive list of scopes to be requested. The &#39;openid&#39; scope is required. | 
**State** | Pointer to **string** | The value used to maintain state between the authentication request and client. | [optional] 

## Methods

### NewOIDCWriteProviderAuthorizeRequest

`func NewOIDCWriteProviderAuthorizeRequest(clientId string, redirectUri string, responseType string, scope string, ) *OIDCWriteProviderAuthorizeRequest`

NewOIDCWriteProviderAuthorizeRequest instantiates a new OIDCWriteProviderAuthorizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCWriteProviderAuthorizeRequestWithDefaults

`func NewOIDCWriteProviderAuthorizeRequestWithDefaults() *OIDCWriteProviderAuthorizeRequest`

NewOIDCWriteProviderAuthorizeRequestWithDefaults instantiates a new OIDCWriteProviderAuthorizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OIDCWriteProviderAuthorizeRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OIDCWriteProviderAuthorizeRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCodeChallenge

`func (o *OIDCWriteProviderAuthorizeRequest) GetCodeChallenge() string`

GetCodeChallenge returns the CodeChallenge field if non-nil, zero value otherwise.

### GetCodeChallengeOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetCodeChallengeOk() (*string, bool)`

GetCodeChallengeOk returns a tuple with the CodeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallenge

`func (o *OIDCWriteProviderAuthorizeRequest) SetCodeChallenge(v string)`

SetCodeChallenge sets CodeChallenge field to given value.

### HasCodeChallenge

`func (o *OIDCWriteProviderAuthorizeRequest) HasCodeChallenge() bool`

HasCodeChallenge returns a boolean if a field has been set.

### GetCodeChallengeMethod

`func (o *OIDCWriteProviderAuthorizeRequest) GetCodeChallengeMethod() string`

GetCodeChallengeMethod returns the CodeChallengeMethod field if non-nil, zero value otherwise.

### GetCodeChallengeMethodOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetCodeChallengeMethodOk() (*string, bool)`

GetCodeChallengeMethodOk returns a tuple with the CodeChallengeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethod

`func (o *OIDCWriteProviderAuthorizeRequest) SetCodeChallengeMethod(v string)`

SetCodeChallengeMethod sets CodeChallengeMethod field to given value.

### HasCodeChallengeMethod

`func (o *OIDCWriteProviderAuthorizeRequest) HasCodeChallengeMethod() bool`

HasCodeChallengeMethod returns a boolean if a field has been set.

### GetMaxAge

`func (o *OIDCWriteProviderAuthorizeRequest) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *OIDCWriteProviderAuthorizeRequest) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *OIDCWriteProviderAuthorizeRequest) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetNonce

`func (o *OIDCWriteProviderAuthorizeRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *OIDCWriteProviderAuthorizeRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *OIDCWriteProviderAuthorizeRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetRedirectUri

`func (o *OIDCWriteProviderAuthorizeRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OIDCWriteProviderAuthorizeRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetResponseType

`func (o *OIDCWriteProviderAuthorizeRequest) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *OIDCWriteProviderAuthorizeRequest) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.


### GetScope

`func (o *OIDCWriteProviderAuthorizeRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OIDCWriteProviderAuthorizeRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetState

`func (o *OIDCWriteProviderAuthorizeRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OIDCWriteProviderAuthorizeRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OIDCWriteProviderAuthorizeRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OIDCWriteProviderAuthorizeRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


