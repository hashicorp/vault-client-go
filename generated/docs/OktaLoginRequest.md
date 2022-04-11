# OktaLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Multi-factor auth method to use (optional) | [optional] 
**Passcode** | Pointer to **string** | One time passcode (optional) | [optional] 
**Password** | Pointer to **string** | Password for this user. | [optional] 
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

### GetMethod

`func (o *OktaLoginRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OktaLoginRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OktaLoginRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OktaLoginRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPasscode

`func (o *OktaLoginRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *OktaLoginRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *OktaLoginRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *OktaLoginRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

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


