# OpenLDAPWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationLdif** | **string** | LDIF string used to create new entities within the LDAP system. This LDIF can be templated. | 
**DefaultTtl** | Pointer to **int32** | Default TTL for dynamic credentials | [optional] 
**DeletionLdif** | **string** | LDIF string used to delete entities created within the LDAP system. This LDIF can be templated. | 
**MaxTtl** | Pointer to **int32** | Max TTL a dynamic credential can be extended to | [optional] 
**RollbackLdif** | Pointer to **string** | LDIF string used to rollback changes in the event of a failure to create credentials. This LDIF can be templated. | [optional] 
**UsernameTemplate** | Pointer to **string** | The template used to create a username | [optional] 



## Methods


### NewOpenLDAPWriteRoleRequest

`func NewOpenLDAPWriteRoleRequest(creationLdif string, deletionLdif string, ) *OpenLDAPWriteRoleRequest`

NewOpenLDAPWriteRoleRequest instantiates a new OpenLDAPWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLDAPWriteRoleRequestWithDefaults

`func NewOpenLDAPWriteRoleRequestWithDefaults() *OpenLDAPWriteRoleRequest`

NewOpenLDAPWriteRoleRequestWithDefaults instantiates a new OpenLDAPWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCreationLdif

`func (o *OpenLDAPWriteRoleRequest) GetCreationLdif() string`

GetCreationLdif returns the CreationLdif field if non-nil, zero value otherwise.

### GetCreationLdifOk

`func (o *OpenLDAPWriteRoleRequest) GetCreationLdifOk() (*string, bool)`

GetCreationLdifOk returns a tuple with the CreationLdif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationLdif

`func (o *OpenLDAPWriteRoleRequest) SetCreationLdif(v string)`

SetCreationLdif sets CreationLdif field to given value.





### GetDefaultTtl

`func (o *OpenLDAPWriteRoleRequest) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *OpenLDAPWriteRoleRequest) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *OpenLDAPWriteRoleRequest) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.


### HasDefaultTtl

`func (o *OpenLDAPWriteRoleRequest) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.




### GetDeletionLdif

`func (o *OpenLDAPWriteRoleRequest) GetDeletionLdif() string`

GetDeletionLdif returns the DeletionLdif field if non-nil, zero value otherwise.

### GetDeletionLdifOk

`func (o *OpenLDAPWriteRoleRequest) GetDeletionLdifOk() (*string, bool)`

GetDeletionLdifOk returns a tuple with the DeletionLdif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionLdif

`func (o *OpenLDAPWriteRoleRequest) SetDeletionLdif(v string)`

SetDeletionLdif sets DeletionLdif field to given value.





### GetMaxTtl

`func (o *OpenLDAPWriteRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *OpenLDAPWriteRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *OpenLDAPWriteRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *OpenLDAPWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetRollbackLdif

`func (o *OpenLDAPWriteRoleRequest) GetRollbackLdif() string`

GetRollbackLdif returns the RollbackLdif field if non-nil, zero value otherwise.

### GetRollbackLdifOk

`func (o *OpenLDAPWriteRoleRequest) GetRollbackLdifOk() (*string, bool)`

GetRollbackLdifOk returns a tuple with the RollbackLdif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackLdif

`func (o *OpenLDAPWriteRoleRequest) SetRollbackLdif(v string)`

SetRollbackLdif sets RollbackLdif field to given value.


### HasRollbackLdif

`func (o *OpenLDAPWriteRoleRequest) HasRollbackLdif() bool`

HasRollbackLdif returns a boolean if a field has been set.




### GetUsernameTemplate

`func (o *OpenLDAPWriteRoleRequest) GetUsernameTemplate() string`

GetUsernameTemplate returns the UsernameTemplate field if non-nil, zero value otherwise.

### GetUsernameTemplateOk

`func (o *OpenLDAPWriteRoleRequest) GetUsernameTemplateOk() (*string, bool)`

GetUsernameTemplateOk returns a tuple with the UsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameTemplate

`func (o *OpenLDAPWriteRoleRequest) SetUsernameTemplate(v string)`

SetUsernameTemplate sets UsernameTemplate field to given value.


### HasUsernameTemplate

`func (o *OpenLDAPWriteRoleRequest) HasUsernameTemplate() bool`

HasUsernameTemplate returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


