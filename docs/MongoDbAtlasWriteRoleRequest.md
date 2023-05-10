# MongoDbAtlasWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlocks** | Pointer to **[]string** | Access list entry in CIDR notation to be added for the API key. Optional for organization and project keys. | [optional] 
**IpAddresses** | Pointer to **[]string** | IP address to be added to the access list for the API key. Optional for organization and project keys. | [optional] 
**MaxTtl** | Pointer to **int32** | The maximum allowed lifetime of credentials issued using this role. | [optional] 
**OrganizationId** | Pointer to **string** | Organization ID required for an organization API key | [optional] 
**ProjectId** | Pointer to **string** | Project ID the project API key belongs to. | [optional] 
**ProjectRoles** | Pointer to **[]string** | Roles assigned when an organization API Key is assigned to a project API key | [optional] 
**Roles** | **[]string** | List of roles that the API Key should be granted. A minimum of one role must be provided. Any roles provided must be valid for the assigned Project, required for organization and project keys. | 
**Ttl** | Pointer to **int32** | Duration in seconds after which the issued credential should expire. Defaults to 0, in which case the value will fallback to the system/mount defaults. | [optional] 



## Methods


### NewMongoDbAtlasWriteRoleRequest

`func NewMongoDbAtlasWriteRoleRequest(roles []string, ) *MongoDbAtlasWriteRoleRequest`

NewMongoDbAtlasWriteRoleRequest instantiates a new MongoDbAtlasWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDbAtlasWriteRoleRequestWithDefaults

`func NewMongoDbAtlasWriteRoleRequestWithDefaults() *MongoDbAtlasWriteRoleRequest`

NewMongoDbAtlasWriteRoleRequestWithDefaults instantiates a new MongoDbAtlasWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCidrBlocks

`func (o *MongoDbAtlasWriteRoleRequest) GetCidrBlocks() []string`

GetCidrBlocks returns the CidrBlocks field if non-nil, zero value otherwise.

### GetCidrBlocksOk

`func (o *MongoDbAtlasWriteRoleRequest) GetCidrBlocksOk() (*[]string, bool)`

GetCidrBlocksOk returns a tuple with the CidrBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlocks

`func (o *MongoDbAtlasWriteRoleRequest) SetCidrBlocks(v []string)`

SetCidrBlocks sets CidrBlocks field to given value.


### HasCidrBlocks

`func (o *MongoDbAtlasWriteRoleRequest) HasCidrBlocks() bool`

HasCidrBlocks returns a boolean if a field has been set.




### GetIpAddresses

`func (o *MongoDbAtlasWriteRoleRequest) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *MongoDbAtlasWriteRoleRequest) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *MongoDbAtlasWriteRoleRequest) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### HasIpAddresses

`func (o *MongoDbAtlasWriteRoleRequest) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.




### GetMaxTtl

`func (o *MongoDbAtlasWriteRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *MongoDbAtlasWriteRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *MongoDbAtlasWriteRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *MongoDbAtlasWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetOrganizationId

`func (o *MongoDbAtlasWriteRoleRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *MongoDbAtlasWriteRoleRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *MongoDbAtlasWriteRoleRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### HasOrganizationId

`func (o *MongoDbAtlasWriteRoleRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.




### GetProjectId

`func (o *MongoDbAtlasWriteRoleRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MongoDbAtlasWriteRoleRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MongoDbAtlasWriteRoleRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### HasProjectId

`func (o *MongoDbAtlasWriteRoleRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.




### GetProjectRoles

`func (o *MongoDbAtlasWriteRoleRequest) GetProjectRoles() []string`

GetProjectRoles returns the ProjectRoles field if non-nil, zero value otherwise.

### GetProjectRolesOk

`func (o *MongoDbAtlasWriteRoleRequest) GetProjectRolesOk() (*[]string, bool)`

GetProjectRolesOk returns a tuple with the ProjectRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoles

`func (o *MongoDbAtlasWriteRoleRequest) SetProjectRoles(v []string)`

SetProjectRoles sets ProjectRoles field to given value.


### HasProjectRoles

`func (o *MongoDbAtlasWriteRoleRequest) HasProjectRoles() bool`

HasProjectRoles returns a boolean if a field has been set.




### GetRoles

`func (o *MongoDbAtlasWriteRoleRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MongoDbAtlasWriteRoleRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MongoDbAtlasWriteRoleRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.





### GetTtl

`func (o *MongoDbAtlasWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MongoDbAtlasWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *MongoDbAtlasWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *MongoDbAtlasWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


