# PkiRevokeIssuerResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | Certificate Authority Chain | [optional] 
**Certificate** | Pointer to **string** | Certificate | [optional] 
**CrlDistributionPoints** | Pointer to **[]string** | Specifies the URL values for the CRL Distribution Points field | [optional] 
**IssuerId** | Pointer to **string** | ID of the issuer | [optional] 
**IssuerName** | Pointer to **string** | Name of the issuer | [optional] 
**IssuingCertificates** | Pointer to **[]string** | Specifies the URL values for the Issuing Certificate field | [optional] 
**KeyId** | Pointer to **string** | ID of the Key | [optional] 
**LeafNotAfterBehavior** | Pointer to **string** |  | [optional] 
**ManualChain** | Pointer to **[]string** | Manual Chain | [optional] 
**OcspServers** | Pointer to **[]string** | Specifies the URL values for the OCSP Servers field | [optional] 
**RevocationSignatureAlgorithm** | Pointer to **string** | Which signature algorithm to use when building CRLs | [optional] 
**RevocationTime** | Pointer to **int64** | Time of revocation | [optional] 
**RevocationTimeRfc3339** | Pointer to **time.Time** | RFC formatted time of revocation | [optional] 
**Revoked** | Pointer to **bool** | Whether the issuer was revoked | [optional] 
**Usage** | Pointer to **string** | Allowed usage | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


