# LdapWriteUserRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **[]string** | Comma-separated list of additional groups associated with the user. | [optional] 
**Policies** | Pointer to **[]string** | Comma-separated list of policies associated with the user. | [optional] 



## Methods


### NewLdapWriteUserRequest

`func NewLdapWriteUserRequest() *LdapWriteUserRequest`

NewLdapWriteUserRequest instantiates a new LdapWriteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapWriteUserRequestWithDefaults

`func NewLdapWriteUserRequestWithDefaults() *LdapWriteUserRequest`

NewLdapWriteUserRequestWithDefaults instantiates a new LdapWriteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetGroups

`func (o *LdapWriteUserRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *LdapWriteUserRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *LdapWriteUserRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### HasGroups

`func (o *LdapWriteUserRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.




### GetPolicies

`func (o *LdapWriteUserRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *LdapWriteUserRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *LdapWriteUserRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *LdapWriteUserRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


