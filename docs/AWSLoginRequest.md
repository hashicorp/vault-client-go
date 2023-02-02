# AWSLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**IamHttpRequestMethod** | Pointer to **string** | HTTP method to use for the AWS request when auth_type is iam. This must match what has been signed in the presigned request. Currently, POST is the only supported value | [optional] 
**IamRequestBody** | Pointer to **string** | Base64-encoded request body when auth_type is iam. This must match the request body included in the signature. | [optional] 
**IamRequestHeaders** | Pointer to **string** | Key/value pairs of headers for use in the sts:GetCallerIdentity HTTP requests headers when auth_type is iam. Can be either a Base64-encoded, JSON-serialized string, or a JSON object of key/value pairs. This must at a minimum include the headers over which AWS has included a signature. | [optional] 
**IamRequestUrl** | Pointer to **string** | Base64-encoded full URL against which to make the AWS request when using iam auth_type. | [optional] 
**Identity** | Pointer to **string** | Base64 encoded EC2 instance identity document. This needs to be supplied along with the &#x27;signature&#x27; parameter. If using &#x27;curl&#x27; for fetching the identity document, consider using the option &#x27;-w 0&#x27; while piping the output to &#x27;base64&#x27; binary. | [optional] 
**Nonce** | Pointer to **string** | The nonce to be used for subsequent login requests when auth_type is ec2. If this parameter is not specified at all and if reauthentication is allowed, then the backend will generate a random nonce, attaches it to the instance&#x27;s identity access list entry and returns the nonce back as part of auth metadata. This value should be used with further login requests, to establish client authenticity. Clients can choose to set a custom nonce if preferred, in which case, it is recommended that clients provide a strong nonce. If a nonce is provided but with an empty value, it indicates intent to disable reauthentication. Note that, when &#x27;disallow_reauthentication&#x27; option is enabled on either the role or the role tag, the &#x27;nonce&#x27; holds no significance. | [optional] 
**Pkcs7** | Pointer to **string** | PKCS7 signature of the identity document when using an auth_type of ec2. | [optional] 
**Role** | Pointer to **string** | Name of the role against which the login is being attempted. If &#x27;role&#x27; is not specified, then the login endpoint looks for a role bearing the name of the AMI ID of the EC2 instance that is trying to login. If a matching role is not found, login fails. | [optional] 
**Signature** | Pointer to **string** | Base64 encoded SHA256 RSA signature of the instance identity document. This needs to be supplied along with &#x27;identity&#x27; parameter. | [optional] 



## Methods


### NewAWSLoginRequest

`func NewAWSLoginRequest() *AWSLoginRequest`

NewAWSLoginRequest instantiates a new AWSLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSLoginRequestWithDefaults

`func NewAWSLoginRequestWithDefaults() *AWSLoginRequest`

NewAWSLoginRequestWithDefaults instantiates a new AWSLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetIamHttpRequestMethod

`func (o *AWSLoginRequest) GetIamHttpRequestMethod() string`

GetIamHttpRequestMethod returns the IamHttpRequestMethod field if non-nil, zero value otherwise.

### GetIamHttpRequestMethodOk

`func (o *AWSLoginRequest) GetIamHttpRequestMethodOk() (*string, bool)`

GetIamHttpRequestMethodOk returns a tuple with the IamHttpRequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamHttpRequestMethod

`func (o *AWSLoginRequest) SetIamHttpRequestMethod(v string)`

SetIamHttpRequestMethod sets IamHttpRequestMethod field to given value.


### HasIamHttpRequestMethod

`func (o *AWSLoginRequest) HasIamHttpRequestMethod() bool`

HasIamHttpRequestMethod returns a boolean if a field has been set.




### GetIamRequestBody

`func (o *AWSLoginRequest) GetIamRequestBody() string`

GetIamRequestBody returns the IamRequestBody field if non-nil, zero value otherwise.

### GetIamRequestBodyOk

`func (o *AWSLoginRequest) GetIamRequestBodyOk() (*string, bool)`

GetIamRequestBodyOk returns a tuple with the IamRequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRequestBody

`func (o *AWSLoginRequest) SetIamRequestBody(v string)`

SetIamRequestBody sets IamRequestBody field to given value.


### HasIamRequestBody

`func (o *AWSLoginRequest) HasIamRequestBody() bool`

HasIamRequestBody returns a boolean if a field has been set.




### GetIamRequestHeaders

`func (o *AWSLoginRequest) GetIamRequestHeaders() string`

GetIamRequestHeaders returns the IamRequestHeaders field if non-nil, zero value otherwise.

### GetIamRequestHeadersOk

`func (o *AWSLoginRequest) GetIamRequestHeadersOk() (*string, bool)`

GetIamRequestHeadersOk returns a tuple with the IamRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRequestHeaders

`func (o *AWSLoginRequest) SetIamRequestHeaders(v string)`

SetIamRequestHeaders sets IamRequestHeaders field to given value.


### HasIamRequestHeaders

`func (o *AWSLoginRequest) HasIamRequestHeaders() bool`

HasIamRequestHeaders returns a boolean if a field has been set.




### GetIamRequestUrl

`func (o *AWSLoginRequest) GetIamRequestUrl() string`

GetIamRequestUrl returns the IamRequestUrl field if non-nil, zero value otherwise.

### GetIamRequestUrlOk

`func (o *AWSLoginRequest) GetIamRequestUrlOk() (*string, bool)`

GetIamRequestUrlOk returns a tuple with the IamRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRequestUrl

`func (o *AWSLoginRequest) SetIamRequestUrl(v string)`

SetIamRequestUrl sets IamRequestUrl field to given value.


### HasIamRequestUrl

`func (o *AWSLoginRequest) HasIamRequestUrl() bool`

HasIamRequestUrl returns a boolean if a field has been set.




### GetIdentity

`func (o *AWSLoginRequest) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AWSLoginRequest) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AWSLoginRequest) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### HasIdentity

`func (o *AWSLoginRequest) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.




### GetNonce

`func (o *AWSLoginRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *AWSLoginRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *AWSLoginRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *AWSLoginRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.




### GetPkcs7

`func (o *AWSLoginRequest) GetPkcs7() string`

GetPkcs7 returns the Pkcs7 field if non-nil, zero value otherwise.

### GetPkcs7Ok

`func (o *AWSLoginRequest) GetPkcs7Ok() (*string, bool)`

GetPkcs7Ok returns a tuple with the Pkcs7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs7

`func (o *AWSLoginRequest) SetPkcs7(v string)`

SetPkcs7 sets Pkcs7 field to given value.


### HasPkcs7

`func (o *AWSLoginRequest) HasPkcs7() bool`

HasPkcs7 returns a boolean if a field has been set.




### GetRole

`func (o *AWSLoginRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AWSLoginRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AWSLoginRequest) SetRole(v string)`

SetRole sets Role field to given value.


### HasRole

`func (o *AWSLoginRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.




### GetSignature

`func (o *AWSLoginRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AWSLoginRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AWSLoginRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.


### HasSignature

`func (o *AWSLoginRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


