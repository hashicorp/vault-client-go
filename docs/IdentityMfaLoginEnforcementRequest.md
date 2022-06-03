# IdentityMfaLoginEnforcementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodAccessors** | Pointer to **[]string** | Array of auth mount accessor IDs | [optional] 
**AuthMethodTypes** | Pointer to **[]string** | Array of auth mount types | [optional] 
**IdentityEntityIds** | Pointer to **[]string** | Array of identity entity IDs | [optional] 
**IdentityGroupIds** | Pointer to **[]string** | Array of identity group IDs | [optional] 
**MfaMethodIds** | **[]string** | Array of Method IDs that determine what methods will be enforced | 

## Methods

### NewIdentityMfaLoginEnforcementRequest

`func NewIdentityMfaLoginEnforcementRequest(mfaMethodIds []string, ) *IdentityMfaLoginEnforcementRequest`

NewIdentityMfaLoginEnforcementRequest instantiates a new IdentityMfaLoginEnforcementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityMfaLoginEnforcementRequestWithDefaults

`func NewIdentityMfaLoginEnforcementRequestWithDefaults() *IdentityMfaLoginEnforcementRequest`

NewIdentityMfaLoginEnforcementRequestWithDefaults instantiates a new IdentityMfaLoginEnforcementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodAccessors

`func (o *IdentityMfaLoginEnforcementRequest) GetAuthMethodAccessors() []string`

GetAuthMethodAccessors returns the AuthMethodAccessors field if non-nil, zero value otherwise.

### GetAuthMethodAccessorsOk

`func (o *IdentityMfaLoginEnforcementRequest) GetAuthMethodAccessorsOk() (*[]string, bool)`

GetAuthMethodAccessorsOk returns a tuple with the AuthMethodAccessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodAccessors

`func (o *IdentityMfaLoginEnforcementRequest) SetAuthMethodAccessors(v []string)`

SetAuthMethodAccessors sets AuthMethodAccessors field to given value.

### HasAuthMethodAccessors

`func (o *IdentityMfaLoginEnforcementRequest) HasAuthMethodAccessors() bool`

HasAuthMethodAccessors returns a boolean if a field has been set.

### GetAuthMethodTypes

`func (o *IdentityMfaLoginEnforcementRequest) GetAuthMethodTypes() []string`

GetAuthMethodTypes returns the AuthMethodTypes field if non-nil, zero value otherwise.

### GetAuthMethodTypesOk

`func (o *IdentityMfaLoginEnforcementRequest) GetAuthMethodTypesOk() (*[]string, bool)`

GetAuthMethodTypesOk returns a tuple with the AuthMethodTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodTypes

`func (o *IdentityMfaLoginEnforcementRequest) SetAuthMethodTypes(v []string)`

SetAuthMethodTypes sets AuthMethodTypes field to given value.

### HasAuthMethodTypes

`func (o *IdentityMfaLoginEnforcementRequest) HasAuthMethodTypes() bool`

HasAuthMethodTypes returns a boolean if a field has been set.

### GetIdentityEntityIds

`func (o *IdentityMfaLoginEnforcementRequest) GetIdentityEntityIds() []string`

GetIdentityEntityIds returns the IdentityEntityIds field if non-nil, zero value otherwise.

### GetIdentityEntityIdsOk

`func (o *IdentityMfaLoginEnforcementRequest) GetIdentityEntityIdsOk() (*[]string, bool)`

GetIdentityEntityIdsOk returns a tuple with the IdentityEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityEntityIds

`func (o *IdentityMfaLoginEnforcementRequest) SetIdentityEntityIds(v []string)`

SetIdentityEntityIds sets IdentityEntityIds field to given value.

### HasIdentityEntityIds

`func (o *IdentityMfaLoginEnforcementRequest) HasIdentityEntityIds() bool`

HasIdentityEntityIds returns a boolean if a field has been set.

### GetIdentityGroupIds

`func (o *IdentityMfaLoginEnforcementRequest) GetIdentityGroupIds() []string`

GetIdentityGroupIds returns the IdentityGroupIds field if non-nil, zero value otherwise.

### GetIdentityGroupIdsOk

`func (o *IdentityMfaLoginEnforcementRequest) GetIdentityGroupIdsOk() (*[]string, bool)`

GetIdentityGroupIdsOk returns a tuple with the IdentityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityGroupIds

`func (o *IdentityMfaLoginEnforcementRequest) SetIdentityGroupIds(v []string)`

SetIdentityGroupIds sets IdentityGroupIds field to given value.

### HasIdentityGroupIds

`func (o *IdentityMfaLoginEnforcementRequest) HasIdentityGroupIds() bool`

HasIdentityGroupIds returns a boolean if a field has been set.

### GetMfaMethodIds

`func (o *IdentityMfaLoginEnforcementRequest) GetMfaMethodIds() []string`

GetMfaMethodIds returns the MfaMethodIds field if non-nil, zero value otherwise.

### GetMfaMethodIdsOk

`func (o *IdentityMfaLoginEnforcementRequest) GetMfaMethodIdsOk() (*[]string, bool)`

GetMfaMethodIdsOk returns a tuple with the MfaMethodIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaMethodIds

`func (o *IdentityMfaLoginEnforcementRequest) SetMfaMethodIds(v []string)`

SetMfaMethodIds sets MfaMethodIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


