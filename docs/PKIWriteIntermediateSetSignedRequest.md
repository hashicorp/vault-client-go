# PKIWriteIntermediateSetSignedRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Certificate** | Pointer to **string** | PEM-format certificate. This must be a CA certificate with a public key matching the previously-generated key from the generation endpoint. Additional parent CAs may be optionally appended to the bundle. | [optional] 



## Methods


### NewPKIWriteIntermediateSetSignedRequest

`func NewPKIWriteIntermediateSetSignedRequest() *PKIWriteIntermediateSetSignedRequest`

NewPKIWriteIntermediateSetSignedRequest instantiates a new PKIWriteIntermediateSetSignedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIWriteIntermediateSetSignedRequestWithDefaults

`func NewPKIWriteIntermediateSetSignedRequestWithDefaults() *PKIWriteIntermediateSetSignedRequest`

NewPKIWriteIntermediateSetSignedRequestWithDefaults instantiates a new PKIWriteIntermediateSetSignedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertificate

`func (o *PKIWriteIntermediateSetSignedRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PKIWriteIntermediateSetSignedRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PKIWriteIntermediateSetSignedRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PKIWriteIntermediateSetSignedRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


