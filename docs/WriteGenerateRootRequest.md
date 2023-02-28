# WriteGenerateRootRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PgpKey** | Pointer to **string** | Specifies a base64-encoded PGP public key. | [optional] 



## Methods


### NewWriteGenerateRootRequest

`func NewWriteGenerateRootRequest() *WriteGenerateRootRequest`

NewWriteGenerateRootRequest instantiates a new WriteGenerateRootRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteGenerateRootRequestWithDefaults

`func NewWriteGenerateRootRequestWithDefaults() *WriteGenerateRootRequest`

NewWriteGenerateRootRequestWithDefaults instantiates a new WriteGenerateRootRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPgpKey

`func (o *WriteGenerateRootRequest) GetPgpKey() string`

GetPgpKey returns the PgpKey field if non-nil, zero value otherwise.

### GetPgpKeyOk

`func (o *WriteGenerateRootRequest) GetPgpKeyOk() (*string, bool)`

GetPgpKeyOk returns a tuple with the PgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpKey

`func (o *WriteGenerateRootRequest) SetPgpKey(v string)`

SetPgpKey sets PgpKey field to given value.


### HasPgpKey

`func (o *WriteGenerateRootRequest) HasPgpKey() bool`

HasPgpKey returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


