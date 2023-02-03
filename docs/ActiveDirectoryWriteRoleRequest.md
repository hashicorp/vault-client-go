# ActiveDirectoryWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**ServiceAccountName** | Pointer to **string** | The username/logon name for the service account with which this role will be associated. | [optional] 
**Ttl** | Pointer to **int32** | In seconds, the default password time-to-live. | [optional] 



## Methods


### NewActiveDirectoryWriteRoleRequest

`func NewActiveDirectoryWriteRoleRequest() *ActiveDirectoryWriteRoleRequest`

NewActiveDirectoryWriteRoleRequest instantiates a new ActiveDirectoryWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryWriteRoleRequestWithDefaults

`func NewActiveDirectoryWriteRoleRequestWithDefaults() *ActiveDirectoryWriteRoleRequest`

NewActiveDirectoryWriteRoleRequestWithDefaults instantiates a new ActiveDirectoryWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetServiceAccountName

`func (o *ActiveDirectoryWriteRoleRequest) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *ActiveDirectoryWriteRoleRequest) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *ActiveDirectoryWriteRoleRequest) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.


### HasServiceAccountName

`func (o *ActiveDirectoryWriteRoleRequest) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.




### GetTtl

`func (o *ActiveDirectoryWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ActiveDirectoryWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ActiveDirectoryWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *ActiveDirectoryWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


