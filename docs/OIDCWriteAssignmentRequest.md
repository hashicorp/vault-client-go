# OIDCWriteAssignmentRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityIds** | Pointer to **[]string** | Comma separated string or array of identity entity IDs | [optional] 
**GroupIds** | Pointer to **[]string** | Comma separated string or array of identity group IDs | [optional] 



## Methods


### NewOIDCWriteAssignmentRequest

`func NewOIDCWriteAssignmentRequest() *OIDCWriteAssignmentRequest`

NewOIDCWriteAssignmentRequest instantiates a new OIDCWriteAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCWriteAssignmentRequestWithDefaults

`func NewOIDCWriteAssignmentRequestWithDefaults() *OIDCWriteAssignmentRequest`

NewOIDCWriteAssignmentRequestWithDefaults instantiates a new OIDCWriteAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEntityIds

`func (o *OIDCWriteAssignmentRequest) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *OIDCWriteAssignmentRequest) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *OIDCWriteAssignmentRequest) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.


### HasEntityIds

`func (o *OIDCWriteAssignmentRequest) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.




### GetGroupIds

`func (o *OIDCWriteAssignmentRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *OIDCWriteAssignmentRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *OIDCWriteAssignmentRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.


### HasGroupIds

`func (o *OIDCWriteAssignmentRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


