# LdapDuoAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Duo api host | [optional] 
**Ikey** | Pointer to **string** | Duo integration key | [optional] 
**Skey** | Pointer to **string** | Duo secret key | [optional] 

## Methods

### NewLdapDuoAccessRequest

`func NewLdapDuoAccessRequest() *LdapDuoAccessRequest`

NewLdapDuoAccessRequest instantiates a new LdapDuoAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapDuoAccessRequestWithDefaults

`func NewLdapDuoAccessRequestWithDefaults() *LdapDuoAccessRequest`

NewLdapDuoAccessRequestWithDefaults instantiates a new LdapDuoAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *LdapDuoAccessRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LdapDuoAccessRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LdapDuoAccessRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LdapDuoAccessRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIkey

`func (o *LdapDuoAccessRequest) GetIkey() string`

GetIkey returns the Ikey field if non-nil, zero value otherwise.

### GetIkeyOk

`func (o *LdapDuoAccessRequest) GetIkeyOk() (*string, bool)`

GetIkeyOk returns a tuple with the Ikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkey

`func (o *LdapDuoAccessRequest) SetIkey(v string)`

SetIkey sets Ikey field to given value.

### HasIkey

`func (o *LdapDuoAccessRequest) HasIkey() bool`

HasIkey returns a boolean if a field has been set.

### GetSkey

`func (o *LdapDuoAccessRequest) GetSkey() string`

GetSkey returns the Skey field if non-nil, zero value otherwise.

### GetSkeyOk

`func (o *LdapDuoAccessRequest) GetSkeyOk() (*string, bool)`

GetSkeyOk returns a tuple with the Skey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkey

`func (o *LdapDuoAccessRequest) SetSkey(v string)`

SetSkey sets Skey field to given value.

### HasSkey

`func (o *LdapDuoAccessRequest) HasSkey() bool`

HasSkey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


