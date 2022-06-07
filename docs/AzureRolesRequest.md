# AzureRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationObjectId** | Pointer to **string** | Application Object ID to use for static service principal credentials. | [optional] 
**AzureGroups** | Pointer to **string** | JSON list of Azure groups to add the service principal to. | [optional] 
**AzureRoles** | Pointer to **string** | JSON list of Azure roles to assign. | [optional] 
**MaxTtl** | Pointer to **int32** | Maximum time a service principal. If not set or set to 0, will use system default. | [optional] 
**Ttl** | Pointer to **int32** | Default lease for generated credentials. If not set or set to 0, will use system default. | [optional] 

## Methods

### NewAzureRolesRequest

`func NewAzureRolesRequest() *AzureRolesRequest`

NewAzureRolesRequest instantiates a new AzureRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRolesRequestWithDefaults

`func NewAzureRolesRequestWithDefaults() *AzureRolesRequest`

NewAzureRolesRequestWithDefaults instantiates a new AzureRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationObjectId

`func (o *AzureRolesRequest) GetApplicationObjectId() string`

GetApplicationObjectId returns the ApplicationObjectId field if non-nil, zero value otherwise.

### GetApplicationObjectIdOk

`func (o *AzureRolesRequest) GetApplicationObjectIdOk() (*string, bool)`

GetApplicationObjectIdOk returns a tuple with the ApplicationObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationObjectId

`func (o *AzureRolesRequest) SetApplicationObjectId(v string)`

SetApplicationObjectId sets ApplicationObjectId field to given value.

### HasApplicationObjectId

`func (o *AzureRolesRequest) HasApplicationObjectId() bool`

HasApplicationObjectId returns a boolean if a field has been set.

### GetAzureGroups

`func (o *AzureRolesRequest) GetAzureGroups() string`

GetAzureGroups returns the AzureGroups field if non-nil, zero value otherwise.

### GetAzureGroupsOk

`func (o *AzureRolesRequest) GetAzureGroupsOk() (*string, bool)`

GetAzureGroupsOk returns a tuple with the AzureGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureGroups

`func (o *AzureRolesRequest) SetAzureGroups(v string)`

SetAzureGroups sets AzureGroups field to given value.

### HasAzureGroups

`func (o *AzureRolesRequest) HasAzureGroups() bool`

HasAzureGroups returns a boolean if a field has been set.

### GetAzureRoles

`func (o *AzureRolesRequest) GetAzureRoles() string`

GetAzureRoles returns the AzureRoles field if non-nil, zero value otherwise.

### GetAzureRolesOk

`func (o *AzureRolesRequest) GetAzureRolesOk() (*string, bool)`

GetAzureRolesOk returns a tuple with the AzureRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRoles

`func (o *AzureRolesRequest) SetAzureRoles(v string)`

SetAzureRoles sets AzureRoles field to given value.

### HasAzureRoles

`func (o *AzureRolesRequest) HasAzureRoles() bool`

HasAzureRoles returns a boolean if a field has been set.

### GetMaxTtl

`func (o *AzureRolesRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *AzureRolesRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *AzureRolesRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *AzureRolesRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetTtl

`func (o *AzureRolesRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AzureRolesRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AzureRolesRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *AzureRolesRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


