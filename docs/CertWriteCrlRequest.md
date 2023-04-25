# CertWriteCrlRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crl** | Pointer to **string** | The public CRL that should be trusted to attest to certificates&#x27; validity statuses. May be DER or PEM encoded. Note: the expiration time is ignored; if the CRL is no longer valid, delete it using the same name as specified here. | [optional] 
**Url** | Pointer to **string** | The URL of a CRL distribution point. Only one of &#x27;crl&#x27; or &#x27;url&#x27; parameters should be specified. | [optional] 



## Methods


### NewCertWriteCrlRequest

`func NewCertWriteCrlRequest() *CertWriteCrlRequest`

NewCertWriteCrlRequest instantiates a new CertWriteCrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertWriteCrlRequestWithDefaults

`func NewCertWriteCrlRequestWithDefaults() *CertWriteCrlRequest`

NewCertWriteCrlRequestWithDefaults instantiates a new CertWriteCrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCrl

`func (o *CertWriteCrlRequest) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *CertWriteCrlRequest) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *CertWriteCrlRequest) SetCrl(v string)`

SetCrl sets Crl field to given value.


### HasCrl

`func (o *CertWriteCrlRequest) HasCrl() bool`

HasCrl returns a boolean if a field has been set.




### GetUrl

`func (o *CertWriteCrlRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CertWriteCrlRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CertWriteCrlRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### HasUrl

`func (o *CertWriteCrlRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


