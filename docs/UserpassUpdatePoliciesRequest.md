# UserpassUpdatePoliciesRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 



## Methods


### NewUserpassUpdatePoliciesRequest

`func NewUserpassUpdatePoliciesRequest() *UserpassUpdatePoliciesRequest`

NewUserpassUpdatePoliciesRequest instantiates a new UserpassUpdatePoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserpassUpdatePoliciesRequestWithDefaults

`func NewUserpassUpdatePoliciesRequestWithDefaults() *UserpassUpdatePoliciesRequest`

NewUserpassUpdatePoliciesRequestWithDefaults instantiates a new UserpassUpdatePoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPolicies

`func (o *UserpassUpdatePoliciesRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *UserpassUpdatePoliciesRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *UserpassUpdatePoliciesRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *UserpassUpdatePoliciesRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *UserpassUpdatePoliciesRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *UserpassUpdatePoliciesRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *UserpassUpdatePoliciesRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *UserpassUpdatePoliciesRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


