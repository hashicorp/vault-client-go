# LdapLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Multi-factor auth method to use (optional) | [optional] 
**Passcode** | Pointer to **string** | One time passcode (optional) | [optional] 
**Password** | Pointer to **string** | Password for this user. | [optional] 

## Methods

### NewLdapLoginRequest

`func NewLdapLoginRequest() *LdapLoginRequest`

NewLdapLoginRequest instantiates a new LdapLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapLoginRequestWithDefaults

`func NewLdapLoginRequestWithDefaults() *LdapLoginRequest`

NewLdapLoginRequestWithDefaults instantiates a new LdapLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *LdapLoginRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LdapLoginRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LdapLoginRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LdapLoginRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPasscode

`func (o *LdapLoginRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *LdapLoginRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *LdapLoginRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *LdapLoginRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### GetPassword

`func (o *LdapLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LdapLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LdapLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LdapLoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


