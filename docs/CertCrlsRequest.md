# CertCrlsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crl** | Pointer to **string** | The public certificate that should be trusted. May be DER or PEM encoded. Note: the expiration time is ignored; if the CRL is no longer valid, delete it using the same name as specified here. | [optional] 

## Methods

### NewCertCrlsRequest

`func NewCertCrlsRequest() *CertCrlsRequest`

NewCertCrlsRequest instantiates a new CertCrlsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertCrlsRequestWithDefaults

`func NewCertCrlsRequestWithDefaults() *CertCrlsRequest`

NewCertCrlsRequestWithDefaults instantiates a new CertCrlsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrl

`func (o *CertCrlsRequest) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *CertCrlsRequest) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *CertCrlsRequest) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *CertCrlsRequest) HasCrl() bool`

HasCrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


