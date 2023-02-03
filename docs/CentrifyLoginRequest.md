# CentrifyLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Mode** | Pointer to **string** | Auth mode (&#x27;ro&#x27; for resource owner, &#x27;cc&#x27; for credential client). | [optional] [default to "ro"]
**Password** | Pointer to **string** | Password for this user. | [optional] 
**Username** | Pointer to **string** | Username of the user. | [optional] 



## Methods


### NewCentrifyLoginRequest

`func NewCentrifyLoginRequest() *CentrifyLoginRequest`

NewCentrifyLoginRequest instantiates a new CentrifyLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentrifyLoginRequestWithDefaults

`func NewCentrifyLoginRequestWithDefaults() *CentrifyLoginRequest`

NewCentrifyLoginRequestWithDefaults instantiates a new CentrifyLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMode

`func (o *CentrifyLoginRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CentrifyLoginRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CentrifyLoginRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### HasMode

`func (o *CentrifyLoginRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.




### GetPassword

`func (o *CentrifyLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CentrifyLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CentrifyLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### HasPassword

`func (o *CentrifyLoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.




### GetUsername

`func (o *CentrifyLoginRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CentrifyLoginRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CentrifyLoginRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### HasUsername

`func (o *CentrifyLoginRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


