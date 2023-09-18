# AwsConfigureCertificateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsPublicCert** | Pointer to **string** | Base64 encoded AWS Public cert required to verify PKCS7 signature of the EC2 instance metadata. | [optional] 
**Type** | Pointer to **string** | Takes the value of either \&quot;pkcs7\&quot; or \&quot;identity\&quot;, indicating the type of document which can be verified using the given certificate. The reason is that the PKCS#7 document will have a DSA digest and the identity signature will have an RSA signature, and accordingly the public certificates to verify those also vary. Defaults to \&quot;pkcs7\&quot;. | [optional] [default to "pkcs7"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


