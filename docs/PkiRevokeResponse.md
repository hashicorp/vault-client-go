# PkiRevokeResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevocationTime** | Pointer to **int64** | Revocation Time | [optional] 
**RevocationTimeRfc3339** | Pointer to **time.Time** | Revocation Time | [optional] 
**State** | Pointer to **string** | Revocation State | [optional] 



## Methods


### NewPkiRevokeResponse

`func NewPkiRevokeResponse() *PkiRevokeResponse`

NewPkiRevokeResponse instantiates a new PkiRevokeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiRevokeResponseWithDefaults

`func NewPkiRevokeResponseWithDefaults() *PkiRevokeResponse`

NewPkiRevokeResponseWithDefaults instantiates a new PkiRevokeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetRevocationTime

`func (o *PkiRevokeResponse) GetRevocationTime() int64`

GetRevocationTime returns the RevocationTime field if non-nil, zero value otherwise.

### GetRevocationTimeOk

`func (o *PkiRevokeResponse) GetRevocationTimeOk() (*int64, bool)`

GetRevocationTimeOk returns a tuple with the RevocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTime

`func (o *PkiRevokeResponse) SetRevocationTime(v int64)`

SetRevocationTime sets RevocationTime field to given value.


### HasRevocationTime

`func (o *PkiRevokeResponse) HasRevocationTime() bool`

HasRevocationTime returns a boolean if a field has been set.




### GetRevocationTimeRfc3339

`func (o *PkiRevokeResponse) GetRevocationTimeRfc3339() time.Time`

GetRevocationTimeRfc3339 returns the RevocationTimeRfc3339 field if non-nil, zero value otherwise.

### GetRevocationTimeRfc3339Ok

`func (o *PkiRevokeResponse) GetRevocationTimeRfc3339Ok() (*time.Time, bool)`

GetRevocationTimeRfc3339Ok returns a tuple with the RevocationTimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimeRfc3339

`func (o *PkiRevokeResponse) SetRevocationTimeRfc3339(v time.Time)`

SetRevocationTimeRfc3339 sets RevocationTimeRfc3339 field to given value.


### HasRevocationTimeRfc3339

`func (o *PkiRevokeResponse) HasRevocationTimeRfc3339() bool`

HasRevocationTimeRfc3339 returns a boolean if a field has been set.




### GetState

`func (o *PkiRevokeResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PkiRevokeResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PkiRevokeResponse) SetState(v string)`

SetState sets State field to given value.


### HasState

`func (o *PkiRevokeResponse) HasState() bool`

HasState returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


