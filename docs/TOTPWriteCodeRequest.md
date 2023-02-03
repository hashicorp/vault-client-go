# TOTPWriteCodeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Code** | Pointer to **string** | TOTP code to be validated. | [optional] 



## Methods


### NewTOTPWriteCodeRequest

`func NewTOTPWriteCodeRequest() *TOTPWriteCodeRequest`

NewTOTPWriteCodeRequest instantiates a new TOTPWriteCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTOTPWriteCodeRequestWithDefaults

`func NewTOTPWriteCodeRequestWithDefaults() *TOTPWriteCodeRequest`

NewTOTPWriteCodeRequestWithDefaults instantiates a new TOTPWriteCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCode

`func (o *TOTPWriteCodeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TOTPWriteCodeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TOTPWriteCodeRequest) SetCode(v string)`

SetCode sets Code field to given value.


### HasCode

`func (o *TOTPWriteCodeRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


