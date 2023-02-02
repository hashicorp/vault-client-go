# GroupWriteByIDRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**MemberEntityIds** | Pointer to **[]string** | Entity IDs to be assigned as group members. | [optional] 
**MemberGroupIds** | Pointer to **[]string** | Group IDs to be assigned as group members. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to be associated with the group. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault &lt;command&gt; &lt;path&gt; metadata&#x3D;key1&#x3D;value1 metadata&#x3D;key2&#x3D;value2 | [optional] 
**Name** | Pointer to **string** | Name of the group. | [optional] 
**Policies** | Pointer to **[]string** | Policies to be tied to the group. | [optional] 
**Type** | Pointer to **string** | Type of the group, &#x27;internal&#x27; or &#x27;external&#x27;. Defaults to &#x27;internal&#x27; | [optional] 



## Methods


### NewGroupWriteByIDRequest

`func NewGroupWriteByIDRequest() *GroupWriteByIDRequest`

NewGroupWriteByIDRequest instantiates a new GroupWriteByIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWriteByIDRequestWithDefaults

`func NewGroupWriteByIDRequestWithDefaults() *GroupWriteByIDRequest`

NewGroupWriteByIDRequestWithDefaults instantiates a new GroupWriteByIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMemberEntityIds

`func (o *GroupWriteByIDRequest) GetMemberEntityIds() []string`

GetMemberEntityIds returns the MemberEntityIds field if non-nil, zero value otherwise.

### GetMemberEntityIdsOk

`func (o *GroupWriteByIDRequest) GetMemberEntityIdsOk() (*[]string, bool)`

GetMemberEntityIdsOk returns a tuple with the MemberEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEntityIds

`func (o *GroupWriteByIDRequest) SetMemberEntityIds(v []string)`

SetMemberEntityIds sets MemberEntityIds field to given value.


### HasMemberEntityIds

`func (o *GroupWriteByIDRequest) HasMemberEntityIds() bool`

HasMemberEntityIds returns a boolean if a field has been set.




### GetMemberGroupIds

`func (o *GroupWriteByIDRequest) GetMemberGroupIds() []string`

GetMemberGroupIds returns the MemberGroupIds field if non-nil, zero value otherwise.

### GetMemberGroupIdsOk

`func (o *GroupWriteByIDRequest) GetMemberGroupIdsOk() (*[]string, bool)`

GetMemberGroupIdsOk returns a tuple with the MemberGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroupIds

`func (o *GroupWriteByIDRequest) SetMemberGroupIds(v []string)`

SetMemberGroupIds sets MemberGroupIds field to given value.


### HasMemberGroupIds

`func (o *GroupWriteByIDRequest) HasMemberGroupIds() bool`

HasMemberGroupIds returns a boolean if a field has been set.




### GetMetadata

`func (o *GroupWriteByIDRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GroupWriteByIDRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GroupWriteByIDRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### HasMetadata

`func (o *GroupWriteByIDRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.




### GetName

`func (o *GroupWriteByIDRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupWriteByIDRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupWriteByIDRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *GroupWriteByIDRequest) HasName() bool`

HasName returns a boolean if a field has been set.




### GetPolicies

`func (o *GroupWriteByIDRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GroupWriteByIDRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GroupWriteByIDRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *GroupWriteByIDRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetType

`func (o *GroupWriteByIDRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupWriteByIDRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupWriteByIDRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *GroupWriteByIDRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


