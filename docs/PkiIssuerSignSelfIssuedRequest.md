# PkiIssuerSignSelfIssuedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | PEM-format self-issued certificate to be signed. | [optional] 
**RequireMatchingCertificateAlgorithms** | Pointer to **bool** | If true, require the public key algorithm of the signer to match that of the self issued certificate. | [optional] [default to false]

## Methods

### NewPkiIssuerSignSelfIssuedRequest

`func NewPkiIssuerSignSelfIssuedRequest() *PkiIssuerSignSelfIssuedRequest`

NewPkiIssuerSignSelfIssuedRequest instantiates a new PkiIssuerSignSelfIssuedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiIssuerSignSelfIssuedRequestWithDefaults

`func NewPkiIssuerSignSelfIssuedRequestWithDefaults() *PkiIssuerSignSelfIssuedRequest`

NewPkiIssuerSignSelfIssuedRequestWithDefaults instantiates a new PkiIssuerSignSelfIssuedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *PkiIssuerSignSelfIssuedRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiIssuerSignSelfIssuedRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiIssuerSignSelfIssuedRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PkiIssuerSignSelfIssuedRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetRequireMatchingCertificateAlgorithms

`func (o *PkiIssuerSignSelfIssuedRequest) GetRequireMatchingCertificateAlgorithms() bool`

GetRequireMatchingCertificateAlgorithms returns the RequireMatchingCertificateAlgorithms field if non-nil, zero value otherwise.

### GetRequireMatchingCertificateAlgorithmsOk

`func (o *PkiIssuerSignSelfIssuedRequest) GetRequireMatchingCertificateAlgorithmsOk() (*bool, bool)`

GetRequireMatchingCertificateAlgorithmsOk returns a tuple with the RequireMatchingCertificateAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMatchingCertificateAlgorithms

`func (o *PkiIssuerSignSelfIssuedRequest) SetRequireMatchingCertificateAlgorithms(v bool)`

SetRequireMatchingCertificateAlgorithms sets RequireMatchingCertificateAlgorithms field to given value.

### HasRequireMatchingCertificateAlgorithms

`func (o *PkiIssuerSignSelfIssuedRequest) HasRequireMatchingCertificateAlgorithms() bool`

HasRequireMatchingCertificateAlgorithms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


