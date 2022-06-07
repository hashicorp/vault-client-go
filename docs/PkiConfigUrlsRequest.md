# PkiConfigUrlsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlDistributionPoints** | Pointer to **[]string** | Comma-separated list of URLs to be used for the CRL distribution points attribute. See also RFC 5280 Section 4.2.1.13. | [optional] 
**IssuingCertificates** | Pointer to **[]string** | Comma-separated list of URLs to be used for the issuing certificate attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 
**OcspServers** | Pointer to **[]string** | Comma-separated list of URLs to be used for the OCSP servers attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 

## Methods

### NewPkiConfigUrlsRequest

`func NewPkiConfigUrlsRequest() *PkiConfigUrlsRequest`

NewPkiConfigUrlsRequest instantiates a new PkiConfigUrlsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigUrlsRequestWithDefaults

`func NewPkiConfigUrlsRequestWithDefaults() *PkiConfigUrlsRequest`

NewPkiConfigUrlsRequestWithDefaults instantiates a new PkiConfigUrlsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrlDistributionPoints

`func (o *PkiConfigUrlsRequest) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *PkiConfigUrlsRequest) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *PkiConfigUrlsRequest) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.

### HasCrlDistributionPoints

`func (o *PkiConfigUrlsRequest) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.

### GetIssuingCertificates

`func (o *PkiConfigUrlsRequest) GetIssuingCertificates() []string`

GetIssuingCertificates returns the IssuingCertificates field if non-nil, zero value otherwise.

### GetIssuingCertificatesOk

`func (o *PkiConfigUrlsRequest) GetIssuingCertificatesOk() (*[]string, bool)`

GetIssuingCertificatesOk returns a tuple with the IssuingCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificates

`func (o *PkiConfigUrlsRequest) SetIssuingCertificates(v []string)`

SetIssuingCertificates sets IssuingCertificates field to given value.

### HasIssuingCertificates

`func (o *PkiConfigUrlsRequest) HasIssuingCertificates() bool`

HasIssuingCertificates returns a boolean if a field has been set.

### GetOcspServers

`func (o *PkiConfigUrlsRequest) GetOcspServers() []string`

GetOcspServers returns the OcspServers field if non-nil, zero value otherwise.

### GetOcspServersOk

`func (o *PkiConfigUrlsRequest) GetOcspServersOk() (*[]string, bool)`

GetOcspServersOk returns a tuple with the OcspServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServers

`func (o *PkiConfigUrlsRequest) SetOcspServers(v []string)`

SetOcspServers sets OcspServers field to given value.

### HasOcspServers

`func (o *PkiConfigUrlsRequest) HasOcspServers() bool`

HasOcspServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


