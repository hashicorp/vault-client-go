# GoogleCloudWriteRoleServiceAccountsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to **[]string** | Service-account emails or IDs to add. | [optional] 
**Remove** | Pointer to **[]string** | Service-account emails or IDs to remove. | [optional] 



## Methods


### NewGoogleCloudWriteRoleServiceAccountsRequest

`func NewGoogleCloudWriteRoleServiceAccountsRequest() *GoogleCloudWriteRoleServiceAccountsRequest`

NewGoogleCloudWriteRoleServiceAccountsRequest instantiates a new GoogleCloudWriteRoleServiceAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudWriteRoleServiceAccountsRequestWithDefaults

`func NewGoogleCloudWriteRoleServiceAccountsRequestWithDefaults() *GoogleCloudWriteRoleServiceAccountsRequest`

NewGoogleCloudWriteRoleServiceAccountsRequestWithDefaults instantiates a new GoogleCloudWriteRoleServiceAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAdd

`func (o *GoogleCloudWriteRoleServiceAccountsRequest) GetAdd() []string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *GoogleCloudWriteRoleServiceAccountsRequest) GetAddOk() (*[]string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *GoogleCloudWriteRoleServiceAccountsRequest) SetAdd(v []string)`

SetAdd sets Add field to given value.


### HasAdd

`func (o *GoogleCloudWriteRoleServiceAccountsRequest) HasAdd() bool`

HasAdd returns a boolean if a field has been set.




### GetRemove

`func (o *GoogleCloudWriteRoleServiceAccountsRequest) GetRemove() []string`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *GoogleCloudWriteRoleServiceAccountsRequest) GetRemoveOk() (*[]string, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *GoogleCloudWriteRoleServiceAccountsRequest) SetRemove(v []string)`

SetRemove sets Remove field to given value.


### HasRemove

`func (o *GoogleCloudWriteRoleServiceAccountsRequest) HasRemove() bool`

HasRemove returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


