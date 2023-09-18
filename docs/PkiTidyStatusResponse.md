# PkiTidyStatusResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeAccountDeletedCount** | Pointer to **int32** | The number of revoked acme accounts removed | [optional] 
**AcmeAccountRevokedCount** | Pointer to **int32** | The number of unused acme accounts revoked | [optional] 
**AcmeAccountSafetyBuffer** | Pointer to **int32** | Safety buffer after creation after which accounts lacking orders are revoked | [optional] 
**AcmeOrdersDeletedCount** | Pointer to **int32** | The number of expired, unused acme orders removed | [optional] 
**CertStoreDeletedCount** | Pointer to **int32** | The number of certificate storage entries deleted | [optional] 
**CrossRevokedCertDeletedCount** | Pointer to **int32** |  | [optional] 
**CurrentCertStoreCount** | Pointer to **int32** | The number of revoked certificate entries deleted | [optional] 
**CurrentRevokedCertCount** | Pointer to **int32** | The number of revoked certificate entries deleted | [optional] 
**Error** | Pointer to **string** | The error message | [optional] 
**InternalBackendUuid** | Pointer to **string** |  | [optional] 
**IssuerSafetyBuffer** | Pointer to **int32** | Issuer safety buffer | [optional] 
**LastAutoTidyFinished** | Pointer to **string** | Time the last auto-tidy operation finished | [optional] 
**Message** | Pointer to **string** | Message of the operation | [optional] 
**MissingIssuerCertCount** | Pointer to **int32** |  | [optional] 
**PauseDuration** | Pointer to **string** | Duration to pause between tidying certificates | [optional] 
**RevocationQueueDeletedCount** | Pointer to **int32** |  | [optional] 
**RevocationQueueSafetyBuffer** | Pointer to **int32** | Revocation queue safety buffer | [optional] 
**RevokedCertDeletedCount** | Pointer to **int32** | The number of revoked certificate entries deleted | [optional] 
**SafetyBuffer** | Pointer to **int32** | Safety buffer time duration | [optional] 
**State** | Pointer to **string** | One of Inactive, Running, Finished, or Error | [optional] 
**TidyAcme** | Pointer to **bool** | Tidy Unused Acme Accounts, and Orders | [optional] 
**TidyCertStore** | Pointer to **bool** | Tidy certificate store | [optional] 
**TidyCrossClusterRevokedCerts** | Pointer to **bool** | Tidy the cross-cluster revoked certificate store | [optional] 
**TidyExpiredIssuers** | Pointer to **bool** | Tidy expired issuers | [optional] 
**TidyMoveLegacyCaBundle** | Pointer to **bool** |  | [optional] 
**TidyRevocationQueue** | Pointer to **bool** |  | [optional] 
**TidyRevokedCertIssuerAssociations** | Pointer to **bool** | Tidy revoked certificate issuer associations | [optional] 
**TidyRevokedCerts** | Pointer to **bool** | Tidy revoked certificates | [optional] 
**TimeFinished** | Pointer to **string** | Time the operation finished | [optional] 
**TimeStarted** | Pointer to **string** | Time the operation started | [optional] 
**TotalAcmeAccountCount** | Pointer to **int32** | Total number of acme accounts iterated over | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


