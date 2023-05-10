# EntityUpdateByIdRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | If set true, tokens tied to this identity will not be able to be used (but will not be revoked). | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to be associated with the entity. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault &lt;command&gt; &lt;path&gt; metadata&#x3D;key1&#x3D;value1 metadata&#x3D;key2&#x3D;value2 | [optional] 
**Name** | Pointer to **string** | Name of the entity | [optional] 
**Policies** | Pointer to **[]string** | Policies to be tied to the entity. | [optional] 



## Methods


### NewEntityUpdateByIdRequest

`func NewEntityUpdateByIdRequest() *EntityUpdateByIdRequest`

NewEntityUpdateByIdRequest instantiates a new EntityUpdateByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityUpdateByIdRequestWithDefaults

`func NewEntityUpdateByIdRequestWithDefaults() *EntityUpdateByIdRequest`

NewEntityUpdateByIdRequestWithDefaults instantiates a new EntityUpdateByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDisabled

`func (o *EntityUpdateByIdRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *EntityUpdateByIdRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *EntityUpdateByIdRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### HasDisabled

`func (o *EntityUpdateByIdRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.




### GetMetadata

`func (o *EntityUpdateByIdRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntityUpdateByIdRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntityUpdateByIdRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### HasMetadata

`func (o *EntityUpdateByIdRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.




### GetName

`func (o *EntityUpdateByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityUpdateByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityUpdateByIdRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *EntityUpdateByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.




### GetPolicies

`func (o *EntityUpdateByIdRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EntityUpdateByIdRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EntityUpdateByIdRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *EntityUpdateByIdRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


