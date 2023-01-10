# JWTWriteOIDCAuthURLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientNonce** | Pointer to **string** | Optional client-provided nonce that must match during callback, if present. | [optional] 
**RedirectUri** | Pointer to **string** | The OAuth redirect_uri to use in the authorization URL. | [optional] 
**Role** | Pointer to **string** | The role to issue an OIDC authorization URL against. | [optional] 

## Methods

### NewJWTWriteOIDCAuthURLRequest

`func NewJWTWriteOIDCAuthURLRequest() *JWTWriteOIDCAuthURLRequest`

NewJWTWriteOIDCAuthURLRequest instantiates a new JWTWriteOIDCAuthURLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWTWriteOIDCAuthURLRequestWithDefaults

`func NewJWTWriteOIDCAuthURLRequestWithDefaults() *JWTWriteOIDCAuthURLRequest`

NewJWTWriteOIDCAuthURLRequestWithDefaults instantiates a new JWTWriteOIDCAuthURLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientNonce

`func (o *JWTWriteOIDCAuthURLRequest) GetClientNonce() string`

GetClientNonce returns the ClientNonce field if non-nil, zero value otherwise.

### GetClientNonceOk

`func (o *JWTWriteOIDCAuthURLRequest) GetClientNonceOk() (*string, bool)`

GetClientNonceOk returns a tuple with the ClientNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNonce

`func (o *JWTWriteOIDCAuthURLRequest) SetClientNonce(v string)`

SetClientNonce sets ClientNonce field to given value.

### HasClientNonce

`func (o *JWTWriteOIDCAuthURLRequest) HasClientNonce() bool`

HasClientNonce returns a boolean if a field has been set.

### GetRedirectUri

`func (o *JWTWriteOIDCAuthURLRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *JWTWriteOIDCAuthURLRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *JWTWriteOIDCAuthURLRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *JWTWriteOIDCAuthURLRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetRole

`func (o *JWTWriteOIDCAuthURLRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *JWTWriteOIDCAuthURLRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *JWTWriteOIDCAuthURLRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *JWTWriteOIDCAuthURLRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


