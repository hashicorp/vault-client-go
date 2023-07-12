# CloudFoundryConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfApiAddr** | Pointer to **string** | CF’s API address. | [optional] 
**CfApiMutualTlsCertificate** | Pointer to **string** | The PEM-format certificates that are presented for mutual TLS with the CloudFoundry API. If not set, mutual TLS is not used | [optional] 
**CfApiMutualTlsKey** | Pointer to **string** | The PEM-format private key that are used for mutual TLS with the CloudFoundry API. If not set, mutual TLS is not used | [optional] 
**CfApiTrustedCertificates** | Pointer to **[]string** | The PEM-format CA certificates that are acceptable for the CF API to present. | [optional] 
**CfClientId** | Pointer to **string** | The client id for CF’s API. | [optional] 
**CfClientSecret** | Pointer to **string** | The client secret for CF’s API. | [optional] 
**CfPassword** | Pointer to **string** | The password for CF’s API. | [optional] 
**CfUsername** | Pointer to **string** | The username for CF’s API. | [optional] 
**IdentityCaCertificates** | Pointer to **[]string** | The PEM-format CA certificates that are required to have issued the instance certificates presented for logging in. | [optional] 
**LoginMaxSecondsNotAfter** | Pointer to **int32** | Duration in seconds for the maximum acceptable length in the future a \&quot;signing_time\&quot; can be. Useful for clock drift. Set low to reduce the opportunity for replay attacks. | [optional] [default to 60]
**LoginMaxSecondsNotBefore** | Pointer to **string** | Duration in seconds for the maximum acceptable age of a \&quot;signing_time\&quot;. Useful for clock drift. Set low to reduce the opportunity for replay attacks. | [optional] [default to "300"]
**PcfApiAddr** | Pointer to **string** | Deprecated. Please use \&quot;cf_api_addr\&quot;. | [optional] 
**PcfApiTrustedCertificates** | Pointer to **[]string** | Deprecated. Please use \&quot;cf_api_trusted_certificates\&quot;. | [optional] 
**PcfPassword** | Pointer to **string** | Deprecated. Please use \&quot;cf_password\&quot;. | [optional] 
**PcfUsername** | Pointer to **string** | Deprecated. Please use \&quot;cf_username\&quot;. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


