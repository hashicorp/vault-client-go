# EntityWriteByNameRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Disabled** | Pointer to **bool** | If set true, tokens tied to this identity will not be able to be used (but will not be revoked). | [optional] 
**Id** | Pointer to **string** | ID of the entity. If set, updates the corresponding existing entity. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to be associated with the entity. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault &lt;command&gt; &lt;path&gt; metadata&#x3D;key1&#x3D;value1 metadata&#x3D;key2&#x3D;value2 | [optional] 
**Policies** | Pointer to **[]string** | Policies to be tied to the entity. | [optional] 



## Methods


### NewEntityWriteByNameRequest

`func NewEntityWriteByNameRequest() *EntityWriteByNameRequest`

NewEntityWriteByNameRequest instantiates a new EntityWriteByNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWriteByNameRequestWithDefaults

`func NewEntityWriteByNameRequestWithDefaults() *EntityWriteByNameRequest`

NewEntityWriteByNameRequestWithDefaults instantiates a new EntityWriteByNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDisabled

`func (o *EntityWriteByNameRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *EntityWriteByNameRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *EntityWriteByNameRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### HasDisabled

`func (o *EntityWriteByNameRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.




### GetId

`func (o *EntityWriteByNameRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityWriteByNameRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityWriteByNameRequest) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *EntityWriteByNameRequest) HasId() bool`

HasId returns a boolean if a field has been set.




### GetMetadata

`func (o *EntityWriteByNameRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntityWriteByNameRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntityWriteByNameRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### HasMetadata

`func (o *EntityWriteByNameRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.




### GetPolicies

`func (o *EntityWriteByNameRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EntityWriteByNameRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EntityWriteByNameRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *EntityWriteByNameRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


