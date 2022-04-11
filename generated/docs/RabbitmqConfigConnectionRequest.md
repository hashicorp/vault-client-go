# RabbitmqConfigConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionUri** | Pointer to **string** | RabbitMQ Management URI | [optional] 
**Password** | Pointer to **string** | Password of the provided RabbitMQ management user | [optional] 
**PasswordPolicy** | Pointer to **string** | Name of the password policy to use to generate passwords for dynamic credentials. | [optional] 
**Username** | Pointer to **string** | Username of a RabbitMQ management administrator | [optional] 
**UsernameTemplate** | Pointer to **string** | Template describing how dynamic usernames are generated. | [optional] 
**VerifyConnection** | Pointer to **bool** | If set, connection_uri is verified by actually connecting to the RabbitMQ management API | [optional] [default to true]

## Methods

### NewRabbitmqConfigConnectionRequest

`func NewRabbitmqConfigConnectionRequest() *RabbitmqConfigConnectionRequest`

NewRabbitmqConfigConnectionRequest instantiates a new RabbitmqConfigConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRabbitmqConfigConnectionRequestWithDefaults

`func NewRabbitmqConfigConnectionRequestWithDefaults() *RabbitmqConfigConnectionRequest`

NewRabbitmqConfigConnectionRequestWithDefaults instantiates a new RabbitmqConfigConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionUri

`func (o *RabbitmqConfigConnectionRequest) GetConnectionUri() string`

GetConnectionUri returns the ConnectionUri field if non-nil, zero value otherwise.

### GetConnectionUriOk

`func (o *RabbitmqConfigConnectionRequest) GetConnectionUriOk() (*string, bool)`

GetConnectionUriOk returns a tuple with the ConnectionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUri

`func (o *RabbitmqConfigConnectionRequest) SetConnectionUri(v string)`

SetConnectionUri sets ConnectionUri field to given value.

### HasConnectionUri

`func (o *RabbitmqConfigConnectionRequest) HasConnectionUri() bool`

HasConnectionUri returns a boolean if a field has been set.

### GetPassword

`func (o *RabbitmqConfigConnectionRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RabbitmqConfigConnectionRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RabbitmqConfigConnectionRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RabbitmqConfigConnectionRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *RabbitmqConfigConnectionRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *RabbitmqConfigConnectionRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *RabbitmqConfigConnectionRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *RabbitmqConfigConnectionRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetUsername

`func (o *RabbitmqConfigConnectionRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RabbitmqConfigConnectionRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RabbitmqConfigConnectionRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RabbitmqConfigConnectionRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUsernameTemplate

`func (o *RabbitmqConfigConnectionRequest) GetUsernameTemplate() string`

GetUsernameTemplate returns the UsernameTemplate field if non-nil, zero value otherwise.

### GetUsernameTemplateOk

`func (o *RabbitmqConfigConnectionRequest) GetUsernameTemplateOk() (*string, bool)`

GetUsernameTemplateOk returns a tuple with the UsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameTemplate

`func (o *RabbitmqConfigConnectionRequest) SetUsernameTemplate(v string)`

SetUsernameTemplate sets UsernameTemplate field to given value.

### HasUsernameTemplate

`func (o *RabbitmqConfigConnectionRequest) HasUsernameTemplate() bool`

HasUsernameTemplate returns a boolean if a field has been set.

### GetVerifyConnection

`func (o *RabbitmqConfigConnectionRequest) GetVerifyConnection() bool`

GetVerifyConnection returns the VerifyConnection field if non-nil, zero value otherwise.

### GetVerifyConnectionOk

`func (o *RabbitmqConfigConnectionRequest) GetVerifyConnectionOk() (*bool, bool)`

GetVerifyConnectionOk returns a tuple with the VerifyConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyConnection

`func (o *RabbitmqConfigConnectionRequest) SetVerifyConnection(v bool)`

SetVerifyConnection sets VerifyConnection field to given value.

### HasVerifyConnection

`func (o *RabbitmqConfigConnectionRequest) HasVerifyConnection() bool`

HasVerifyConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


