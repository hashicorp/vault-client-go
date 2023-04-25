# OidcProviderAuthorize2Request


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The ID of the requesting client. | 
**CodeChallenge** | Pointer to **string** | The code challenge derived from the code verifier. | [optional] 
**CodeChallengeMethod** | Pointer to **string** | The method that was used to derive the code challenge. The following methods are supported: &#x27;S256&#x27;, &#x27;plain&#x27;. Defaults to &#x27;plain&#x27;. | [optional] [default to "plain"]
**MaxAge** | Pointer to **int32** | The allowable elapsed time in seconds since the last time the end-user was actively authenticated. | [optional] 
**Nonce** | Pointer to **string** | The value that will be returned in the ID token nonce claim after a token exchange. | [optional] 
**RedirectUri** | **string** | The redirection URI to which the response will be sent. | 
**ResponseType** | **string** | The OIDC authentication flow to be used. The following response types are supported: &#x27;code&#x27; | 
**Scope** | **string** | A space-delimited, case-sensitive list of scopes to be requested. The &#x27;openid&#x27; scope is required. | 
**State** | Pointer to **string** | The value used to maintain state between the authentication request and client. | [optional] 



## Methods


### NewOidcProviderAuthorize2Request

`func NewOidcProviderAuthorize2Request(clientId string, redirectUri string, responseType string, scope string, ) *OidcProviderAuthorize2Request`

NewOidcProviderAuthorize2Request instantiates a new OidcProviderAuthorize2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcProviderAuthorize2RequestWithDefaults

`func NewOidcProviderAuthorize2RequestWithDefaults() *OidcProviderAuthorize2Request`

NewOidcProviderAuthorize2RequestWithDefaults instantiates a new OidcProviderAuthorize2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientId

`func (o *OidcProviderAuthorize2Request) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OidcProviderAuthorize2Request) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OidcProviderAuthorize2Request) SetClientId(v string)`

SetClientId sets ClientId field to given value.





### GetCodeChallenge

`func (o *OidcProviderAuthorize2Request) GetCodeChallenge() string`

GetCodeChallenge returns the CodeChallenge field if non-nil, zero value otherwise.

### GetCodeChallengeOk

`func (o *OidcProviderAuthorize2Request) GetCodeChallengeOk() (*string, bool)`

GetCodeChallengeOk returns a tuple with the CodeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallenge

`func (o *OidcProviderAuthorize2Request) SetCodeChallenge(v string)`

SetCodeChallenge sets CodeChallenge field to given value.


### HasCodeChallenge

`func (o *OidcProviderAuthorize2Request) HasCodeChallenge() bool`

HasCodeChallenge returns a boolean if a field has been set.




### GetCodeChallengeMethod

`func (o *OidcProviderAuthorize2Request) GetCodeChallengeMethod() string`

GetCodeChallengeMethod returns the CodeChallengeMethod field if non-nil, zero value otherwise.

### GetCodeChallengeMethodOk

`func (o *OidcProviderAuthorize2Request) GetCodeChallengeMethodOk() (*string, bool)`

GetCodeChallengeMethodOk returns a tuple with the CodeChallengeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethod

`func (o *OidcProviderAuthorize2Request) SetCodeChallengeMethod(v string)`

SetCodeChallengeMethod sets CodeChallengeMethod field to given value.


### HasCodeChallengeMethod

`func (o *OidcProviderAuthorize2Request) HasCodeChallengeMethod() bool`

HasCodeChallengeMethod returns a boolean if a field has been set.




### GetMaxAge

`func (o *OidcProviderAuthorize2Request) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *OidcProviderAuthorize2Request) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *OidcProviderAuthorize2Request) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.


### HasMaxAge

`func (o *OidcProviderAuthorize2Request) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.




### GetNonce

`func (o *OidcProviderAuthorize2Request) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *OidcProviderAuthorize2Request) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *OidcProviderAuthorize2Request) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *OidcProviderAuthorize2Request) HasNonce() bool`

HasNonce returns a boolean if a field has been set.




### GetRedirectUri

`func (o *OidcProviderAuthorize2Request) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OidcProviderAuthorize2Request) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OidcProviderAuthorize2Request) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.





### GetResponseType

`func (o *OidcProviderAuthorize2Request) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *OidcProviderAuthorize2Request) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *OidcProviderAuthorize2Request) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.





### GetScope

`func (o *OidcProviderAuthorize2Request) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OidcProviderAuthorize2Request) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OidcProviderAuthorize2Request) SetScope(v string)`

SetScope sets Scope field to given value.





### GetState

`func (o *OidcProviderAuthorize2Request) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OidcProviderAuthorize2Request) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OidcProviderAuthorize2Request) SetState(v string)`

SetState sets State field to given value.


### HasState

`func (o *OidcProviderAuthorize2Request) HasState() bool`

HasState returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


