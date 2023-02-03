# OIDCWriteProviderTokenRequest


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


### NewOIDCWriteProviderTokenRequest

`func NewOIDCWriteProviderTokenRequest(code string, grantType string, redirectUri string, ) *OIDCWriteProviderTokenRequest`

NewOIDCWriteProviderTokenRequest instantiates a new OIDCWriteProviderTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCWriteProviderTokenRequestWithDefaults

`func NewOIDCWriteProviderTokenRequestWithDefaults() *OIDCWriteProviderTokenRequest`

NewOIDCWriteProviderTokenRequestWithDefaults instantiates a new OIDCWriteProviderTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientId

`func (o *OIDCWriteProviderTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OIDCWriteProviderTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OIDCWriteProviderTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### HasClientId

`func (o *OIDCWriteProviderTokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.




### GetClientSecret

`func (o *OIDCWriteProviderTokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OIDCWriteProviderTokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OIDCWriteProviderTokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### HasClientSecret

`func (o *OIDCWriteProviderTokenRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.




### GetCode

`func (o *OIDCWriteProviderTokenRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OIDCWriteProviderTokenRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OIDCWriteProviderTokenRequest) SetCode(v string)`

SetCode sets Code field to given value.





### GetCodeVerifier

`func (o *OIDCWriteProviderTokenRequest) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *OIDCWriteProviderTokenRequest) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *OIDCWriteProviderTokenRequest) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.


### HasCodeVerifier

`func (o *OIDCWriteProviderTokenRequest) HasCodeVerifier() bool`

HasCodeVerifier returns a boolean if a field has been set.




### GetGrantType

`func (o *OIDCWriteProviderTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OIDCWriteProviderTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OIDCWriteProviderTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.





### GetRedirectUri

`func (o *OIDCWriteProviderTokenRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OIDCWriteProviderTokenRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OIDCWriteProviderTokenRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.










[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


