# AwsLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IamHttpRequestMethod** | Pointer to **string** | HTTP method to use for the AWS request when auth_type is iam. This must match what has been signed in the presigned request. | [optional] 
**IamRequestBody** | Pointer to **string** | Base64-encoded request body when auth_type is iam. This must match the request body included in the signature. | [optional] 
**IamRequestHeaders** | Pointer to **string** | Key/value pairs of headers for use in the sts:GetCallerIdentity HTTP requests headers when auth_type is iam. Can be either a Base64-encoded, JSON-serialized string, or a JSON object of key/value pairs. This must at a minimum include the headers over which AWS has included a signature. | [optional] 
**IamRequestUrl** | Pointer to **string** | Base64-encoded full URL against which to make the AWS request when using iam auth_type. | [optional] 
**Identity** | Pointer to **string** | Base64 encoded EC2 instance identity document. This needs to be supplied along with the &#x27;signature&#x27; parameter. If using &#x27;curl&#x27; for fetching the identity document, consider using the option &#x27;-w 0&#x27; while piping the output to &#x27;base64&#x27; binary. | [optional] 
**Nonce** | Pointer to **string** | The nonce to be used for subsequent login requests when auth_type is ec2. If this parameter is not specified at all and if reauthentication is allowed, then the backend will generate a random nonce, attaches it to the instance&#x27;s identity access list entry and returns the nonce back as part of auth metadata. This value should be used with further login requests, to establish client authenticity. Clients can choose to set a custom nonce if preferred, in which case, it is recommended that clients provide a strong nonce. If a nonce is provided but with an empty value, it indicates intent to disable reauthentication. Note that, when &#x27;disallow_reauthentication&#x27; option is enabled on either the role or the role tag, the &#x27;nonce&#x27; holds no significance. | [optional] 
**Pkcs7** | Pointer to **string** | PKCS7 signature of the identity document when using an auth_type of ec2. | [optional] 
**Role** | Pointer to **string** | Name of the role against which the login is being attempted. If &#x27;role&#x27; is not specified, then the login endpoint looks for a role bearing the name of the AMI ID of the EC2 instance that is trying to login. If a matching role is not found, login fails. | [optional] 
**Signature** | Pointer to **string** | Base64 encoded SHA256 RSA signature of the instance identity document. This needs to be supplied along with &#x27;identity&#x27; parameter. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


