# TransitGenerateHmacRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \&quot;sha2-256\&quot;. | [optional] [default to "sha2-256"]
**BatchInput** | Pointer to **[]map[string]interface{}** | Specifies a list of items to be processed in a single batch. When this parameter is set, if the parameter &#x27;input&#x27; is also set, it will be ignored. Any batch output will preserve the order of the batch input. | [optional] 
**Input** | Pointer to **string** | The base64-encoded input data | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for generating the HMAC. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**Urlalgorithm** | Pointer to **string** | Algorithm to use (POST URL parameter) | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


