# TransitVerifyWithAlgorithmRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Deprecated: use \&quot;hash_algorithm\&quot; instead. | [optional] [default to "sha2-256"]
**BatchInput** | Pointer to **[]map[string]interface{}** | Specifies a list of items for processing. When this parameter is set, any supplied &#x27;input&#x27;, &#x27;hmac&#x27; or &#x27;signature&#x27; parameters will be ignored. Responses are returned in the &#x27;batch_results&#x27; array component of the &#x27;data&#x27; element of the response. Any batch output will preserve the order of the batch input | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys. | [optional] 
**HashAlgorithm** | Pointer to **string** | Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 * none Defaults to \&quot;sha2-256\&quot;. Not valid for all key types. See note about none on signing path. | [optional] [default to "sha2-256"]
**Hmac** | Pointer to **string** | The HMAC, including vault header/key version | [optional] 
**Input** | Pointer to **string** | The base64-encoded input data to verify | [optional] 
**MarshalingAlgorithm** | Pointer to **string** | The method by which to unmarshal the signature when verifying. The default is &#x27;asn1&#x27; which is used by openssl and X.509; can also be set to &#x27;jws&#x27; which is used for JWT signatures in which case the signature is also expected to be url-safe base64 encoding instead of standard base64 encoding. Currently only valid for ECDSA P-256 key types\&quot;. | [optional] [default to "asn1"]
**Prehashed** | Pointer to **bool** | Set to &#x27;true&#x27; when the input is already hashed. If the key type is &#x27;rsa-2048&#x27;, &#x27;rsa-3072&#x27; or &#x27;rsa-4096&#x27;, then the algorithm used to hash the input should be indicated by the &#x27;algorithm&#x27; parameter. | [optional] 
**SaltLength** | Pointer to **string** | The salt length used to sign. Currently only applies to the RSA PSS signature scheme. Options are &#x27;auto&#x27; (the default used by Golang, causing the salt to be as large as possible when signing), &#x27;hash&#x27; (causes the salt length to equal the length of the hash used in the signature), or an integer between the minimum and the maximum permissible salt lengths for the given RSA key size. Defaults to &#x27;auto&#x27;. | [optional] [default to "auto"]
**Signature** | Pointer to **string** | The signature, including vault header/key version | [optional] 
**SignatureAlgorithm** | Pointer to **string** | The signature algorithm to use for signature verification. Currently only applies to RSA key types. Options are &#x27;pss&#x27; or &#x27;pkcs1v15&#x27;. Defaults to &#x27;pss&#x27; | [optional] 



## Methods


### NewTransitVerifyWithAlgorithmRequest

`func NewTransitVerifyWithAlgorithmRequest() *TransitVerifyWithAlgorithmRequest`

NewTransitVerifyWithAlgorithmRequest instantiates a new TransitVerifyWithAlgorithmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitVerifyWithAlgorithmRequestWithDefaults

`func NewTransitVerifyWithAlgorithmRequestWithDefaults() *TransitVerifyWithAlgorithmRequest`

NewTransitVerifyWithAlgorithmRequestWithDefaults instantiates a new TransitVerifyWithAlgorithmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TransitVerifyWithAlgorithmRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### HasAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.




### GetBatchInput

`func (o *TransitVerifyWithAlgorithmRequest) GetBatchInput() []map[string]interface{}`

GetBatchInput returns the BatchInput field if non-nil, zero value otherwise.

### GetBatchInputOk

`func (o *TransitVerifyWithAlgorithmRequest) GetBatchInputOk() (*[]map[string]interface{}, bool)`

GetBatchInputOk returns a tuple with the BatchInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchInput

`func (o *TransitVerifyWithAlgorithmRequest) SetBatchInput(v []map[string]interface{})`

SetBatchInput sets BatchInput field to given value.


### HasBatchInput

`func (o *TransitVerifyWithAlgorithmRequest) HasBatchInput() bool`

