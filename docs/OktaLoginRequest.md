# OktaLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Nonce** | Pointer to **string** | Nonce provided if performing login that requires number verification challenge. Logins through the vault login CLI command will automatically generate a nonce. | [optional] 
**Password** | Pointer to **string** | Password for this user. | [optional] 
**Provider** | Pointer to **string** | Preferred factor provider. | [optional] 
**Totp** | Pointer to **string** | TOTP passcode. | [optional] 



## Methods


### NewOktaLoginRequest

`func NewOktaLoginRequest() *OktaLoginRequest`

NewOktaLoginRequest instantiates a new OktaLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaLoginRequestWithDefaults

`func NewOktaLoginRequestWithDefaults() *OktaLoginRequest`

NewOktaLoginRequestWithDefaults instantiates a new OktaLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetNonce

`func (o *OktaLoginRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *OktaLoginRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *OktaLoginRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *OktaLoginRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.




### GetPassword

`func (o *OktaLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OktaLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OktaLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### HasPassword

`func (o *OktaLoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.




### GetProvider

`func (o *OktaLoginRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OktaLoginRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OktaLoginRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### HasProvider

`func (o *OktaLoginRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.




### GetTotp

`func (o *OktaLoginRequest) GetTotp() string`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *OktaLoginRequest) GetTotpOk() (*string, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *OktaLoginRequest) SetTotp(v string)`

SetTotp sets Totp field to given value.


### HasTotp

`func (o *OktaLoginRequest) HasTotp() bool`

HasTotp returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


