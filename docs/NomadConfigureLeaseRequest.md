# NomadConfigureLeaseRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTtl** | Pointer to **string** | Duration after which the issued token should not be allowed to be renewed | [optional] 
**Ttl** | Pointer to **string** | Duration before which the issued token needs renewal | [optional] 



## Methods


### NewNomadConfigureLeaseRequest

`func NewNomadConfigureLeaseRequest() *NomadConfigureLeaseRequest`

NewNomadConfigureLeaseRequest instantiates a new NomadConfigureLeaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNomadConfigureLeaseRequestWithDefaults

`func NewNomadConfigureLeaseRequestWithDefaults() *NomadConfigureLeaseRequest`

NewNomadConfigureLeaseRequestWithDefaults instantiates a new NomadConfigureLeaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMaxTtl

`func (o *NomadConfigureLeaseRequest) GetMaxTtl() string`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *NomadConfigureLeaseRequest) GetMaxTtlOk() (*string, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *NomadConfigureLeaseRequest) SetMaxTtl(v string)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *NomadConfigureLeaseRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetTtl

`func (o *NomadConfigureLeaseRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *NomadConfigureLeaseRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *NomadConfigureLeaseRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *NomadConfigureLeaseRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