HasBatchInput returns a boolean if a field has been set.




### GetContext

`func (o *TransitVerifyWithAlgorithmRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitVerifyWithAlgorithmRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitVerifyWithAlgorithmRequest) SetContext(v string)`

SetContext sets Context field to given value.


### HasContext

`func (o *TransitVerifyWithAlgorithmRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.




### GetHashAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *TransitVerifyWithAlgorithmRequest) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.


### HasHashAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.




### GetHmac

`func (o *TransitVerifyWithAlgorithmRequest) GetHmac() string`

GetHmac returns the Hmac field if non-nil, zero value otherwise.

### GetHmacOk

`func (o *TransitVerifyWithAlgorithmRequest) GetHmacOk() (*string, bool)`

GetHmacOk returns a tuple with the Hmac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmac

`func (o *TransitVerifyWithAlgorithmRequest) SetHmac(v string)`

SetHmac sets Hmac field to given value.


### HasHmac

`func (o *TransitVerifyWithAlgorithmRequest) HasHmac() bool`

HasHmac returns a boolean if a field has been set.




### GetInput

`func (o *TransitVerifyWithAlgorithmRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TransitVerifyWithAlgorithmRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TransitVerifyWithAlgorithmRequest) SetInput(v string)`

SetInput sets Input field to given value.


### HasInput

`func (o *TransitVerifyWithAlgorithmRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.




### GetMarshalingAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) GetMarshalingAlgorithm() string`

GetMarshalingAlgorithm returns the MarshalingAlgorithm field if non-nil, zero value otherwise.

### GetMarshalingAlgorithmOk

`func (o *TransitVerifyWithAlgorithmRequest) GetMarshalingAlgorithmOk() (*string, bool)`

GetMarshalingAlgorithmOk returns a tuple with the MarshalingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarshalingAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) SetMarshalingAlgorithm(v string)`

SetMarshalingAlgorithm sets MarshalingAlgorithm field to given value.


### HasMarshalingAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) HasMarshalingAlgorithm() bool`

HasMarshalingAlgorithm returns a boolean if a field has been set.




### GetPrehashed

`func (o *TransitVerifyWithAlgorithmRequest) GetPrehashed() bool`

GetPrehashed returns the Prehashed field if non-nil, zero value otherwise.

### GetPrehashedOk

`func (o *TransitVerifyWithAlgorithmRequest) GetPrehashedOk() (*bool, bool)`

GetPrehashedOk returns a tuple with the Prehashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashed

`func (o *TransitVerifyWithAlgorithmRequest) SetPrehashed(v bool)`

SetPrehashed sets Prehashed field to given value.


### HasPrehashed

`func (o *TransitVerifyWithAlgorithmRequest) HasPrehashed() bool`

HasPrehashed returns a boolean if a field has been set.




### GetSaltLength

`func (o *TransitVerifyWithAlgorithmRequest) GetSaltLength() string`

GetSaltLength returns the SaltLength field if non-nil, zero value otherwise.

### GetSaltLengthOk

`func (o *TransitVerifyWithAlgorithmRequest) GetSaltLengthOk() (*string, bool)`

GetSaltLengthOk returns a tuple with the SaltLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLength

`func (o *TransitVerifyWithAlgorithmRequest) SetSaltLength(v string)`

SetSaltLength sets SaltLength field to given value.


### HasSaltLength

`func (o *TransitVerifyWithAlgorithmRequest) HasSaltLength() bool`

HasSaltLength returns a boolean if a field has been set.




### GetSignature

`func (o *TransitVerifyWithAlgorithmRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TransitVerifyWithAlgorithmRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TransitVerifyWithAlgorithmRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.


### HasSignature

`func (o *TransitVerifyWithAlgorithmRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.




### GetSignatureAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *TransitVerifyWithAlgorithmRequest) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### HasSignatureAlgorithm

`func (o *TransitVerifyWithAlgorithmRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


