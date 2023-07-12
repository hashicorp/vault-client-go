# CertWriteCrlRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crl** | Pointer to **string** | The public CRL that should be trusted to attest to certificates&#x27; validity statuses. May be DER or PEM encoded. Note: the expiration time is ignored; if the CRL is no longer valid, delete it using the same name as specified here. | [optional] 
**Url** | Pointer to **string** | The URL of a CRL distribution point. Only one of &#x27;crl&#x27; or &#x27;url&#x27; parameters should be specified. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


