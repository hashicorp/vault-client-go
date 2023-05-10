# OidcWriteAssignmentRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityIds** | Pointer to **[]string** | Comma separated string or array of identity entity IDs | [optional] 
**GroupIds** | Pointer to **[]string** | Comma separated string or array of identity group IDs | [optional] 



## Methods


### NewOidcWriteAssignmentRequest

`func NewOidcWriteAssignmentRequest() *OidcWriteAssignmentRequest`

NewOidcWriteAssignmentRequest instantiates a new OidcWriteAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcWriteAssignmentRequestWithDefaults

`func NewOidcWriteAssignmentRequestWithDefaults() *OidcWriteAssignmentRequest`

NewOidcWriteAssignmentRequestWithDefaults instantiates a new OidcWriteAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEntityIds

`func (o *OidcWriteAssignmentRequest) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *OidcWriteAssignmentRequest) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *OidcWriteAssignmentRequest) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.


### HasEntityIds

`func (o *OidcWriteAssignmentRequest) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.




### GetGroupIds

`func (o *OidcWriteAssignmentRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *OidcWriteAssignmentRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *OidcWriteAssignmentRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.


### HasGroupIds

`func (o *OidcWriteAssignmentRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


