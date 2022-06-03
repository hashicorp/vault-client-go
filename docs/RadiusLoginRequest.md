# RadiusLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Multi-factor auth method to use (optional) | [optional] 
**Passcode** | Pointer to **string** | One time passcode (optional) | [optional] 
**Password** | Pointer to **string** | Password for this user. | [optional] 
**Urlusername** | Pointer to **string** | Username to be used for login. (URL parameter) | [optional] 
**Username** | Pointer to **string** | Username to be used for login. (POST request body) | [optional] 

## Methods

### NewRadiusLoginRequest

`func NewRadiusLoginRequest() *RadiusLoginRequest`

NewRadiusLoginRequest instantiates a new RadiusLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusLoginRequestWithDefaults

`func NewRadiusLoginRequestWithDefaults() *RadiusLoginRequest`

NewRadiusLoginRequestWithDefaults instantiates a new RadiusLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *RadiusLoginRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RadiusLoginRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RadiusLoginRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RadiusLoginRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPasscode

`func (o *RadiusLoginRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *RadiusLoginRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *RadiusLoginRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *RadiusLoginRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### GetPassword

`func (o *RadiusLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RadiusLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RadiusLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RadiusLoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrlusername

`func (o *RadiusLoginRequest) GetUrlusername() string`

GetUrlusername returns the Urlusername field if non-nil, zero value otherwise.

### GetUrlusernameOk

`func (o *RadiusLoginRequest) GetUrlusernameOk() (*string, bool)`

GetUrlusernameOk returns a tuple with the Urlusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlusername

`func (o *RadiusLoginRequest) SetUrlusername(v string)`

SetUrlusername sets Urlusername field to given value.

### HasUrlusername

`func (o *RadiusLoginRequest) HasUrlusername() bool`

HasUrlusername returns a boolean if a field has been set.

### GetUsername

`func (o *RadiusLoginRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RadiusLoginRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RadiusLoginRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RadiusLoginRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


