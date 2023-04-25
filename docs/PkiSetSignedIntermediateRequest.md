# PkiSetSignedIntermediateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | PEM-format certificate. This must be a CA certificate with a public key matching the previously-generated key from the generation endpoint. Additional parent CAs may be optionally appended to the bundle. | [optional] 



## Methods


### NewPkiSetSignedIntermediateRequest

`func NewPkiSetSignedIntermediateRequest() *PkiSetSignedIntermediateRequest`

NewPkiSetSignedIntermediateRequest instantiates a new PkiSetSignedIntermediateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiSetSignedIntermediateRequestWithDefaults

`func NewPkiSetSignedIntermediateRequestWithDefaults() *PkiSetSignedIntermediateRequest`

NewPkiSetSignedIntermediateRequestWithDefaults instantiates a new PkiSetSignedIntermediateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertificate

`func (o *PkiSetSignedIntermediateRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiSetSignedIntermediateRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiSetSignedIntermediateRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiSetSignedIntermediateRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


