# OpenldapLibraryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableCheckInEnforcement** | Pointer to **bool** | Disable the default behavior of requiring that check-ins are performed by the entity that checked them out. | [optional] [default to false]
**MaxTtl** | Pointer to **int32** | In seconds, the max amount of time a check-out&#39;s renewals should last. Defaults to 24 hours. | [optional] [default to 86400]
**ServiceAccountNames** | Pointer to **[]string** | The username/logon name for the service accounts with which this set will be associated. | [optional] 
**Ttl** | Pointer to **int32** | In seconds, the amount of time a check-out should last. Defaults to 24 hours. | [optional] [default to 86400]

## Methods

### NewOpenldapLibraryRequest

`func NewOpenldapLibraryRequest() *OpenldapLibraryRequest`

NewOpenldapLibraryRequest instantiates a new OpenldapLibraryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenldapLibraryRequestWithDefaults

`func NewOpenldapLibraryRequestWithDefaults() *OpenldapLibraryRequest`

NewOpenldapLibraryRequestWithDefaults instantiates a new OpenldapLibraryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableCheckInEnforcement

`func (o *OpenldapLibraryRequest) GetDisableCheckInEnforcement() bool`

GetDisableCheckInEnforcement returns the DisableCheckInEnforcement field if non-nil, zero value otherwise.

### GetDisableCheckInEnforcementOk

`func (o *OpenldapLibraryRequest) GetDisableCheckInEnforcementOk() (*bool, bool)`

GetDisableCheckInEnforcementOk returns a tuple with the DisableCheckInEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCheckInEnforcement

`func (o *OpenldapLibraryRequest) SetDisableCheckInEnforcement(v bool)`

SetDisableCheckInEnforcement sets DisableCheckInEnforcement field to given value.

### HasDisableCheckInEnforcement

`func (o *OpenldapLibraryRequest) HasDisableCheckInEnforcement() bool`

HasDisableCheckInEnforcement returns a boolean if a field has been set.

### GetMaxTtl

`func (o *OpenldapLibraryRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *OpenldapLibraryRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *OpenldapLibraryRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *OpenldapLibraryRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetServiceAccountNames

`func (o *OpenldapLibraryRequest) GetServiceAccountNames() []string`

GetServiceAccountNames returns the ServiceAccountNames field if non-nil, zero value otherwise.

### GetServiceAccountNamesOk

`func (o *OpenldapLibraryRequest) GetServiceAccountNamesOk() (*[]string, bool)`

GetServiceAccountNamesOk returns a tuple with the ServiceAccountNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountNames

`func (o *OpenldapLibraryRequest) SetServiceAccountNames(v []string)`

SetServiceAccountNames sets ServiceAccountNames field to given value.

### HasServiceAccountNames

`func (o *OpenldapLibraryRequest) HasServiceAccountNames() bool`

HasServiceAccountNames returns a boolean if a field has been set.

### GetTtl

`func (o *OpenldapLibraryRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *OpenldapLibraryRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *OpenldapLibraryRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *OpenldapLibraryRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


