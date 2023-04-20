# PkiReadUrlsConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlDistributionPoints** | Pointer to **[]string** | Comma-separated list of URLs to be used for the CRL distribution points attribute. See also RFC 5280 Section 4.2.1.13. | [optional] 
**EnableTemplating** | Pointer to **bool** | Whether or not to enable templating of the above AIA fields. When templating is enabled the special values &#x27;{{issuer_id}}&#x27; and &#x27;{{cluster_path}}&#x27; are available, but the addresses are not checked for URI validity until issuance time. This requires /config/cluster&#x27;s path to be set on all PR Secondary clusters. | [optional] 
**IssuingCertificates** | Pointer to **[]string** | Comma-separated list of URLs to be used for the issuing certificate attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 
**OcspServers** | Pointer to **[]string** | Comma-separated list of URLs to be used for the OCSP servers attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 



## Methods


### NewPkiReadUrlsConfigurationResponse

`func NewPkiReadUrlsConfigurationResponse() *PkiReadUrlsConfigurationResponse`

NewPkiReadUrlsConfigurationResponse instantiates a new PkiReadUrlsConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiReadUrlsConfigurationResponseWithDefaults

`func NewPkiReadUrlsConfigurationResponseWithDefaults() *PkiReadUrlsConfigurationResponse`

NewPkiReadUrlsConfigurationResponseWithDefaults instantiates a new PkiReadUrlsConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCrlDistributionPoints

`func (o *PkiReadUrlsConfigurationResponse) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *PkiReadUrlsConfigurationResponse) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *PkiReadUrlsConfigurationResponse) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.


### HasCrlDistributionPoints

`func (o *PkiReadUrlsConfigurationResponse) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.




### GetEnableTemplating

`func (o *PkiReadUrlsConfigurationResponse) GetEnableTemplating() bool`

GetEnableTemplating returns the EnableTemplating field if non-nil, zero value otherwise.

### GetEnableTemplatingOk

`func (o *PkiReadUrlsConfigurationResponse) GetEnableTemplatingOk() (*bool, bool)`

GetEnableTemplatingOk returns a tuple with the EnableTemplating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTemplating

`func (o *PkiReadUrlsConfigurationResponse) SetEnableTemplating(v bool)`

SetEnableTemplating sets EnableTemplating field to given value.


### HasEnableTemplating

`func (o *PkiReadUrlsConfigurationResponse) HasEnableTemplating() bool`

HasEnableTemplating returns a boolean if a field has been set.




### GetIssuingCertificates

`func (o *PkiReadUrlsConfigurationResponse) GetIssuingCertificates() []string`

GetIssuingCertificates returns the IssuingCertificates field if non-nil, zero value otherwise.

### GetIssuingCertificatesOk

`func (o *PkiReadUrlsConfigurationResponse) GetIssuingCertificatesOk() (*[]string, bool)`

GetIssuingCertificatesOk returns a tuple with the IssuingCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificates

`func (o *PkiReadUrlsConfigurationResponse) SetIssuingCertificates(v []string)`

SetIssuingCertificates sets IssuingCertificates field to given value.


### HasIssuingCertificates

`func (o *PkiReadUrlsConfigurationResponse) HasIssuingCertificates() bool`

HasIssuingCertificates returns a boolean if a field has been set.




### GetOcspServers

`func (o *PkiReadUrlsConfigurationResponse) GetOcspServers() []string`

GetOcspServers returns the OcspServers field if non-nil, zero value otherwise.

### GetOcspServersOk

`func (o *PkiReadUrlsConfigurationResponse) GetOcspServersOk() (*[]string, bool)`

GetOcspServersOk returns a tuple with the OcspServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServers

`func (o *PkiReadUrlsConfigurationResponse) SetOcspServers(v []string)`

SetOcspServers sets OcspServers field to given value.


### HasOcspServers

`func (o *PkiReadUrlsConfigurationResponse) HasOcspServers() bool`

HasOcspServers returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


