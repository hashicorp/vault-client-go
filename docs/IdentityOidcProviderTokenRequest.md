# IdentityOidcProviderTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID of the requesting client. | [optional] 
**Code** | **string** | The authorization code received from the provider&#39;s authorization endpoint. | 
**CodeVerifier** | Pointer to **string** | The code verifier associated with the authorization code. | [optional] 
**GrantType** | **string** | The authorization grant type. The following grant types are supported: &#39;authorization_code&#39;. | 
**RedirectUri** | **string** | The callback location where the authentication response was sent. | 

## Methods

### NewIdentityOidcProviderTokenRequest

`func NewIdentityOidcProviderTokenRequest(code string, grantType string, redirectUri string, ) *IdentityOidcProviderTokenRequest`

NewIdentityOidcProviderTokenRequest instantiates a new IdentityOidcProviderTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityOidcProviderTokenRequestWithDefaults

`func NewIdentityOidcProviderTokenRequestWithDefaults() *IdentityOidcProviderTokenRequest`

NewIdentityOidcProviderTokenRequestWithDefaults instantiates a new IdentityOidcProviderTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IdentityOidcProviderTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityOidcProviderTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityOidcProviderTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdentityOidcProviderTokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCode

`func (o *IdentityOidcProviderTokenRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IdentityOidcProviderTokenRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IdentityOidcProviderTokenRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeVerifier

`func (o *IdentityOidcProviderTokenRequest) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *IdentityOidcProviderTokenRequest) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *IdentityOidcProviderTokenRequest) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.

### HasCodeVerifier

`func (o *IdentityOidcProviderTokenRequest) HasCodeVerifier() bool`

HasCodeVerifier returns a boolean if a field has been set.

### GetGrantType

`func (o *IdentityOidcProviderTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *IdentityOidcProviderTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *IdentityOidcProviderTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetRedirectUri

`func (o *IdentityOidcProviderTokenRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *IdentityOidcProviderTokenRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *IdentityOidcProviderTokenRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


