# PkiConfigureCrlRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRebuild** | Pointer to **bool** | If set to true, enables automatic rebuilding of the CRL | [optional] 
**AutoRebuildGracePeriod** | Pointer to **string** | The time before the CRL expires to automatically rebuild it, when enabled. Must be shorter than the CRL expiry. Defaults to 12h. | [optional] [default to "12h"]
**CrossClusterRevocation** | Pointer to **bool** | Whether to enable a global, cross-cluster revocation queue. Must be used with auto_rebuild&#x3D;true. | [optional] 
**DeltaRebuildInterval** | Pointer to **string** | The time between delta CRL rebuilds if a new revocation has occurred. Must be shorter than the CRL expiry. Defaults to 15m. | [optional] [default to "15m"]
**Disable** | Pointer to **bool** | If set to true, disables generating the CRL entirely. | [optional] 
**EnableDelta** | Pointer to **bool** | Whether to enable delta CRLs between authoritative CRL rebuilds | [optional] 
**Expiry** | Pointer to **string** | The amount of time the generated CRL should be valid; defaults to 72 hours | [optional] [default to "72h"]
**OcspDisable** | Pointer to **bool** | If set to true, ocsp unauthorized responses will be returned. | [optional] 
**OcspExpiry** | Pointer to **string** | The amount of time an OCSP response will be valid (controls the NextUpdate field); defaults to 12 hours | [optional] [default to "1h"]
**UnifiedCrl** | Pointer to **bool** | If set to true enables global replication of revocation entries, also enabling unified versions of OCSP and CRLs if their respective features are enabled. disable for CRLs and ocsp_disable for OCSP. | [optional] [default to false]
**UnifiedCrlOnExistingPaths** | Pointer to **bool** | If set to true, existing CRL and OCSP paths will return the unified CRL instead of a response based on cluster-local data | [optional] [default to false]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


