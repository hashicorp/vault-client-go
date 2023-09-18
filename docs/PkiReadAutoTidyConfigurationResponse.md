# PkiReadAutoTidyConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeAccountSafetyBuffer** | Pointer to **int32** | Safety buffer after creation after which accounts lacking orders are revoked | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether automatic tidy is enabled or not | [optional] 
**IntervalDuration** | Pointer to **int32** | Specifies the duration between automatic tidy operation | [optional] 
**IssuerSafetyBuffer** | Pointer to **int32** | Issuer safety buffer | [optional] 
**MaintainStoredCertificateCounts** | Pointer to **bool** |  | [optional] 
**PauseDuration** | Pointer to **string** | Duration to pause between tidying certificates | [optional] 
**PublishStoredCertificateCountMetrics** | Pointer to **bool** |  | [optional] 
**RevocationQueueSafetyBuffer** | Pointer to **int32** |  | [optional] 
**SafetyBuffer** | Pointer to **int32** | Safety buffer time duration | [optional] 
**TidyAcme** | Pointer to **bool** | Tidy Unused Acme Accounts, and Orders | [optional] 
**TidyCertStore** | Pointer to **bool** | Specifies whether to tidy up the certificate store | [optional] 
**TidyCrossClusterRevokedCerts** | Pointer to **bool** |  | [optional] 
**TidyExpiredIssuers** | Pointer to **bool** | Specifies whether tidy expired issuers | [optional] 
**TidyMoveLegacyCaBundle** | Pointer to **bool** |  | [optional] 
**TidyRevocationQueue** | Pointer to **bool** |  | [optional] 
**TidyRevokedCertIssuerAssociations** | Pointer to **bool** | Specifies whether to associate revoked certificates with their corresponding issuers | [optional] 
**TidyRevokedCerts** | Pointer to **bool** | Specifies whether to remove all invalid and expired certificates from storage | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


