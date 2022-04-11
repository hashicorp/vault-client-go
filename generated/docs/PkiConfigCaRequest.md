# PkiConfigCaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PemBundle** | Pointer to **string** | PEM-format, concatenated unencrypted secret key and certificate. | [optional] 

## Methods

### NewPkiConfigCaRequest

`func NewPkiConfigCaRequest() *PkiConfigCaRequest`

NewPkiConfigCaRequest instantiates a new PkiConfigCaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigCaRequestWithDefaults

`func NewPkiConfigCaRequestWithDefaults() *PkiConfigCaRequest`

NewPkiConfigCaRequestWithDefaults instantiates a new PkiConfigCaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPemBundle

`func (o *PkiConfigCaRequest) GetPemBundle() string`

GetPemBundle returns the PemBundle field if non-nil, zero value otherwise.

### GetPemBundleOk

`func (o *PkiConfigCaRequest) GetPemBundleOk() (*string, bool)`

GetPemBundleOk returns a tuple with the PemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemBundle

`func (o *PkiConfigCaRequest) SetPemBundle(v string)`

SetPemBundle sets PemBundle field to given value.

### HasPemBundle

`func (o *PkiConfigCaRequest) HasPemBundle() bool`

HasPemBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


