# OidcProviderTokenRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID of the requesting client. | [optional] 
**ClientSecret** | Pointer to **string** | The secret of the requesting client. | [optional] 
**Code** | **string** | The authorization code received from the provider&#x27;s authorization endpoint. | 
**CodeVerifier** | Pointer to **string** | The code verifier associated with the authorization code. | [optional] 
**GrantType** | **string** | The authorization grant type. The following grant types are supported: &#x27;authorization_code&#x27;. | 
**RedirectUri** | **string** | The callback location where the authentication response was sent. | 



## Methods


### NewOidcProviderTokenRequest

`func NewOidcProviderTokenRequest(code string, grantType string, redirectUri string, ) *OidcProviderTokenRequest`

NewOidcProviderTokenRequest instantiates a new OidcProviderTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcProviderTokenRequestWithDefaults

`func NewOidcProviderTokenRequestWithDefaults() *OidcProviderTokenRequest`

NewOidcProviderTokenRequestWithDefaults instantiates a new OidcProviderTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientId

`func (o *OidcProviderTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OidcProviderTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OidcProviderTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### HasClientId

`func (o *OidcProviderTokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.




### GetClientSecret

`func (o *OidcProviderTokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OidcProviderTokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OidcProviderTokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### HasClientSecret

`func (o *OidcProviderTokenRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.




### GetCode

`func (o *OidcProviderTokenRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OidcProviderTokenRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OidcProviderTokenRequest) SetCode(v string)`

SetCode sets Code field to given value.





### GetCodeVerifier

`func (o *OidcProviderTokenRequest) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *OidcProviderTokenRequest) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *OidcProviderTokenRequest) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.


### HasCodeVerifier

`func (o *OidcProviderTokenRequest) HasCodeVerifier() bool`

HasCodeVerifier returns a boolean if a field has been set.




### GetGrantType

`func (o *OidcProviderTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OidcProviderTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OidcProviderTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.





### GetRedirectUri

`func (o *OidcProviderTokenRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OidcProviderTokenRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OidcProviderTokenRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.










[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


