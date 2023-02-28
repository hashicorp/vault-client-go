# PKIIssuerSignSelfIssuedRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | PEM-format self-issued certificate to be signed. | [optional] 
**RequireMatchingCertificateAlgorithms** | Pointer to **bool** | If true, require the public key algorithm of the signer to match that of the self issued certificate. | [optional] [default to false]



## Methods


### NewPKIIssuerSignSelfIssuedRequest

`func NewPKIIssuerSignSelfIssuedRequest() *PKIIssuerSignSelfIssuedRequest`

NewPKIIssuerSignSelfIssuedRequest instantiates a new PKIIssuerSignSelfIssuedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIIssuerSignSelfIssuedRequestWithDefaults

`func NewPKIIssuerSignSelfIssuedRequestWithDefaults() *PKIIssuerSignSelfIssuedRequest`

NewPKIIssuerSignSelfIssuedRequestWithDefaults instantiates a new PKIIssuerSignSelfIssuedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertificate

`func (o *PKIIssuerSignSelfIssuedRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PKIIssuerSignSelfIssuedRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PKIIssuerSignSelfIssuedRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PKIIssuerSignSelfIssuedRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetRequireMatchingCertificateAlgorithms

`func (o *PKIIssuerSignSelfIssuedRequest) GetRequireMatchingCertificateAlgorithms() bool`

GetRequireMatchingCertificateAlgorithms returns the RequireMatchingCertificateAlgorithms field if non-nil, zero value otherwise.

### GetRequireMatchingCertificateAlgorithmsOk

`func (o *PKIIssuerSignSelfIssuedRequest) GetRequireMatchingCertificateAlgorithmsOk() (*bool, bool)`

GetRequireMatchingCertificateAlgorithmsOk returns a tuple with the RequireMatchingCertificateAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMatchingCertificateAlgorithms

`func (o *PKIIssuerSignSelfIssuedRequest) SetRequireMatchingCertificateAlgorithms(v bool)`

SetRequireMatchingCertificateAlgorithms sets RequireMatchingCertificateAlgorithms field to given value.


### HasRequireMatchingCertificateAlgorithms

`func (o *PKIIssuerSignSelfIssuedRequest) HasRequireMatchingCertificateAlgorithms() bool`

HasRequireMatchingCertificateAlgorithms returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


