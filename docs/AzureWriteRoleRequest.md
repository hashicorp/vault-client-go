# AzureWriteRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationObjectId** | Pointer to **string** | Application Object ID to use for static service principal credentials. | [optional] 
**AzureGroups** | Pointer to **string** | JSON list of Azure groups to add the service principal to. | [optional] 
**AzureRoles** | Pointer to **string** | JSON list of Azure roles to assign. | [optional] 
**MaxTtl** | Pointer to **int32** | Maximum time a service principal. If not set or set to 0, will use system default. | [optional] 
**PermanentlyDelete** | Pointer to **bool** | Indicates whether new application objects should be permanently deleted. If not set, objects will not be permanently deleted. | [optional] [default to false]
**Ttl** | Pointer to **int32** | Default lease for generated credentials. If not set or set to 0, will use system default. | [optional] 

## Methods

### NewAzureWriteRoleRequest

`func NewAzureWriteRoleRequest() *AzureWriteRoleRequest`

NewAzureWriteRoleRequest instantiates a new AzureWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureWriteRoleRequestWithDefaults

`func NewAzureWriteRoleRequestWithDefaults() *AzureWriteRoleRequest`

NewAzureWriteRoleRequestWithDefaults instantiates a new AzureWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationObjectId

`func (o *AzureWriteRoleRequest) GetApplicationObjectId() string`

GetApplicationObjectId returns the ApplicationObjectId field if non-nil, zero value otherwise.

### GetApplicationObjectIdOk

`func (o *AzureWriteRoleRequest) GetApplicationObjectIdOk() (*string, bool)`

GetApplicationObjectIdOk returns a tuple with the ApplicationObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationObjectId

`func (o *AzureWriteRoleRequest) SetApplicationObjectId(v string)`

SetApplicationObjectId sets ApplicationObjectId field to given value.

### HasApplicationObjectId

`func (o *AzureWriteRoleRequest) HasApplicationObjectId() bool`

HasApplicationObjectId returns a boolean if a field has been set.

### GetAzureGroups

`func (o *AzureWriteRoleRequest) GetAzureGroups() string`

GetAzureGroups returns the AzureGroups field if non-nil, zero value otherwise.

### GetAzureGroupsOk

`func (o *AzureWriteRoleRequest) GetAzureGroupsOk() (*string, bool)`

GetAzureGroupsOk returns a tuple with the AzureGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureGroups

`func (o *AzureWriteRoleRequest) SetAzureGroups(v string)`

SetAzureGroups sets AzureGroups field to given value.

### HasAzureGroups

`func (o *AzureWriteRoleRequest) HasAzureGroups() bool`

HasAzureGroups returns a boolean if a field has been set.

### GetAzureRoles

`func (o *AzureWriteRoleRequest) GetAzureRoles() string`

GetAzureRoles returns the AzureRoles field if non-nil, zero value otherwise.

### GetAzureRolesOk

`func (o *AzureWriteRoleRequest) GetAzureRolesOk() (*string, bool)`

GetAzureRolesOk returns a tuple with the AzureRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRoles

`func (o *AzureWriteRoleRequest) SetAzureRoles(v string)`

SetAzureRoles sets AzureRoles field to given value.

### HasAzureRoles

`func (o *AzureWriteRoleRequest) HasAzureRoles() bool`

HasAzureRoles returns a boolean if a field has been set.

### GetMaxTtl

`func (o *AzureWriteRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *AzureWriteRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *AzureWriteRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *AzureWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPermanentlyDelete

`func (o *AzureWriteRoleRequest) GetPermanentlyDelete() bool`

GetPermanentlyDelete returns the PermanentlyDelete field if non-nil, zero value otherwise.

### GetPermanentlyDeleteOk

`func (o *AzureWriteRoleRequest) GetPermanentlyDeleteOk() (*bool, bool)`

GetPermanentlyDeleteOk returns a tuple with the PermanentlyDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentlyDelete

`func (o *AzureWriteRoleRequest) SetPermanentlyDelete(v bool)`

SetPermanentlyDelete sets PermanentlyDelete field to given value.

### HasPermanentlyDelete

`func (o *AzureWriteRoleRequest) HasPermanentlyDelete() bool`

HasPermanentlyDelete returns a boolean if a field has been set.

### GetTtl

`func (o *AzureWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AzureWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AzureWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *AzureWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


