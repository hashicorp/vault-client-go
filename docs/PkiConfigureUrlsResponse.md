# PkiConfigureUrlsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlDistributionPoints** | Pointer to **[]string** | Comma-separated list of URLs to be used for the CRL distribution points attribute. See also RFC 5280 Section 4.2.1.13. | [optional] 
**EnableTemplating** | Pointer to **bool** | Whether or not to enabling templating of the above AIA fields. When templating is enabled the special values &#x27;{{issuer_id}}&#x27; and &#x27;{{cluster_path}}&#x27; are available, but the addresses are not checked for URI validity until issuance time. This requires /config/cluster&#x27;s path to be set on all PR Secondary clusters. | [optional] [default to false]
**IssuingCertificates** | Pointer to **[]string** | Comma-separated list of URLs to be used for the issuing certificate attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 
**OcspServers** | Pointer to **[]string** | Comma-separated list of URLs to be used for the OCSP servers attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


