# TransitSignWithAlgorithmRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Deprecated: use \&quot;hash_algorithm\&quot; instead. | [optional] [default to "sha2-256"]
**BatchInput** | Pointer to **[]map[string]interface{}** | Specifies a list of items for processing. When this parameter is set, any supplied &#x27;input&#x27; or &#x27;context&#x27; parameters will be ignored. Responses are returned in the &#x27;batch_results&#x27; array component of the &#x27;data&#x27; element of the response. Any batch output will preserve the order of the batch input | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys. | [optional] 
**HashAlgorithm** | Pointer to **string** | Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 * none Defaults to \&quot;sha2-256\&quot;. Not valid for all key types, including ed25519. Using none requires setting prehashed&#x3D;true and signature_algorithm&#x3D;pkcs1v15, yielding a PKCSv1_5_NoOID instead of the usual PKCSv1_5_DERnull signature. | [optional] [default to "sha2-256"]
**Input** | Pointer to **string** | The base64-encoded input data | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for signing. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**MarshalingAlgorithm** | Pointer to **string** | The method by which to marshal the signature. The default is &#x27;asn1&#x27; which is used by openssl and X.509. It can also be set to &#x27;jws&#x27; which is used for JWT signatures; setting it to this will also cause the encoding of the signature to be url-safe base64 instead of using standard base64 encoding. Currently only valid for ECDSA P-256 key types\&quot;. | [optional] [default to "asn1"]
**Prehashed** | Pointer to **bool** | Set to &#x27;true&#x27; when the input is already hashed. If the key type is &#x27;rsa-2048&#x27;, &#x27;rsa-3072&#x27; or &#x27;rsa-4096&#x27;, then the algorithm used to hash the input should be indicated by the &#x27;algorithm&#x27; parameter. | [optional] 
**SaltLength** | Pointer to **string** | The salt length used to sign. Currently only applies to the RSA PSS signature scheme. Options are &#x27;auto&#x27; (the default used by Golang, causing the salt to be as large as possible when signing), &#x27;hash&#x27; (causes the salt length to equal the length of the hash used in the signature), or an integer between the minimum and the maximum permissible salt lengths for the given RSA key size. Defaults to &#x27;auto&#x27;. | [optional] [default to "auto"]
**SignatureAlgorithm** | Pointer to **string** | The signature algorithm to use for signing. Currently only applies to RSA key types. Options are &#x27;pss&#x27; or &#x27;pkcs1v15&#x27;. Defaults to &#x27;pss&#x27; | [optional] 



## Methods


### NewTransitSignWithAlgorithmRequest

`func NewTransitSignWithAlgorithmRequest() *TransitSignWithAlgorithmRequest`

NewTransitSignWithAlgorithmRequest instantiates a new TransitSignWithAlgorithmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitSignWithAlgorithmRequestWithDefaults

`func NewTransitSignWithAlgorithmRequestWithDefaults() *TransitSignWithAlgorithmRequest`

NewTransitSignWithAlgorithmRequestWithDefaults instantiates a new TransitSignWithAlgorithmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAlgorithm

`func (o *TransitSignWithAlgorithmRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TransitSignWithAlgorithmRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TransitSignWithAlgorithmRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### HasAlgorithm

`func (o *TransitSignWithAlgorithmRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.




### GetBatchInput

`func (o *TransitSignWithAlgorithmRequest) GetBatchInput() []map[string]interface{}`

GetBatchInput returns the BatchInput field if non-nil, zero value otherwise.

### GetBatchInputOk

`func (o *TransitSignWithAlgorithmRequest) GetBatchInputOk() (*[]map[string]interface{}, bool)`

GetBatchInputOk returns a tuple with the BatchInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchInput

`func (o *TransitSignWithAlgorithmRequest) SetBatchInput(v []map[string]interface{})`

SetBatchInput sets BatchInput field to given value.


### HasBatchInput

`func (o *TransitSignWithAlgorithmRequest) HasBatchInput() bool`

HasBatchInput returns a boolean if a field has been set.




### GetContext

`func (o *TransitSignWithAlgorithmRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitSignWithAlgorithmRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitSignWithAlgorithmRequest) SetContext(v string)`

SetContext sets Context field to given value.


### HasContext

`func (o *TransitSignWithAlgorithmRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.




### GetHashAlgorithm

`func (o *TransitSignWithAlgorithmRequest) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *TransitSignWithAlgorithmRequest) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *TransitSignWithAlgorithmRequest) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.


### HasHashAlgorithm

`func (o *TransitSignWithAlgorithmRequest) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.




### GetInput

`func (o *TransitSignWithAlgorithmRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TransitSignWithAlgorithmRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TransitSignWithAlgorithmRequest) SetInput(v string)`

SetInput sets Input field to given value.


### HasInput

`func (o *TransitSignWithAlgorithmRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.




### GetKeyVersion

`func (o *TransitSignWithAlgorithmRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *TransitSignWithAlgorithmRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *TransitSignWithAlgorithmRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.


### HasKeyVersion

`func (o *TransitSignWithAlgorithmRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.




### GetMarshalingAlgorithm

`func (o *TransitSignWithAlgorithmRequest) GetMarshalingAlgorithm() string`

GetMarshalingAlgorithm returns the MarshalingAlgorithm field if non-nil, zero value otherwise.

### GetMarshalingAlgorithmOk

`func (o *TransitSignWithAlgorithmRequest) GetMarshalingAlgorithmOk() (*string, bool)`

GetMarshalingAlgorithmOk returns a tuple with the MarshalingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarshalingAlgorithm

`func (o *TransitSignWithAlgorithmRequest) SetMarshalingAlgorithm(v string)`

SetMarshalingAlgorithm sets MarshalingAlgorithm field to given value.


### HasMarshalingAlgorithm

`func (o *TransitSignWithAlgorithmRequest) HasMarshalingAlgorithm() bool`

HasMarshalingAlgorithm returns a boolean if a field has been set.




### GetPrehashed

`func (o *TransitSignWithAlgorithmRequest) GetPrehashed() bool`

GetPrehashed returns the Prehashed field if non-nil, zero value otherwise.

### GetPrehashedOk

`func (o *TransitSignWithAlgorithmRequest) GetPrehashedOk() (*bool, bool)`

GetPrehashedOk returns a tuple with the Prehashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashed

`func (o *TransitSignWithAlgorithmRequest) SetPrehashed(v bool)`

SetPrehashed sets Prehashed field to given value.


### HasPrehashed

`func (o *TransitSignWithAlgorithmRequest) HasPrehashed() bool`

HasPrehashed returns a boolean if a field has been set.




### GetSaltLength

`func (o *TransitSignWithAlgorithmRequest) GetSaltLength() string`

GetSaltLength returns the SaltLength field if non-nil, zero value otherwise.

### GetSaltLengthOk

`func (o *TransitSignWithAlgorithmRequest) GetSaltLengthOk() (*string, bool)`

GetSaltLengthOk returns a tuple with the SaltLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLength

`func (o *TransitSignWithAlgorithmRequest) SetSaltLength(v string)`

SetSaltLength sets SaltLength field to given value.


### HasSaltLength

`func (o *TransitSignWithAlgorithmRequest) HasSaltLength() bool`

HasSaltLength returns a boolean if a field has been set.




### GetSignatureAlgorithm

`func (o *TransitSignWithAlgorithmRequest) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *TransitSignWithAlgorithmRequest) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *TransitSignWithAlgorithmRequest) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### HasSignatureAlgorithm

`func (o *TransitSignWithAlgorithmRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


