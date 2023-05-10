# PkiConfigureCaRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PemBundle** | Pointer to **string** | PEM-format, concatenated unencrypted secret key and certificate. | [optional] 



## Methods


### NewPkiConfigureCaRequest

`func NewPkiConfigureCaRequest() *PkiConfigureCaRequest`

NewPkiConfigureCaRequest instantiates a new PkiConfigureCaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigureCaRequestWithDefaults

`func NewPkiConfigureCaRequestWithDefaults() *PkiConfigureCaRequest`

NewPkiConfigureCaRequestWithDefaults instantiates a new PkiConfigureCaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPemBundle

`func (o *PkiConfigureCaRequest) GetPemBundle() string`

GetPemBundle returns the PemBundle field if non-nil, zero value otherwise.

### GetPemBundleOk

`func (o *PkiConfigureCaRequest) GetPemBundleOk() (*string, bool)`

GetPemBundleOk returns a tuple with the PemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemBundle

`func (o *PkiConfigureCaRequest) SetPemBundle(v string)`

SetPemBundle sets PemBundle field to given value.


### HasPemBundle

`func (o *PkiConfigureCaRequest) HasPemBundle() bool`

HasPemBundle returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


