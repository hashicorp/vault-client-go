# OpenldapRoleRequest

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

### NewOpenldapRoleRequest

`func NewOpenldapRoleRequest(creationLdif string, deletionLdif string, ) *OpenldapRoleRequest`

NewOpenldapRoleRequest instantiates a new OpenldapRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenldapRoleRequestWithDefaults

`func NewOpenldapRoleRequestWithDefaults() *OpenldapRoleRequest`

NewOpenldapRoleRequestWithDefaults instantiates a new OpenldapRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationLdif

`func (o *OpenldapRoleRequest) GetCreationLdif() string`

GetCreationLdif returns the CreationLdif field if non-nil, zero value otherwise.

### GetCreationLdifOk

`func (o *OpenldapRoleRequest) GetCreationLdifOk() (*string, bool)`

GetCreationLdifOk returns a tuple with the CreationLdif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationLdif

`func (o *OpenldapRoleRequest) SetCreationLdif(v string)`

SetCreationLdif sets CreationLdif field to given value.


### GetDefaultTtl

`func (o *OpenldapRoleRequest) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *OpenldapRoleRequest) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *OpenldapRoleRequest) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *OpenldapRoleRequest) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetDeletionLdif

`func (o *OpenldapRoleRequest) GetDeletionLdif() string`

GetDeletionLdif returns the DeletionLdif field if non-nil, zero value otherwise.

### GetDeletionLdifOk

`func (o *OpenldapRoleRequest) GetDeletionLdifOk() (*string, bool)`

GetDeletionLdifOk returns a tuple with the DeletionLdif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionLdif

`func (o *OpenldapRoleRequest) SetDeletionLdif(v string)`

SetDeletionLdif sets DeletionLdif field to given value.


### GetMaxTtl

`func (o *OpenldapRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *OpenldapRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *OpenldapRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *OpenldapRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetRollbackLdif

`func (o *OpenldapRoleRequest) GetRollbackLdif() string`

GetRollbackLdif returns the RollbackLdif field if non-nil, zero value otherwise.

### GetRollbackLdifOk

`func (o *OpenldapRoleRequest) GetRollbackLdifOk() (*string, bool)`

GetRollbackLdifOk returns a tuple with the RollbackLdif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackLdif

`func (o *OpenldapRoleRequest) SetRollbackLdif(v string)`

SetRollbackLdif sets RollbackLdif field to given value.

### HasRollbackLdif

`func (o *OpenldapRoleRequest) HasRollbackLdif() bool`

HasRollbackLdif returns a boolean if a field has been set.

### GetUsernameTemplate

`func (o *OpenldapRoleRequest) GetUsernameTemplate() string`

GetUsernameTemplate returns the UsernameTemplate field if non-nil, zero value otherwise.

### GetUsernameTemplateOk

`func (o *OpenldapRoleRequest) GetUsernameTemplateOk() (*string, bool)`

GetUsernameTemplateOk returns a tuple with the UsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameTemplate

`func (o *OpenldapRoleRequest) SetUsernameTemplate(v string)`

SetUsernameTemplate sets UsernameTemplate field to given value.

### HasUsernameTemplate

`func (o *OpenldapRoleRequest) HasUsernameTemplate() bool`

HasUsernameTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


