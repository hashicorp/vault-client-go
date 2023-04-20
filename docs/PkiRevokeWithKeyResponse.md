# PkiRevokeWithKeyResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevocationTime** | Pointer to **int32** | Revocation Time | [optional] 
**RevocationTimeRfc3339** | Pointer to **time.Time** | Revocation Time | [optional] 
**State** | Pointer to **string** | Revocation State | [optional] 



## Methods


### NewPkiRevokeWithKeyResponse

`func NewPkiRevokeWithKeyResponse() *PkiRevokeWithKeyResponse`

NewPkiRevokeWithKeyResponse instantiates a new PkiRevokeWithKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiRevokeWithKeyResponseWithDefaults

`func NewPkiRevokeWithKeyResponseWithDefaults() *PkiRevokeWithKeyResponse`

NewPkiRevokeWithKeyResponseWithDefaults instantiates a new PkiRevokeWithKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetRevocationTime

`func (o *PkiRevokeWithKeyResponse) GetRevocationTime() int32`

GetRevocationTime returns the RevocationTime field if non-nil, zero value otherwise.

### GetRevocationTimeOk

`func (o *PkiRevokeWithKeyResponse) GetRevocationTimeOk() (*int32, bool)`

GetRevocationTimeOk returns a tuple with the RevocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTime

`func (o *PkiRevokeWithKeyResponse) SetRevocationTime(v int32)`

SetRevocationTime sets RevocationTime field to given value.


### HasRevocationTime

`func (o *PkiRevokeWithKeyResponse) HasRevocationTime() bool`

HasRevocationTime returns a boolean if a field has been set.




### GetRevocationTimeRfc3339

`func (o *PkiRevokeWithKeyResponse) GetRevocationTimeRfc3339() time.Time`

GetRevocationTimeRfc3339 returns the RevocationTimeRfc3339 field if non-nil, zero value otherwise.

### GetRevocationTimeRfc3339Ok

`func (o *PkiRevokeWithKeyResponse) GetRevocationTimeRfc3339Ok() (*time.Time, bool)`

GetRevocationTimeRfc3339Ok returns a tuple with the RevocationTimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimeRfc3339

`func (o *PkiRevokeWithKeyResponse) SetRevocationTimeRfc3339(v time.Time)`

SetRevocationTimeRfc3339 sets RevocationTimeRfc3339 field to given value.


### HasRevocationTimeRfc3339

`func (o *PkiRevokeWithKeyResponse) HasRevocationTimeRfc3339() bool`

HasRevocationTimeRfc3339 returns a boolean if a field has been set.




### GetState

`func (o *PkiRevokeWithKeyResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PkiRevokeWithKeyResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PkiRevokeWithKeyResponse) SetState(v string)`

SetState sets State field to given value.


### HasState

`func (o *PkiRevokeWithKeyResponse) HasState() bool`

HasState returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


