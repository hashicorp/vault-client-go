# LDAPWriteStaticRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | Pointer to **string** | The distinguished name of the entry to manage. | [optional] 
**RotationPeriod** | Pointer to **int32** | Period for automatic credential rotation of the given entry. | [optional] 
**Username** | Pointer to **string** | The username/logon name for the entry with which this role will be associated. | [optional] 

## Methods

### NewLDAPWriteStaticRoleRequest

`func NewLDAPWriteStaticRoleRequest() *LDAPWriteStaticRoleRequest`

NewLDAPWriteStaticRoleRequest instantiates a new LDAPWriteStaticRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPWriteStaticRoleRequestWithDefaults

`func NewLDAPWriteStaticRoleRequestWithDefaults() *LDAPWriteStaticRoleRequest`

NewLDAPWriteStaticRoleRequestWithDefaults instantiates a new LDAPWriteStaticRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *LDAPWriteStaticRoleRequest) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *LDAPWriteStaticRoleRequest) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *LDAPWriteStaticRoleRequest) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *LDAPWriteStaticRoleRequest) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRotationPeriod

`func (o *LDAPWriteStaticRoleRequest) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *LDAPWriteStaticRoleRequest) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *LDAPWriteStaticRoleRequest) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *LDAPWriteStaticRoleRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### GetUsername

`func (o *LDAPWriteStaticRoleRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LDAPWriteStaticRoleRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LDAPWriteStaticRoleRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LDAPWriteStaticRoleRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


