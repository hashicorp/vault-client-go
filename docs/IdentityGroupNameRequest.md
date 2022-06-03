# IdentityGroupNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the group. If set, updates the corresponding existing group. | [optional] 
**MemberEntityIds** | Pointer to **[]string** | Entity IDs to be assigned as group members. | [optional] 
**MemberGroupIds** | Pointer to **[]string** | Group IDs to be assigned as group members. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to be associated with the group. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault &lt;command&gt; &lt;path&gt; metadata&#x3D;key1&#x3D;value1 metadata&#x3D;key2&#x3D;value2 | [optional] 
**Policies** | Pointer to **[]string** | Policies to be tied to the group. | [optional] 
**Type** | Pointer to **string** | Type of the group, &#39;internal&#39; or &#39;external&#39;. Defaults to &#39;internal&#39; | [optional] 

## Methods

### NewIdentityGroupNameRequest

`func NewIdentityGroupNameRequest() *IdentityGroupNameRequest`

NewIdentityGroupNameRequest instantiates a new IdentityGroupNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityGroupNameRequestWithDefaults

`func NewIdentityGroupNameRequestWithDefaults() *IdentityGroupNameRequest`

NewIdentityGroupNameRequestWithDefaults instantiates a new IdentityGroupNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityGroupNameRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityGroupNameRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityGroupNameRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityGroupNameRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemberEntityIds

`func (o *IdentityGroupNameRequest) GetMemberEntityIds() []string`

GetMemberEntityIds returns the MemberEntityIds field if non-nil, zero value otherwise.

### GetMemberEntityIdsOk

`func (o *IdentityGroupNameRequest) GetMemberEntityIdsOk() (*[]string, bool)`

GetMemberEntityIdsOk returns a tuple with the MemberEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEntityIds

`func (o *IdentityGroupNameRequest) SetMemberEntityIds(v []string)`

SetMemberEntityIds sets MemberEntityIds field to given value.

### HasMemberEntityIds

`func (o *IdentityGroupNameRequest) HasMemberEntityIds() bool`

HasMemberEntityIds returns a boolean if a field has been set.

### GetMemberGroupIds

`func (o *IdentityGroupNameRequest) GetMemberGroupIds() []string`

GetMemberGroupIds returns the MemberGroupIds field if non-nil, zero value otherwise.

### GetMemberGroupIdsOk

`func (o *IdentityGroupNameRequest) GetMemberGroupIdsOk() (*[]string, bool)`

GetMemberGroupIdsOk returns a tuple with the MemberGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroupIds

`func (o *IdentityGroupNameRequest) SetMemberGroupIds(v []string)`

SetMemberGroupIds sets MemberGroupIds field to given value.

### HasMemberGroupIds

`func (o *IdentityGroupNameRequest) HasMemberGroupIds() bool`

HasMemberGroupIds returns a boolean if a field has been set.

### GetMetadata

`func (o *IdentityGroupNameRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentityGroupNameRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentityGroupNameRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IdentityGroupNameRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPolicies

`func (o *IdentityGroupNameRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *IdentityGroupNameRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *IdentityGroupNameRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *IdentityGroupNameRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetType

`func (o *IdentityGroupNameRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityGroupNameRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityGroupNameRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityGroupNameRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


