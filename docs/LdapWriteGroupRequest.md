# LdapWriteGroupRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to **[]string** | Comma-separated list of policies associated to the group. | [optional] 



## Methods


### NewLdapWriteGroupRequest

`func NewLdapWriteGroupRequest() *LdapWriteGroupRequest`

NewLdapWriteGroupRequest instantiates a new LdapWriteGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapWriteGroupRequestWithDefaults

`func NewLdapWriteGroupRequestWithDefaults() *LdapWriteGroupRequest`

NewLdapWriteGroupRequestWithDefaults instantiates a new LdapWriteGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPolicies

`func (o *LdapWriteGroupRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *LdapWriteGroupRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *LdapWriteGroupRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *LdapWriteGroupRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


