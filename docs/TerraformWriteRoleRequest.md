# TerraformWriteRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTtl** | Pointer to **int32** | Maximum time for role. If not set or set to 0, will use system default. | [optional] 
**Organization** | Pointer to **string** | Name of the Terraform Cloud or Enterprise organization | [optional] 
**TeamId** | Pointer to **string** | ID of the Terraform Cloud or Enterprise team under organization (e.g., settings/teams/team-xxxxxxxxxxxxx) | [optional] 
**Ttl** | Pointer to **int32** | Default lease for generated credentials. If not set or set to 0, will use system default. | [optional] 
**UserId** | Pointer to **string** | ID of the Terraform Cloud or Enterprise user (e.g., user-xxxxxxxxxxxxxxxx) | [optional] 

## Methods

### NewTerraformWriteRoleRequest

`func NewTerraformWriteRoleRequest() *TerraformWriteRoleRequest`

NewTerraformWriteRoleRequest instantiates a new TerraformWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformWriteRoleRequestWithDefaults

`func NewTerraformWriteRoleRequestWithDefaults() *TerraformWriteRoleRequest`

NewTerraformWriteRoleRequestWithDefaults instantiates a new TerraformWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTtl

`func (o *TerraformWriteRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *TerraformWriteRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *TerraformWriteRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *TerraformWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetOrganization

`func (o *TerraformWriteRoleRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TerraformWriteRoleRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TerraformWriteRoleRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *TerraformWriteRoleRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetTeamId

`func (o *TerraformWriteRoleRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TerraformWriteRoleRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TerraformWriteRoleRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TerraformWriteRoleRequest) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTtl

`func (o *TerraformWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *TerraformWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *TerraformWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *TerraformWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUserId

`func (o *TerraformWriteRoleRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TerraformWriteRoleRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TerraformWriteRoleRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TerraformWriteRoleRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


