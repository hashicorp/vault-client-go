# OpenldapStaticRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | Pointer to **string** | The distinguished name of the entry to manage. | [optional] 
**RotationPeriod** | Pointer to **int32** | Period for automatic credential rotation of the given entry. | [optional] 
**Username** | Pointer to **string** | The username/logon name for the entry with which this role will be associated. | [optional] 

## Methods

### NewOpenldapStaticRoleRequest

`func NewOpenldapStaticRoleRequest() *OpenldapStaticRoleRequest`

NewOpenldapStaticRoleRequest instantiates a new OpenldapStaticRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenldapStaticRoleRequestWithDefaults

`func NewOpenldapStaticRoleRequestWithDefaults() *OpenldapStaticRoleRequest`

NewOpenldapStaticRoleRequestWithDefaults instantiates a new OpenldapStaticRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *OpenldapStaticRoleRequest) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *OpenldapStaticRoleRequest) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *OpenldapStaticRoleRequest) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *OpenldapStaticRoleRequest) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRotationPeriod

`func (o *OpenldapStaticRoleRequest) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *OpenldapStaticRoleRequest) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *OpenldapStaticRoleRequest) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *OpenldapStaticRoleRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### GetUsername

`func (o *OpenldapStaticRoleRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OpenldapStaticRoleRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OpenldapStaticRoleRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OpenldapStaticRoleRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


