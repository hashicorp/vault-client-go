# NomadWriteLeaseConfigRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTtl** | Pointer to **int32** | Duration after which the issued token should not be allowed to be renewed | [optional] 
**Ttl** | Pointer to **int32** | Duration before which the issued token needs renewal | [optional] 



## Methods


### NewNomadWriteLeaseConfigRequest

`func NewNomadWriteLeaseConfigRequest() *NomadWriteLeaseConfigRequest`

NewNomadWriteLeaseConfigRequest instantiates a new NomadWriteLeaseConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNomadWriteLeaseConfigRequestWithDefaults

`func NewNomadWriteLeaseConfigRequestWithDefaults() *NomadWriteLeaseConfigRequest`

NewNomadWriteLeaseConfigRequestWithDefaults instantiates a new NomadWriteLeaseConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMaxTtl

`func (o *NomadWriteLeaseConfigRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *NomadWriteLeaseConfigRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *NomadWriteLeaseConfigRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *NomadWriteLeaseConfigRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetTtl

`func (o *NomadWriteLeaseConfigRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *NomadWriteLeaseConfigRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *NomadWriteLeaseConfigRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *NomadWriteLeaseConfigRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


