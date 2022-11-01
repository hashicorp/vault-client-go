# LdapStaticRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | Pointer to **string** | The distinguished name of the entry to manage. | [optional] 
**RotationPeriod** | Pointer to **int32** | Period for automatic credential rotation of the given entry. | [optional] 
**Username** | Pointer to **string** | The username/logon name for the entry with which this role will be associated. | [optional] 

## Methods

### NewLdapStaticRoleRequest

`func NewLdapStaticRoleRequest() *LdapStaticRoleRequest`

NewLdapStaticRoleRequest instantiates a new LdapStaticRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapStaticRoleRequestWithDefaults

`func NewLdapStaticRoleRequestWithDefaults() *LdapStaticRoleRequest`

NewLdapStaticRoleRequestWithDefaults instantiates a new LdapStaticRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *LdapStaticRoleRequest) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *LdapStaticRoleRequest) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *LdapStaticRoleRequest) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *LdapStaticRoleRequest) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRotationPeriod

`func (o *LdapStaticRoleRequest) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *LdapStaticRoleRequest) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *LdapStaticRoleRequest) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *LdapStaticRoleRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### GetUsername

`func (o *LdapStaticRoleRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LdapStaticRoleRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LdapStaticRoleRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LdapStaticRoleRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


