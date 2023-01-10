# RadiusLoginWithUsernameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password for this user. | [optional] 
**Username** | Pointer to **string** | Username to be used for login. (POST request body) | [optional] 

## Methods

### NewRadiusLoginWithUsernameRequest

`func NewRadiusLoginWithUsernameRequest() *RadiusLoginWithUsernameRequest`

NewRadiusLoginWithUsernameRequest instantiates a new RadiusLoginWithUsernameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusLoginWithUsernameRequestWithDefaults

`func NewRadiusLoginWithUsernameRequestWithDefaults() *RadiusLoginWithUsernameRequest`

NewRadiusLoginWithUsernameRequestWithDefaults instantiates a new RadiusLoginWithUsernameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *RadiusLoginWithUsernameRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RadiusLoginWithUsernameRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RadiusLoginWithUsernameRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RadiusLoginWithUsernameRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *RadiusLoginWithUsernameRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RadiusLoginWithUsernameRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RadiusLoginWithUsernameRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RadiusLoginWithUsernameRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


