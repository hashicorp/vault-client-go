# MongoDBAtlasWriteRoleRequest


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


### NewMongoDBAtlasWriteRoleRequest

`func NewMongoDBAtlasWriteRoleRequest(roles []string, ) *MongoDBAtlasWriteRoleRequest`

NewMongoDBAtlasWriteRoleRequest instantiates a new MongoDBAtlasWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBAtlasWriteRoleRequestWithDefaults

`func NewMongoDBAtlasWriteRoleRequestWithDefaults() *MongoDBAtlasWriteRoleRequest`

NewMongoDBAtlasWriteRoleRequestWithDefaults instantiates a new MongoDBAtlasWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCidrBlocks

`func (o *MongoDBAtlasWriteRoleRequest) GetCidrBlocks() []string`

GetCidrBlocks returns the CidrBlocks field if non-nil, zero value otherwise.

### GetCidrBlocksOk

`func (o *MongoDBAtlasWriteRoleRequest) GetCidrBlocksOk() (*[]string, bool)`

GetCidrBlocksOk returns a tuple with the CidrBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlocks

`func (o *MongoDBAtlasWriteRoleRequest) SetCidrBlocks(v []string)`

SetCidrBlocks sets CidrBlocks field to given value.


### HasCidrBlocks

`func (o *MongoDBAtlasWriteRoleRequest) HasCidrBlocks() bool`

HasCidrBlocks returns a boolean if a field has been set.




### GetIpAddresses

`func (o *MongoDBAtlasWriteRoleRequest) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *MongoDBAtlasWriteRoleRequest) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *MongoDBAtlasWriteRoleRequest) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### HasIpAddresses

`func (o *MongoDBAtlasWriteRoleRequest) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.




### GetMaxTtl

`func (o *MongoDBAtlasWriteRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *MongoDBAtlasWriteRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *MongoDBAtlasWriteRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *MongoDBAtlasWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetOrganizationId

`func (o *MongoDBAtlasWriteRoleRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *MongoDBAtlasWriteRoleRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *MongoDBAtlasWriteRoleRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### HasOrganizationId

`func (o *MongoDBAtlasWriteRoleRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.




### GetProjectId

`func (o *MongoDBAtlasWriteRoleRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MongoDBAtlasWriteRoleRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MongoDBAtlasWriteRoleRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### HasProjectId

`func (o *MongoDBAtlasWriteRoleRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.




### GetProjectRoles

`func (o *MongoDBAtlasWriteRoleRequest) GetProjectRoles() []string`

GetProjectRoles returns the ProjectRoles field if non-nil, zero value otherwise.

### GetProjectRolesOk

`func (o *MongoDBAtlasWriteRoleRequest) GetProjectRolesOk() (*[]string, bool)`

GetProjectRolesOk returns a tuple with the ProjectRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoles

`func (o *MongoDBAtlasWriteRoleRequest) SetProjectRoles(v []string)`

SetProjectRoles sets ProjectRoles field to given value.


### HasProjectRoles

`func (o *MongoDBAtlasWriteRoleRequest) HasProjectRoles() bool`

HasProjectRoles returns a boolean if a field has been set.




### GetRoles

`func (o *MongoDBAtlasWriteRoleRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MongoDBAtlasWriteRoleRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MongoDBAtlasWriteRoleRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.





### GetTtl

`func (o *MongoDBAtlasWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MongoDBAtlasWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *MongoDBAtlasWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *MongoDBAtlasWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


