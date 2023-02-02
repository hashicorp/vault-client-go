# OpenLDAPWriteStaticRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Dn** | Pointer to **string** | The distinguished name of the entry to manage. | [optional] 
**RotationPeriod** | Pointer to **int32** | Period for automatic credential rotation of the given entry. | [optional] 
**Username** | Pointer to **string** | The username/logon name for the entry with which this role will be associated. | [optional] 



## Methods


### NewOpenLDAPWriteStaticRoleRequest

`func NewOpenLDAPWriteStaticRoleRequest() *OpenLDAPWriteStaticRoleRequest`

NewOpenLDAPWriteStaticRoleRequest instantiates a new OpenLDAPWriteStaticRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLDAPWriteStaticRoleRequestWithDefaults

`func NewOpenLDAPWriteStaticRoleRequestWithDefaults() *OpenLDAPWriteStaticRoleRequest`

NewOpenLDAPWriteStaticRoleRequestWithDefaults instantiates a new OpenLDAPWriteStaticRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDn

`func (o *OpenLDAPWriteStaticRoleRequest) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *OpenLDAPWriteStaticRoleRequest) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *OpenLDAPWriteStaticRoleRequest) SetDn(v string)`

SetDn sets Dn field to given value.


### HasDn

`func (o *OpenLDAPWriteStaticRoleRequest) HasDn() bool`

HasDn returns a boolean if a field has been set.




### GetRotationPeriod

`func (o *OpenLDAPWriteStaticRoleRequest) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *OpenLDAPWriteStaticRoleRequest) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *OpenLDAPWriteStaticRoleRequest) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.


### HasRotationPeriod

`func (o *OpenLDAPWriteStaticRoleRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.




### GetUsername

`func (o *OpenLDAPWriteStaticRoleRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OpenLDAPWriteStaticRoleRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OpenLDAPWriteStaticRoleRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### HasUsername

`func (o *OpenLDAPWriteStaticRoleRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


