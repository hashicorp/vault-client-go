# GoogleCloudEditLabelsForRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to **[]string** | BoundLabels to add (in $key:$value) | [optional] 
**Remove** | Pointer to **[]string** | Label key values to remove | [optional] 



## Methods


### NewGoogleCloudEditLabelsForRoleRequest

`func NewGoogleCloudEditLabelsForRoleRequest() *GoogleCloudEditLabelsForRoleRequest`

NewGoogleCloudEditLabelsForRoleRequest instantiates a new GoogleCloudEditLabelsForRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudEditLabelsForRoleRequestWithDefaults

`func NewGoogleCloudEditLabelsForRoleRequestWithDefaults() *GoogleCloudEditLabelsForRoleRequest`

NewGoogleCloudEditLabelsForRoleRequestWithDefaults instantiates a new GoogleCloudEditLabelsForRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAdd

`func (o *GoogleCloudEditLabelsForRoleRequest) GetAdd() []string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *GoogleCloudEditLabelsForRoleRequest) GetAddOk() (*[]string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *GoogleCloudEditLabelsForRoleRequest) SetAdd(v []string)`

SetAdd sets Add field to given value.


### HasAdd

`func (o *GoogleCloudEditLabelsForRoleRequest) HasAdd() bool`

HasAdd returns a boolean if a field has been set.




### GetRemove

`func (o *GoogleCloudEditLabelsForRoleRequest) GetRemove() []string`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *GoogleCloudEditLabelsForRoleRequest) GetRemoveOk() (*[]string, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *GoogleCloudEditLabelsForRoleRequest) SetRemove(v []string)`

SetRemove sets Remove field to given value.


### HasRemove

`func (o *GoogleCloudEditLabelsForRoleRequest) HasRemove() bool`

HasRemove returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


