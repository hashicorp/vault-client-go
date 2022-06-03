# UserpassUsersPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 

## Methods

### NewUserpassUsersPoliciesRequest

`func NewUserpassUsersPoliciesRequest() *UserpassUsersPoliciesRequest`

NewUserpassUsersPoliciesRequest instantiates a new UserpassUsersPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserpassUsersPoliciesRequestWithDefaults

`func NewUserpassUsersPoliciesRequestWithDefaults() *UserpassUsersPoliciesRequest`

NewUserpassUsersPoliciesRequestWithDefaults instantiates a new UserpassUsersPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *UserpassUsersPoliciesRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *UserpassUsersPoliciesRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *UserpassUsersPoliciesRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *UserpassUsersPoliciesRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *UserpassUsersPoliciesRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *UserpassUsersPoliciesRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *UserpassUsersPoliciesRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *UserpassUsersPoliciesRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


